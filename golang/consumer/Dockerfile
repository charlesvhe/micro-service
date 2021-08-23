# Build stage
FROM golang:alpine as builder
WORKDIR /app
COPY . .
#RUN apk update && apk add bash gcc musl-dev
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
#RUN go mod download
RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go
#Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

CMD ["/app/main"]