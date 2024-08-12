FROM golang:1.22.6 as builder
WORKDIR /app
COPY . .
RUN make build

FROM gcr.io/distroless/base
WORKDIR /root/
COPY --from=builder /app/app .

CMD ["./app", "serve"]
