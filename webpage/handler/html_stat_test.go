package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chutte2012de/web-page-analyzer/webpage/model"
	mock "github.com/chutte2012de/web-page-analyzer/webpage/usecase/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateHtmlStat_Success(t *testing.T) {
	htmlElement := new(mock.MockHtmlElement)
	httpMethod := http.MethodPost
	routeUrl := "/webpage/htmlstat"
	requestBody := `{"url": "https://riyasewana.com/login.php"}`

	htmlStatHandler := HtmlStat{
		HtmlElement: htmlElement,
	}

	resHtmlStat := model.HtmlStat{
		Id:    uint64(9988321),
		Url:   "https://riyasewana.com/login.php",
		Title: "Title of the Web Page 123",
	}
	htmlElement.On("ExtractFromUrl", "https://riyasewana.com/login.php").Return(resHtmlStat, nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(httpMethod, routeUrl, bytes.NewBuffer([]byte(requestBody)))

	htmlStatHandler.Create(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var actual model.HtmlStat
	json.Unmarshal(recorder.Body.Bytes(), &actual)
	assert.Equal(t, uint64(9988321), actual.Id)
	assert.Equal(t, "https://riyasewana.com/login.php", actual.Url)
	assert.Equal(t, "Title of the Web Page 123", actual.Title)
}
