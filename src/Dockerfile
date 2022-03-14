FROM golang:alpine as builder
# first (build) stage
WORKDIR /app
COPY start.go .
RUN go build -o start start.go

# final (target) stage

FROM alpine:3.10
LABEL author="leechao" mail="lic0914@163.com" MAINTAINER="leechao"
COPY --from=builder /app/start /
COPY index.html /
COPY index.txt /
CMD ["/start"]




