package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/api/io"
)

// shortenUrl shortens URLs
func shortenUrl(context *gin.Context) {
	// receive request
	var req io.ShortenUrlRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
	}

	// TODO: shorten url

	// TODO: store in database

	// create response
	response := io.ShortenUrlResponse{
		OriginalUrl: req.URL,
		ShortUrl:    "short url",
	}

	// send response to client
	context.JSON(http.StatusOK, response)
}

// errResponse converts a error instance to a gin.H instance
func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
