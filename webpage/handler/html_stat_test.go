package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chutte2012de/web-page-analyzer/webpage/model"
	mock "github.com/chutte2012de/web-page-analyzer/webpage/usecase/mock"
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
		Id:  uint64(9988321),
		Url: "https://riyasewana.com/login.php",
	}
	htmlElement.On("ExtractFromUrl", "https://riyasewana.com/login.php").Return(resHtmlStat, nil)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(httpMethod, routeUrl, bytes.NewBuffer([]byte(requestBody)))

	htmlStatHandler.Create(recorder, request)
	fmt.Println("StatusCode: ", recorder.Result().StatusCode)
	fmt.Println("Body: ", recorder.Result().Body)

	var actual model.HtmlStat
	json.Unmarshal(recorder.Body.Bytes(), &actual)
	fmt.Println("actual: ", actual)
	fmt.Println("Id: ", actual.Id)
}
