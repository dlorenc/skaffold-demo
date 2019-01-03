FROM golang:1.11
COPY . /go/src/github.com/dlorenc/skaffold-demo/microservices/backend
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' github.com/dlorenc/skaffold-demo/microservices/backend

FROM scratch
EXPOSE 8000
COPY --from=0 /go/backend /
CMD ["/backend"]