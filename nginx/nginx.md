# Nginx 配置

### try_files
`以 try_files $uri $uri/ /index.php; 为例。 首先查找 $uri 地址 ，然后$url/ $url/ 查找的$url/文件夹 ， 最后找 /index.php`
``` 当用户请求http://servers.blog.ustc.edu.cn/example  时，这里的 $uri就是 /example。try_files 会到硬盘里尝试找这个文件。如果存在名为 /$root/example（其中 $root 是 WordPress 的安装目录）的文件，就直接把这个文件的内容发送给用户。显然，目录中没有叫 example 的文件。然后就看 $uri/，增加了一个 /，也就是看有没有名为 /$root/example/ 的目录。又找不到，就会 fallback 到 try_files 的最后一个选项 /index.php，发起一个内部 “子请求”，也就是相当于 nginx 发起一个 HTTP 请求到 http://servers.blog.ustc.edu.cn/index.php。这个请求会被 location ~ \.php$ { ... } catch 住，也就是进入 FastCGI 的处理程序。而具体的 URI 及参数是在 REQUEST_URI 中传递给 FastCGI 和 WordPress 程序的，因此不受 URI 变化的影响。```

###  








# nginx 跨域 

## go语言 服务 。使用转发 nginx为docker
- 注意事项。 go语言启动时 必须和 proxy_pass 的ip一致
```nginx
server
{
    listen       7091;
    #server_name ad-admin-test.tvblack.com;
    server_name localhost;
    error_log /var/log/nginx/ssp_admin.log;

    location /
    {
        if ($request_method = OPTIONS) {
            return 200;
        }

        add_header 'Access-Control-Allow-Origin' 'http://localhost:1015';
        add_header 'Access-Control-Allow-Credentials' 'true';
        add_header 'Access-Control-Allow-Headers' 'X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,X-Token';
        add_header 'Access-Control-Allow-Methods' 'GET,POST';

        proxy_pass http://192.168.1.15:8092;
        proxy_redirect off;
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Protocol $scheme;
        proxy_set_header X-Url-Scheme $scheme;
    }
}

```


## php 后台 nginx为docker
- 注意事项。  fastcgi_pass  adx_admin_php:9000;  adx_admin_php 必须和docker-compose.yml 一致

```nginx
server
{
    listen       8089;
    #server_name ad-admin-test.tvblack.com;
    server_name localhost;
    root   /var/www/adx_admin/public;
    charset utf-8;
    access_log /var/log/nginx/adx_admin_access_log;
    error_log /var/log/nginx/adx_admin_error_log;
    index index.php;

    location /
    {
        if (!-e $request_filename) {
             rewrite  ^/(.*)$  /index.php?s=$1  last;
             break;
        }
    }
    location ~ \.php$
    {
        if ($request_method = OPTIONS) {
            return 200;
        }

        add_header 'Access-Control-Allow-Origin' 'http://localhost:8608';
        add_header 'Access-Control-Allow-Credentials' 'true';
        add_header 'Access-Control-Allow-Headers' 'X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,X-Token';
        add_header 'Access-Control-Allow-Methods' 'GET,POST';
        fastcgi_pass   adx_admin_php:9000;
        fastcgi_index  index.php;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }
}
```

总结跨域
add_header 'Access-Control-Allow-Origin' 'http://localhost:8608';  // 只允许此地址访问接口
add_header 'Access-Control-Allow-Credentials' 'true';  // 允许携带凭证（cookie） 访问接口
add_header 'Access-Control-Allow-Headers' 'X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,X-Token';  // 只允许携带的header头
add_header 'Access-Control-Allow-Methods' 'GET,POST';  // 只允许访问的方法





#react 打包后，服务器nginx配置

```nginx

server {
 listen 80;
 server_name www.tophudong.com;
 error_log /data/logs/tophudong/www.tophudong.com.log;
 root /data/projects/dingdian;

 location / {
    try_files $uri $uri/ @fallback;
    index index.html;
 }

 location @fallback {
    rewrite ^.*$ /index.html break;
 }
}
```
