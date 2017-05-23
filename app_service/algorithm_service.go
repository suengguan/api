package app_service

import (
	"encoding/json"
	"fmt"
	"model"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

type AlgorithmService struct {
	baseUrl string
}

var AlgorithmApi AlgorithmService

func (this *AlgorithmService) Init(ipAddr string) {
	this.baseUrl = ipAddr + "/v1/algorithm"
}

func (this *AlgorithmService) Create(data []byte) ([]byte, error) {
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

func (this *AlgorithmService) Delete(data []byte) error {
	var result []byte
	var err error

	url := this.baseUrl + "/"
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

func (this *AlgorithmService) Update(data []byte) ([]byte, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/"
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

func (this *AlgorithmService) GetAll() ([]byte, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/"
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
