package middlewares

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

var totalCount = 0

func makeCall(count int, wg *sync.WaitGroup) {
	res, err := http.Get("http://localhost:8080/users")
	if err != nil {
		fmt.Printf("error making API call %e", err)
		wg.Done()
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Failed call number %d with status code %d\n", count, res.StatusCode)
		wg.Done()
		return
	}
	fmt.Printf("Succeeded call number %d\n", count)
	totalCount++
	wg.Done()
}

func TestRateLimit(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go makeCall(i, &wg)
	}

	wg.Wait()
	if totalCount > MAX_REQUESTS {
		t.Fatalf("More requests executed than max allowed - %d", totalCount)
	}

	fmt.Printf("Total success count - %d", totalCount)
}
