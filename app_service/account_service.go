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
	// ACCOUNT_SERVICE = "http://192.168.0.206:9000/v1/account"
	ACCOUNT_SERVICE = "http://account-service.pme-system-beta1:8080/v1/account"
	//ACCOUNT_SERVICE = "account-service"
)

func ApiRegisterAccount(data []byte) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(ACCOUNT_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := ACCOUNT_SERVICE + "/register/"
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

func ApiDeleteAccounts(data []byte) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(ACCOUNT_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := ACCOUNT_SERVICE + "/"
	result, err = httpclient.Delete(url, data)
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

func ApiUpdateAccount(data []byte) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(ACCOUNT_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := ACCOUNT_SERVICE + "/"
	result, err = httpclient.Put(url, data)
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

func ApiGetAllAccount() ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(ACCOUNT_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := ACCOUNT_SERVICE + "/"
	result, err = httpclient.Get(url)
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
