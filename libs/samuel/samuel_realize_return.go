package samuel

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sectask/domain/samuel"
	"strconv"
	"time"
)

var DefaultTimeout = time.Duration(120)

func (cred *SamuelConfig) GetRealizeGainLoss(clientCode, startDate, endDate string) (res samuel.ResponseRealizeGainLoss, err error) {
	fullURL := cred.BaseURL2 + `/api/sr/v1/balpos/realizegl`

	u, err := url.Parse(fullURL)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Samuel_GetRealizeGainLoss %s %s\n", fullURL, err.Error())
		return res, err
	}
	q := u.Query()
	q.Set("clientcode", clientCode)
	q.Set("fromdate", startDate)
	q.Set("todate", endDate)
	u.RawQuery = q.Encode()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	fmt.Printf("Samuel_GetRealizeGainLoss %s\n", u.String())

	client := &http.Client{Transport: tr, Timeout: DefaultTimeout * time.Second}
	r, _ := http.NewRequest("GET", u.String(), nil)
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	status, _ := strconv.Atoi(resp.Status)
	if status >= 400 {
		fmt.Printf("Samuel_GetRealizeGainLoss Error %s %s\n", fullURL, resp)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	fmt.Printf("Samuel_GetRealizeGainLoss %s %s\n", fullURL, string([]byte(body)))

	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}
