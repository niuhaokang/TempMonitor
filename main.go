package main

import (
	"TempMonitor/model"
	"TempMonitor/util"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"time"
)

func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(path + "/conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	InitConfig()
	mailman := model.NewMailMan()

	mailman.SetSender(viper.GetString("Email.Sender"))
	mailman.SetReceiver(viper.GetString("Email.Receiver"))
	mailman.SetHost(viper.GetString("Email.Host"))
	mailman.SetPort(viper.GetInt("Email.Port"))
	mailman.SetPassword(viper.GetString("Email.Password"))
	mailman.SetEmail(&model.Email{
		Subject: "",
		Body: "",
	})

	machine := model.NewMachine()

	machine.SetServerName(viper.GetString("Machine.ServerName"))
	machine.SetSafeTemp(viper.GetInt("Machine.SafeTemp"))

	err := util.StartServer(mailman, machine)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	for {
		cmd := exec.Command(viper.GetString("Machine.Command"))
		resp, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("err: %v\n", err)
		}
		infos := util.ParseCMD(string(resp))
		for _, info := range infos {
			if info.Temp >= machine.SafeTemp {
				subject := fmt.Sprintf("%s机器显卡温度警告", machine.ServerName)
				body := fmt.Sprintf("%s号显卡温度为%d℃超过安全温度(%d℃), 请管理员立刻检查必要时关机处理。", info.Label, info.Temp, machine.SafeTemp)
				err := util.SentEmail(mailman, subject, body)
				if err != nil {
					log.Fatalf("err : %v\n", err)
				}
			}
		}
		time.Sleep(30 * time.Minute)
	}
}
