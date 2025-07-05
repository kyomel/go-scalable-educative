package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/kyomel/go-scalable-educative/services"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	var results sync.Map
	var tasks sync.WaitGroup

	apis := map[string]func() (res *http.Response, err error){
		"users":    services.GetUsers,
		"products": services.GetProducts,
		"books":    services.GetBooks,
	}
	for name := range apis {
		tasks.Add(1)
		go func(resource string) {
			res, err := apis[resource]()
			if err != nil {
				results.Store(resource, nil)
			} else {
				defer res.Body.Close()
				var resData interface{}
				err = json.NewDecoder(res.Body).Decode(&resData)
				if err != nil {
					results.Store(resource, nil)
				} else {
					results.Store(resource, resData)
				}
			}
			tasks.Done()
		}(name)
	}
	tasks.Wait()
	data := map[string]interface{}{}
	for name := range apis {
		apiRes, _ := results.Load(name)
		if apiRes == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data[name] = apiRes
	}
	res, _ := json.Marshal(data)
	w.Write(res)
}
