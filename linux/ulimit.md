- 修改当前交互终端的limit值 

查询当前终端的文件句柄数： ulimit -n 回车，一般的系统默认的1024. 

修改文件句柄数为65535，ulimit -n 65535.此时系统的文件句柄数为65535. 
- 将ulimit 值添加到/etc/profile文件中（适用于有root权限登录的系统） 

为了每次系统重新启动时，都可以获取更大的ulimit值，将ulimit 加入到/etc/profile 文件底部。 

echo ulimit -n 65535 >>/etc/profile 

source /etc/profile    #加载修改后的profile 

ulimit -n     #显示65535，修改完毕！

- 以为大功告成了，可突然发现自己再次登录进来的时候，ulimit的值还是1024，这是为什么呢？ 
关键的原因是你登录的用户是什么身份，是不是root用户，由于服务器的root用户权限很大，一般是不能用来登录的，都是通过自己本人的登录权限进行登录，并通过sudo方式切换到root用户下进行工作。 用户登录的时候执行sh脚本的顺序： 
/etc/profile.d/file 
/etc/profile 
/etc/bashrc 
/mingjie/.bashrc 
/mingjie/.bash_profile 

由于ulimit -n的脚本命令加载在第二部分，用户登录时由于权限原因在第二步还不能完成ulimit的修改，所以ulimit的值还是系统默认的1024。 

解决办法： 
修改linux的软硬件限制文件/etc/security/limits.conf. 

在文件尾部添加如下代码： 
* soft nofile 65536
* hard nofile 131072
* soft nproc 2048
* hard nproc 4096