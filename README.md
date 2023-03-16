# get-dat-money

Welcome to `get-dat-money`, where we get dat money. This service is a simple API that fetches price data for any ticker, for any date range. If the database does not yet have a particular ticker's data, it will fetch it from an external source and save it to the db. If it does have a ticker's data, it will fetch directly from the db.

___
## Endpoints

### `GET /api/v1/prices`

There are three query params expected:
- `ticker` (required)
- `start_date` (optional; defaults to "1900-01-01")
    - Expected format: YYYY-MM-DD
- `end_date` (optional; defaults to today's date)
    - Expected format: YYYY-MM-DD
    - Note that the data returned is up to (but not including) the `end_date` requested. e.g. sending a request with `end_date=2023-03-15` will return data for `2023-03-14`, but not for `2023-03-15`.
- `interval` (optional; defaults to "1d")
    - Currently, "1d" is the only valid input until we beef out the fetcher(s).

---

## Running Locally

There are two `make` commands:
- `make server detach=<true,false>`
    - Spins up a database as well as the API server.
    - Optional argument `detach` defaults to `false` so you can see the logs.
- `make kill`
    - Kills the containers with a simple `docker compose down`.
- `make test`
    - Runs all unit tests.

Once the environment is healthy and the server is up, you can run `curl` commands to see the magic work, such as:

```
curl 'http://localhost:8080/api/v1/prices?ticker=AAPL&start_date=2023-03-01'
curl 'http://localhost:8080/api/v1/prices?ticker=MSFT&start_date=2022-01-01&end_date=2023-03-04'
curl 'http://localhost:8080/api/v1/prices?ticker=TSLA&interval=1d'
```

---

## CI

We use Github Actions for our CI workflow. You can find the simple job in `.github/workflows/ci.yml`.

It simply sets up two Go version environments: 1.19 and 1.20. Then, it installs dependencies, builds the module, and runs
```
go test ./... -cover
```
to recursively find and run tests, as well as print out the test coverage. You can find all the runs [here](https://github.com/Margin-Call-Trading-Services/get-dat-money/actions/workflows/ci.yml).

---

## TODO

This is a young buck that is still a work in progress to beef up. Tasks on the horizon are:

- Adding terraform (or AWS CDK?) IaC. Ideally, we'd have this running as an ECS service in a private subnet, with the access path being API Gateway, Network Load Balancer, and finally Application Load Balancer.
- See what other endpoints and functionalities make sense for this service.
