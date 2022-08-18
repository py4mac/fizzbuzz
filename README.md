<H1>Go Fizzbuzz implementation with PostgreSQL</H1>

# Project details

Exercise: Write a simple fizz-buzz REST server.

"The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by ""fizz"", all multiples of 5 by ""buzz"", and all multiples of 15 by ""fizzbuzz"".
The output would look like this: ""1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...""."

Your goal is to implement a web server that will expose a REST API endpoint that:

- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:

- Ready for production
- Easy to maintain by other developers

Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:

- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request

# Setup local development

## Clone repository

```sh
❯ git clone git@github.com:py4mac/fizzbuzz.git
```

## Install tools

The following tools are required to be installed on local machine.

[Golang](https://go.dev/) in version 1.19<br/>
[Docker](https://www.docker.com/products/docker-desktop/) docker and docker-compose<br/>
[Linter](https://golangci-lint.run/usage/install/) for linter source code<br/>
[Swag](https://github.com/swaggo/swag) for generating swagger documentation<br/>
[Migrate](https://github.com/golang-migrate/migrate) for database migrations<br/>
[Mock](https://github.com/golang/mock) for mocking<br/>

## Other stacks

These other stacks are already pre-installed inside differents containers. No installation is required here, it's just for information.

[PostgeSQL](https://github.com/jackc/pgx) as fizzbuzz user request store<br/>
[Jaeger](https://www.jaegertracing.io/) open source, end-to-end distributed [tracing](https://opentracing.io/)<br/>
[Prometheus](https://prometheus.io/) monitoring and alerting<br/>
[Grafana](https://grafana.com/) for to compose observability dashboards with everything from Prometheus<br/>
[Echo](https://github.com/labstack/echo) web framework<br/>

## Getting Started

### Run unittest

```sh
❯ make test
```

### Run local containers

The following command starts PostgreSQL, Jaeger, Prometheus and Grafana containers.

```sh
❯ make start_local
```

### Initialize database

Onece PostgreSQL container is running, the following command initialize the database.

```sh
❯ make migrate_up
```

### Run local application

```sh
❯ make run
```

# Services availables

Once local containers and local application is started, the following endpoints are availables.

## Swagger UI

It exposes RestAPI handlers
http://localhost:8000/swagger/index.html

### Fizzbuzz

**URL** : `/api/v1/fizzbuzz`

**Method** : `GET`

```sh
❯ curl -i http://localhost:8000/api/v1/fizzbuzz\?int1\=3\&int2\=5\&limit\=100\&str1\=fizz\&str2\=buzz
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Vary: Origin
Date: Thu, 18 Aug 2022 09:14:59 GMT
Content-Length: 415

"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz,31,32,fizz,34,buzz,fizz,37,38,fizz,buzz,41,fizz,43,44,fizzbuzz,46,47,fizz,49,buzz,fizz,52,53,fizz,buzz,56,fizz,58,59,fizzbuzz,61,62,fizz,64,buzz,fizz,67,68,fizz,buzz,71,fizz,73,74,fizzbuzz,76,77,fizz,79,buzz,fizz,82,83,fizz,buzz,86,fizz,88,89,fizzbuzz,91,92,fizz,94,buzz,fizz,97,98,fizz,buzz"
```

### Stats

**URL** : `/api/v1/stats`

**Method** : `GET`

```sh
❯ curl -i http://localhost:8000/api/v1/stats
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Vary: Origin
Date: Thu, 18 Aug 2022 09:16:09 GMT
Content-Length: 70

{"hits":20,"int1":3,"int2":5,"limit":100,"str1":"fizz","str2":"buzz"}
```

## Metrics Endpoint

http://localhost:8080/metrics

## Jaeger UI

http://localhost:16686

## Prometheus UI

http://localhost:9090

## Grafana UI

http://localhost:3000/d/3bB90wi4z/fuzzbuzz

## Stop local containers

```sh
❯ make stop_local
```

# Setup production environnement

### Run production containers

```sh
❯ make run
```

and database migration script

```sh
❯ make migrate_up
```

# Services availables

Once production containers are started, the same services are started with the API. The same endpoints are also availables.

# TODO

- [] Test
  - [x] Generate mock
  - [x] Complete validate params sequence
  - [x] Fix fizzbuzz return value should be a single string a not an array
  - [x] Errors cleanup
  - [] Unittest
  - [] Integration test?
- [] Documentation
  - [x] Cleanup readme.md
  - [] Code cleanup
  - [] Godoc internal?
- [] Build
  - [] Cleanup golangci linter
  - [] Remove Jaeger?
  - [] Generate vendor folder
