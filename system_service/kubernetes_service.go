package system_service

import (
	kubeApi "system-service/kubernetes"                         // todo
	"system-service/kubernetes/KubeRESTfulClient/ResourceModel" // todo

	"github.com/astaxie/beego"
)

func ApiCreateNamespace(namespace string) error {
	var err error

	var svc kubeApi.Service
	err = svc.CreateNamespace(namespace)
	if err != nil {
		beego.Debug(err)
		return err
	}

	return err
}

func ApiDeleteNamespace(namespace string) error {
	var err error

	var svc kubeApi.Service
	err = svc.DeleteNamespace(namespace)
	if err != nil {
		beego.Debug(err)
		return err
	}

	return err
}

func ApiCreateQuota(namespace, name, cpu, memory string) error {
	var err error

	var svc kubeApi.Service
	err = svc.CreateResourceQuota(namespace, name, cpu, memory)
	if err != nil {
		beego.Debug(err)
		return err
	}

	return err
}

// return cpu,memory,error
func ApiGetTotalCpuAndMemory() (float64, float64, error) {
	var err error
	var totalCpu float64
	var totalMemory float64

	var svc kubeApi.Service
	totalCpu, totalMemory, err = svc.GetTotalCpuAndMemory()
	if err != nil {
		beego.Debug(err)
		return 0.0, 0.0, err
	}

	return totalCpu, totalMemory, err
}

func ApiCreateRc(namespace string, label string, image string, containerPort int,
	limitCpu, limitMemory string, envList []*ResourceModel.Env, hostPathList []*ResourceModel.ContainerHostVolume) error {
	var err error

	var svc kubeApi.Service
	err = svc.CreateRc(namespace, label, image, containerPort, limitCpu, limitMemory, envList, hostPathList)
	if err != nil {
		beego.Debug(err)
		return err
	}

	return err
}

func ApiGetPodName(namespace, label string) (string, error) {
	var podName string
	var err error

	var svc kubeApi.Service
	podName, err = svc.GetPodName(namespace, label)
	if err != nil {
		return "", err
	}

	return podName, err
}

func ApiCreateSvc(namespace string, name string, label string, containerPort int) error {
	var err error

	var svc kubeApi.Service
	err = svc.CreateSvc(namespace, name, label, containerPort)
	if err != nil {
		beego.Debug(err)
		return err
	}

	return err
}

func ApiGetPodsByNamespace(namespace string) ([]string, error) {
	var err error
	var podNames []string

	var svc kubeApi.Service
	podNames, err = svc.GetPodsByNamespace(namespace)
	if err != nil {
		return nil, err
	}

	return podNames, err
}

func ApiDeleteRc(namespace string, rcName string) error {
	var err error

	var svc kubeApi.Service
	err = svc.DeleteRc(namespace, rcName)
	if err != nil {
		return err
	}

	return err
}

func ApiDeleteSvc(namespace string, svcName string) error {
	var err error

	var svc kubeApi.Service
	err = svc.DeleteSvc(namespace, svcName)
	if err != nil {
		return err
	}

	return err
}
