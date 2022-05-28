package utils_test

import (
	"testing"

	"github.com/Ashilesh/load-balancer/utils"
)

type searchTestData struct {
	searchTestArr []int
	searchTestVal int
	searchTestResult
}

type searchTestResult struct {
	ind     int
	isFound bool
}

func getTestData() []searchTestData {
	data := []searchTestData{
		{
			searchTestArr: []int{1, 10, 20},
			searchTestVal: 5,
			searchTestResult: searchTestResult{
				ind:     0,
				isFound: false,
			},
		},
		{
			searchTestArr: []int{1, 10, 20, 23, 30},
			searchTestVal: 20,
			searchTestResult: searchTestResult{
				ind:     2,
				isFound: true,
			},
		},
		{
			searchTestArr: []int{1, 10, 20, 23, 30},
			searchTestVal: 50,
			searchTestResult: searchTestResult{
				ind:     4,
				isFound: false,
			},
		},
		{
			searchTestArr: []int{7, 10, 20, 23, 30},
			searchTestVal: 2,
			searchTestResult: searchTestResult{
				ind:     4,
				isFound: false,
			},
		}}

	return data
}

func TestSearch(t *testing.T) {
	testData := getTestData()

	for _, val := range testData {
		ind, isFound := utils.Search(val.searchTestArr, val.searchTestVal)
		if ind != val.ind {
			t.Errorf("expected to get %d but got %d", val.ind, ind)
		}
		if isFound != val.isFound {
			t.Errorf("expected value found to be %t but got %t", val.isFound, isFound)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	testData := getTestData()

	for _, val := range testData {
		ind, isFound := utils.BinarySearch(val.searchTestArr, val.searchTestVal)
		if ind != val.ind {
			t.Errorf("expected to get %d but got %d", val.ind, ind)
		}
		if isFound != val.isFound {
			t.Errorf("expected value found to be %t but got %t", val.isFound, isFound)
		}
	}
}
