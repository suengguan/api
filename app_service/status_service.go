package app_service

import (
	"encoding/json"
	"fmt"
	"model"
	"strconv"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

type StatusService struct {
	baseUrl string
}

var StatusApi StatusService

func (this *StatusService) Init(ipAddr string) {
	this.baseUrl = ipAddr + "/v1/status"
}

func (this *StatusService) GetAllJobStatus(userId int64) ([]byte, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/job/" + strconv.FormatInt(userId, 10)
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
