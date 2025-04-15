package mock

import (
	"github.com/chutte2012de/web-page-analyzer/webpage/model"
	"github.com/stretchr/testify/mock"
)

type MockLinkManager struct {
	mock.Mock
}

func (_m *MockLinkManager) GetLinksInfo(inputUrlStr string, linksInPage []string) (model.LinksInfo, error) {
	ret := _m.Called(inputUrlStr, linksInPage)

	var r0 model.LinksInfo
	if rf, ok := ret.Get(0).(func(string, []string) model.LinksInfo); ok {
		r0 = rf(inputUrlStr, linksInPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.LinksInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(inputUrlStr, linksInPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockLinkManager) GetHostAndWebsiteBaseNames(inputUrlStr string) (string, string, error) {
	ret := _m.Called(inputUrlStr)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(inputUrlStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(inputUrlStr)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(inputUrlStr)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

func (_m *MockLinkManager) CategorizeLinks(websiteBaseName string, links []string) ([]string, []string, []string, error) {
	ret := _m.Called(websiteBaseName, links)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, []string) []string); ok {
		r0 = rf(websiteBaseName, links)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(string, []string) []string); ok {
		r1 = rf(websiteBaseName, links)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 []string
	if rf, ok := ret.Get(2).(func(string, []string) []string); ok {
		r2 = rf(websiteBaseName, links)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]string)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string, []string) error); ok {
		r3 = rf(websiteBaseName, links)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

func (_m *MockLinkManager) GetLinksSummaryStat(links []model.Link) (model.SummaryStat, error) {
	ret := _m.Called(links)

	var r0 model.SummaryStat
	if rf, ok := ret.Get(0).(func([]model.Link) model.SummaryStat); ok {
		r0 = rf(links)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.SummaryStat)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]model.Link) error); ok {
		r1 = rf(links)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockLinkManager) ValidateLinks(preffix string, links []string) ([]model.Link, error) {
	ret := _m.Called(preffix, links)

	var r0 []model.Link
	if rf, ok := ret.Get(0).(func(string, []string) []model.Link); ok {
		r0 = rf(preffix, links)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(preffix, links)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
