FROM golang:1.18-buster AS build

WORKDIR /go/src/lstrgiang/ascenda

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/app ./main.go


###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/app /usr/local/bin/app
RUN apk add --no-cache ca-certificates
COPY ./data /data

WORKDIR /

ENTRYPOINT ["app","server", "--supplier", "/data/suppliers.json"]
