FROM golang:1.12-alpine AS build
WORKDIR /frisk
RUN apk --no-cache add curl openssh git mercurial bzr subversion

ENV GOPROXY https://proxy.golang.org

COPY go.mod .
#COPY go.sum .
RUN go mod download
COPY . .
RUN go install -v ./...

FROM alpine:latest as production
RUN apk add --no-cache --update ca-certificates curl openssh
COPY --from=build /go/bin/frisk /frisk
ENTRYPOINT ["/frisk"]
CMD ["-h"]
