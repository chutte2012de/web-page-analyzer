package mock

import (
	"github.com/chutte2012de/web-page-analyzer/webpage/model"
	"github.com/stretchr/testify/mock"
)

type MockHtmlElement struct {
	mock.Mock
}

func (_m *MockHtmlElement) ExtractFromUrl(url string) (model.HtmlStat, error) {
	ret := _m.Called(url)

	var r0 model.HtmlStat
	if rf, ok := ret.Get(0).(func(string) model.HtmlStat); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.HtmlStat)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
