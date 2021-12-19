# paack-go
We are PaackAndGo, a small bus company. We want to create a REST API that helps us manage the trips that we offer.

# Directory tree

- `pkg`: Here is the main source code.
- `challenge`: Here is the initial challenge
- `logs`: Here is the logs directory, this directory needs to be created manually.
- `postman`: Here is a `json` file to import to your local Postman.

# 1. Requirements

| Software          | Version  | Importance |
|-------------------|----------|------------|
| 🐳 Docker         | 20.10.12 | Required   |
| 🐙 Docker Compose | 1.29.2   | Required   |
| 🐃 GNU Make       | 4.2.1    | Optional   |
| ‍🚀 Postman       | 9.1.5    | Optional   |

# 2. Before run

## With Makefile

```shell
make l/init
```

## Manually

1. Copy the `.env-example` file to `./pkg/.env` file:

```shell
cp -n .env.example ./pkg/.env
```

2. Create the `logs` directory

```shell
mkdir -p logs
```

## Fill out the variables
- Use your own credentials. If you want to start quickly, these are the default values (do not modify):

```
# Use to initialize the application
GLOBAL_ENV=develop
GLOBAL_LOG_FILE=paack-go.log

# Use for the app (you can change by the AWS secretes storage or by Secrets on GCP)
APP_PORT=8087

POSTGRES_DB=db
POSTGRES_USER=user
POSTGRES_PASSWORD=secret
POSTGRES_DRIVER=postgres
POSTGRES_HOST=127.0.0.1

```
