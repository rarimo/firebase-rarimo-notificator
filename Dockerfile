FROM golang:1.22.3-alpine AS buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokend/firebase-rarimo-notificator
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/firebase-rarimo-notificator /go/src/gitlab.com/tokend/firebase-rarimo-notificator


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/firebase-rarimo-notificator /usr/local/bin/firebase-rarimo-notificator
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["firebase-rarimo-notificator"]
