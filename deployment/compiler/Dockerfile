FROM golang:latest

RUN export GO111MODULE=on \
    && export GOPROXY=https://goproxy.cn/,direct \
    && go install github.com/zeromicro/go-zero/tools/goctl@latest \
    && goctl env check -i -f --verbose

CMD /bin/bash
