# Stage 1: Build Go application
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files into the container
COPY . .

# Build the Go application inside the container
RUN go build -o main .

# Stage 2: Set up PostgreSQL
# Use the official PostgreSQL image as the database
# You can specify the desired version, e.g., postgres:13
FROM postgres:latest

# Set environment variables for the PostgreSQL container
ENV POSTGRES_USER ganesh
ENV POSTGRES_PASSWORD Libyar
ENV POSTGRES_DB useradmin

# Expose the PostgreSQL default port (5432)
EXPOSE 5432

# Move back to the first stage and set environment variables for the Go application
ENV PORT 5051

# Expose the Go application port (5051)
EXPOSE 5051

# Set the working directory to /app in the second stage
WORKDIR /app

# Copy the Go application binary into the container
COPY --from=0 /app/main .

# Run the Go application when the container starts
CMD ["./main"]
