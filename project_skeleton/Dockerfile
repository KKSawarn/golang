FROM golang:alpine

LABEL maintainer "Quick programmer "

# Move to working directory /build
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

ENV port 8000
# Build the application
RUN go build

# Command to run when starting the container
CMD ["./usercurd"]