# trimly
## What is Trimly?
Trimly is a URL shortener which means it takes a long URL and makes it shorter.

## Features
- [x] User Authentication
- [x] Hashing Algorithm (MD5 + base64 URL Encoding + Random counter)
- [x] API Testing
- [x] Caching with go-cache
- [] Swagger Documentation
- [x] Postman Documentation


## Requires
`go 1.17+` `MongoDB`  `docker`

## Setup

- Clone project using:
```bash
    git clone https://github/meshachdamilare/trimly.git
```

- Create your env file named `your_env_name.env`, and write your env variables according to
  `sample.env` in the projects root directory

- Run unit tests using:
```bash
    go test -tags=unit ./...
```

- Start the server using:
```bash
    docker-compose up
```

## Documentation

- View Postman documentation for the API at
  [Postman link](https://documenter.getpostman.com/view/27840229/2s93zCZgNg)

## Hosting

- Trimly API is hosted on render
  [Link](https://trimly-86ks.onrender.com/)
