package httpclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kosha/airtable-connector/pkg/models"
)

func tokenBasedAuth(token string) string {
	return token
}

func makeHttpReq(token string, req *http.Request) []byte {
	req.Header.Add("Authorization", "Bearer "+tokenBasedAuth(token))
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}

func ListRecords(url string, apiKey string) []*models.Records {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var records []*models.Records

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &records)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return records
}

func GetRecords(url string, apiKey string) *models.Records {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var records *models.Records

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &records)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return records
}

func ListComments(url string, apiKey string) *models.CommentStruct {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var comments *models.CommentStruct

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &comments)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return comments
}

func GetBases(url string, apiKey string) *models.BasesStruct {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var bases *models.BasesStruct

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &bases)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return bases
}