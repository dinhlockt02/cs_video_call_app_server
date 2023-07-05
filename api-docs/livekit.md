# How to run Livekit correctly at localhost
## Install Livekit Server
You can install livekit server at `https://docs.livekit.io/getting-started/server-setup/`
## Download `livekit.yaml`
You can copy the content of the file [here](https://github.com/dinhlockt02/cs_video_call_app_server/blob/main/livekit.yaml)
### Execute command
`livekit-server --dev --config /path/to/file/livekit.yaml --bind 0.0.0.0`
### Why
Because I used the webhook to do some use cases, we have to configure the livekit to specify the webhook url
