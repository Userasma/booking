# Start with the official Go image
FROM golang:1.16-alpine

# Set the working directory
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the necessary port
EXPOSE 8080

# Run the Go application
CMD ["./main"]
