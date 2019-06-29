# 图片/视频等直接访问的配置

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

    location /uploaded/ {
    #   autoindex on;    // 是否支持目录文件浏览
       alias /data/share_tasks/share_tasks_upload/uploaded/ ;
     }
}
```

#### 解释
    - alias: 等待有缘人