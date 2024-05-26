package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/db/service"
	"url-shortener/util"
)

// shortenUrl shortens URLs
func (s *server) shortenUrl(context *gin.Context) {
	// receive request
	var req ShortenUrlRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	// generate a short url
	shortUrl := util.GenerateShortUrl()

	// service record in the database
	res, err := s.dbService.SaveUrl(service.NewSaveUrlRequest(req.URL, shortUrl))
	if err != nil {
		context.JSON(http.StatusInternalServerError, errResponse(fmt.Errorf("something went wrong. try again")))
		return
	}

	// create response
	response := ShortenUrlResponse{
		OriginalUrl: res.OriginalUrl(),
		ShortUrl:    res.ShortUrl(),
	}

	// send response to the client
	context.JSON(http.StatusOK, response)
}

// errResponse converts a error instance to a gin.H instance
func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
