FROM golang:1.14.2-alpine as build

WORKDIR /root
RUN apk update && apk add busybox-extras git
ADD . .

# RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-extldflags "-static"'

# FROM alpine:latest as alpine
# RUN apk --no-cache add tzdata zip ca-certificates
# WORKDIR /usr/share/zoneinfo
# # -0 means no compression.  Needed because go's
# # tz loader doesn't handle compressed data.
# RUN zip -r -0 /zoneinfo.zip .

# FROM balenalib/raspberrypi4-64-alpine as alpine



FROM scratch 

# COPY --from=alpine /dev /dev
COPY --from=build /root/fanshim /

# ENTRYPOINT ["/bin/sleep", "6000"]
ENTRYPOINT ["/fanshim"]