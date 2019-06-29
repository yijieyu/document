# 跨域

## 产生跨域原因
  - 页面的域名与请求服务的域名不一致
  - 域名不一致时，浏览器会先发送一个option请求方法的请求，查看请求的服务是不是支持option的方法。
  
  
## 解决跨域
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
        # 判断是否跨域请求，返回200 代表请求成功，浏览器才会进行真正的请求：如GET，POST
        if ($request_method = OPTIONS) {
            return 200;
        }

        # 跨域参数
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

#### 总结跨域
    add_header 'Access-Control-Allow-Origin' 'http://localhost:8608';  // 只允许此地址访问接口
    add_header 'Access-Control-Allow-Credentials' 'true';  // 允许携带凭证（cookie） 访问接口
    add_header 'Access-Control-Allow-Headers' 'X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,X-Token';  // 只允许携带的header头
    add_header 'Access-Control-Allow-Methods' 'GET,POST';  // 只允许访问的方法
