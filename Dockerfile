FROM golang:1.24.6-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tasks-api /app/cmd/main.go
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/ .

EXPOSE 8080

# CMD ["/cmd/main"]
CMD ["./tasks-api"]