package yavoice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Voice struct {
	Result       string `json:"result"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func Voice2Text(file string, token string) (string, error) {

	fstr, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fstr.Close()
	req, err := http.NewRequest("POST", "https://stt.api.cloud.yandex.net/speech/v1/stt:recognize?topic=general", fstr)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Api-Key "+token)
	req.Header.Set("Transfer-Encoding", "chunked")
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	fmt.Println(responseString)
	var ctask Voice

	jsonErr := json.Unmarshal(responseData, &ctask)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return "", fmt.Errorf("error parce json")
	}
	if ctask.Result != "" {
		return ctask.Result, nil
	} else {
		return "", fmt.Errorf(ctask.ErrorCode + " " + ctask.ErrorMessage)
	}

}
