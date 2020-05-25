# samba简单使用

## 安装

```bash
sudo yum install samba
```

## 配置

#### 配置文件

配置文件位置：/etc/samba/smb.conf

主要有以上三个部分：[global], [homes], [printers]

###### [global]

```toml
[global]
	workgroup = SAMBA
	security = user
	passdb backend = tdbsam

  # 设置打印机相关参数
  printing = cups
	printcap name = cups
	load printers = yes
	cups options = raw
	
	#设置允许打开软链
  wide links = yes
  symlinks = yes
  unix extensions = no

```

workgroup 用来定义工作组

security = user #这里指定samba的安全等级。关于安全等级有四种：

- share：用户不需要账户及密码即可登录samba服务器
- user：由提供服务的samba服务器负责检查账户及密码（默认）
- server：检查账户及密码的工作由另一台windows或samba服务器负责
- domain：指定windows域控制服务器来验证用户的账户及密码。

passdb backend = tdbsam 

- passdb backend （用户后台），
- samba有三种用户后台：smbpasswd, tdbsam和ldapsam。

###### [home]

```toml
[homes]
	comment = Home Directories # 设置的samba用户的家目录
	valid users = %S, %D%w%S
	browseable = No
	read only = No
	inherit acls = Yes
```

###### [printers]

 该部分内容设置打印机共享。

### 新增分享目录

```toml
[shared]
        comment = Shared Directories # 简单描述
        path = /tmp/share # 分享目录
        public = yes # 是否
#        admin users = smbuser # 设置用户名
#        valid users = @smbuser #表示允许smbuser连接到此samba服务
				# 读写权限
 				browseable = yes 
        writable = yes
        create mask = 0777
        directory mask = 0777
        force directory mode = 0777
        force create mode = 0777
```

## 设置用户

```bash
useradd smbuser
smbpasswd -a smbuser # 设置 用户密码
```

## 启动
```bash
# 重启
systemctl restart smb

#启动
systemctl start smb

```

## 事项

如果并不涉及加密， 可以在全局设置[global]  security = share
以及设置的分享目录 注释掉 admin users 和 valid users