# Twitter Service

A simple JSON micro-service for fetching tweets.

## Configuration

Configuraton is handled through the following environment varibles:

### Twitter Timeline
| Environment Variable |                                        Description                                         | Default Value |
|----------------------|--------------------------------------------------------------------------------------------|---------------|
| `USERNAME`           | The twitter user name for the timeline. Must match the twitter credentials. (**Required**) |               |
| `TWEET_COUNT`        | The number of tweets to keep.                                                              | `10`          |
| `TIMELINE`           | The twitter timeline to include. Either `"user"` or `"home"`.                              | `"home"`      |

### API Configuration
|  Environment Variable |         Description         | Default Value |
|-----------------------|-----------------------------|---------------|
| `HOST`                | The host to listen on.      | `""` (all)    |
| `PORT`                | The HTTP port to listen on. | `3000`        |

### Twitter OAuth Credentials
| Environment Variable |                    Description                    |
|----------------------|---------------------------------------------------|
| `CONSUMER_KEY`       | Twitter application consumer key (API Key).       |
| `CONSUMER_SECRET`    | Twitter application consumer secret (API Secret). |
| `ACCESS_TOKEN`       | Twitter user access token.                        |
| `ACCESS_SECRET`      | Twitter user access token scret.                  |


## Run

```shell
env PORT=8000 \
    USERNAME=handle \
    TIMELINE=home \
    CONSUMER_KEY=XXXXXXXXXXXXXXXXXXXXXXXX \
    CONSUMER_SECRET=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX \
    ACCESS_TOKEN=XXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX \
    ACCESS_SECRET=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX \
    twitter-service
```
