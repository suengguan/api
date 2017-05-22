package dao_service

import (
	"encoding/json"
	"fmt"
	"model"
	"strconv"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

type AlgorithmDaoService struct {
	baseUrl string
}

var AlgorithmDaoApi AlgorithmDaoService

func (this *AlgorithmDaoService) Init(ipAddr string) {
	this.baseUrl = ipAddr + "/v1/algorithm"
}

func (this *AlgorithmDaoService) GetAll() ([]*model.Algorithm, error) {
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

	var algorithms []*model.Algorithm
	err = json.Unmarshal(([]byte)(response.Result), &algorithms)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return algorithms, err
}

func (this *AlgorithmDaoService) GetById(id int64) (*model.Algorithm, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/id/" + strconv.FormatInt(id, 10)
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

	var a model.Algorithm
	err = json.Unmarshal(([]byte)(response.Result), &a)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &a, err
}

func (this *AlgorithmDaoService) GetByNameAndVersion(name string, version string) (*model.Algorithm, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/" + name + "/" + version
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

	var a model.Algorithm
	err = json.Unmarshal(([]byte)(response.Result), &a)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &a, err
}

func (this *AlgorithmDaoService) Create(algorithm *model.Algorithm) (*model.Algorithm, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(algorithm)
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

	var a model.Algorithm
	err = json.Unmarshal(([]byte)(response.Result), &a)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &a, err
}

func (this *AlgorithmDaoService) Update(algorithm *model.Algorithm) (*model.Algorithm, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(algorithm)
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

	var a model.Algorithm
	err = json.Unmarshal(([]byte)(response.Result), &a)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &a, err
}

func (this *AlgorithmDaoService) DeleteById(id int64) error {
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
