package samuel

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sectask/libs/httpresponse"
	"strconv"
	"time"

	"sectask/domain/samuel"
)

func (s *SamuelConfig) GetSamuelStockPositions(clientCode string) (samuel.SamuelStockPositionData, httpresponse.ErrMessage) {
	var result samuel.SamuelResponse
	var stockPositionData samuel.SamuelStockPositionData
	var listStockPositionB []samuel.SamuelStockPositionB

	fullURL := s.BaseURL + URLStockPosition

	payload := map[string]interface{}{
		"clientCode": clientCode,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return stockPositionData, httpresponse.ErrMessage{
			ErrMapping: errors.New(SamuelError),
			Message:    err.Error(),
		}
	}
	pBody := []byte(string(b))
	fmt.Printf("Samuel_GetSamuelStockPositions body %s %s\n", fullURL, string(pBody))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: DefaultTimeout * time.Second}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("token", os.Getenv("CA_SAMUEL_TOKEN"))
	r.Header.Add("access-token", s.SamuelAccessToken)
	resp, err := client.Do(r)
	if err != nil {
		return stockPositionData, httpresponse.ErrMessage{
			ErrMapping: errors.New(SamuelError),
			Message:    err.Error(),
		}
	}

	log.Println("token", os.Getenv("CA_SAMUEL_TOKEN"), s.SamuelAccessToken)

	status, _ := strconv.Atoi(resp.Status)
	if status >= 400 {
		fmt.Printf("Samuel_GetSamuelStockPositions Error %s %s %s\n", clientCode, fullURL, resp)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return stockPositionData, httpresponse.ErrMessage{
			ErrMapping: errors.New(SamuelError),
			Message:    err.Error(),
		}
	}

	// fmt.Printf("Samuel_GetSamuelStockPositions %s %s %s\n", clientCode, fullURL, string([]byte(body)))

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return stockPositionData, httpresponse.ErrMessage{
			ErrMapping: errors.New(SamuelError),
			Message:    err.Error(),
		}
	}

	if errSession := s.HandleExpiredSession(string([]byte(body))); errSession.ErrMapping != nil {
		return stockPositionData, errSession
	} else if !result.Ok {
		return stockPositionData, httpresponse.ErrMessage{
			ErrMapping: errors.New(SamuelError),
			Message:    strconv.FormatBool(result.Ok),
		}
	}

	s.ConvertSamuelData(result.Data, &stockPositionData)
	s.ConvertSamuelDataToArray(stockPositionData.A.B, &listStockPositionB)
	stockPositionData.A.ListB = listStockPositionB

	return stockPositionData, httpresponse.ErrMessage{}
}

// func (s *SamuelConfig) GetSamuelStockPosition(clientCode string) (samuel.SamuelStockPositionData, httpresponse.ErrMessage) {
// 	var result samuel.SamuelResponse
// 	var stockPositionData samuel.SamuelStockPositionData
// 	var listStockPositionB []samuel.SamuelStockPositionB

// 	body := map[string]interface{}{
// 		"clientCode": clientCode,
// 	}

// 	headers := map[string]string{
// 		"token":        os.Getenv("CA_SAMUEL_TOKEN"),
// 		"access-token": s.SamuelAccessToken,
// 	}
// 	log.Println(s.SamuelClient.BaseURL)

// 	resp, err := s.SamuelClient.R().
// 		SetHeaders(headers).
// 		SetBody(body).
// 		SetResult(&result).
// 		Post(URLStockPosition)
// 	if err != nil {
// 		return stockPositionData, httpresponse.ErrMessage{
// 			ErrMapping: errors.New(SamuelError),
// 			Message:    err.Error(),
// 		}
// 	}
// 	log.Println(URLStockPosition+" HEADER = ", headers)
// 	log.Println(URLStockPosition+" BODY = ", body)
// 	log.Println(URLStockPosition+" RESP = ", resp.String())

// 	if errSession := s.HandleExpiredSession(resp.String()); errSession.ErrMapping != nil {
// 		return stockPositionData, errSession
// 	} else if !result.Ok {
// 		return stockPositionData, httpresponse.ErrMessage{
// 			ErrMapping: errors.New(SamuelError),
// 			Message:    strconv.FormatBool(result.Ok),
// 		}
// 	}

// 	s.ConvertSamuelData(result.Data, &stockPositionData)
// 	s.ConvertSamuelDataToArray(stockPositionData.A.B, &listStockPositionB)
// 	stockPositionData.A.ListB = listStockPositionB
// 	return stockPositionData, httpresponse.ErrMessage{}
// }
