FROM golang:alpine as builder
# first (build) stage
#RUN go mod download
WORKDIR /app
ENV GOPROXY=https://proxy.golang.com.cn,direct
COPY go.mod .
COPY go.sum .
COPY start.go .
RUN go get github.com/joho/godotenv
RUN go build -o start start.go

# final (target) stage

FROM alpine:3.10
WORKDIR /app
LABEL author="leechao" mail="lic0914@163.com" MAINTAINER="leechao"
RUN apk --no-cache add curl
COPY --from=builder /app/start .
COPY .env .
COPY index.html .
COPY index.txt .
CMD ["./start"]




