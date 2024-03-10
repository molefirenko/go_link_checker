package controllers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type link_status struct {
	Link         string
	Error        bool
	StatusCode   int
	StatusString string
}

func ProcessLinks(context *gin.Context) {
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var request struct {
		Links []string
	}

	var results []link_status

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

func get_link_status(link string) link_status {
	var response link_status

	resp, err := http.Get(link)

	if err != nil {
		response = link_status{
			Link:         link,
			Error:        true,
			StatusCode:   0,
			StatusString: err.Error(),
		}

		return response
	}

	response = link_status{
		Link:         link,
		Error:        false,
		StatusCode:   resp.StatusCode,
		StatusString: http.StatusText(resp.StatusCode),
	}

	return response
}
