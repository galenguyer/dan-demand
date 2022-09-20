FROM docker.io/golang:1.19-alpine
WORKDIR /go/src/github.com/galenguyer/dan-demand/
RUN apk add bash ca-certificates git gcc g++ libc-dev

ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=1 go build -a -v -tags netgo -ldflags '-w -extldflags "-static"'

FROM docker.io/alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=0 /go/src/github.com/galenguyer/dan-demand/dan-demand ./
VOLUME ["/config"]
ENTRYPOINT ["/root/dan-demand"]
