FROM golang:1.17-alpine as builder
WORKDIR /app
COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /bin/
COPY --from=builder /app/app .


EXPOSE 9090
CMD [ "./app" ]

