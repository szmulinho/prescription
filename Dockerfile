FROM golang:1.21.1-alpine AS build

WORKDIR /prescription
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o prescription

FROM alpine:latest

WORKDIR /prescription
COPY --from=build /prescription/prescription /prescription/prescription

EXPOSE 8080

CMD ["./prescription"]
