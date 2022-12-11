CURDIR      = $(shell pwd)
APIDIR      = ${CURDIR}/api
BINDIR      = ${APIDIR}/bin
MODELDIR    = ${CURDIR}/model
GOCTL       = goctl
TARGET      = fms
TAG         = latest
DOCKERFILE  = fms.Dockerfile
IMG_CONTEXT = .

all: tidy codegen modelgen
	cd ${APIDIR} && go build -o ${BINDIR}/${TARGET}

tidy:
	cd ${CURDIR}../../ && go mod tidy

codegen:
	cd ${APIDIR} && ${GOCTL} api go -api fms.api -dir .

modelgen:
	cd ${MODELDIR} && ${GOCTL} model mysql ddl -src fms.sql -dir .