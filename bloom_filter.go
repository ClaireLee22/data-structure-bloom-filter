package main

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	diffSizefilters := []*bloom.BloomFilter{bloom.New(10, 3), bloom.New(20, 3), bloom.New(30, 3)}
	diffKhashFunctionsfilters := []*bloom.BloomFilter{bloom.New(20, 2), bloom.New(20, 3), bloom.New(20, 4)}
	urls := createUrls(1, 10)
	notExistedUrls := createUrls(11, 10)

	// different filter size
	compareFilters(diffSizefilters, urls, notExistedUrls)
	// different # of hash functions
	compareFilters(diffKhashFunctionsfilters, urls, notExistedUrls)

}

func compareFilters(filters []*bloom.BloomFilter, urls []string, notExistedUrls []string) {
	for _, filter := range filters {
		fmt.Printf("filter size: %d; # of hash functions: %d\n", filter.Cap(), filter.K())
		insertElements(filter, urls)

		// check false positives
		fmt.Println("------------------------------------------")
		fmt.Printf("| %25s | %10s |\n", "test url", "is exist")
		for _, url := range notExistedUrls {
			fmt.Println("------------------------------------------")
			isExisted(filter, url)
		}
		fmt.Printf("------------------------------------------\n\n")
	}
}

func createUrls(start int, numOfUrls int) []string {
	urls := []string{}
	for i := start; i < start+numOfUrls; i++ {
		urls = append(urls, fmt.Sprintf("https://testurl%d.com", i))
	}
	return urls
}

func insertElements(filter *bloom.BloomFilter, urls []string) {
	for _, url := range urls {
		filter.Add([]byte(url))
	}
}

func isExisted(filter *bloom.BloomFilter, url string) {
	if filter.Test([]byte(url)) {
		fmt.Printf("| %25s | %10s |\n", url, "Yes")
	} else {
		fmt.Printf("| %25s | %10s |\n", url, "No")
	}
}

/* output
filter size: 10; # of hash functions: 3
------------------------------------------
|                  test url |   is exist |
------------------------------------------
|     https://testurl11.com |        Yes |
------------------------------------------
|     https://testurl12.com |        Yes |
------------------------------------------
|     https://testurl13.com |        Yes |
------------------------------------------
|     https://testurl14.com |        Yes |
------------------------------------------
|     https://testurl15.com |        Yes |
------------------------------------------
|     https://testurl16.com |        Yes |
------------------------------------------
|     https://testurl17.com |        Yes |
------------------------------------------
|     https://testurl18.com |        Yes |
------------------------------------------
|     https://testurl19.com |        Yes |
------------------------------------------
|     https://testurl20.com |        Yes |
------------------------------------------

filter size: 20; # of hash functions: 3
------------------------------------------
|                  test url |   is exist |
------------------------------------------
|     https://testurl11.com |        Yes |
------------------------------------------
|     https://testurl12.com |         No |
------------------------------------------
|     https://testurl13.com |         No |
------------------------------------------
|     https://testurl14.com |         No |
------------------------------------------
|     https://testurl15.com |        Yes |
------------------------------------------
|     https://testurl16.com |         No |
------------------------------------------
|     https://testurl17.com |        Yes |
------------------------------------------
|     https://testurl18.com |         No |
------------------------------------------
|     https://testurl19.com |         No |
------------------------------------------
|     https://testurl20.com |        Yes |
------------------------------------------

filter size: 30; # of hash functions: 3
------------------------------------------
|                  test url |   is exist |
------------------------------------------
|     https://testurl11.com |         No |
------------------------------------------
|     https://testurl12.com |        Yes |
------------------------------------------
|     https://testurl13.com |         No |
------------------------------------------
|     https://testurl14.com |         No |
------------------------------------------
|     https://testurl15.com |        Yes |
------------------------------------------
|     https://testurl16.com |         No |
------------------------------------------
|     https://testurl17.com |         No |
------------------------------------------
|     https://testurl18.com |         No |
------------------------------------------
|     https://testurl19.com |         No |
------------------------------------------
|     https://testurl20.com |        Yes |
------------------------------------------

filter size: 20; # of hash functions: 2
------------------------------------------
|                  test url |   is exist |
------------------------------------------
|     https://testurl11.com |         No |
------------------------------------------
|     https://testurl12.com |         No |
------------------------------------------
|     https://testurl13.com |         No |
------------------------------------------
|     https://testurl14.com |         No |
------------------------------------------
|     https://testurl15.com |        Yes |
------------------------------------------
|     https://testurl16.com |        Yes |
------------------------------------------
|     https://testurl17.com |         No |
------------------------------------------
|     https://testurl18.com |         No |
------------------------------------------
|     https://testurl19.com |         No |
------------------------------------------
|     https://testurl20.com |        Yes |
------------------------------------------

filter size: 20; # of hash functions: 3
------------------------------------------
|                  test url |   is exist |
------------------------------------------
|     https://testurl11.com |        Yes |
------------------------------------------
|     https://testurl12.com |         No |
------------------------------------------
|     https://testurl13.com |         No |
------------------------------------------
|     https://testurl14.com |         No |
------------------------------------------
|     https://testurl15.com |        Yes |
------------------------------------------
|     https://testurl16.com |         No |
------------------------------------------
|     https://testurl17.com |        Yes |
------------------------------------------
|     https://testurl18.com |         No |
------------------------------------------
|     https://testurl19.com |         No |
------------------------------------------
|     https://testurl20.com |        Yes |
------------------------------------------

filter size: 20; # of hash functions: 4
------------------------------------------
|                  test url |   is exist |
------------------------------------------
|     https://testurl11.com |        Yes |
------------------------------------------
|     https://testurl12.com |         No |
------------------------------------------
|     https://testurl13.com |         No |
------------------------------------------
|     https://testurl14.com |        Yes |
------------------------------------------
|     https://testurl15.com |        Yes |
------------------------------------------
|     https://testurl16.com |         No |
------------------------------------------
|     https://testurl17.com |        Yes |
------------------------------------------
|     https://testurl18.com |         No |
------------------------------------------
|     https://testurl19.com |         No |
------------------------------------------
|     https://testurl20.com |        Yes |
------------------------------------------
*/
