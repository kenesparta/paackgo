# paack-go

We are PaackAndGo, a small bus company. We want to create a REST API that helps us manage the trips that we offer.

# Directory tree

- `pkg`: Here is the main source code.
- `challenge`: Here is the initial challenge
- `pkg/logs`: Here is the log file, this directory needs to be created manually.
- `postman`: Here is a `json` file to import to your local Postman.

# 1. Requirements

| Software          | Version  | Importance                        |
|-------------------|----------|-----------------------------------|
| 🦫 Golang         | 1.17     | Required                          |
| 🐳 Docker         | 20.10.12 | Required (only if not use Golang) |
| 🐙 Docker Compose | 1.29.2   | Required (only if not use Golang) |
| 🐃 GNU Make       | 4.2.1    | Optional                          |
| ‍🚀 Postman       | 9.1.5    | Optional                          |

# 2. Before run

## With Makefile

Run this command:

```shell
make l/init
```

## Manually

1. Copy the `.env-example` file to `./pkg/.env` file.


2. Copy the `./challenge/cities.txt` file to `./pkg/cities.txt` file.

3. Create the `./pkg/logs` directory.

## Fill out the variables

- Use your own credentials. If you want to start quickly, these are the default values (do not modify):

```
# Use to initialize the application
GLOBAL_ENV=develop

# Use for the app (you can change by the AWS secretes storage or by Secrets on GCP)
APP_PORT=8080

# Persistence
FILE_NAME=./cities.txt

# Log
LOGS_DIR=./logs

```

# 3. Run the app

- Quick run:
```shell
make l/init
make l/run
```

- Using docker:
```shell
make l/docker
```

- Using golang in local machine:
```shell
make l/run
```

- To run test use the command:
```shell
make l/test
```
