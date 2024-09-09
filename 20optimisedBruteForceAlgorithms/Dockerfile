FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o optimised-brute-force-attack main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /mygolang-ico/20optimisedBruteForceAlgorithms .

EXPOSE 8088

CMD [ "./optimised-brute-force-attack" ]