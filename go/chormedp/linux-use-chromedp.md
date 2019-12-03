## 流程

    1. 安装桌面
    2. 安装google-chrome
    3. 安装x11vnc
    4. Xvfb
    5. 配置

## 安装桌面
```bash
yum -y groups install "GNOME Desktop"
```

## centos 安装 google-chrome

## [官网下载](#https://www.google.com/intl/zh-CN/chrome/)
    centos选在[.rpm文件](https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm) 下载
    
## 下载好后使用 yum install 下载的.rpm文件
    如果遇到错误， 则使用更新yum update 基本就可以解决问题
    

## 安装[x11vnc](#https://centos.pkgs.org/7/epel-x86_64/x11vnc-0.9.13-11.el7.x86_64.rpm.html)
    1. 可以直接 yum install https://centos.pkgs.org/7/epel-x86_64/x11vnc-0.9.13-11.el7.x86_64.rpm.html
    2. 也可以 yum install http://download-ib01.fedoraproject.org/pub/epel/7/x86_64/Packages/x/x11vnc-0.9.13-11.el7.x86_64.rpm
    3. 也可以 http://ftp.tu-chemnitz.de/pub/linux/dag/redhat/el7/en/x86_64/rpmforge/RPMS/x11vnc-0.9.13-1.el7.rf.x86_64.rpm
   
## 安装 Xvfb
    yum install Xvfb
    
#### 虚拟客户端
    1. /usr/local/x11vnc/bin/x11vnc -rfbport 12345 -passwd 123456 -create -forever
    2. x11vnc -rfbport 12345 -passwd 123456 -create -forever
    
#### 启动客户端窗口命令 

   ```bash
Xvfb :2 -screen 0 1024x768x24&
```
   :2 代表窗口号

#### 命令行启动谷歌浏览器 
   逻辑就是让谷歌浏览器基于虚拟窗口号
   
```bash
# -display 代表第几个窗口展示
google-chrome-stable -disable-gpu --screenshot -display 0 https://www.suning.com/
```
   后续使用shell脚本将-display 设置为环境变量


#### 综合以上结合go程序
   shell 脚本
```bash
#!/bin/env bash

export DISPLAY=:0

exec google-chrome-stable ${@:1}
```
   GO程序去调用shell脚本
```go
package main

import (
 "github.com/chromedp/cdproto/network"
 	"github.com/chromedp/chromedp"
 	"github.com/davecgh/go-spew/spew"
)


func main(){
	test()
}

func test(){
	options := []chromedp.ExecAllocatorOption{
    		// 开启GUI来debug
    		chromedp.Flag("mute-audio", false),
    		chromedp.Flag("headless", false),
    		
    		//此处调用shell脚本
    		chromedp.ExecPath("/tmp/a.sh"),
    		//设置UA
    		chromedp.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36`),
    	}
    
    	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
    
    	c, cc := chromedp.NewExecAllocator(context.Background(), options...)
    	defer cc()
    	spew.Dump(c)
    
    	ctxt, cancelCtxt := chromedp.NewContext(c) // create new tab
    	defer cancelCtxt()                         // close tab afterwards
    
    	var title string
    	var ok bool
    	if err := chromedp.Run(ctxt,
    		//chromedp.Navigate("https://example.com"),
    		//chromedp.Title(&title),
    
    		chromedp.Navigate(`https://v.qq.com/x/search/?q=&stag=&smartbox_ab=`),
    		chromedp.WaitVisible(`#searchForm`),
    		chromedp.SendKeys(`#keywords`, `哪吒`, chromedp.ByID),
    		chromedp.Click(`.search_btn`, chromedp.NodeVisible),
    		chromedp.WaitVisible(`.wrapper_main .result_item `),
    		chromedp.AttributeValue(`.wrapper_main .result_item:nth-child(2) ._infos .figure_pic`, `alt`, &title, &ok),
    	); err != nil {
    		log.Fatalf("Failed: %v", err)
    	}
    
    	spew.Dump(title)
}

```

## docker 镜像安装

#### 拉取centos:centos7.7.1908镜像
```bash
# 获取systemctl权限
docker run --privileged -d -ti -e "container=docker"  -v /sys/fs/cgroup:/sys/fs/cgroup -p 5901:5901 centos:centos7.7.1908 /usr/sbin/init

# 进去容器 找到容器id
docker ps

docker exec -ti 容器id bash

# 容器内操作
# 1. 更新
yum update -y

# 2.安装桌面系统
yum groupinstall "GNOME Desktop" "X Window System" "Desktop"
yum install gnome-classic-session gnome-terminal nautilus-open-terminal control-center liberation-mono-fonts

# 3. 安装google
yum install https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm

# 4. 安装vncserver
yum install tigervnc-server tigervnc vnc vnc-server

# 5. 配置vncserver
cp /usr/lib/systemd/system/vncserver@.service /etc/systemd/system/vncserver@.service

#启动服务的用户为 root，添加 User=root, 这样 VNC Client 访问时可以看到菜单栏(Menu bar);
#将 <USER> 替换为 admin (本机的非 root 用户), 这样用户登录到 admin 的界面;
vi /etc/systemd/system/vncserver@.service

# vncserver 设置密码
vim /etc/libvirt/qemu.conf
# 找到这个
vnc_password = "123456"
vnc_listen = "0.0.0.0"

# 另外设置vncsever 密码 ， 输入两次相同的密码, 遇到view-only 输入 n
vncpasswd

# 重新加载配置
systemctl daemon-reload

# 设置虚拟机窗口
vncserver :1 -name test -depth 32 -geometry 1920x1080 -xstartup /etc/systemd/system/vncserver@.service

# 谷歌浏览器利用虚拟机窗口访问
google-chrome-stable DISPLAY=:1

#
sudo docker run --privileged -d -ti -p 5901:5901 -p 8081:8081 yijieyu/centos_gnome_vnc:v1 /bin/bash


sudo docker cp chromedp_exec.sh 52ff097d6bf1:/home/ops
sudo docker cp conf.yml 52ff097d6bf1:/home/ops
sudo docker cp cookies/ 52ff097d6bf1:/home/ops
sudo docker cp media_collection_login_server 52ff097d6bf1:/home/ops


```

