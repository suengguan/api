package app_service

import (
	"encoding/json"
	"fmt"
	"model"
	"strconv"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

type LogService struct {
	baseUrl string
}

var LogApi LogService

func (this *LogService) Init(ipAddr string) {
	this.baseUrl = ipAddr + "/v1"
}

func (this *LogService) WriteAction(data []byte) error {
	var result []byte
	var err error

	url := this.baseUrl + "/action/"
	result, err = httpclient.Post(url, data)
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

func (this *LogService) GetAllActions(userId int64) ([]byte, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/action/" + strconv.FormatInt(userId, 10)
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

func (this *LogService) GetPodLog(podId int64) ([]byte, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/log/pod/" + strconv.FormatInt(podId, 10)
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
