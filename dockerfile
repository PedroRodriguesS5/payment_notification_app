FROM golang:1.23.2

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code (including subdirectories)
COPY . .

# Build the Go application
RUN go build -v -o /usr/local/bin/app .

# Expose port 8080
EXPOSE 8080

# Run the built binary
CMD ["air","usr", "/usr/local/bin/app", ".air.toml"]
