build-images:
	docker build -t sidecar-app:v1 app
	docker build -t sidecar-proxy:v1 sidecar_proxy