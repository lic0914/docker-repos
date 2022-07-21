FROM golang:alpine as build
# first (build) stage
#RUN go mod download
WORKDIR /app
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.com.cn,direct
COPY . .
RUN go get 
RUN go build -o main main.go

# final (target) stage

FROM alpine:3.10
WORKDIR /app
LABEL author="leechao" mail="lic0914@163.com" MAINTAINER="leechao"
RUN apk --no-cache add curl
COPY --from=build /app/main .
COPY .env .
CMD ["./main"]




