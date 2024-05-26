package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	mockdb "url-shortener/db/mock"
	"url-shortener/db/models"
	"url-shortener/db/service"
	"url-shortener/util"
)

// //////////////////////////////////////////////////////////////
// saveUrlRequestMatcher is a custom gomock matcher
type saveUrlRequestMatcher struct {
	originalUrl string
}

func (s saveUrlRequestMatcher) Matches(x any) bool {
	inputSession, ok := x.(*service.SaveUrlRequest)
	if !ok {
		return false
	}

	return s.originalUrl == inputSession.OriginalUrl()
}

func (s saveUrlRequestMatcher) String() string {
	return fmt.Sprintf("is equal to %v (%T)", s.originalUrl, s.originalUrl)
}

func newSaveUrlRequestMatcher(originalUrl string) gomock.Matcher {
	return saveUrlRequestMatcher{
		originalUrl: originalUrl,
	}
}

////////////////////////////////////////////////////////////////

// TestShortenUrl tests shortenUrl handler function
func TestShortenUrl(t *testing.T) {
	// assume we have this record in the database in the urls table
	urlRecord := models.Url{
		OriginalUrl: "http://example.test",
		ShortUrl:    util.GenerateShortUrl(),
	}

	testCases := []struct {
		name          string
		req           ShortenUrlRequest
		buildStubs    func(dbService *mockdb.MockDBService, req ShortenUrlRequest)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			req:  ShortenUrlRequest{URL: urlRecord.OriginalUrl},
			buildStubs: func(dbService *mockdb.MockDBService, req ShortenUrlRequest) {
				dbService.
					EXPECT().
					SaveUrl(newSaveUrlRequestMatcher(req.URL)).
					Return(service.NewSaveUrlResponse(urlRecord.OriginalUrl, urlRecord.ShortUrl), nil).
					Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check status code
				require.Equal(t, http.StatusOK, recorder.Code)

				// convert body from json to ShortenUrlResponse
				var response ShortenUrlResponse

				jsonBody, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				err = json.Unmarshal(jsonBody, &response)
				require.NoError(t, err)

				// check response body
				require.Equal(t, urlRecord.OriginalUrl, response.OriginalUrl)
				require.Equal(t, urlRecord.ShortUrl, response.ShortUrl)
			},
		},
		{
			name: "BadRequest",
			req:  ShortenUrlRequest{URL: "example.test"}, // invalid url
			buildStubs: func(dbService *mockdb.MockDBService, req ShortenUrlRequest) {
				dbService.EXPECT().SaveUrl(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check status code
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			req:  ShortenUrlRequest{URL: urlRecord.OriginalUrl},
			buildStubs: func(dbService *mockdb.MockDBService, req ShortenUrlRequest) {
				dbService.
					EXPECT().
					SaveUrl(newSaveUrlRequestMatcher(req.URL)).
					Return(nil, sql.ErrConnDone).
					Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check status code
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// get a new gomock controller
			controller := gomock.NewController(t)
			defer controller.Finish()

			// get an instance of the MockDBService (mock implementation of service.DBService interface)
			dbService := mockdb.NewMockDBService(controller)

			// build stubs
			testCase.buildStubs(dbService, testCase.req)

			// get a new test server
			server := newTestServer(dbService)

			// convert request body to json
			requestBody, err := json.Marshal(testCase.req)
			require.NoError(t, err)

			// create http request
			httpReq, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(requestBody))

			recorder := httptest.NewRecorder()
			require.NotEmpty(t, recorder)

			server.router.ServeHTTP(recorder, httpReq)

			testCase.checkResponse(t, recorder)
		})
	}
}
