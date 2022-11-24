package basalt

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getPackages(name string) (resp Response) {
	http, err := http.Get("https://rdb.altlinux.org/api/export/branch_binary_packages/" + name)
	if err != nil {
		log.Fatal(err)
	}

	defer http.Body.Close()
	body, err := io.ReadAll(http.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &resp)

	return resp
}
