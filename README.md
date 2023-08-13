# Stream Consumer API

The Stream Consumer API is a Go-based application that provides functionality to consume data streams, calculate statistics, and expose the results through a simple API. This README will guide you through the usage of the API and provide insights into the implementation.

## Getting Started

To use the Stream Consumer API, follow these steps:

1. Clone the repository to your local machine:

```
git clone https://github.com/asadrizv/social-media-analyser
```


2. Start the API server locally:
```
make run
```

3. To start the API server in a Dockerized manner:
```
docker compose up
```

4. To run unit tests
````shell
make test
````



The API server will be up and running, ready to accept incoming requests.

## API Endpoints

The API exposes the following endpoint:

- **GET /analysis?duration={duration}&dimension={dimension}**

This endpoint accepts two query parameters:
- `duration`: The duration for which the stream should be analyzed (e.g., "30s" for 30 seconds, "30m" for 30 minutes).
- `dimension`: The dimension of interest (e.g., "likes", "comments", "saves").

The API will block for the specified duration, analyze the stream data, and return statistics related to the provided dimension.

## Implementation Overview

The Stream Consumer API is implemented in Go and leverages the standard library to handle HTTP requests, SSE (Server-Sent Events) streams, and JSON parsing. It consumes a stream of data and calculates statistics, including max timestamp, minimum timestamp, and various percentiles.

The core components include:
- `StreamConsumerService`: Consumes the stream data, analyzes it, and calculates statistics.
- `MetricsStatistics`: Struct to hold the calculated statistics.
- `calculatePercentile`: Utility function to calculate percentiles from a slice of float64 numbers.

## Usage

1. Start the API server as mentioned in the "Getting Started" section.
2. Use a tool like `curl` or a browser to send a GET request to the API endpoint:

```shell
curl "http://localhost:8080/analysis?duration=30s&dimension=likes"

```