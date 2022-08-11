package routers

import (
	"Go-web-Scraping/controllers"
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//var URL mod.UrlString
type URL struct {
	UrlStr string `json:"urlStr"`
}

func TestSetupRouters(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/result", controllers.GetResult)

	urls := []URL{
		{UrlStr: "https://www.google.com"},
		{UrlStr: "https://news.ycombinator.com/"},
		{UrlStr: "http://www.slbfe.lk/"},
		{UrlStr: "http://www.softwareqatest.com"},
	}

	for i := 0; i < len(urls); i++ {
		var request = `{"urlStr":"` + urls[i].UrlStr + `"}`
		jsonValue, _ := json.Marshal(request)
		body := bytes.NewBuffer([]byte(jsonValue))
		req, err := http.NewRequest("POST", "/result", body)
		if err != nil {
			t.Fatalf("Couldn't create request: %v\n", err)
		}
		w := httptest.NewRecorder()
		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

	}

}
