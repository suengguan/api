package app_service

import (
	//"config"
	"encoding/json"
	"fmt"
	"model"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

const (
	//LOGIN_SERVICE = "http://192.168.0.206:9005/v1/login"
	//LOGIN_SERVICE = "login-service"
	LOGIN_SERVICE = "http://login-service.pme-system-beta1:8080/v1/login"
)

func ApiLogin(data []byte) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(LOGIN_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := LOGIN_SERVICE + "/"
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
