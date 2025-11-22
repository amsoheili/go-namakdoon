package main

import (
	"encoding/json"
	"net/http"
)

type Quarter struct {
	a Side
	b Side
}

type Side struct {
	imageUrl string
}

type NamakdoonRequestModel struct {
	Quarter int `json:"quarter"`
	Count   int `json:"count"`
}

func main() {
	http.HandleFunc("/", namakdoonHandler)
	http.ListenAndServe(":8080", nil)

}

func namakdoonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var bodyMarshaled NamakdoonRequestModel
	if err := json.NewDecoder(r.Body).Decode(&bodyMarshaled); err != nil {
		http.Error(w, "body is not specified correctly", http.StatusBadRequest)
		return
	}
	http.ServeFile(w, r, getImageUrl(bodyMarshaled.Quarter, bodyMarshaled.Count))
}

func getImageUrl(quarter int, count int) string {
	quarter1 := Quarter{
		a: Side{"camel"},
		b: Side{"cat"},
	}
	quarter2 := Quarter{
		a: Side{"dog"},
		b: Side{"lion"},
	}
	quarter3 := Quarter{
		a: Side{"monkey"},
		b: Side{"owl"},
	}
	quarter4 := Quarter{
		a: Side{"shark"},
		b: Side{"snake"},
	}
	dataMap := map[int]Quarter{
		1: quarter1,
		2: quarter2,
		3: quarter3,
		4: quarter4,
	}
	var resultImageUrl string
	if count%2 == 0 {
		resultImageUrl = dataMap[quarter].a.imageUrl
	} else {
		resultImageUrl = dataMap[quarter].b.imageUrl
	}

	return `./assets/` + resultImageUrl + `.jpeg`
}
