FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.50.0

COPY go.mod go.sum ./

RUN go mod download

COPY . .

#CMD ["go", "run", "main.go"] #production
CMD ["air", "-c", ".air.toml"]