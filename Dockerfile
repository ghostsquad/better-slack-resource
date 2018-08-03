# stage: base
FROM alpine:3.8 as base

RUN set -ex; \
    apk add --update --no-cache \
        bash \
    ; \
    rm -rf /var/cache/apk/*;

# stage: build
FROM golang:alpine3.8 as build

COPY --from=instrumentisto/dep:0.5.0 /usr/local/bin/dep /usr/local/bin/dep

WORKDIR /go/src/app

COPY Gopkg.* .

RUN dep ensure -vendor-only

COPY . .

RUN set -ex; \
    go build -o dist/out cmd/out/out.go

# stage: final
FROM base

COPY --from=build /go/src/app/cmd/check/check /opt/resource/check
COPY --from=build /go/src/app/cmd/in/in /opt/resource/in
COPY --from=build /go/src/app/dist/out /opt/resource/out

