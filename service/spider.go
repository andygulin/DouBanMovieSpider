package service

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Spider interface {
	SpiderSubject() (SubjectResponse, error)
	SpiderComment() ([]CommentResponse, error)
	SpiderReview() ([]ReviewResponse, error)
	SpiderPhoto() ([]PhotoResponse, error)
}

type Request struct {
	SubjectId string `json:"subject_id" bson:"subject_id"`
}

func (obj *Request) SpiderSubject() (SubjectResponse, error) {
	url := fmt.Sprintf("%s/subject/%s/", Domain, obj.SubjectId)
	client := http.Client{Timeout: time.Second * 15}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return SubjectResponse{}, err
	}
	req.Header.Set("User-Agent", UserAgent)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	if res.StatusCode != http.StatusOK {
		log.Printf(fmt.Sprintf("status code error: %d %s\n", res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return SubjectResponse{}, err
	}

	response := SubjectResponse{}
	response.SubjectId = obj.SubjectId

	title := doc.Find("h1 span[property='v:itemreviewed']").Text()
	response.Title = title

	coverImage, _ := doc.Find("#mainpic a img").Attr("src")
	response.CoverImage = coverImage

	foo := doc.Find("h1 span[class='year']").Text()
	foo = strings.ReplaceAll(foo, "(", "")
	foo = strings.ReplaceAll(foo, ")", "")
	year, _ := strconv.Atoi(foo)
	response.Year = year

	info := doc.Find("div#info").Text()
	infos := strings.Split(info, "\n")
	for _, s := range infos {
		if idx := strings.Index(s, "导演: "); idx != -1 {
			t := s[idx+len("导演: "):]
			response.Director = strings.Split(t, " / ")
		}
		if idx := strings.Index(s, "编剧: "); idx != -1 {
			t := s[idx+len("编剧: "):]
			response.Writer = strings.Split(t, " / ")
		}
		if idx := strings.Index(s, "主演: "); idx != -1 {
			t := s[idx+len("主演: "):]
			response.LeadingRole = strings.Split(t, " / ")
		}
		if idx := strings.Index(s, "类型: "); idx != -1 {
			t := s[idx+len("类型: "):]
			response.Type = strings.Split(t, " / ")
		}

		if idx := strings.Index(s, "制片国家/地区: "); idx != -1 {
			t := s[idx+len("制片国家/地区: "):]
			response.Nation = strings.Split(t, " / ")
		}
		if idx := strings.Index(s, "语言: "); idx != -1 {
			t := s[idx+len("语言: "):]
			response.Language = strings.Split(t, " / ")
		}
		if idx := strings.Index(s, "上映日期: "); idx != -1 {
			t := s[idx+len("上映日期: "):]
			response.ReleaseDate = strings.Split(t, " / ")
		}
		if idx := strings.Index(s, "片长: "); idx != -1 {
			t := s[idx+len("片长: "):]
			response.Length = t
		}
		if idx := strings.Index(s, "又名: "); idx != -1 {
			t := s[idx+len("又名: "):]
			response.AlternateName = strings.Split(t, " / ")
		} else {
			response.AlternateName = make([]string, 0)
		}
		if idx := strings.Index(s, "IMDb: "); idx != -1 {
			t := s[idx+len("IMDb: "):]
			response.IMDB = t
		}
	}

	intro := doc.Find("span[property='v:summary']").Text()
	response.Intro = intro

	foo = doc.Find("strong[property='v:average']").Text()
	ratingNum, _ := strconv.ParseFloat(foo, 32)
	response.RatingNum = float32(ratingNum)

	doc.Find("div[class='ratings-on-weight'] div[class='item'] span[class='rating_per']").Each(func(i int, selection *goquery.Selection) {
		bar, _ := strconv.ParseFloat(strings.ReplaceAll(selection.Text(), "%", ""), 32)
		switch i {
		case 0:
			response.RatingNumPercent.RatingNum5 = float32(bar)
		case 1:
			response.RatingNumPercent.RatingNum4 = float32(bar)
		case 2:
			response.RatingNumPercent.RatingNum3 = float32(bar)
		case 3:
			response.RatingNumPercent.RatingNum2 = float32(bar)
		case 4:
			response.RatingNumPercent.RatingNum1 = float32(bar)
		default:

		}
	})

	foo = doc.Find("span[property='v:votes']").Text()
	commentCount, _ := strconv.Atoi(foo)
	response.CommentCount = commentCount

	return response, nil
}

func (obj *Request) SpiderComment() ([]CommentResponse, error) {
	var rets []CommentResponse

	const statusP = "P"
	const statusF = "F"

	var err error
	var pRets []CommentResponse
	var fRets []CommentResponse
	pRets, err = spiderComment0(obj.SubjectId, statusP)
	if err == nil {
		rets = append(rets, pRets...)
	}
	fRets, err = spiderComment0(obj.SubjectId, statusF)
	if err == nil {
		rets = append(rets, fRets...)
	}
	return rets, nil
}

func spiderComment0(subjectId string, status string) ([]CommentResponse, error) {
	var rets []CommentResponse

	start := 0
	limit := 20
	for {
		url := fmt.Sprintf("%s/subject/%s/comments?start=%d&limit=%d&status=%s&sort=new_score", Domain, subjectId, start, limit, status)
		client := http.Client{Timeout: time.Second * 15}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", UserAgent)
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			log.Printf(fmt.Sprintf("status code error: %d %s\n", res.StatusCode, res.Status))
			break
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}

		itemCount := doc.Find(".comment-item").Length()
		if itemCount > 0 {
			doc.Find(".comment-item").Each(func(i int, selection *goquery.Selection) {
				var commentResponse CommentResponse

				commentResponse.SubjectId = subjectId
				if status == "P" {
					commentResponse.HaveSee = true
				}
				if status == "F" {
					commentResponse.HaveSee = false
				}

				avatar, _ := selection.Find(".avatar a img").Attr("src")
				commentResponse.Avatar = avatar

				userName := selection.Find(".comment-info a").Text()
				commentResponse.UserName = userName

				star1 := selection.Find(".comment-info .rating").HasClass("allstar10")
				star2 := selection.Find(".comment-info .rating").HasClass("allstar20")
				star3 := selection.Find(".comment-info .rating").HasClass("allstar30")
				star4 := selection.Find(".comment-info .rating").HasClass("allstar40")
				star5 := selection.Find(".comment-info .rating").HasClass("allstar50")
				if star1 {
					commentResponse.Star = 1
				} else if star2 {
					commentResponse.Star = 2
				} else if star3 {
					commentResponse.Star = 3
				} else if star4 {
					commentResponse.Star = 4
				} else if star5 {
					commentResponse.Star = 5
				} else {
					commentResponse.Star = 0
				}

				content := selection.Find(".comment-content span.short").Text()
				commentResponse.Content = content

				commentDate := selection.Find(".comment-info span.comment-time").Text()
				commentDate = strings.ReplaceAll(commentDate, "\n", "")
				commentDate = strings.Trim(commentDate, " ")
				commentResponse.CommentDate = commentDate

				rets = append(rets, commentResponse)
			})
		} else {
			break
		}
		start += limit
	}
	return rets, nil
}

func (obj *Request) SpiderReview() ([]ReviewResponse, error) {
	var rets []ReviewResponse

	start := 0
	limit := 20
	for {
		url := fmt.Sprintf("%s/subject/%s/reviews?start=%d", Domain, obj.SubjectId, start)
		client := http.Client{Timeout: time.Second * 15}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", UserAgent)
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			log.Printf(fmt.Sprintf("status code error: %d %s\n", res.StatusCode, res.Status))
			break
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}

		itemCount := doc.Find(".review-item").Length()
		if itemCount > 0 {
			doc.Find(".review-item").Each(func(i int, selection *goquery.Selection) {
				var reviewResponse ReviewResponse

				reviewResponse.SubjectId = obj.SubjectId
				reviewResponse.HaveSee = true

				avatar, _ := selection.Find(".main-hd .avator img").Attr("src")
				reviewResponse.Avatar = avatar

				userName := selection.Find(".main-hd .name").Text()
				reviewResponse.UserName = userName

				star1 := selection.Find(".main-title-rating").HasClass("allstar10")
				star2 := selection.Find(".main-title-rating").HasClass("allstar20")
				star3 := selection.Find(".main-title-rating").HasClass("allstar30")
				star4 := selection.Find(".main-title-rating").HasClass("allstar40")
				star5 := selection.Find(".main-title-rating").HasClass("allstar50")
				if star1 {
					reviewResponse.Star = 1
				} else if star2 {
					reviewResponse.Star = 2
				} else if star3 {
					reviewResponse.Star = 3
				} else if star4 {
					reviewResponse.Star = 4
				} else if star5 {
					reviewResponse.Star = 5
				} else {
					reviewResponse.Star = 0
				}

				commentDate := selection.Find(".main-meta").Text()
				reviewResponse.CommentDate = commentDate

				title := selection.Find(".main-bd h2 a").Text()
				reviewResponse.Title = title

				summary := selection.Find(".short-content").Text()
				reviewResponse.Summary = summary

				reviewId, _ := selection.Attr("id")
				rUrl := fmt.Sprintf("%s/j/review/%s/full", Domain, reviewId)

				req2, _ := http.NewRequest("GET", rUrl, nil)
				req2.Header.Set("User-Agent", UserAgent)
				res2, _ := client.Do(req2)

				b, _ := io.ReadAll(res2.Body)
				type ReviewContent struct {
					Body string `json:"body"`
					Html string `json:"html"`
				}
				var reviewContent ReviewContent
				_ = json.Unmarshal(b, &reviewContent)
				reviewResponse.Content = reviewContent.Html

				rets = append(rets, reviewResponse)
			})
		} else {
			break
		}

		start += limit
	}
	return rets, nil
}

func (obj *Request) SpiderPhoto() ([]PhotoResponse, error) {
	var rets []PhotoResponse

	start := 0
	limit := 30
	for {
		url := fmt.Sprintf("%s/subject/%s/photos?type=S&start=%d&sortby=like&size=a&subtype=a", Domain, obj.SubjectId, start)
		client := http.Client{Timeout: time.Second * 15}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", UserAgent)
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			log.Printf(fmt.Sprintf("status code error: %d %s\n", res.StatusCode, res.Status))
			break
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}

		itemCount := doc.Find(".poster-col3 li").Length()
		if itemCount > 0 {
			doc.Find(".poster-col3 li").Each(func(i int, selection *goquery.Selection) {
				var photoResponse PhotoResponse

				photoResponse.SubjectId = obj.SubjectId
				thumbnailImage, _ := selection.Find(".cover a img").Attr("src")
				photoResponse.ThumbnailImage = thumbnailImage

				photoResponse.LargeImage = strings.ReplaceAll(thumbnailImage, "/m/", "/l/")
				photoResponse.RawImage = strings.ReplaceAll(thumbnailImage, "/m/", "/raw/")

				foo := selection.Find(".prop").Text()
				foo = strings.ReplaceAll(foo, " ", "")
				foo = strings.ReplaceAll(foo, "\n", "")
				bar := strings.Split(foo, "x")
				photoResponse.RawSize.Width, _ = strconv.Atoi(bar[0])
				photoResponse.RawSize.Height, _ = strconv.Atoi(bar[1])

				rets = append(rets, photoResponse)
			})
		} else {
			break
		}

		start += limit
	}
	return rets, nil
}
