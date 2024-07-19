FROM golang:1.22.3-bookworm as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

FROM debian:bookworm-slim

COPY --from=build /app/app .
COPY --from=build /app/go.mod .

RUN apt-get update

CMD ["./app"]
