package dao_service

import (
	"encoding/json"
	"fmt"
	"model"
	"strconv"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

type ActionDaoService struct {
	baseUrl string
}

var ActionDaoApi ActionDaoService

func (this *ActionDaoService) Init(ipAddr string) {
	this.baseUrl = ipAddr + "/v1/action"
}

func (this *ActionDaoService) Create(action *model.Action) error {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(action)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return err
	}

	url := this.baseUrl + "/"
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

func (this *ActionDaoService) GetAll(userId int64) ([]*model.Action, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/" + strconv.FormatInt(userId, 10)
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

	var actions []*model.Action
	err = json.Unmarshal(([]byte)(response.Result), &actions)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return actions, err
}
