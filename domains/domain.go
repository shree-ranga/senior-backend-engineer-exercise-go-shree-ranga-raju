package domains

type EmployeeModel struct {
	Id       int
	Gender   string
	JobTitle *string
}

type JobTitleModel struct {
	Id         int
	Department string
	Title      string
}

type EmploymentModel struct {
	Id         int
	EmployeeId int
	JobId      int
}
