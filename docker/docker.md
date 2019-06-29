# docker 原理
Docker基于Linux容器技术（LXC），Namespace，Cgroup，UnionFS（联合文件系统）等技术实现应用隔离虚拟化方案，
采用是C/S的架构，Docker客户端与Docker daemon进行交互，daemon负责构建、运行和发布Docker容器。
Docker的客户端的daemon通过RESTful API进行socket通信。
Docker包含镜像和容器，镜像是通过Dockerfile或者容器构建的，镜像是只读的
容器时通过镜像启动运行的，在镜像上有一个可写层

# docker常用命令

- docker build -t tagname <Dockerfile path> 编译镜像
- docker images 查看镜像
- docker rmi    删除镜像
- docker run 创建并启动容器
- docker start 启动容器
- docker stop 停止容器
- docker ps 查看正在运行的容器
- docker ps -a 查看所有容器（包括停止运行的）
- docker rm 删除容器
- docker rm `docker ps -a | grep Exited | awk '{print $1}'`   删除异常停止的docker容器
- docker rmi -f  `docker images | grep '<none>' | awk '{print $3}'`   删除名称或标签为none的镜像
- docker run -d --name docker_gitlab-runner --privileged=true yijieyu/docker-gitlab-runner:latest /usr/sbin/init 后台启动一个容器
- docker run -it -d --privileged --name=docker1 -d docker:latest 启动包含docker的jing xiang


# Dcokerfile 常用命令

- FROM 父镜像
- MAINTAINER 作者信息
- ENV 环境变量
- RUN 运行一条shell指令
- WORKDIR 容器运行时的工作目录， 容器运行的默认目录
- CP 拷贝当前主机文件到镜像内
- ADD 拷贝主机，网络地址等信息到镜像内，会自动解压压缩包
- EXPOSE 容器暴露的端口
- ENTRYPOINT 容器启动后第一个执行命令，这个是在容器启动时不可替换的
- CMD 容器启动后执行的命令，如果定义了ENTRYPOINT， CMD就是ENTRYPOINT，命令的参数，容器启动时最后一个参数可以替换CMD指令


# docker compose介绍
docker compose是docker的容器编排服务，使用yml配置容器状态与依赖，每个docker compose启动后都会创建一个独立的虚拟网络
配置中所有的容器都运行在这个子网里面，可以实现容器通信

# docker compose 常用命令

- docker-compose up -d 后台启动容器
- docker-compose down 停止并删除容器
- docker-compose start 启动容器
- docker-compose stop 关闭容器
- docker-compose logs -f 查看容器日志

# docker compose yml配置常用指令

- version 配置文件版本
- services 服务配置项
- image 镜像名称
- container_name 容器名称, 默认是{配置文件所在目录名称}_{服务名称}-init-replica_{编号}，例如：chatops_mongo-init-replica_1
- privileged linux内核权限问题，设置为true，这样容器内就可以对挂载的主机目录写数据
- restart 容器异常结束后自动重启（不是通过docker命令关闭，是容器内部出现错误导致的，只要容器内常驻的进程被关闭容器就会关闭），开启自动启动
- environment 定义容器的环境变量
- ports 主机与容器的端口映射
- volumes 挂载主机目录和容器卷
- networks 网络设置



# docker安装
- 下载docker文件
- 安装 rpm -i docker-ce-17.06.0.ce-1.el7.centos.x86_64.rpm
- `yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo`
-- `sudo yum install -y yum-utils  device-mapper-persistent-data lvm2`
-- `container-selinux >= 2.9 错误` 解决：wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo   yum install epel-release   #阿里云上的epel源，然后yum install container-selinux

-- 修改docker配置存储
 --graph=/opt/docker

- docker-composer安装
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose

docker-compose --version