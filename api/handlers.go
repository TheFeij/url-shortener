package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// redirectShortUrl redirects client to the original url of the given short url
func (s *server) redirectShortUrl(context *gin.Context) {
	var req RedirectShortUrl

	// get the request uri
	if err := context.ShouldBindUri(&req); err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	// retrieve the original url of the short url
	res, err := s.dbService.GetOriginalUrl(service.NewGetOriginalUrlRequest(req.ShortUrl))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, "url does not exist")
			return
		}

		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	// redirect the client to the original url
	context.Redirect(http.StatusPermanentRedirect, res.OriginalUrl())
}

// errResponse converts a error instance to a gin.H instance
func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
