FROM golang:1.20-alpine as builder

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /newordsbe-digistar

FROM alpine:latest

WORKDIR /app

COPY .env .env

RUN export $(cat .env | xargs)

COPY --from=builder /newordsbe-digistar /newordsbe-digistar

EXPOSE 8099

ENTRYPOINT [ "/newordsbe-digistar" ]