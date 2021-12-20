# paack-go

We are PaackAndGo, a small bus company. We want to create a REST API that helps us manage the trips that we offer.

# Directory tree

- `pkg`: Here is the main source code.
- `challenge`: Here is the initial challenge
- `pkg/logs`: Here are the log files, this directory needs to be created manually.
- `postman`: Here is a `json` file to import to your local Postman.

# 1. Requirements

| Software          | Version  | Importance                        |
|-------------------|----------|-----------------------------------|
| 🦫 Golang         | 1.17.3   | Required                          |
| 🐳 Docker         | 20.10.12 | Required (only if not use Golang) |
| 🐙 Docker Compose | 1.29.2   | Required (only if not use Golang) |
| 🐃 GNU Make       | 4.2.1    | Optional                          |
| ‍🚀 Postman       | 9.1.5    | Optional                          |

# 2. Before run

## With Makefile

- Run without Docker:

```shell
make l/init
make l/run
```

- Run with Docker:

```shell
make l/init
make l/docker
```

- Run with the tests:

```shell
make l/init
make l/docker
```

## Manually

1. Copy the `.env-example` file to `./pkg/.env` file:

```shell
cp -n .env.example ./pkg/.env
```

2. Create the `logs` directory

```shell
mkdir -p ./pkg/logs
```

## Fill out the variables

- Use your own credentials. If you want to start quickly, these are the default values (do not modify):

```
# Use to initialize the application
GLOBAL_ENV=develop
GLOBAL_LOG_FILE=paack-go.log

# Use for the app (you can change by the AWS secretes storage or by Secrets on GCP)
APP_PORT=8080

# Persistence
FILE_NAME=cities.txt

```
