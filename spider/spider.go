package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	SubjectId string `json:"subject_id"`
}

type SubjectResponse struct {
	SubjectId  string `json:"subject_id"`
	Title      string `json:"title"`
	CoverImage []byte `json:"cover_image"`
	Year       int    `json:"year"`

	Director    []string `json:"director"`
	Writer      []string `json:"writer"`
	LeadingRole []string `json:"leading_role"`
	Type        []string `json:"type"`

	Nation        []string `json:"nation"`
	Language      []string `json:"language"`
	ReleaseDate   []string `json:"release_date"`
	Length        string   `json:"length"`
	AlternateName []string `json:"alternate_name"`
	IMDB          string   `json:"imdb"`

	Intro string `json:"intro"`

	CommentCount      int     `json:"comment_count"`
	RatingNum         float32 `json:"rating_num"`
	RatingNum5Percent float32 `json:"rating_num_5_percent"`
	RatingNum4Percent float32 `json:"rating_num_4_percent"`
	RatingNum3Percent float32 `json:"rating_num_3_percent"`
	RatingNum2Percent float32 `json:"rating_num_2_percent"`
	RatingNum1Percent float32 `json:"rating_num_1_percent"`
}

type CommentResponse struct {
	SubjectId   string `json:"subject_id"`
	UserId      int    `json:"user_id"`
	Avatar      []byte `json:"avatar"`
	UserName    string `json:"user_name"`
	Star        int    `json:"star"`
	CommentDate string `json:"comment_date"`
	Content     string `json:"content"`
}

type ReviewsResponse struct {
	CommentResponse
	Title string `json:"title"`
}

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36"

func (obj *Request) SpiderSubject() (SubjectResponse, error) {
	url := fmt.Sprintf("https://movie.douban.com/subject/%s/", obj.SubjectId)
	client := http.Client{Timeout: time.Second * 15}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return SubjectResponse{}, err
	}
	req.Header.Set("User-Agent", userAgent)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return SubjectResponse{}, err
	}

	response := SubjectResponse{}
	response.SubjectId = obj.SubjectId

	title := doc.Find("h1 span[property='v:itemreviewed']").Text()
	response.Title = title

	// TODO
	response.CoverImage = []byte{}

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
		if i == 0 {
			response.RatingNum5Percent = float32(bar)
		}
		if i == 1 {
			response.RatingNum4Percent = float32(bar)
		}
		if i == 2 {
			response.RatingNum3Percent = float32(bar)
		}
		if i == 3 {
			response.RatingNum2Percent = float32(bar)
		}
		if i == 4 {
			response.RatingNum1Percent = float32(bar)
		}
	})

	foo = doc.Find("span[property='v:votes']").Text()
	commentCount, _ := strconv.Atoi(foo)
	response.CommentCount = commentCount

	return response, nil
}

func (obj *Request) SpiderComment() (CommentResponse, error) {

	return CommentResponse{}, nil
}

func (obj *Request) SpiderReviews() (ReviewsResponse, error) {

	return ReviewsResponse{}, nil
}
