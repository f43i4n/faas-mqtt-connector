FROM golang:1.13 as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go test ./...

RUN go build -o mqtt-connector .

FROM debian

WORKDIR /app

COPY --from=builder /build/mqtt-connector /app/

CMD [ "./mqtt-connector" ]
