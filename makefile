APP_NAME = "app"
GO_MOD_LINE = $(shell head -n 1 go.mod | cut -c 8-)
GO_MOD_NAME = ${GO_MOD_LINE}
CONF_PATH = ${GO_MOD_NAME}/internal/conf
VERSION = $(shell cat VERSION)
BUILD_TIME = $(shell date +'%Y-%m-%d_%T')
SHA1 = $(shell git rev-parse HEAD)

BUILD_DIR=build
update-doc: 
	swag init -d api
build:
	go build -ldflags '-X ${CONF_PATH}.Version=${VERSION} -X ${CONF_PATH}.Sha1=${SHA1} -X ${CONF_PATH}.BuildTime=${BUILD_TIME}' -o ${BUILD_DIR}/${APP_NAME} .
clean:
	rm -rf ${BUILD_DIR}