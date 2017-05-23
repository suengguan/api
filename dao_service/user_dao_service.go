package dao_service

import (
	"encoding/json"
	"fmt"
	"model"
	"strconv"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

type UserDaoService struct {
	baseUrl string
}

var UserDaoApi UserDaoService

func (this *UserDaoService) Init(ipAddr string) {
	this.baseUrl = ipAddr + "/v1/user"
	beego.Debug(this.baseUrl)
}

func (this *UserDaoService) GetAll() ([]*model.User, error) {
	return nil, nil
}

func (this *UserDaoService) Create(user *model.User) (*model.User, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(user)
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

	var u model.User
	err = json.Unmarshal(([]byte)(response.Result), &u)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &u, err
}

func (this *UserDaoService) Update(user *model.User) (*model.User, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(user)
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

	var u model.User
	err = json.Unmarshal(([]byte)(response.Result), &u)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &u, err
}

func (this *UserDaoService) GetAllExcludeOneById(id int64) ([]*model.User, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/exclude/id/" + strconv.FormatInt(id, 10)
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

	var users []*model.User
	err = json.Unmarshal(([]byte)(response.Result), &users)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return users, err
}

func (this *UserDaoService) DeleteById(id int64) error {
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

func (this *UserDaoService) GetById(id int64) (*model.User, error) {
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

	var user model.User
	err = json.Unmarshal(([]byte)(response.Result), &user)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &user, err
}

func (this *UserDaoService) GetByName(name string) (*model.User, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/name/" + name
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

	var user model.User
	err = json.Unmarshal(([]byte)(response.Result), &user)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &user, err
}

func (this *UserDaoService) GetByRole(role int) (*model.User, error) {
	return nil, nil
}
