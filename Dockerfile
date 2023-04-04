FROM golang:alpine AS build

WORKDIR /app
COPY *.go .

RUN go build -o webserver main.go


FROM alpine:latest

WORKDIR /app
COPY --from=build /app/webserver .

RUN ls -la

EXPOSE 8080
CMD ["./webserver"]
