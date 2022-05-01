package main

import (
	"github.com/dev-eunjiLee/learngo/scrapper"
	"github.com/labstack/echo"
	"os"
	"strings"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

const fileName string = "jobs.csv"

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	// echo.Context.Attachment(): 첨부파일 리턴
	return c.Attachment(fileName, fileName)
}

func main() {

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))

}
