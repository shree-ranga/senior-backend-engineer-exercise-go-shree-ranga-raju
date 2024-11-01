package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"syndio/contracts"
	"syndio/domains"
	"syndio/repo"
)

type SyndioService interface {
	CreateEmployments(ctx context.Context, jobs []contracts.EmploymentContract) error
}

type syndioService struct {
	repo repo.SyndioRepo
}

func NewSyndioService(repo repo.SyndioRepo) SyndioService {
	return &syndioService{
		repo: repo,
	}
}

func (s *syndioService) CreateEmployments(ctx context.Context, jobs []contracts.EmploymentContract) error {
	if len(jobs) == 0 {
		return nil
	}

	var failedJobs []contracts.EmploymentContract

	for _, job := range jobs {
		// Normalize department and job title
		department := strings.ToLower(job.Department)
		title := strings.ToLower(job.JobTitle)

		// get or create jobId
		jobModel := domains.JobTitleModel{
			Department: department,
			Title:      title,
		}
		jobId, err := s.repo.GetOrCreateJob(jobModel)
		if err != nil {
			log.Printf("Failed to insert job %v: %v\n", job, err)
			failedJobs = append(failedJobs, job)
			continue // skip creating employment for this job and continue
		}

		// create new employment
		employmentModel := domains.EmploymentModel{
			EmployeeId: job.EmployeeId,
			JobId:      jobId,
		}
		if err := s.repo.CreateEmployment(employmentModel); err != nil {
			failedJobs = append(failedJobs, job)
			log.Printf("Failed to create employment for job %v: %v\n", job, err)
		}
	}

	if len(failedJobs) > 0 {
		return fmt.Errorf("failed to create employment records for %d job(s): %v", len(failedJobs), failedJobs)
	}

	return nil
}
