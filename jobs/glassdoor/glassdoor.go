package glassdoor

import (
	"Scraper/jobs"
	"fmt"

	"github.com/gocolly/colly/v2"
)

type scraper struct {
	Collector *colly.Collector
}

func (s scraper) Scrape() (docs []interface{}, err error) {
	fmt.Println("running glassdoor scraper")

	s.Collector.OnHTML("li.react-job-listing", func(e *colly.HTMLElement) {

		sourceDomain := "https://www.glassdoor.com"
		jobTitle := e.Attr("data-normalize-job-title")
		companyName := e.ChildText("div.d-flex.flex-column div.d-flex.justify-content-between a.jobLink span")
		location := e.Attr("data-job-loc")
		salary := e.ChildText("span[data-test=detailSalary]")
		summary := e.Attr("data-normalize-job-title")
		jobUrl := fmt.Sprintf("%s%s", sourceDomain, e.ChildAttr("a.jobLink", "href"))

		jobPost := jobs.JobPost{
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
		"https://www.glassdoor.com/Job/jobs.htm?sc.keyword=golang&locT=N&locId=217",
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
		colly.AllowedDomains("www.glassdoor.com"),
	)

	s := scraper{
		Collector: c,
	}

	return s
}
