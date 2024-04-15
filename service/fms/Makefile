CURDIR      = $(shell pwd)
APIDIR      = ${CURDIR}/api
BINDIR      = ${APIDIR}/bin
MODELDIR    = ${CURDIR}/model
GOCTL       = goctl
TARGET      = fms
TAG         = latest
DOCKERFILE  = fms.Dockerfile
IMG_CONTEXT = .
PG_URL      = "postgres://postgres:postgres@127.0.0.1:5432/fms?sslmode=disable"

all: tidy codegen modelgen
	cd ${APIDIR} && rm -f ${BINDIR}/${TARGET} && go build -o ${BINDIR}/${TARGET}

tidy:
	cd ${CURDIR}../../ && rm -f go.sum && go mod tidy

codegen:
	cd ${APIDIR} && ${GOCTL} api go -api fms.api -dir .

modelgen:
	# cd ${MODELDIR} && ${GOCTL} model mysql ddl -src fms.sql -dir .
	cd ${MODELDIR} && rm -f ${MODELDIR}/* && ${GOCTL} model pg datasource --table="*" --url ${PG_URL} -d .