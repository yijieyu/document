# 文档
    - 所有文档使用md方式（github直接可以渲染）
    - 不要在根目录添加md文件，尽量放在目录里，更清晰
    - 新增加文档，记得更新下面的目录结构

## 目录结构 
```
├── LICENSE
├── README.md
├── docker
│   ├── docker-compose              docker-compose 各种实例
│   │   ├── mysql                   mysql 实例
│   │   │   └── docker-compose.yml
│   │   ├── php                     php使用docker-compose 运行实例
│   │   │   ├── docker-compose.yml
│   │   │   ├── nginx
│   │   │   │   ├── conf.d
│   │   │   │   │   └── tvblack_adv.conf
│   │   │   │   └── nginx.conf
│   │   │   ├── ott_admin
│   │   │   │   └── docker-entrypoint.sh
│   │   │   └── php7.1.8
│   │   │       ├── php.conf.d
│   │   │       │   ├── docker.conf
│   │   │       │   ├── www.conf
│   │   │       │   └── zz-docker.conf
│   │   │       ├── php.ini
│   │   │       └── php.ini.bak
│   │   └── redis                   redis 实例
│   │       └── docker-compose.yml
│   ├── docker-file                 docker file 各种实例
│   │   └── php7.1.8                php7.1.8 实例
│   │       └── Dockerfile
│   └── docker.md                   docker常用命令
├── e-book                          电子书
│   ├── go
│   └── mysql
│       └── SQL反模式中文完整版.pdf
├── git
│   └── git.md                      git 常用命令操作
├── idea
│   └── idea_keyboard_shortcuts.md  idea快捷键
└── nginx
    ├── examples.md                 前后分离部署
    ├── nginx.md                    nginx文档
    ├── nginx_cross_domain.md       跨域配置
    ├── nginx_go.md                 go-nginx配置
    ├── nginx_php.md                php-nginx配置
    ├── nginx_static_resource.md    静态资源配置
    └── react_vue_angular.md        前端框架打包后配置
```