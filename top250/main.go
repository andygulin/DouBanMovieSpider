package main

import (
	"DouBanMovieSpider/service"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var subjectIds []string

	for i := 1; i < 10; i++ {
		start := i * 25

		url := fmt.Sprintf("%s/top250?start=%d&filter=", service.Domain, start)
		client := http.Client{Timeout: time.Second * 15}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("User-Agent", service.UserAgent)
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			fmt.Printf("status code error: %d %s\n", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".item .pic a").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists {
				subjectId := strings.Split(href, "/")[4]
				subjectIds = append(subjectIds, subjectId)
			}
		})
	}

	fmt.Println(len(subjectIds))

	storeHandle := new(service.StoreHandle)
	for _, subjectId := range subjectIds {
		obj := &service.Request{SubjectId: subjectId}

		subjectResponse, _ := obj.SpiderSubject()
		_, _ = storeHandle.HandleSubject(subjectResponse)

		commentResponses, _ := obj.SpiderComment()
		_, _ = storeHandle.HandleComment(commentResponses)

		photoResponses, _ := obj.SpiderPhoto()
		_, _ = storeHandle.HandlePhoto(photoResponses)

		reviewResponses, _ := obj.SpiderReview()
		_, _ = storeHandle.HandleReview(reviewResponses)

		fmt.Printf("subjectId %s Success:", subjectId)
	}
}
