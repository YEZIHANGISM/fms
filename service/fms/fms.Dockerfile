FROM golang:latest

RUN mkdir -p /ism/fms/bin \
    && mkdir -p /ism/fms/etc

WORKDIR /ism/fms

COPY api/bin/fms /ism/fms/bin/
COPY api/etc/fms.yaml /ism/fms/etc/

ENTRYPOINT [ "./bin/fms", "-f", "etc/fms.yaml" ]