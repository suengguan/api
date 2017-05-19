package system_service

import (
	"encoding/json"
	"fmt"
	"model"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

const (
	//EMAIL_SERVICE = "http://192.168.0.206:7000/v1/email"
	//EMAIL_SERVICE = "http://email-service.pme-system-beta1:8080/v1/email"
	EMAIL_SERVICE = "http://192.168.0.21:8080/v1/email"
)

func ApiSendEmails(emails []*model.Email) error {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(emails)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return err
	}

	url := EMAIL_SERVICE + "/send/"
	result, err = httpclient.Post(url, requestData)
	if err != nil {
		return err
	}

	// check result
	var response model.Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return err
	}

	if response.Status != model.MSG_RESULTCODE_SUCCESS {
		err = fmt.Errorf("%s", response.Reason)
		return err
	}

	return err
}
