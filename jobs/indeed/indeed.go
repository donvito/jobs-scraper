package indeed

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

	fmt.Println("running indeed scraper")

	var jobPost jobs.JobPost
	sourceDomain := "https://sg.indeed.com"

	s.Collector.OnHTML("div.slider_item", func(e *colly.HTMLElement) {

		jobTitle := strings.TrimLeft(e.ChildText("h2.jobTitle"), "new")
		companyName := e.ChildText("span.companyName")
		location := e.ChildText("div.companyLocation")
		salary := e.ChildText("div.salary-snippet")
		summary := e.ChildText("div.job-snippet")
		jobUrl := fmt.Sprintf("%s%s", sourceDomain, e.ChildAttr("a", "href"))

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
		"https://sg.indeed.com/golang-jobs-in-Singapore",
		"https://sg.indeed.com/jobs?q=golang&l=Singapore&start=10",
		"https://sg.indeed.com/jobs?q=golang&l=Singapore&start=20",
		"https://sg.indeed.com/jobs?q=golang&l=Singapore&start=30",
		"https://sg.indeed.com/jobs?q=golang&l=Singapore&start=40",
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
		colly.AllowedDomains("www.indeed.com.sg", "indeed.com.sg", "sg.indeed.com"),
	)

	s := scraper{
		Collector: c,
	}

	return s
}
