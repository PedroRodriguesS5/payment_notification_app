FROM golang:1.23-alpine as builer

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./cmd/api

# Port to be expose
EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]