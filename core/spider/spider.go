package spider

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
)

type UrlInfo struct {
	Url    string
	Method string
	Data   string
}

var visited = map[string]bool{}

func Crawl(allowdomain string, spiderurl string) []UrlInfo {
	var SpiderData []UrlInfo

	c := colly.NewCollector(
		colly.AllowedDomains(allowdomain),
		colly.Async(true),
		//colly.Debugger(&debug.LogDebugger{}),
		//colly.MaxDepth(10),
	)
	_ = c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})

	c.OnHTML("form", func(e *colly.HTMLElement) {
		inputNames := e.ChildAttrs("input", "name")

		for k, v := range inputNames {
			inputNames[k] = v + "="
		}
		data := strings.Join(inputNames, "&")
		method := strings.ToUpper(e.Attr("method"))
		link := e.Request.AbsoluteURL(e.Attr("action"))

		u, err := url.Parse(link)
		if err != nil {
			panic(err)
		}

		if (allowdomain != "") && (u.Host == allowdomain) {
			SpiderData = append(SpiderData, UrlInfo{link, method, data})

		}

	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))

		// 禁止访问相同url
		if visited[link] {
			return
		}
		// TODO ignore file

		visited[link] = true

		//SpiderData = append(SpiderData, UrlInfo{link, "GET", ""})

		c.Visit(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		SpiderData = append(SpiderData, UrlInfo{r.URL.String(), r.Method, ""})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	c.Visit(spiderurl)

	c.Wait()
	return SpiderData
}
