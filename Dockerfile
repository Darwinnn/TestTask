FROM golang:1.13
COPY . /go/src/github.com/darwinnn/TestTask
WORKDIR /go/src/github.com/darwinnn/TestTask
RUN make


FROM alpine:3.10  
RUN apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null
COPY --from=0 /go/src/github.com/darwinnn/TestTask/bin/app .
EXPOSE 8080
CMD /app
