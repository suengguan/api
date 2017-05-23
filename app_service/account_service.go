package app_service

import (
	"encoding/json"
	"fmt"
	"model"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

type AccountService struct {
	baseUrl string
}

var AccountApi AccountService

func (this *AccountService) Init(ipAddr string) {
	this.baseUrl = ipAddr + "/v1/account"
}

func (this *AccountService) Register(data []byte) ([]byte, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/register/"
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

// delete accounts
func (this *AccountService) Delete(data []byte) ([]byte, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/"
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

func (this *AccountService) Update(data []byte) ([]byte, error) {
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

func (this *AccountService) GetAll() ([]byte, error) {
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
