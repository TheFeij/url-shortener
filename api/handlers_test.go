package api

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	apiio "url-shortener/api/io"
)

// TestShortenUrl tests shortenUrl handler function
func TestShortenUrl(t *testing.T) {
	// get the singleton server instance
	server := GetServer()

	testCases := []struct {
		name          string
		req           apiio.ShortenUrlRequest
		checkResponse func(t *testing.T, req apiio.ShortenUrlRequest, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			req:  apiio.ShortenUrlRequest{URL: "https://example.test"},
			checkResponse: func(t *testing.T, req apiio.ShortenUrlRequest, recorder *httptest.ResponseRecorder) {
				// check status code
				require.Equal(t, http.StatusOK, recorder.Code)

				// convert body from json to apiio.ShortenUrlResponse
				var response apiio.ShortenUrlResponse

				jsonBody, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				err = json.Unmarshal(jsonBody, &response)
				require.NoError(t, err)

				// check response body
				require.Equal(t, req.URL, response.OriginalUrl)
				require.NotEmpty(t, response.ShortUrl)
			},
		},
		{
			name: "BadRequest",
			req:  apiio.ShortenUrlRequest{URL: "example.test"}, // invalid url
			checkResponse: func(t *testing.T, req apiio.ShortenUrlRequest, recorder *httptest.ResponseRecorder) {
				// check status code
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// convert request body to json
			requestBody, err := json.Marshal(testCase.req)
			require.NoError(t, err)

			// create http request
			httpReq, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(requestBody))

			recorder := httptest.NewRecorder()
			require.NotEmpty(t, recorder)

			server.router.ServeHTTP(recorder, httpReq)

			testCase.checkResponse(t, testCase.req, recorder)
		})
	}
}
