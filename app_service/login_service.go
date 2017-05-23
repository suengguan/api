package app_service

import (
	"encoding/json"
	"fmt"
	"model"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

type LoginService struct {
	baseUrl string
}

var LoginApi LoginService

func (this *LoginService) Init(ipAddr string) {
	this.baseUrl = ipAddr + "/v1/login"
}

func (this *LoginService) Login(data []byte) ([]byte, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/"
	result, err = httpclient.Post(url, data)
	if err != nil {
		return nil, err
	}

	// check result
	var response model.Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	if response.Status != model.MSG_RESULTCODE_SUCCESS {
		err = fmt.Errorf("%s", response.Reason)
		return nil, err
	}

	return ([]byte)(response.Result), err
}
