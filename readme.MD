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

# Project Folders

This project is organized into several directories, each containing a different aspect of the code base.
## [gcommon](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gcommon)
This directory contains common utility functions that can be used throughout the project. It includes functionalities like:

- [Ternary functions](https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcommon/ternary.go)
- [Panic related utilities](https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcommon/common.go)
- [Random string creation](https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcommon/random.go)

## [gcrypto](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gcrypto)

The `gcrypto` directory is home to all cryptographic related functions. It includes but not limited to:

- [AES encryption and decryption](https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcrypto/aes.go)
- [HMAC hash generation](https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcrypto/hmac.go)
- [bcrypt for password hashing](https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcrypto/bcyrpt.go)

## [gdb](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gdb)

This directory contains the database-related functionalities of the project. It can be further divided into:

- [gsql](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gdb/gsql): Contains functions related to SQL based databases. SQLx and PGX are majorly used here. and codegen for struct generation.
- [gPOSTGRES](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gdb/gsql/gPOSTGRE): A dedicated directory for PostgreSQL related operations.

## [gdir](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gdir)

The `gdir` directory contains functions related to:

- Finding path from go module
- Find go module & getting module name

## [genv](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/genv)

This directory is utilized for:

- Initializing and handling .env variables

## [glog](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/glog)

Under this directory, you will find:

- Canonical logging line utilities, predominantly using zerolog

## [gmap](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gmap)

This directory simply contains:

- Various map related utility functions

## [gstr](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gstr)

Here you have:

- Various string manipulation and handling functions

## [gstruct](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gstruct)

It provides functionalities to:

- Get struct tag values
- Get struct field names
- Get struct type

## [gtime](https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gtime)

This directory is dedicated to time related utilities:

- Tracking execution time of a function
- Formatting time duration from nanoseconds to days.