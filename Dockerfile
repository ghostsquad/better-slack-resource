# stage: base
FROM alpine:3.8 as base

RUN set -ex; \
    apk add --update --no-cache \
        bash \
    ; \
    rm -rf /var/cache/apk/*;

# stage: build
FROM golang:alpine3.8 as build

RUN set -ex; \
    apk add --update --no-cache \
        git \
    ;

COPY --from=instrumentisto/dep:0.5.0 /usr/local/bin/dep /usr/local/bin/dep

WORKDIR /go/src/github.com/ghostsquad/slack-off

COPY Gopkg.* ./

RUN dep ensure -vendor-only

COPY . .

RUN go test ./...

RUN go build -o dist/out cmd/out/out.go

RUN ln -s /go/src/github.com/ghostsquad/slack-off /app

# stage: final
FROM base

COPY --from=build /app/cmd/check/check /opt/resource/check
COPY --from=build /app/cmd/in/in /opt/resource/in
COPY --from=build /app/dist/out /opt/resource/out

