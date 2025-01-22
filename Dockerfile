FROM golang:alpine
RUN apk add build-base

ENV CGO_ENABLED=1
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the app
RUN go build -a -o /app/hotel-data-merge

# Run the compiled app
CMD ["/app/hotel-data-merge"]
EXPOSE 80
