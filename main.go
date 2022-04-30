package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var baseUrl string = "https://kr.indeed.com/%EC%B7%A8%EC%97%85?q=python&limit=50&vjk=6dd48f3771d01215"

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

func main() {
	totalPages := getPages()

	var jobs []extractedJob

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}

	fmt.Println(jobs)
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseUrl + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // res.Body는 IO(입력과 출력)으로 함수가 끝났을 때 닫아야 한다

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs

}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Find("div>h2>a").Attr("data-jk")
	//text := s.Find("h2>a").Attr('"da')

	title := cleanString(card.Find("a>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	salary := cleanString(card.Find(".salary-snippet").Text())
	summary := cleanString(card.Find(".job-snippet").Text())

	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}

}

// getPages: 페이지 수 리턴하는 함수
func getPages() int {

	pages := 0

	res, err := http.Get(baseUrl)

	// 에러가 난 경우 프로그램 종료
	checkErr(err)

	// statusCode도 체크 후 이상이 있는 경우 프로그램 종료
	checkCode(res)

	defer res.Body.Close() // res.Body는 IO(입력과 출력)으로 함수가 끝났을 때 닫아야 한다

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination ").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

func cleanString(str string) string {
	// strings.Fields: string을 모든 단어마다 분리해 string 배열로 반환
	// strings.TrimSpace:  string의 앞 뒤 공백 제거
	// strings.Join(string, 합칠 때 사용할 구분자): string 배열 합치서 하나의 string으로 리턴
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
