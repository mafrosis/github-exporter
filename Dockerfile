FROM hub.artifactory.gcp.anz/golang:1.15-alpine as build
RUN sed -i -e "s~http://dl-cdn.alpinelinux.org~https://artifactory.gcp.anz/artifactory/alpinelinux~g" /etc/apk/repositories && \
	apk add --no-cache git make

WORKDIR /go/exporter
COPY . /go/exporter

ENV GOPROXY=https://artifactory.gcp.anz/artifactory/go
ENV GO111MODULE=on

RUN make build


# Minimal base image with ca-certs and tzdata
# https://github.com/GoogleContainerTools/distroless/blob/master/base/README.md
FROM gcr.artifactory.gcp.anz/distroless/static

COPY --from=build /go/exporter/bin/ghe-exporter /usr/bin/ghe-exporter

EXPOSE 9504

CMD ["ghe-exporter"]
