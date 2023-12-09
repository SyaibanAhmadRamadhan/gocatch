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

10. [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - The Go driver for MongoDB. This driver provides a simple, consistent API to MongoDB from Go applications.

11. [rueidis](https://github.com/redis/rueidis) - A Redis client library for Go that focuses on simplicity, correctness, and performance.
## Releases

Releases should follow the [Semantic Versioning](http://semver.org/) guidelines. The first couple of releases are `v1` and `v1.1`, with subsequent examples like `v1.2.1`.

[//]: # (# Project Folders)

[//]: # ()
[//]: # (This project is organized into several directories, each containing a different aspect of the code base.)

[//]: # (## [gcommon]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gcommon&#41;)

[//]: # (This directory contains common utility functions that can be used throughout the project. It includes functionalities like:)

[//]: # ()
[//]: # (- [Ternary functions]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcommon/ternary.go&#41;)

[//]: # (- [Panic related utilities]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcommon/common.go&#41;)

[//]: # (- [Random string creation]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcommon/random.go&#41;)

[//]: # ()
[//]: # (## [gcrypto]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gcrypto&#41;)

[//]: # ()
[//]: # (The `gcrypto` directory is home to all cryptographic related functions. It includes but not limited to:)

[//]: # ()
[//]: # (- [AES encryption and decryption]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcrypto/aes.go&#41;)

[//]: # (- [HMAC hash generation]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcrypto/hmac.go&#41;)

[//]: # (- [bcrypt for password hashing]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/gcrypto/bcyrpt.go&#41;)

[//]: # ()
[//]: # (## [ginfra]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/ginfra&#41;)

[//]: # ()
[//]: # (This directory contains the database-related functionalities of the project. It can be further divided into:)

[//]: # ()
[//]: # (- [gsql]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gdb/gsql&#41;: Contains functions related to SQL based databases. SQLx and PGX are majorly used here. and codegen for struct generation.)

[//]: # (- [gpostgre]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gdb/gsql/gpostgre&#41;: A dedicated directory for PostgreSQL related operations.)

[//]: # ()
[//]: # (## [gdir]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gdir&#41;)

[//]: # ()
[//]: # (The `gdir` directory contains functions related to:)

[//]: # ()
[//]: # (- Finding path from go module)

[//]: # (- Find go module & getting module name)

[//]: # ()
[//]: # (## [genv]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/genv&#41;)

[//]: # ()
[//]: # (This directory is utilized for:)

[//]: # ()
[//]: # (- Initializing and handling .env variables)

[//]: # ()
[//]: # (## [glog]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/glog&#41;)

[//]: # ()
[//]: # (Under this directory, you will find:)

[//]: # ()
[//]: # (- Canonical logging line utilities, predominantly using zerolog)

[//]: # ()
[//]: # (## [gmap]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gmap&#41;)

[//]: # ()
[//]: # (This directory simply contains:)

[//]: # ()
[//]: # (- Various map related utility functions)

[//]: # ()
[//]: # (## [gstr]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gstr&#41;)

[//]: # ()
[//]: # (Here you have:)

[//]: # ()
[//]: # (- Various string manipulation and handling functions)

[//]: # ()
[//]: # (## [gstruct]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gstruct&#41;)

[//]: # ()
[//]: # (It provides functionalities to:)

[//]: # ()
[//]: # (- Get struct tag values)

[//]: # (- Get struct field names)

[//]: # (- Get struct type)

[//]: # ()
[//]: # (## [gtime]&#40;https://github.com/SyaibanAhmadRamadhan/gocatch/tree/main/gtime&#41;)

[//]: # ()
[//]: # (This directory is dedicated to time related utilities:)

[//]: # ()
[//]: # (- Tracking execution time of a function)

[//]: # (- Formatting time duration from nanoseconds to days.)