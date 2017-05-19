package system_service

import (
	"model"
	"testing"
	"time"
)

func Test_ApiSendEmails(t *testing.T) {
	// var emails []*model.Email
	// var e1 model.Email
	// e1.To = "suengguan@126.com"
	// e1.Subject = "test"
	// e1.Body = "Api ApiSendEmails test"
	// emails = append(emails, &e1)

	// var e2 model.Email
	// e2.To = "qsg@corex-tek.com"
	// e2.Subject = "test"
	// e2.Body = "Api ApiSendEmails test"
	// emails = append(emails, &e2)

	// err := ApiSendEmails(emails)

	var emails []*model.Email
	var e1 model.Email
	e1.To = "suengguan@126.com"
	e1.Subject = "登录提醒"
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	e1.Body = "【科思世纪官网登录提醒】欢迎您于 " + currentTime + " 登录PME系统。特此提醒。【科思世纪】"
	emails = append(emails, &e1)

	err := ApiSendEmails(emails)

	t.Log(err)
}
