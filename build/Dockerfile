# -----
# BUILD
# -----

FROM golang:1.24.2-alpine AS build

# set up workdir
WORKDIR /go/src

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# copy project files to container
COPY . .

# compile app
RUN go build -o ./go_app ./cmd/app/main.go

# ---
# RUN
# ---

FROM alpine:3.21 AS run

WORKDIR /app

RUN mkdir /app/db

# copy compiled file and web-files
COPY --from=build /go/src/go_app .
COPY --from=build /go/src/web ./web

# run app
CMD ["/bin/sh", "-c", "/app/go_app"]
