Go Fizzbuzz implementation with PostgreSQL

# Project details

# Getting Started

## Dependencies

[PostgeSQL](https://github.com/jackc/pgx) as event store<br/>
[Jaeger](https://www.jaegertracing.io/) open source, end-to-end distributed [tracing](https://opentracing.io/)<br/>
[Prometheus](https://prometheus.io/) monitoring and alerting<br/>
[Grafana](https://grafana.com/) for to compose observability dashboards with everything from Prometheus<br/>
[Echo](https://github.com/labstack/echo) web framework<br/>
[Migrate](https://github.com/golang-migrate/migrate) for migrations<br/>
[Mock](https://github.com/golang/mock) for mocking<br/>

# Services availables

## Swagger UI

http://localhost:8000/swagger/index.html

## Metrics Endpoint

http://localhost:8080/metrics

## Jaeger UI

http://localhost:16686

## Prometheus UI

http://localhost:9090

## Grafana UI

http://localhost:3000/d/3bB90wi4z/fuzzbuzz

# TODO

- [] Fix fizzbuzz return value should be a single string a not an array
- [] Cleanup golangci linter
- [x] Timeout handler management
- [] Documentation
- [] Test coverage
  - [x] Generate mock
- [x] Generate swagger documentation
- [x] Add postgres support
  - [x] Add sql migration scripts
  - [x] Add pg support to repository
- [x] Prometheus dashboard
- [] Add docker prometheus, postgres & jaeger
- [] Cleanup readme.md
