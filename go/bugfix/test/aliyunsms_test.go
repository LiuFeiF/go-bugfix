package test

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client" //阿里云提供的
	"log"
	"testing"
)

func SendRefundOrderArrive(t *testing.T) error {
	var tel, shopTitle, productTitle string
	var money float64                 //此为我在阿里云控制台配置的消息模板
	var signName, templateCode string //此为在阿里云申请的信息签名以及审核通过后阿里云提供的信息模板id
	accessKeyId := ""                 //阿里云开通账号id
	accessKeySecret := ""             //阿里云凭证
	endpoint := ""                    //此为在阿里云地址
	c := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &endpoint,
	}
	newClient, err := dysmsapi20170525.NewClient(c) //链接阿里云
	if err != nil {
		log.Println("NewClient err:" + err.Error())
		return err
	}
	param := fmt.Sprintf("{\"shopName\":\"%s\", \"productName\":\"%s\", \"money\":%.2f}", shopTitle, productTitle, money) //和阿里云的模板字段需要对应
	request := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &tel,
		TemplateCode:  &templateCode,
		SignName:      &signName,
		TemplateParam: &param,
	}
	_, err = newClient.SendSms(request)
	if err != nil {
		log.Println("SendSms err:" + err.Error())
		return err
	}
	return nil
}
