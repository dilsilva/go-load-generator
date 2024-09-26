# Go Load Generator

This project is a simple and efficient HTTP load generator written in Go. It allows you to test the performance of your web services by sending concurrent HTTP requests to a specified URL. The project also provides metrics like success/failure rates, average response time, and 95th percentile response time.

## Features

- **Concurrency**: Simulate multiple users by sending requests concurrently.
- **Metrics Collection**: Tracks and reports success/failure counts, average response times, and 95th percentile response time.
- **Lightweight**: Built with Go, utilizing minimal resources with Docker.
- **Containerized**: Easily containerized using a secure, multi-stage Dockerfile for safe deployment.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Container Setup](#container-setup)
- [Flags](#flags)
- [Development](#development)
- [License](#license)

---

## Prerequisites

Before running the project, ensure you have the following installed:

- **Go**: Version 1.20 or higher.
- **Docker**: Latest stable version to build and run the containerized application.

## Installation

### Cloning the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/yourusername/go-load-generator.git
cd go-load-generator
```

### Building the Project

You can build the Go binary directly on your local machine:

```bash
go build -o loadgen ./cmd/loadgen/
```

This will create the `loadgen` executable in the root directory.

### Running Locally

After building, you can run the application on your local machine:

```bash
./loadgen -url http://example.com -c 10 -r 100
```

## Docker Setup

This project is containerized using Docker. You can build and run the container with the following commands.

### Build the Docker Image

```bash
container build -t loadgen .
```

### Running the Docker Container

```bash
container run --rm loadgen -url http://example.com -c 10 -r 100
```

You can pass custom flags for concurrency, total requests, and the URL directly into the `container run` command.

## Flags

The load generator supports several flags for customizing your load tests:

- `-url`: The URL to load test (default: `http://google.com`).
- `-c`: Number of concurrent requests (default: `10`).
- `-r`: Total number of requests to send (default: `100`).

### Example Usage

```bash
./loadgen -url http://example.com -c 50 -r 500
```

This will send 500 requests to `http://example.com` with a concurrency level of 50.

## Development

To contribute or modify the project, follow these steps:

### Prerequisites

- **Go**: Ensure Go is installed on your machine.
- **Docker**: Docker should be installed for containerized builds.

### Running in Development

1. Clone the repository.
2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build and test changes locally using Go or Docker.

### Linting and Testing

Use Go tools to lint and test the project:

```bash
go test ./...
golangci-lint run
```

## Metrics

The load generator reports the following metrics after the test is complete:

- **Total Requests**: The total number of requests sent.
- **Successful Requests**: Number of successful responses (status code < 400).
- **Failed Requests**: Number of failed responses (status code >= 400).
- **Average Response Time**: The average time it took to get a response.
- **95th Percentile Response Time**: Time within which 95% of the requests were completed.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to contribute or report issues via GitHub. Happy load testing!