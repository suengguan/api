package app_service

import (
	//"config"
	"encoding/json"
	"fmt"
	"model"
	"strconv"
	"utility/httpclient"

	"github.com/astaxie/beego"
)

const (
	// SUMMARY_SERVICE = "http://192.168.0.206:9007/v1/summary"
	//SUMMARY_SERVICE = "summary-service"
	SUMMARY_SERVICE = "http://summary-service.pme-system-beta1:8080/v1/summary"
)

func ApiGetSummary(userId int64) ([]byte, error) {
	var result []byte
	var err error

	// var serviceBaseUrl string
	// serviceBaseUrl, err = config.GetUrlByServiceName(SUMMARY_SERVICE)
	// if err != nil {
	// 	return nil, err
	// }

	url := SUMMARY_SERVICE + "/" + strconv.FormatInt(userId, 10)
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
