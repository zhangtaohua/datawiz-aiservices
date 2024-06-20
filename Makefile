.DEFAULT_GOAL := linux_local_producer

COMPILE_LDFLAGS="-s -w"
APP_NAME=starwiz_ai_service
IMAGE_NAME=starwiz_ai_go
TEST_CONTAINER_NAME=test_starwiz_ai


PACKAGES=`go list ./... | grep -v vendor | grep -v mocks`

fmt:
	for pkg in ${PACKAGES}; do \
		go fmt $$pkg; \
	done;

docker_builder:
	docker build --target builder -t ${IMAGE_NAME}:build -f ./build/docker/Dockerfile.build .

docker_producer:
	docker build -t ${IMAGE_NAME}:prod -f ./build/docker/Dockerfile.build .
	docker save -o ./build/images/${IMAGE_NAME}.tar.gz ${IMAGE_NAME}:prod

windows_local_producer:windows_local_builder
	docker build -t ${IMAGE_NAME}:local_prod -f ./build/docker/Dockerfile.local .
	docker save -o ./build/images/${IMAGE_NAME}.tar.gz ${IMAGE_NAME}:local_prod

linux_local_producer:linux_local_builder
	docker build -t ${IMAGE_NAME}:local_prod -f ./build/docker/Dockerfile.local .
	docker save -o ./build/images/${IMAGE_NAME}.tar.gz ${IMAGE_NAME}:local_prod

macos_local_producer:macos_local_builder
	docker build -t ${IMAGE_NAME}:local_prod -f ./build/docker/Dockerfile.local .
	docker save -o ./build/images/${IMAGE_NAME}.tar.gz ${IMAGE_NAME}:local_prod

windows_local_builder:
	$(info "start build...")
	set CGO_ENABLED=0\
	&& set GOOS=windows\
	&& set GOARCH=amd64\
	&& go build -ldflags=${COMPILE_LDFLAGS} -o build/bin/${APP_NAME}.exe

linux_local_builder:
	$(info "start build...")
	set CGO_ENABLED=0\
	&& set GOOS=linux\
	&& set GOARCH=amd64\
	&& go build -ldflags=${COMPILE_LDFLAGS} -o build/bin/${APP_NAME}

macos_local_builder:
	set CGO_ENABLED=0\
	&& set GOOS=darwin\
	&& GOARCH=amd64\
	&& go build -ldflags=${COMPILE_LDFLAGS} -o build/bin/${APP_NAME}

test:
	docker run -d --name ${TEST_CONTAINER_NAME} -v /d/Work/Golang/run/aigo/.env:/app/.env -p 8088:8088  ${IMAGE_NAME}:local_prod

clean_test:
	docker stop ${TEST_CONTAINER_NAME}
	docker container rm ${TEST_CONTAINER_NAME}
	docker rmi ${IMAGE_NAME}:local_prod
	docker rmi ${IMAGE_NAME}:prod

clean_bin:
	rm -rf ./build/bin/*

clean_docker_cache:
	docker builder prune

clean:clean_test clean_bin clean_docker_cache
