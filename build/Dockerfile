# -----
# BUILD
# -----

FROM golang:1.24.2-alpine AS build

# set up workdir
WORKDIR /go/src

# install required packages for supporting CGO
RUN apk add --no-cache build-base sqlite-dev

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# copy project files to container
COPY . .

# tells the C compiler to enable support for large files, allowing the use
# of 64-bit versions of file handling functions such as pread64 and pwrite64
ENV CGO_CFLAGS="-D_LARGEFILE64_SOURCE"
ENV CGO_LDFLAGS="-D_LARGEFILE64_SOURCE"
ENV CGO_ENABLED=1

# compile migrator
RUN go build -o ./migrator ./cmd/migrator/main.go
# compile app
RUN go build -o ./app ./cmd/app/main.go

# ---
# RUN
# ---

FROM alpine:3.21 AS run

WORKDIR /app

RUN mkdir -p /app/media/excel

# copy compiled files
COPY --from=build /go/src/migrator .
COPY --from=build /go/src/app .

# run app
CMD ["/bin/sh", "-c", "/app/app"]
