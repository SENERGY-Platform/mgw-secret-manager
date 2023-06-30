FROM golang:1.18 AS builder
ENV GO111MODULE=on

COPY . /app
WORKDIR /app
RUN git log -1 --oneline > version.txt

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o ../app

FROM alpine:3.18 AS prod
WORKDIR /root
COPY --from=builder /app/app .
COPY --from=builder /app/version.txt .

EXPOSE 8080

CMD ["./app"]

FROM golang:1.18 AS test 
COPY . /app
WORKDIR /app
CMD ["go", "test",  "./...", "-coverprofile=/coverage/cov.xml", "-coverpkg=./..."]