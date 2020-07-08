ARG GO_VERSION_ARG=1.14
ARG ALPINE_VERSION_ARG=3.12

FROM golang:${GO_VERSION_ARG}-alpine${ALPINE_VERSION_ARG} AS gobuildder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/checode
COPY . .
RUN go get
RUN go build -o /go/dist/checode ./checode.go

FROM alpine:${ALPINE_VERSION_ARG} AS checode
COPY --from=gobuildder /go/dist/checode /bin/checode
RUN chmod +x /bin/checode
WORKDIR /opt/code
CMD ["/bin/checode", "/opt/code"]
