package dao_service

import (
	"encoding/json"
	"fmt"
	"model"
	"strconv"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

type ResourceDaoService struct {
	baseUrl string
}

var ResourceDaoApi ResourceDaoService

func (this *ResourceDaoService) Init(ipAddr string) {
	this.baseUrl = ipAddr + "/v1/resource"
}

func (this *ResourceDaoService) Create(resource *model.Resource) (*model.Resource, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(resource)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return nil, err
	}

	url := this.baseUrl + "/"
	result, err = httpclient.Post(url, requestData)
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

	var newResource model.Resource
	err = json.Unmarshal(([]byte)(response.Result), &newResource)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &newResource, err
}

func (this *ResourceDaoService) Update(resource *model.Resource) (*model.Resource, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(resource)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return nil, err
	}

	url := this.baseUrl + "/"
	result, err = httpclient.Put(url, requestData)
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

	var newResource model.Resource
	err = json.Unmarshal(([]byte)(response.Result), &newResource)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &newResource, err
}

func (this *ResourceDaoService) DeleteById(id int64) error {
	var result []byte
	var err error

	url := this.baseUrl + "/id/" + strconv.FormatInt(id, 10)
	result, err = httpclient.Delete(url, nil)
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

func (this *ResourceDaoService) GetByUserId(userId int64) (*model.Resource, error) {
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

	var resource model.Resource
	err = json.Unmarshal(([]byte)(response.Result), &resource)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &resource, err
}
