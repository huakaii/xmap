package main

import (
	"fmt"
	"xmap/core/spider"
)

func main() {
	fmt.Println(spider.Crawl("www.qq.com", "https://www.qq.com"))
}
