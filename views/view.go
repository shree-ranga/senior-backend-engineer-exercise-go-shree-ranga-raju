package views

type Employee struct {
	Id       int    `json:"id"`
	Gender   string `json:"gender"`
	JobTitle string `json:"job_title"`
}

type JobTitle struct {
	Id         int    `json:"id"`
	Department string `json:"department"`
	Title      string `json:"job_title"`
}

type JobTitles struct {
	Titles []JobTitle `json:"job_titles"`
}
