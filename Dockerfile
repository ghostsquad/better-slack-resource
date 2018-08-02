# stage: base
FROM alpine:3.8 as base

RUN set -ex; \
    apk add --update --no-cache \
        bash \
    ; \
    rm -rf /var/cache/apk/*;

# stage: build
FROM golang:alpine3.8 as build

WORKDIR /go/src/app
COPY glide.yaml .

RUN set -ex; \
    glide install

RUN set -ex; \
    go build -o dist/out cmd/out/out.go

# stage: final
FROM base

COPY --from=build /go/src/app/cmd/check/check /opt/resource/check
COPY --from=build /go/src/app/cmd/in/in /opt/resource/in
COPY --from=build /go/src/app/dist/out /opt/resource/out

