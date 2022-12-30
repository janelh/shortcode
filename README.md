## Shortcode - Go URL Shortener
This is a simple URL shortener written in Go that stores shortcodes in Redis and can be run using Docker Compose.

## Requirements
* Go (tested with version 1.19)
* Docker
* Docker Compose

## Getting Started
1. Clone this repository: git clone https://github.com/your-username/go-url-shortener.git
2. Build the Docker image: docker-compose build
3. Run the Docker containers: docker-compose up

The URL shortener should now be running at http://localhost:8080

## Usage
To shorten a URL, send a POST request to the server with the original URL in the request body:

```
curl -X POST -d '{"url":"http://example.com"}' http://localhost:8080/urls
```
The server will respond with the shortened URL, which can be accessed by visiting the shortcode URL:
```
http://localhost:8080/urls/abc123
```

http://localhost:8080/urls/abc123 will redirect to http://example.com

## Configuration
The following environment variables can be used to configure the behavior of the URL shortener:

* **RDS_HOST**: the hostname or IP address of the Redis server
* **RDS_PASSWORD**: the password for the Redis server (default: "")
* **BASE_URL**: the base URL of the shortener (default: "http://localhost:8080")

These variables can be set in the .env file in the root of the project.
