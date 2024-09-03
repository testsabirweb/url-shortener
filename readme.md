# URL Shortener Service

This is a simple URL shortener service built with Go (Golang) and Gin, which allows users to shorten URLs and retrieve the original URLs using the shortened ones. The service also provides metrics for the top 3 domains that have been shortened the most.

## Features

- Shorten long URLs into short, easy-to-share links.
- Redirect users to the original URL using the short link.
- Metrics API to get the top 3 domains shortened the most.

## Project Structure

- `main.go`: The entry point of the application.
- `store/`: Contains the in-memory storage implementation.
- `handler/`: Contains the HTTP handlers for creating short URLs and redirecting to the original URLs.
- `shortener/`: Contains the acutal algorithmic implementation for shortening of url
- `Dockerfile`: Used for containerizing the application.

## Prerequisites

- Go 1.22.2 or later
- Docker (if you want to containerize the application)

## Getting Started

### Clone the Repository

```bash
git clone git@github.com:testsabirweb/url-shortener.git
cd url-shortener
```

## Testing(Postman)
```text
I have attached infracloud.postman_collection.json file, one can import and test using postman. examples are also added for reference
```

## Dockerhub image
https://hub.docker.com/repository/docker/sabir9644/url-shortener/general

