# Use the official Go image as the base image
FROM golang:1.16-alpine

# Set environment variables for the application
ENV STREAMING_URL=https://stream.upfluence.co/stream
ENV PORT=8089

# Set the working directory inside the container
WORKDIR /app

# Copy the source code to the container's working directory
COPY . .

# Build the application inside the container
RUN go build -o app cmd/main.go

# Expose the port the application listens on
EXPOSE ${PORT}

# Command to run the application
CMD ["./app"]
