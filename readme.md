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