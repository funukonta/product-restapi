FROM golang:alpine

# Install git (required for fetching dependencies)
RUN apk update && apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy the entire project
COPY . .

# Run 'go mod tidy' to clean up the module file
RUN go mod tidy

# Build the Go application
RUN go build -o binary

# Expose port 8080 (the port your application will run on)
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["/app/binary"]
