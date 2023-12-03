# GOCATCH

[//]: # (This is a Go-based project developed to perform several database operations and migrations featuring an intuitive command line interface.)

## Libraries Used

This project utilises the following external Go libraries:

1. [goose](https://github.com/pressly/goose) - A library for managing database migrations, version 3. This utility creates discrete SQL migration files that get sequentially applied to transition your database schema to the desired state.

2. [pgx](https://github.com/jackc/pgx) - A PostgreSQL driver and toolkit, version 5. This tool provides robust, feature-rich, and high performance connectivity solutions for accessing PostgreSQL databases.

3. [godotenv](https://github.com/joho/godotenv) - A Go port of the Ruby dotenv library (which loads environment variables from a .env file into ENV when the environment is bootstrapped).

4. [dockertest](https://github.com/ory/dockertest) - A Go library for executing Docker-based integration testing. It provides easy access to Docker containers to set up isolated environments for running tests.

5. [ulid](https://github.com/oklog/ulid) - A Go library for generating ULIDs (Universally Unique Lexicographically Sortable Identifier).

6. [zerolog](https://github.com/rs/zerolog) - A Go library for logging.

7. [mysql](https://github.com/go-sql-driver/mysql) - A MySQL driver for Go.

8. [sqlx](https://github.com/jmoiron/sqlx) - An extension to the standard Go "database/sql" library.

9. [pq](https://github.com/lib/pq) - A driver and toolkit for PostgreSQL.

## Releases

Releases should follow the [Semantic Versioning](http://semver.org/) guidelines. The first couple of releases are `v1` and `v1.1`, with subsequent examples like `v1.2.1`.