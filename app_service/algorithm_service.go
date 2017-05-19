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
	// ALGORITHM_SERVICE = "http://192.168.0.206:9001/v1/algorithm"
	//ALGORITHM_SERVICE = "algorithm-service"
	ALGORITHM_SERVICE = "http://algorithm-service.pme-system-beta1:8080/v1/algorithm"
)

func ApiCreateAlgorithm(data []byte) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(ALGORITHM_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := ALGORITHM_SERVICE + "/"
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

func ApiDeleteAlgorithm(data []byte) error {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(ALGORITHM_SERVICE)
	// if err != nil {
	// 	return err
	// }

	url := ALGORITHM_SERVICE + "/"
	result, err = httpclient.Delete(url, data)
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

func ApiUpdateAlgorithm(data []byte) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(ALGORITHM_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := ALGORITHM_SERVICE + "/"
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

func ApiGetAllAlgorithm() ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(ALGORITHM_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := ALGORITHM_SERVICE + "/"
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
