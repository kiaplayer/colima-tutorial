FROM golang:1.21-alpine

COPY . /app

WORKDIR /app

RUN go build -v -o /usr/local/bin/kubeapp ./...

ENTRYPOINT ["kubeapp"]