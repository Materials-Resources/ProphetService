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

