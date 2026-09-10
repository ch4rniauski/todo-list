FROM golang:1.26.5-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY .. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/todo-list ./cmd/todo-list

FROM alpine:3
WORKDIR /app
COPY --from=build /app/todo-list /app/todo-list

EXPOSE 8080
CMD ["/app/todo-list"]
