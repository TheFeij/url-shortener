# URL Shortener

A simple URL shortener service built with Golang.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)

## Introduction

This project is a URL shortener service implemented in Golang. It allows users to shorten long URLs and redirect to the original URLs using the shortened version.

## Features

- Shorten long URLs
- Redirect to the original URL using the shortened URL
- RESTful API

## Getting Started


### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/TheFeij/url-shortener.git
    cd url-shortener
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Start the server:
    ```bash
    make start
    ```

## Configuration

Configuration is managed via a `config.json` file. Here is an example configuration:

```json
{
   "DATABASE_ADDRESS": "postgresql://****:****@localhost:5432/url_shortener?sslmode=disable",
   "SERVER_ADDRESS": ":8080"
}
```

## Usage

### Shorten a URL
Send a POST request to /shorten with the URL to be shortened.

Request:
```json
{
  "url": "http://example.com"
}
```

Response:
```json
{
   "short_url": "http://localhost:8080/abc123"
}
```

### Redirect to the Original URL
Access the shortened URL in the browser or via a GET request to be redirected.

### API Endpoints
- POST /shorten: Shorten a URL.
- GET /:shortUrl: Redirect to the original URL