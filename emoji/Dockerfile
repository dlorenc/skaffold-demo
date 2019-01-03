FROM golang:1.11
COPY *.go /go/src/github.com/dlorenc/skaffold-demo/emoji/
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' github.com/dlorenc/skaffold-demo/emoji

FROM scratch
EXPOSE 8080
COPY --from=0 /go/emoji /
CMD ["/emoji"]