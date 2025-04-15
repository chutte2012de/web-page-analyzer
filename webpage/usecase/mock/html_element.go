package mock

import (
	"io"

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

func (_m *MockHtmlElement) ExtractFromIoReader(r io.Reader) (model.HtmlStat, []string, error) {
	ret := _m.Called(r)

	var r0 model.HtmlStat
	if rf, ok := ret.Get(0).(func(io.Reader) model.HtmlStat); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.HtmlStat)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(io.Reader) []string); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(io.Reader) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
