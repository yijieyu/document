# nginx安装

### 下载（网址：http://nginx.org）
```bash
wget http://nginx.org/download/nginx-1.12.2.tar.gz
```

#### 解压 
```bash
tar -zxvf nginx-1.12.2.tar.gz
```

#### 配置
```bash
./configure \
--prefix=/data/server/nginx \
--pid-path=/var/run/nginx/nginx.pid \
--lock-path=/var/lock/nginx.lock \
--error-log-path=/var/log/nginx/error.log \
--http-log-path=/var/log/nginx/access.log \
--with-http_gzip_static_module \
--http-client-body-temp-path=/var/temp/nginx/client \
--http-proxy-temp-path=/var/temp/nginx/proxy \
--http-fastcgi-temp-path=/var/temp/nginx/fastcgi \
--http-uwsgi-temp-path=/var/temp/nginx/uwsgi \
--http-scgi-temp-path=/var/temp/nginx/scgi
```

#### 编译安装
```bash
make && make install
```