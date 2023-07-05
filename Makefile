build:
	@docker buildx build --push --platform linux/arm64 -t dinhlockt02/cs_video_call_app_server:dev-alpine . \
	&& docker buildx build --push --platform linux/amd64 -t dinhlockt02/cs_video_call_app_server:dev-alpine-amd .
livekit:
	livekit-server --dev --config livekit.yaml
