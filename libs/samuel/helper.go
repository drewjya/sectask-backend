package samuel

import (
	"encoding/json"
	"errors"
	"os"
	errs "sectask/libs/errors"
	"sectask/libs/httpresponse"
	"strings"
)

func (s *SamuelConfig) HandleExpiredSession(body string) httpresponse.ErrMessage {
	body = strings.ToLower(body)
	if strings.Contains(body, Expired) {
		return httpresponse.ErrMessage{
			ErrMapping: errors.New(InvalidToken),
			Message:    errs.PleaseReloginSamuelMSG,
		}
	} else if strings.Contains(body, PleaseInputPin) {
		return httpresponse.ErrMessage{
			ErrMapping: errors.New(PinRequired),
			Message:    errs.PleaseReinputPINMSG,
		}
	}
	return httpresponse.ErrMessage{}
}

func (s *SamuelConfig) HandleMaintenanceSamuel(body string) httpresponse.ErrMessage {
	body = strings.ToLower(body)
	if os.Getenv("PRO_API_APP_ENV") == "dev" {
		return httpresponse.ErrMessage{}
	}
	if strings.Contains(body, PleaseCheckCS) || strings.Contains(body, ProcessingBackOfficeData) || strings.Contains(body, NoData) {
		return httpresponse.ErrMessage{
			ErrMapping: errors.New(BackOfficeMaintenance),
			Message:    errs.MaintenanceSamuel,
		}
	}
	return httpresponse.ErrMessage{}
}

func (s *SamuelConfig) ConvertSamuelData(samuelData interface{}, data interface{}) {
	byteData, _ := json.Marshal(samuelData)
	json.Unmarshal(byteData, &data)
}

func (s *SamuelConfig) ConvertSamuelDataToArray(samuelData interface{}, arrData interface{}) {
	var result []interface{}
	switch sd := samuelData.(type) {
	case []interface{}:
		byteData, _ := json.Marshal(sd)
		json.Unmarshal(byteData, &result)
	case map[string]interface{}:
		var data interface{}
		byteData, _ := json.Marshal(sd)
		json.Unmarshal(byteData, &data)
		result = append(result, data)
	}

	byteData, _ := json.Marshal(result)
	json.Unmarshal(byteData, &arrData)
}
