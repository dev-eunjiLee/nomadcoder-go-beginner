package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

var baseUrl string = "https://kr.indeed.com/%EC%B7%A8%EC%97%85?q=python&limit=50&vjk=6dd48f3771d01215"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseUrl + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // res.Body는 IO(입력과 출력)으로 함수가 끝났을 때 닫아야 한다

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, s *goquery.Selection) {
		id, _ := s.Find("div>h2>a").Attr("data-jk")
		//text := s.Find("h2>a").Attr('"da')
		fmt.Println(id)
	})
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
