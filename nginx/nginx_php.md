# php后台 由docker启动

- 注意事项。  fastcgi_pass  adx_admin_php:9000
- adx_admin_php 必须和docker-compose.yml 一致

```nginx
server
{
    listen       8089;
    #server_name xx.xx.com;
    server_name localhost;
    root   /var/www/xxx/public;
    charset utf-8;
    access_log /var/log/nginx/xxx_access_log;
    error_log /var/log/nginx/xxx_error_log;
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
        fastcgi_pass   adx_admin_php:9000;
        fastcgi_index  index.php;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }
}
```