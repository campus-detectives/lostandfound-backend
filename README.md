# Lost and Found Backend

This repository contains the backend for a Lost and Found application.

## Usage

To run the API, make sure you have [Go](https://go.dev) installed and then use the following command:

```bash
go run ./cmd/api
```
or
```bash
make run/api
```

## Configuration

The API app uses environment variables for configuration. You can set these variables directly in your environment or use a `.env` file placed at the project root.

### Configuration Variables

The following configuration variables are used:

- **LAF_DSN**: PostgreSQL DSN (Data Source Name) to connect to the database server.

- **LAF_HTTP_ADDR**: The address on which the HTTP server will run. This should be specified in the form `host:port`.

### Setting Environment Variables

You can set these environment variables in your terminal or in a `.env` file. If using a `.env` file, make sure it is placed in the root directory of the project. An example `.env` file might look like this:

```
LAF_DSN=postgres://myusername:mypassword@localhost:5432/mydatabase
LAF_HTTP_ADDR=localhost:8000
```
