FROM golang:1.20-alpine3.17 AS dev

WORKDIR /app/payment_transfer_api
COPY go.* /app/payment_transfer_api/
RUN go mod download
COPY internal /app/payment_transfer_api/internal
COPY cmd /app/payment_transfer_api/cmd
ENV PAYMENT_DB=${PAYMENT_DB}
RUN go build -o api ./cmd/api.go
EXPOSE 5000
CMD [ "go", "run", "cmd/api.go" ]

FROM alpine:3.17.2 AS prod

WORKDIR /app/payment_transfer_api
COPY --from=dev /app/payment_transfer_api/api .
CMD [ "./api" ]