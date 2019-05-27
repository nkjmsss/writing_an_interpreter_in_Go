FROM golang:1.12-rc-alpine3.9
ENV GO111MODULE=on
WORKDIR /go/src/github.com/nkjmsss/writing_an_interpreter_in_Go
RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
      git
# download go modules (with trying to use cache)
COPY go.mod .
COPY go.sum .
RUN go mod download
# COPY the source code as the last step
COPY . .
