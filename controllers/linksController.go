package controllers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

type TlinkStatus struct {
	Link   string
	Error  bool
	Status whoisparser.WhoisInfo
}

func ProcessLinks(context *gin.Context) {
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var request struct {
		Links []string
	}

	var results []TlinkStatus

	var wg sync.WaitGroup

	err := context.Bind(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	for _, value := range request.Links {
		wg.Add(1)
		go func(value string) {
			defer wg.Done()
			r := get_link_status(value)
			results = append(results, r)
		}(value)
	}

	wg.Wait()

	context.JSON(http.StatusOK, gin.H{
		"links": results,
		"count": len(results),
	})
}

func get_link_status(link string) TlinkStatus {
	var response TlinkStatus
	var resultParseError bool

	result, err := whois.Whois(link)

	if err != nil {
		domain := whoisparser.WhoisInfo{}
		response = TlinkStatus{
			Link:   link,
			Error:  true,
			Status: domain,
		}

		return response
	}

	parsedResult, err := whoisparser.Parse(result)

	if err != nil {
		resultParseError = true
	} else {
		resultParseError = false
	}

	response = TlinkStatus{
		Link:   link,
		Error:  resultParseError,
		Status: parsedResult,
	}

	return response
}
