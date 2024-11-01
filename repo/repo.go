package repo

import (
	"database/sql"
	"syndio/domains"
)

// For simplicity purposes, single repo for handling all the entities
type SyndioRepo interface {
	GetOrCreateJob(model domains.JobTitleModel) (int, error)
	CreateEmployment(model domains.EmploymentModel) error
}

type syndioRepo struct {
	db *sql.DB
}

func NewSyndioRepo(db *sql.DB) SyndioRepo {
	return &syndioRepo{
		db: db,
	}
}

func (r *syndioRepo) GetOrCreateJob(model domains.JobTitleModel) (int, error) {
	var jobId int64

	// Get the job if exists
	err := r.db.QueryRow("SELECT id FROM jobs WHERE job_title = ?", model.Title).Scan(&jobId)
	if err != nil {
		// check if no rows were returned
		if err == sql.ErrNoRows {
			// insert the new job
			query := "INSERT INTO jobs (department, job_title) VALUES (?, ?)"
			result, err := r.db.Exec(query, model.Department, model.Title)
			if err != nil {
				return 0, err
			}

			// get the id of the newly inserted job
			jobId, err = result.LastInsertId()
			if err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	return int(jobId), nil
}

func (r *syndioRepo) CreateEmployment(model domains.EmploymentModel) error {
	// create a new employment record
	query := "INSERT INTO employment (employee_id, job_id) VALUES (?, ?)"
	_, err := r.db.Exec(query, model.EmployeeId, model.JobId)
	return err
}
