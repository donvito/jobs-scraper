package main

import (
	"Scraper/jobs"
	"Scraper/jobs/indeed"
	"Scraper/jobs/jobsdb"
	"Scraper/jobs/linkedin"
	"fmt"
)

func main() {

	scrapers := []jobs.Scraper{
		indeed.Scraper(),
		linkedin.Scraper(),
		jobsdb.Scraper(),
	}

	var documents []interface{}

	for _, s := range scrapers {
		docs, err := s.Scrape()
		documents = append(documents, docs...)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	for _, d := range documents {
		fmt.Printf("%v\n\n", d)
	}

	err := jobs.SaveToRockset(documents)
	if err != nil {
		fmt.Println(err.Error())
	}

}
