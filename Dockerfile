# compiler image
FROM golang:alpine3.22 AS build-env

# set workdir where the app will work from.
WORKDIR /app

# download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# add all files to image.
ADD . ./

# disable cgo to build on alpine.
ENV CGO_ENABLED 0

# run tests.
RUN go test ./...

# build the binary.
RUN GOOS=linux GOARCH=amd64 go build -a -o hellosvc ./cmd/server/main.go

# create final application image.
FROM alpine:3.22
WORKDIR /app
COPY --from=build-env /app/hellosvc .
ENTRYPOINT ./hellosvc
