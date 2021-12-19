# paack-go
We are PaackAndGo, a small bus company. We want to create a REST API that helps us manage the trips that we offer.

# Directory tree

- `pkg`: Here is the main source code.
- `challenge`: Here is the initial challenge
- `postman`: Here is a `json` file to import to your local Postman.

# 1. Requirements

| Software         | Version | Importance                   |
| ---------------- | ------- | ---------------------------- |
| 🐳 Docker         | 20.10.12 | Required                     |
| 🐙 Docker Compose | 1.29.2  | Required                     |
| 🐃 GNU Make       | 4.2.1   | Optional                     |
| ‍🚀 Postman        | 9.1.5   | Optional                     |

# 2. Before run

1. Copy the `.env-example` file to `.env` file:

```shell
cp .env-example .env
```

2. Create the `logs` directory

```shell
mkdir logs
```


3. Fill out the variables with your own credentials. If you want to start quickly,
   these are the default values (do not modify):
```
API_PORT=8087

POSTGRES_DB=qas_db
POSTGRES_USER=qas
POSTGRES_PASSWORD=secret
POSTGRES_DRIVER=postgres
POSTGRES_HOST=qas_pgsql
```
