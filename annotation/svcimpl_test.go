package service

import (
	"net/http"
	"testing"
)

func BenchmarkHttpRouter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		request, err := http.NewRequest("GET", "http://localhost:6060/guest", nil)
		if err != nil {
			panic(err)
		}
		_, err = http.DefaultClient.Do(request)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGorillaMux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		request, err := http.NewRequest("GET", "http://localhost:6061/guest", nil)
		if err != nil {
			panic(err)
		}
		_, err = http.DefaultClient.Do(request)
		if err != nil {
			panic(err)
		}
	}
}

//func Test1(t *testing.T) {
//	request, err := http.NewRequest("GET", "http://localhost:6060/guest", nil)
//	if err != nil {
//		panic(err)
//	}
//	resp, err := http.DefaultClient.Do(request)
//	if err != nil {
//		panic(err)
//	}
//	result := make(map[string]interface{})
//	_ = json.NewDecoder(resp.Body).Decode(&result)
//	fmt.Println(result)
//}
