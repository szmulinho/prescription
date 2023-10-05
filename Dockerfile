FROM golang:alpine AS build
WORKDIR /prescription
LABEL maintainer="szmulinho"
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
FROM alpine:latest
WORKDIR /root/
COPY --from=build /prescription .
EXPOSE 8095
CMD ["./prescription"]

