#stage 1:
FROM golang:1.23-alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

#stage 2:
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/server .

COPY configs/ configs/

EXPOSE 8080

CMD ["./server"]

#docker build -t track-backend .
#docker run --env-file .env -p 8080:8080 track-backend