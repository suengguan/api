package dao_service

import (
	"encoding/json"
	"fmt"
	"model"
	"strconv"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

type BussinessDaoService struct {
	baseUrl string
}

var BussinessDaoApi BussinessDaoService

func (this *BussinessDaoService) Init(ipAddr string) {
	this.baseUrl = ipAddr + "/v1"
}

func (this *BussinessDaoService) CreateProject(project *model.Project) (*model.Project, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(project)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return nil, err
	}

	url := this.baseUrl + "/project/"
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

	var p model.Project
	err = json.Unmarshal(([]byte)(response.Result), &p)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &p, err
}

func (this *BussinessDaoService) UpdateProject(project *model.Project) (*model.Project, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(project)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return nil, err
	}

	url := this.baseUrl + "/project/"
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

	var p model.Project
	err = json.Unmarshal(([]byte)(response.Result), &p)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &p, err
}

func (this *BussinessDaoService) GetProjectById(projectId int64) (*model.Project, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/project/id/" + strconv.FormatInt(projectId, 10)
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

	var p model.Project
	err = json.Unmarshal(([]byte)(response.Result), &p)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &p, err
}

func (this *BussinessDaoService) GetAllProjects(userId int64) ([]*model.Project, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/project/" + strconv.FormatInt(userId, 10)
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

	var projects []*model.Project
	err = json.Unmarshal(([]byte)(response.Result), &projects)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return projects, err
}

func (this *BussinessDaoService) DeleteProjectById(projectId int64) error {
	var result []byte
	var err error

	url := this.baseUrl + "/project/id/" + strconv.FormatInt(projectId, 10)
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

func (this *BussinessDaoService) DeleteAllProjects(userId int64) error {
	var result []byte
	var err error

	url := this.baseUrl + "/project/" + strconv.FormatInt(userId, 10)
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

func (this *BussinessDaoService) CreateJob(job *model.Job) (*model.Job, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(job)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return nil, err
	}

	url := this.baseUrl + "/job/"
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

	var j model.Job
	err = json.Unmarshal(([]byte)(response.Result), &j)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &j, err
}

func (this *BussinessDaoService) UpdateJob(job *model.Job) (*model.Job, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(job)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return nil, err
	}

	url := this.baseUrl + "/job/"
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

	var j model.Job
	err = json.Unmarshal(([]byte)(response.Result), &j)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &j, err
}

func (this *BussinessDaoService) UpdatePod(pod *model.Pod) (*model.Pod, error) {
	var result []byte
	var err error

	var requestData []byte
	requestData, err = json.Marshal(pod)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return nil, err
	}

	url := this.baseUrl + "/pod/"
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

	var p model.Pod
	err = json.Unmarshal(([]byte)(response.Result), &p)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &p, err
}

func (this *BussinessDaoService) GetJobById(jobId int64) (*model.Job, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/job/id/" + strconv.FormatInt(jobId, 10)
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

	var j model.Job
	err = json.Unmarshal(([]byte)(response.Result), &j)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &j, err
}

func (this *BussinessDaoService) GetModuleById(moduleId int64) (*model.Module, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/module/id/" + strconv.FormatInt(moduleId, 10)
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

	var m model.Module
	err = json.Unmarshal(([]byte)(response.Result), &m)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &m, err
}

func (this *BussinessDaoService) GetPodById(podId int64) (*model.Pod, error) {
	var result []byte
	var err error

	url := this.baseUrl + "/pod/id/" + strconv.FormatInt(podId, 10)
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

	var p model.Pod
	err = json.Unmarshal(([]byte)(response.Result), &p)
	if err != nil {
		beego.Debug("json Unmarshal data failed!", err)
		return nil, err
	}

	return &p, err
}
