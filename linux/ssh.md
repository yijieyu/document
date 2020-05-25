# ssh 使用

`ssh 简介：`

## 免密码登录

#### 1.创建公钥

```bash
# 使用 ssh-keygen 设置公钥，私钥，如果没有特别设置的，一路回车
ssh-keygen 
```

#### 2.服务器设置公钥

```bash
# 使用cat 将刚才生成的*.pub（公钥） 打印出来
cat id_rsa.pub
# 将公钥放置在远程服务器的用户的家目录， 可以放置多个用户
# 比如放置在root用户下， authorized_keys 代表允许登录的公钥，写入到authorized_keys
vi /root/.ssh/authorized_keys 
```

#### 3. 本地设置免密登录

```
# 编辑用户下.ssh/config 文件
HOST test
    HostName 远程服务器地址
    Port ssh端口
    ServerAliveInterval 120
    User 远程服务器公钥存的用户
    ServerAliveCountMax 3
```

#### 4.测试登录

```bash
ssh test 
# 没有报错，恭喜成功
```

## ssh 搭建通道

#### 配置

目的 解决两个内网通过ssh直接访问,条件 需要一个公网服务器，两个内网服务器

```
本地服务器：A(192.168.1.1)
目标服务器：B(192.168.88.1)
远程服务器：C(123.123.123.123)
通过服务器C 达到A->C->B 
```

配置远程服务器ssh 支持端口转发

ssh 配置目录 /etc/ssh/sshd_config

找到GatewayPorts no

no修改为 yes

```bash
#目标服务器B 设置
ssh -qTfNn -R 2222:localhost:22 tunnel-user@123.123.123.123 -p 22
# 访问远程服务器C的2222 端口时， 会转发到目标服务器B的22 端口上
# 远程服务器测试
ssh -p 2222 用户名@localhost # 用户名是目标服务器B的用户

# 然后本地服务器登录测试
# ssh -p 2222 用户名@远程服务器ip 
ssh -p 2222 root@123.123.123.123 # 用户是目标服务器B的用户
```

#### 通道安全

设置用户

```bash
# 添加用户转发的用户
sudo useradd tunnel-user 
# 可以设置密码， 最好设置ssh免密登录
# 设置密码
# 生成随机字符串
pwgen -sy1 16 10
sudo passwd tunnel-user
# 设置免密码登录 则是将目标服务器B的公钥放到远程服务器C的tunnel-user 的authorized_keys

# 禁用用户执行shell
sudo chsh -s /bin/false tunnel-user
```

#### 免密登录

本地服务器A的公钥保存在目标服务器B的登录用下.ssh/authorized_keys中

本地服务器A配置~/.ssh/config

```
HOST 目标服务器B的简称
    Port 52100 # 远程服务器端口
    User root  # 目标服务器B的用户
    HostName 103.102.192.54 # 远程服务器 IP
    ServerAliveInterval 60 # 活跃时间
```

#### 保持通道打开

通过脚本的方式 + crontab的方式

```bash
vi tunnel
#复制下面的shell代码
chmox +x tunnel
```

```bash
#!/bin/bash
pid=`ps -ef | grep 远程服务器监听端口 | grep ssh | awk -F' ' '{print $2}'`

if [ ! -n "$pid" ];then
	`ssh -qTfNn -R 2222:localhost:22 tunnel-user@123.123.123.123 -p 22`
fi
```

Crontab 命令

```bash
# 每分钟执行tunnel-ssh
* * * * * /root/tunnel
```

也可以直接在脚本中执行 while

#### 问题

本地服务器A连接不上目标服务器B，但是目标服务器B的ssh是正常的，此时就需要检查远程服务器C的sshd是否是正常

## 事项

参数含义

```bash
ssh -qTfNn -R 2222:localhost:22 tunnel-user@123.123.123.123 -p 22  -o ServerAliveInterval=120
```

- `-q:`  quiet模式，忽视大部分的警告和诊断信息（比如端口转发时的各种连接错误）
- `-T: ` 禁用tty分配(pseudo-terminal allocation)
- `-f:`登录成功后即转为后台任务执行
- `-N:`不执行远程命令（专门做端口转发）
- `-n:`重定向stdin为`/dev/null`，用于配合`-f`后台任务
- `2222`  代表远程服务器的端口
- `22`  代表目标服务器的接收端口

