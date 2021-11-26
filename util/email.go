package util

import (
	"TempMonitor/model"
	"fmt"
	"gopkg.in/gomail.v2"
)

func SentEmail(mailman *model.MailMan, subject string, body string) error {
	mailman.Update(subject, body)
	msg := gomail.NewMessage()
	msg.SetHeader("From", mailman.Sender)
	msg.SetHeader("To", mailman.Receiver)
	msg.SetHeader("Subject", mailman.Subject)

	msg.SetBody("text", mailman.Body)

	d := gomail.NewDialer(mailman.Host, mailman.Port, mailman.Sender, mailman.Password)
	if err := d.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}

func StartServer(mailman *model.MailMan, machine *model.Machine) error {
	subject := fmt.Sprintf("%s机器GPU温度监控报告", machine.ServerName)
	body := fmt.Sprintf("本机器安全温度为%d, 当显卡温度超过时会邮件报警, 请管理员及时处理。", machine.SafeTemp)
	err := SentEmail(mailman, subject, body)
	return err
}


