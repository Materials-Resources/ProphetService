# Getting Started

# Deploying

Dependencies

- [mssql](https://www.microsoft.com/en-us/sql-server/sql-server-downloads)
- [redpanda](https://redpanda.com)
- A migrated Prophet 21 Database is required to use this as this service does not provide database migrations

### Standalone

Container images are built for each release tag and can be pulled from the [GitHub Container Registry](https://github.com/Materials-Resources/s_prophet/pkgs/container/s_prophet)

### Kubernetes

Helm charts are provided in the deployment folder and can be installed using the following commands:


# Developing

`make dev/setup`: This will install dependencies required for development.

`make qc`: Before pushing changes, run this to ensure that the code is error free and formatted correctly.


### Protobuf

The following proto schemas are used to interact with this microservice:
- catalog
- customer
- inventory
- order
- supplier

It is hosted on [buf.build](https://buf.build/materials-resources/s-prophet) which provides client libraries for various languages.

### Configuration

| YAML Key                | Environment Variable            | Description                                                           |
|-------------------------|---------------------------------|-----------------------------------------------------------------------|
| `database.host`         | `PROPHET_DATABASE_HOST`         | Hostname of the database                                              |
| `database.user`         | `PROPHET_DATABASE_USER`         | Username to connect to the database                                   |
| `database.password`     | `PROPHET_DATABASE_PASSWORD`     | Password to connect to the database                                   |
| `database.name`         | `PROPHET_DATABASE_NAME`         | Name of the database                                                  |
| `observability.service` | `PROPHET_OBSERVABILITY_SERVICE` | Name to use for tracing spans                                         |
| `observability.id`      | `PROPHET_OBSERVABILITY_ID`      | Identifier for application instance for tracing                       |
| `kafka.brokers`         | `PROPHET_KAFKA_BROKERS`         | URLs of the Kafka brokers                                             |
| `kafka.registry`        | `PROPHET_KAFKA_REGISTRY`        | URLs of the Kafka Schema Registry                                     |
| `enviroment`            | `PROPHET_ENVIRONMENT`           | Environment the service is running in (`Production` or `Development`) |
| `defaults`              | `PROPHET_DEFAULTS`              | Map of default values to use when interacting the Prophet21 database  |
