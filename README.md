# get-dat-money

Welcome to `get-dat-money`, where we get dat money. This service is a simple API that fetches price data for any ticker, for any date range. If the database does not yet have a particular ticker's data, it will fetch it from an external source and save it to the db. If it does have a ticker's data, it will fetch directly from the db.

___
## Endpoints

### `GET /api/v1/prices`

There are three query params expected:
- `ticker` (required)
- `start_date` (optional; defaults to 1900-01-01)
    - Expected format: YYYY-MM-DD
- `end_date` (optional; defaults to today's date)
    - Expected format: YYYY-MM-DD

---

## Running Locally

There are two `make` commands:
- `make server detach=<true,false>`
    - Spins up a database as well as the API server.
    - Optional argument `detach` defaults to `false` so you can see the logs.
- `make kill`
    - Kills the containers with a simple `docker compose down`.

Once the environment is healthy and the server is up, you can run `curl` commands to see the magic work, such as:

```
curl 'http://localhost:8080/api/v1/prices?ticker=AAPL&start_date=2023-03-01'
curl 'http://localhost:8080/api/v1/prices?ticker=MSFT&start_date=2022-01-01&end_date=2023-03-04'
curl 'http://localhost:8080/api/v1/prices?ticker=TSLA'
```

---

## TODO

This is a young buck that is still a work in progress to beef up. Tasks on the horizon are:

- Adding test coverage.
- Adding CI/CD GH Action workflows.
- Adding terraform (or AWS CDK?) IaC. Ideally, we'd have this running as an ECS service in a private subnet, with the access path being API Gateway, Network Load Balancer, and finally Application Load Balancer.
- See what other endpoints and functionalities make sense for this service.
