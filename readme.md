# GPU显卡温度监控脚本

## 工作步骤：
1. 首先在服务器上执行 pip install gpustat 命令，用于之后的温度解析。
2. 在qq邮箱开启smtp服务，并记录smtp密码
3. 配置 ./conf/config.yaml 文件
```yaml
Email:
  # qq邮箱，将作为发送方发送邮件  
  Sender: xxxxxxxxx@qq.com 
  # 管理员邮箱，将作为接收方接收邮件
  Receiver: xxxxxxxxxx@163.com
  # 配置smtp密码
  Password: xxxxxxxxxx
  # smtp主机(如果使用qq邮箱作为发送方则不需改变)
  Host: smtp.qq.com
  # 发送端口号(如果使用qq邮箱作为发送方则不需改变)
  Port: 465

Machine:
  # 机器名
  ServerName: g3090
  # 安全温度(可根据实际情况设置)
  SafeTemp: 85
  # 在服务器端执行的命令
  Command: gpustat
```
4. 编译文件
```shell
# 当开发机器操作系统为MacOS
# 部署机器操作系统为Linux时
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
# 部署机器操作系统为Windows时
GOOS=windows GOARCH=amd64 go build main.go

# 当开发机器操作系统为Linux
# 部署机器操作系统为MacOS时
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go 
# 部署机器操作系统为Windows时
GOOS=windows GOARCH=amd64 go build main.go

# 当开发机器操作系统为Windows
# 部署机器操作系统为Mac时
SET CGO_ENABLED=0
SET GOOS=darwin3
SET GOARCH=amd64
go build main.go
# 部署机器操作系统为Windows时
SET CGO_ENABLED=0
SET GOOS=darwin3
SET GOARCH=amd64
go build main.go

# 当部署至本机同类型的操作系统时
go build main.go
```
5. 上传至服务器
```shell
scp -r local/src/TempMonitor username@host:remote/src/TempMonitor
```

6. 运行文件
```shell
cd remote/src/TempMonitor
./main
```