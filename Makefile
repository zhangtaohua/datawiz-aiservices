.DEFAULT_GOAL := build


export DATE := $(shell date +%Y%m%d-%H:%M:%S)
export LATEST_COMMIT := $(shell git log --pretty=format:'%h' -n 1)
export BRANCH := $(shell git branch |grep -v "no branch"| grep \*|cut -d ' ' -f2)
export BUILT_ON_IP := $(shell [ $$(uname) = Linux ] && hostname -i || hostname )
export RUNTIME_VER := $(shell go version)
export BUILT_ON_OS=$(shell uname -a)
ifeq ($(BRANCH),)
BRANCH := master
endif

export COMMIT_CNT := $(shell git rev-list HEAD | wc -l | sed 's/ //g' )
export BUILD_NUMBER := ${BRANCH}-${COMMIT_CNT}
export COMPILE_LDFLAGS='-s -w \
                          -X "main.BuildDate=${DATE}" \
                          -X "main.LatestCommit=${LATEST_COMMIT}" \
                          -X "main.BuildNumber=${BUILD_NUMBER}" \
                          -X "main.BuiltOnIP=${BUILT_ON_IP}" \
                          -X "main.BuiltOnOs=${BUILT_ON_OS}" \
                          -X "main.Branch=${BRANCH}" \
                          -X "main.CommitCnt=${COMMIT_CNT}" \
                          -X "main.RuntimeVer=${RUNTIME_VER}" '


PACKAGES=`go list ./... | grep -v vendor | grep -v mocks`

fmt:
	for pkg in ${PACKAGES}; do \
		go fmt $$pkg; \
	done;

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server -ldflags=${COMPILE_LDFLAGS} main.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server-db-migrate -ldflags=${COMPILE_LDFLAGS} cmd/dbmigrate/main.go

air:
	env GOOS=linux GOARCH=amd64 go build -o tmp/main -ldflags=${COMPILE_LDFLAGS} main.go

macos:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/server-darwin -ldflags=${COMPILE_LDFLAGS} main.go

swag:
	swag init -g main.go
	
clean:
	rm -rf ./bin/*