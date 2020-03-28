package jobs

const (
	JobCategoryGo = "go"
	RocksetWorkspace = "jobportal"
	RocksetCollection = "jobs"
)

type JobPost struct {
	Title       string
	CompanyName string
	Location    string
	Salary      string
	Summary     string
	Category    string
	JobUrl      string
	Source      string
}

type Scraper interface {
	Scrape() (docs []interface{}, err error)
}