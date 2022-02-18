# syntax=docker/dockerfile:1

# Build Stage
FROM golang:1.16-buster as Build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /unicornvert

# Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /unicornvert /unicornvert

EXPOSE 8689

USER nonroot:nonroot

ENTRYPOINT ["/unicornvert"]