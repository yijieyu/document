## 自动部署说明

## 部署所需环境
  - python 必须是python3+ 版本，命令查看python -V Python 3.7.4
  - 安装fabric ，命令 pip install fabric （python3+）
  - 安装PyYAML ，命令 pip install PyYAML （python3+）

## 部署配置文件
```
例子：fab_config.yml。 文件名必须是fab_config.yml
```

#### 字段描述
| 字段 | 描述| 更新人 | 
|---| --- | --- |
| app_name | 应用名称 | fy |
| hosts | [多组服务器信息](#服务器信息)-对应prod，test,dev环境等 | fy |
| macro | 配置文件需要替换的变量，键值对 | fy |
| pre_cmd | 发送前本地执行命令，多条，从上至下执行 | fy |
| exec_files | 往服务器发送文件，键值对 | fy |
| exec_cmd | 服务器执行命令，多条，从上至下执行 | fy |
| deploy_env | 部署启动的配置-对应prod，test,dev环境等的端口 | fy |

#### 服务器信息
| 字段 | 描述| 例子 |
|---| --- | --- |
| host | ip地址/或者ssh配置的Host | user@127.0.0.1:22/ test |
| password | 当host为ip时，必须填写用户登录密码 |
| remote_path | 程序存放目录 | 绝对路径 /data/projects/.... |
| back_log | 程序部署备份存放目录 | 绝对路径 /data/projects/.... |


## 使用说明
```bash
# 必须把fabfile.py 和fab_config.yml放在项目根目录下

# 查看fab 可执行命令
fab -l

#结果如下
Available tasks:
  deploy
  rollback 

# deploy 部署 env为配置的环境值：prod，dev，test... 
fab deploy env
fab deploy prod

# rollback 部署 env为配置的环境值：prod，dev，test... 
fab rollback env
fab rollback dev
  
```