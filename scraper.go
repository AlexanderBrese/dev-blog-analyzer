package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

// todo use domains in colly.AllowedDomains()
// var domains = [4]string{"dev.to", "hackernoon.com", "realpython.com", "hacker.io"}

type article struct {
	title        string
	author       string
	published_at string
	thumbnail    string
	video_url    string
	likes        int
	content      string
	comments     string
	word_count   int
	tags         string
}

func main() {
	articles := []article{}

	c := colly.NewCollector(
		colly.UserAgent("dev-blog-scraperv1"),
		colly.AllowedDomains("dev.to", "hackernoon.com", "realpython.com", "hacker.io"),
		colly.Async(true),
	)

	// getting title & storing into article.title
	c.OnHTML("div[id=page-content]", func(e *colly.HTMLElement) {
		temp := article{}
		temp.title = e.ChildText("div[class=crayons-article__header__meta]")
		articles = append(articles, temp)

	})

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		Parallelism: 4,
		RandomDelay: 5 * time.Second,
	})

	// print visiting when making request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://https://dev.to/")
	c.Visit("https://hackr.io/blog")
	c.Visit("https://realpython.com/")
	c.Visit("https://hackernoon.com/")
	c.Wait()

	fmt.Println(articles)
}
