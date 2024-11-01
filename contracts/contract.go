package contracts

type EmploymentContract struct {
	EmployeeId int    `json:"id"`
	Department string `json:"department"`
	JobTitle   string `json:"job_title"`
}
