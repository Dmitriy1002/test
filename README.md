
# Proof of Work TCP Server

## Description

This project implements a TCP server protected from DDoS attacks using a Proof of Work (PoW) algorithm. After successful PoW verification, the server sends one of the quotes from the "words of wisdom" collection.

## Project Structure

- `server.go`: The main server file that handles incoming TCP connections, issues a PoW challenge, and verifies the solution.
- `client.go`: The client file that connects to the server, solves the PoW challenge, and receives a quote.
- `Dockerfile.server`: Dockerfile for building and running the server.
- `Dockerfile.client`: Dockerfile for building and running the client.
- `README.md`: This file with a description of the project and usage instructions.
- `server_test.go`: Tests for the server.
- `client_test.go`: Tests for the client.

## Usage Instructions

### Server

1. Build and run the server:

```sh
docker build -t pow-server -f Dockerfile.server .
docker run -p 5000:5000 pow-server
```

### Client

1. Build and run the client:

```sh
docker build -t pow-client -f Dockerfile.client .
docker run pow-client
```

## Testing

To run the tests, use the following command:

```sh
go test
```
