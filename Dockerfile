# build binary
FROM golang:1.22.0 AS builder
WORKDIR /build

ENV CGO_ENABLE=0

COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -gcflags "all=-N -l" -o bootstrap #  -gcflags "all=-N -l" ... disable compiler optimization

# Install delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080
EXPOSE 2345
ENTRYPOINT ["/go/bin/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/build/bootstrap", "--continue"]
