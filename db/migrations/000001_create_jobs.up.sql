CREATE TABLE IF NOT EXISTS jobs (
    id INTEGER PRIMARY KEY,
    department TEXT NOT NULL,
    job_title TEXT UNIQUE NOT NULL
);