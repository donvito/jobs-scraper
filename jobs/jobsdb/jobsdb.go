package jobsdb

import (
	"Scraper/jobs"
	"fmt"
	"github.com/gocolly/colly/v2"
)

type scraper struct {
	Collector *colly.Collector
}

func (s scraper) Scrape() (docs []interface{}, err error) {

	fmt.Println("running jobsdb scraper")

	var jobPost jobs.JobPost
	sourceDomain := "https://sg.indeed.com"

	s.Collector.OnHTML("div.job", func(e *colly.HTMLElement) {

		jobTitle := e.ChildText("a.jobtitle")
		companyName := e.ChildText("span.company")
		location := e.ChildText("span.location")
		salary := e.ChildText("span.salaryText")
		summary := e.ChildText("div.summary")
		jobUrl := fmt.Sprintf("%s%s", "https://sg.jobsdb.com",  e.ChildAttr("a", "href"))

		jobPost = jobs.JobPost{
			Title:       jobTitle,
			CompanyName: companyName,
			Location:    location,
			Salary:      salary,
			Summary:     summary,
			JobUrl:      jobUrl,
			Category:    jobs.JobCategoryGo,

			Source: sourceDomain,
		}

		docs = append(docs, jobPost)

	})

	pagesToVisit := []string{
		"https://sg.jobsdb.com/j?q=golang&l=&sp=homepage",
		"https://sg.jobsdb.com/j?l=&p=2&q=golang&sp=homepage",
	}

	for _, page := range pagesToVisit {
		s.Collector.Visit(page)
		if err != nil {
			continue
		}
	}

	return
}


func Scraper() scraper {

	c := colly.NewCollector(
		colly.AllowedDomains("www.indeed.com.sg", "indeed.com.sg", "sg.indeed.com"),
	)

	s := scraper{
		Collector:  c,
	}

	return s
}