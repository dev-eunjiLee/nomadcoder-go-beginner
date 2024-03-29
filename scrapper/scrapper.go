package scrapper

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

// Scrape Indeed by a term
func Scrape(term string) {

	var baseUrl = "https://kr.indeed.com/%EC%B7%A8%EC%97%85?q=" + term + "&limit=50&vjk=6dd48f3771d01215"

	totalPages := getPages(baseUrl)

	var jobs []extractedJob

	c := make(chan []extractedJob)

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseUrl, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
}

func getPage(page int, url string, mainC chan<- []extractedJob) {

	var jobs []extractedJob

	c := make(chan extractedJob)

	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // res.Body는 IO(입력과 출력)으로 함수가 끝났을 때 닫아야 한다

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)

	}
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Find("div>h2>a").Attr("data-jk")
	//text := s.Find("h2>a").Attr('"da')

	title := CleanString(card.Find("a>span").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	salary := CleanString(card.Find(".salary-snippet").Text())
	summary := CleanString(card.Find(".job-snippet").Text())

	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}

}

// getPages: 페이지 수 리턴하는 함수
func getPages(baseUrl string) int {

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

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)

	// Writer.Flush(): 이 시점에 파일에 데이터가 입력된다
	// for문까지 다 돌고 나면 Write된 데이터가 파일에 입력!
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}

	fmt.Println("Done, extracted", len(jobs))
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

// CleanString cleans a string
func CleanString(str string) string {
	// strings.Fields: string을 모든 단어마다 분리해 string 배열로 반환
	// strings.TrimSpace:  string의 앞 뒤 공백 제거
	// strings.Join(string, 합칠 때 사용할 구분자): string 배열 합치서 하나의 string으로 리턴
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
