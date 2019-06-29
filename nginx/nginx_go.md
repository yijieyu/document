# go语言 作为服务

- 注意事项。 go语言启动时 必须和 proxy_pass 的ip一致

```nginx
server
{
    listen       7091;
    #server_name xx.xxx.com;
    server_name localhost;
    access_log /var/log/nginx/xxx_access_log;
    error_log /var/log/nginx/xxx_error.log;

    location /
    {
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