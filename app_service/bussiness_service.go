package app_service

import (
	//"config"
	"encoding/json"
	"fmt"
	"model"
	"strconv"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

const (
	//BUSSINESS_SERVICE = "http://192.168.0.206:9002/v1"
	//BUSSINESS_SERVICE = "bussiness-service"
	BUSSINESS_SERVICE = "http://bussiness-service.pme-system-beta1:8080/v1"
)

func ApiCreateProject(data []byte) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(BUSSINESS_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := BUSSINESS_SERVICE + "/project/"
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

func ApiCreateJob(data []byte) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(BUSSINESS_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := BUSSINESS_SERVICE + "/job/"
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

func ApiUpdateJob(data []byte) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(BUSSINESS_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := BUSSINESS_SERVICE + "/job/"
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

func ApiGetCurrentPods(userId int64) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(BUSSINESS_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := BUSSINESS_SERVICE + "/pod/current/" + strconv.FormatInt(userId, 10)
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
