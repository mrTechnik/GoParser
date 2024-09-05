package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MyRequest(cli http.Client, path string) (string, error) {

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0")
	req.Header.Add("accept", "*/*")

	res, err := cli.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Status)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	body_string := string(body)
	return body_string, nil
}
