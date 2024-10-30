# Syndio Backend App

Using the `employees.db` sqlite database in this repository with the following table/data:

```
sqlite> .open employees.db
sqlite> .schema employees
CREATE TABLE employees (id INTEGER PRIMARY KEY, gender TEXT not null);
sqlite> SELECT * FROM employees;
1|male
2|male
3|male
4|female
5|female
6|female
7|non-binary
```

1) create an api with an endpoint that saves the incoming job data for the corresponding employee.
   2) The incoming data will look as follows:
```
[
  { "id": 1, department": "Engineering", "job_title": "Senior Enginer" },
  { "id": 2, department": "Engineering", "job_title": "Super Senior Enginer" },
  { "id": 3, "department": "Sales", "job_title": "Head of Sales"},
  { "id": 4, "department": "Support", "job_title": "Tech Support" },
  { "id": 5, "department": "Engineering", "job_title": "Junior Enginer" },
  { "id": 6, "department": "Sales", "job_title": "Sales Rep" },
  { "id": 7, "department": "Marketing", "job_title": "Senior Marketer" },
]
```

## Requirements

- The api must take an environment variable `PORT` and respond to requests on that port.
- Provide basic setup instructions for:
  - What is required to run the api
  - how to ingest the data through the new endpoint,
  - a way to update the existing database given to you.

## Success

- We can run the api and ingest data from your setup instructions
- The api is written in Python or Go

## Not Required

- Tests
- Logging, monitoring, or anything more than basic error handling

## Submission

- Respond to the email you received giving you this with:
  - a zip file, or link to a git repo
  - instructions on how to setup and run the code (could be included w/ zip/git)
- We'll follow the setup instructions to test it on a local machine, then we'll get back to you.

## Notes

- Keep it simple
- We expect this to take less than an hour, please try and limit your effort to that window.
- If the api works, and returns what we requested, its a success.
- Anything extra (tests, other endpoints, ...) is not worth bonus/etc.
- We truly value your time and just want a basic benchmark and common piece of code to use in future interviews.
- If we bring you in for in-person interviews we'll expand on this submission.
