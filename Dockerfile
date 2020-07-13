ARG GO_VERSION_ARG=1.14
ARG ALPINE_VERSION_ARG=3.12

FROM golang:${GO_VERSION_ARG}-alpine${ALPINE_VERSION_ARG} AS gobuildder
WORKDIR /go/src
COPY . .
RUN go get
RUN ./script/build /go/checode

FROM alpine:${ALPINE_VERSION_ARG} AS checode
COPY --from=gobuildder /go/checode /bin/checode
RUN chmod +x /bin/checode
WORKDIR /opt/code
CMD ["/bin/checode", "/opt/code"]
