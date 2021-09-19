build-proto:
	protoc -I=go_app/transport --go_out=plugins=grpc:go_app/transport go_app/transport/go_app.proto

build-images-go:
	docker build -t sidecar-go-app:v1 go_app
	docker build -t sidecar-gp-proxy:v1 go_sidecar_proxy	

build-images:
	docker build -t sidecar-app:v1 app
	docker build -t sidecar-proxy:v1 sidecar_proxy