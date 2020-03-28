package linkedin

import (
	"Scraper/jobs"
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

type scraper struct {
	Collector *colly.Collector
}

func (s scraper) Scrape() (docs []interface{}, err error) {

	fmt.Println("running linkedin scraper")

	var jobPost jobs.JobPost
	sourceDomain := "https://sg.indeed.com"

	s.Collector.OnHTML("li.job-result-card", func(e *colly.HTMLElement) {

		jobTitle := e.ChildText("span.screen-reader-text")
		companyName := e.ChildText("a.job-result-card__subtitle-link")
		location := e.ChildText("span.job-result-card__location")
		salary := e.ChildText("span.salaryText")
		summary := e.ChildText("p.job-result-card__snippet")
		jobUrl := fmt.Sprintf("%s", e.ChildAttr("a", "href"))
		jobUrl =  strings.Replace(jobUrl, "(MISSING)", "", -1)

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
		"https://www.linkedin.com/jobs/golang-jobs/?originalSubdomain=sg&position=1&pageNum=0",
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
		colly.AllowedDomains("www.linkedin.com"),
	)

	s := scraper{
		Collector:  c,
	}

	return s
}