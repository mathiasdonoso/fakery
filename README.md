# fakery
Dummy HTTP server generator

## Warning :warning:
This project is still in WIP

## TOC
* [The Problem](#The-Problem)
* [The Solution](#The-Solution)
* [Installation](#Installation)
* [Configuration](#Configuration)
* [Getting Started](#Getting-Started)
* [API](#API)
* [Contributors](#Contributors)

## The Problem
Often, when developing or testing applications, you need to interact with services that provide specific responses. However, directly using these services might not be feasible due to downtime, additional costs, or other constraints.

## The Solution
**fakery** offers a solution by allowing you to effortlessly create lightweight servers with custom dummy responses for each endpoint you require. While this approach isn't groundbreaking, the inclusion of an executable or CLI makes it accessible beyond just developers, enabling quality assurance teams, among others, to utilize it effectively.

## Installation
-- Pending --

## Configuration
To configure fakery, you simply need a JSON configuration file (e.g., config.json) with the following structure:
```json
[
  {
    "request": {
      "url": "/path/to/something",
      "post": "this is some post data in textual format",
      "headers": {
         "authorization": "Basic usernamez:passwordinBase64"
      },
      "method": "POST"
    },
    "response": {
      "status": 200,
      "headers": {
        "Content-Type": "application/json"
      },
      "latency": 1000,
      "body": "Your request was successfully processed!"
    }
  },
  {
    "request": {
      "url": "/path/to/a/json/response",
      "headers": {
        "Content-Type": "application/json"
      },
      "method": "GET"
    },
    "response": {
      "status": 201,
      "headers": {
        "Content-Type": "application/json"
      },
      "body": "{\"message\": \"Item created\", \"item\": {\"id\": \"1\", \"name\": \"Ball\", \"price\": 9.99}}"
    }
  },
  {
    "request": {
      "url": "/path/to/thing",
      "headers": {
        "Content-Type": "application/json"
      },
      "post": "this is some post data in textual format",
      "method": "POST"
    },
    "response": {
      "status": 304,
      "headers": {
        "Content-Type": "application/json"
      }
    }
  }
]
```

Each object consists of:
- request: Request Object
- response: Response Object

Request object:
- url: string
- headers: Object with key and values as strings.
- method: string

Response object:
- status: number
- headers: Object with key and values as strings.
- body: string
- latency: number

## Getting Started
```bash
fakery -p 42069 -d service1-data.json
```

To build from source:
```bash
go build
./fakery -p 42069 -d service1-data.json
```

## API
The command accepts the following flags:
- p: The port where the server will run, default is port 8000.
- d: The data used by the server to create endpoints; by default, it looks for **config.json**.

## Contributors
This project is open source. Contributions are welcomed through PRs.
