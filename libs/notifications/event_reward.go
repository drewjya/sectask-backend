package notifications

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type RegistTradingNotifPayload struct {
	UserID   string
	Username string
	Email    string
}

func NotifDoneRegistTradingCompetition(data RegistTradingNotifPayload) error {
	url := os.Getenv("NOTIF_URL") + "/v1/other/done-regist-trad-comp"
	topic := "NOTIF_API_DONE_REGIST"
	fmt.Println(topic + "_Request : " + url)

	payload := map[string]interface{}{
		"user_id":  data.UserID,
		"username": data.Username,
		"email":    data.Email,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(topic + "_Err_Marshal : " + err.Error())
		return err
	}
	fmt.Println(topic + "_Payload : " + string(b))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 60 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(topic + "_Err_New_Request : " + err.Error())
		return err
	}
	basic := base64.StdEncoding.EncodeToString([]byte(os.Getenv("NOTIF_AUTH_USERNAME") + ":" + os.Getenv("NOTIF_AUTH_PASSWORD")))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basic)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(topic + "_Err_Do_Request : " + err.Error())
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(topic + "_Err_Do_Request : " + err.Error())
		return err
	}
	fmt.Println(topic + "_Body : " + string(body))
	if res.StatusCode != 200 {
		fmt.Println(topic + "_Code : " + fmt.Sprintf("%v", res.StatusCode))
		return errors.New("invalid status code")
	}
	return nil
}
