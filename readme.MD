# GOCATCH

[//]: # (This is a Go-based project developed to perform several database operations and migrations featuring an intuitive command line interface.)

## Libraries Used

This project makes use of the following external Go Libraries:

1. [goose](https://github.com/pressly/goose) - A database migration handling library, version 3. Goose creates separate SQL migration files that get applied in order to bring your database schema to the desired state.

2. [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit, version 5. It provides robust, full-featured, and high performance connectivity solutions for accessing PostgreSQL databases.

3. [godotenv](https://github.com/joho/godotenv) - A Go port of the Ruby dotenv library (which loads environment variables from a .env file into ENV when the environment is bootstrapped).
   
4. [dockertest](https://github.com/ory/dockertest) - A Go library for Docker integration testing. It provides easy access to Docker containers for creating isolated environments to run tests in.

5. [ulid](https://github.com/oklog/ulid) - A Go library for generete ulid, click link for detail information.

6. [zerolog](https://github.com/rs/zerolog) - A Go library for Logging, click link for detail information.

## Releases

Releases should follow [Semver](http://semver.org/) though the first couple of releases are `v1` and `v1.1`.  Example `v1.2.1`