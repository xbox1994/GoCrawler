package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	body, _ := ioutil.ReadFile("citylist_test_data.html")

	list := ParseCityList(body)

	const resultSize = 470
	if len(list.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(list.Requests))
	}
}
