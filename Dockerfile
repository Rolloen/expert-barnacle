FROM golang:1.22-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN GOOS=linux go build  -o /app ./cmd/app

EXPOSE 15442

ENTRYPOINT ["/app"] 