# Syndio Backend App Solution

### Pre-requisites

- macOS Monterey or later
- Go version 1.23
- golang-migrate

### Folder Structure

```bash
senior-backend-engineer-exercise-go-shree-ranga-raju
├── README.md
├── appContext
│   └── appContext.go
├── contracts
│   └── contract.go
├── db
│   └── migrations
│       ├── 000001_create_jobs.down.sql
│       ├── 000001_create_jobs.up.sql
│       ├── 000002_create_employment.down.sql
│       ├── 000002_create_employment.up.sql
│       └── migrate.go
├── domains
│   └── domain.go
├── employees.db
├── go.mod
├── go.sum
├── handlers
│   └── handler.go
├── main.go
├── repo
│   └── repo.go
├── router
│   └── router.go
├── run.sh
├── services
│   ├── dependencies.go
│   └── service.go
├── test.db
└── views
    └── view.go

11 directories, 21 files
```

### Project setup

1. Open a terminal and clone the repository

```bash
git clone https://github.com/shree-ranga/senior-backend-engineer-exercise-go-shree-ranga-raju.git
```

2. Navivgate to the project directory

```bash
cd senior-backend-engineer-exercise-go-shree-ranga-raju
```

### Run Server

1. Set the desired server port:

```bash
export SYNDIO_PORT=5000 #replace with your desired port number
```

2. Start the server:

```bash
go mod tidy
sh run.sh
```

3. Curl <br>
   Open up a new terminal window

```bash
export SYNDIO_PORT=5000 #make sure it's the same port from step 1
curl -X POST http://127.0.0.1:${SYNDIO_PORT}/employments \
     -H "Content-Type: application/json" \
     -d '[
         { "id": 1, "department": "Engineering", "job_title": "Senior Engineer" },
         { "id": 2, "department": "Engineering", "job_title": "Super Senior Engineer" },
         { "id": 3, "department": "Sales", "job_title": "Head of Sales" },
         { "id": 4, "department": "Support", "job_title": "Tech Support" },
         { "id": 5, "department": "Engineering", "job_title": "Junior Engineer" },
         { "id": 6, "department": "Sales", "job_title": "Sales Rep" },
         { "id": 7, "department": "Marketing", "job_title": "Senior Marketer" }
     ]'
```

### Migrations

Using a [standalone migration tool](https://github.com/golang-migrate/migrate) to manage migrations.
With this tool, need to create the migration files manually.

#### Setup

```
brew install golang-migrate
```

#### Create Migration files

To create migration files manually, run the below command in the project dir.

```sh
migrate create -ext sql -dir ./db/migrations -seq [MIGRATION_NAME]
# -ext sql: Specifies that the migration file extension should be .sql.
# -dir ./migrations: Specifies the target directory for migration files.
# -seq: Creates the migration with a sequential number (e.g., 000001_sample.up.sql and 000001_sample.down.sql).
```

This command will create `.up.sql`, `.down.sql` files in the `migrations` dir.

#### Adding SQL to migration files

Add SQL statements directly into the `.up.sql` file (and corresponding rollback statements in `.down.sql`) in the `db/migrations` directory.

### Database Schema

Added two new tables:

1. `jobs`
2. `employment`
3. `schema_migrations` (added by golang-migrate)
   <br><br>
   `employees.db` new schema as follows:

```bash
sqlite> .open employees.db
sqlite> .schema
CREATE TABLE employees (id INTEGER PRIMARY KEY, gender TEXT not null);
CREATE TABLE schema_migrations (version uint64,dirty bool);
CREATE UNIQUE INDEX version_unique ON schema_migrations (version);
CREATE TABLE employment (
    id INTEGER PRIMARY KEY,
    employee_id INTEGER NOT NULL,
    job_id INTEGER NOT NULL
);
CREATE TABLE jobs (    id INTEGER PRIMARY KEY,    department TEXT NOT NULL,    job_title TEXT UNIQUE NOT NULL);
```
