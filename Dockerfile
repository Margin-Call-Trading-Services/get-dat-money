FROM golang:1.19-alpine3.17 as builder

COPY go.mod go.sum /build/
WORKDIR /build
RUN go mod download

COPY . /build
RUN go build -o service main.go

FROM alpine:3.17

COPY --from=builder /build/service /app/
WORKDIR /app

EXPOSE 8080

ENTRYPOINT [ "./service" ]