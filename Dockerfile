ARG GO_VERSION=1.13.6
ARG ALPINE_VERSION=3.11.3

FROM golang:${GO_VERSION}-alpine as builder

# ENV PORT
# ENV USERNAME
# ENV TWEET_COUNT
# ENV TIMELINE
# ENV HOST
# ENV PORT
# ENV USERNAME

# Download dependencies (this is done in a seperate layer to take advantage
# of Docker layer caching to decrease build times)
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy source files
COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix 'static' -o twitter-service .

# Use alpine linux base, with latest CA roots installed
FROM alpine:${ALPINE_VERSION}
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /src/twitter-service /usr/local/bin/twitter-service

EXPOSE 80
USER nobody
CMD /usr/local/bin/twitter-service
