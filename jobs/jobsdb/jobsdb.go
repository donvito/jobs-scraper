package jobsdb

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

	fmt.Println("running jobsdb scraper")

	var jobPost jobs.JobPost
	sourceDomain := "https://sg.jobsdb.com"

	s.Collector.OnHTML("article.job-card", func(e *colly.HTMLElement) {

		jobTitle := e.ChildText("a.job-link")
		companyName := e.ChildText("span.job-company")
		location := e.ChildText("span.job-location")

		var salary string
		badge := strings.TrimLeft(e.ChildText("div.badge"), "new")
		if strings.ToLower(badge) != "quick apply" {
			salary = badge
		}

		summary := e.ChildText("div.job-abstract")
		jobUrl := fmt.Sprintf("%s%s", "https://sg.jobsdb.com", e.ChildAttr("a", "href"))

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
		err = s.Collector.Visit(page)
		if err != nil {
			continue
		}
	}

	return
}

func Scraper() scraper {

	c := colly.NewCollector(
		colly.AllowedDomains("sg.jobsdb.com"),
	)

	s := scraper{
		Collector: c,
	}

	return s
}
