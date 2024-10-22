FROM golang:1.23 AS builder
ENV GO111MODULE=on

COPY . /app
WORKDIR /app
#RUN git log -1 --oneline > version.txt
ARG VERSION=dev

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o ../app -ldflags="-X 'main.version=$VERSION'"

FROM alpine:3.20 AS prod
WORKDIR /root
COPY --from=builder /app/app .
#COPY --from=builder /app/version.txt .

EXPOSE 8080

HEALTHCHECK --interval=10s --timeout=5s --retries=3 CMD wget -nv -t1 --spider 'http://localhost/health-check' || exit 1

CMD ["./app"]

FROM golang:1.23 AS test
COPY . /app
WORKDIR /app
CMD ["go", "test",  "./...", "-coverprofile=/coverage/cov.xml", "-coverpkg=./..."]