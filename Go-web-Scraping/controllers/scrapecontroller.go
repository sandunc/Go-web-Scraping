package controllers

import (
	mod "Go-web-Scraping/models"

	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

var HtmlScraping mod.Scarping
var urlString mod.UrlString

func GetResult(c *gin.Context) {
	urlString.Reset()
	HtmlScraping.Reset()

	// Bind request payload with our model
	if binderr := c.ShouldBindJSON(&urlString); binderr != nil {
		var msgString = "Error occurred while binding request data"
		log.Fatal(msgString)
		return
	}

	log.Println(urlString.UrlStr + "------- is Scraping-----")

	// Request the HTML page.
	res, err := http.Get(urlString.UrlStr)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//find the title --------------
	text := doc.Find("title").Text()
	re := regexp.MustCompile("\\s{2,}")
	text = re.ReplaceAllString(text, "\n")
	HtmlScraping.Title = text

	//find the version --------------
	doc.Find("html").Each(func(index int, selector *goquery.Selection) {
		link, _ := selector.Attr("version") // get the links
		if len(link) != 0 {
			HtmlScraping.Version = link
		}
	})

	// Find the headers items-----------------
	for i := 1; i <= 6; i++ {
		var htag string
		var headerIndex int
		headerIndex = 0
		s2 := strconv.Itoa(i)
		htag = "h" + s2

		doc.Find(htag).Each(func(i int, s *goquery.Selection) {
			headerIndex = headerIndex + 1
		})
		if headerIndex != 0 {
			headers := strconv.Itoa(headerIndex)
			row := []string{htag, headers}
			HtmlScraping.Headings = append(HtmlScraping.Headings, row)
		}
	}

	// find the links --------------------
	var externalURL int
	var externalURLvaild int
	var internalURL int
	externalURL = 0
	externalURLvaild = 0
	internalURL = 0
	doc.Find("a").Each(func(index int, selector *goquery.Selection) {
		link, _ := selector.Attr("href") // get the links
		var pattern = regexp.MustCompile(`(?:https?):\/\/(\w+:?\w*)?(\S+)(:\d+)?(\/|\/([\w#!:.?+=&%!\-\/]))?`)
		matched := pattern.MatchString(link)
		if matched == true { // external links
			u, err := url.ParseRequestURI(link) // check the validity
			if u != nil {
				externalURLvaild = externalURLvaild + 1
			} else {
				panic(err)
			}
			externalURL = externalURL + 1
		} else { // internal links
			internalURL = internalURL + 1
		}
		matched = false
	})
	if externalURL != 0 {
		externalURLCount := strconv.Itoa(externalURL)
		externalURLvaildCount := strconv.Itoa(externalURLvaild)
		row := []string{"External", externalURLCount, externalURLvaildCount}
		HtmlScraping.Links = append(HtmlScraping.Links, row)
	}
	if internalURL != 0 {
		internalURLCount := strconv.Itoa(internalURL)
		row := []string{"Internal", internalURLCount, internalURLCount}
		HtmlScraping.Links = append(HtmlScraping.Links, row)
	}

	//find the login from --------------
	doc.Find("input").Each(func(index int, selector *goquery.Selection) {
		link, _ := selector.Attr("type") // get the links
		if link == "password" {
			HtmlScraping.Loginform = true
		}
	})

	c.JSON(http.StatusOK, gin.H{
		"result": HtmlScraping,
	})

}
