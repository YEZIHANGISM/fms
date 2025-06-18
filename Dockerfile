# 第一阶段：构建
FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o fms .

# 第二阶段：精简运行环境
FROM alpine:3.20

WORKDIR /ism

COPY --from=builder /build/fms .
COPY --from=builder /build/configs/fms.yaml ./configs/fms.yaml

EXPOSE 8080

ENTRYPOINT ["/ism/fms"]
