CURDIR      = $(shell pwd)
APIDIR      = ${CURDIR}/api
BINDIR      = ${APIDIR}/bin
GOCTL       = goctl
TARGET      = fms
TAG	        = latest
DOCKERFILE  = fms.Dockerfile
IMG_CONTEXT = .

all: codegen
	cd ${APIDIR} && go build -o ${BINDIR}/${TARGET}

codegen:
	cd ${APIDIR} && ${GOCTL} api go -api fms.api -dir .

${TARGET}-image:
	cd ${CURDIR} && docker build -t ${TARGET}:${TAG} -f ${DOCKERFILE} ${IMG_CONTEXT}