package parser

import (
	"go_dev/day9/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents,error := fetcher.Fetch("https//www.zhenai.com/zhenhun")
	if error != nil{
		panic(error)
	}
	expectedUrls := []string{
		"","","",
	}

	expectedCities := []string{
		"","","",
	}






	result := ParseCityList(contents)
	const resultSzie = 470
	if len(result.Requests) != resultSzie{
		t.Errorf("result should have %d requests",len(result.Requests))

	}

	if len(result.Items) != resultSzie{
		t.Errorf("BAD")
	}

	for i,Url := range expectedUrls{
		if result.Requests[i].Url != Url{
			t.Errorf("bad")
		}
	}

	for i,Url := range expectedCities{
		if result.Items[i].(string) != Url{
			t.Errorf("bad")
		}
	}








}



