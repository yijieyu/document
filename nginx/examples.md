# 前后台使用同一个域名

```nginx
server {
 listen 80;
 server_name www.xxx.com;
 error_log /data/logs/nginx/www.xxx.com.log;
 root /data/share_tasks/share_tasks_admin_ui;

 # 后台接口地址
 location /api {
    proxy_pass http://127.0.0.1:8100/;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
    proxy_set_header Host $http_host;

    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forward-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forward-Proto http;
    proxy_set_header X-Nginx-Proxy true;

    proxy_redirect off;
 }

 # 上传地址
 location /uploaded/ {
#   autoindex on;
   alias /data/share_tasks/share_tasks_upload/uploaded/ ;
 }

 # 后台页面资源
 location / {
    try_files $uri $uri/ @fallback;
    index index.html;
 }
 location @fallback {
    rewrite ^.*$ /index.html break;
 }

}
```

# 总结
    - /api 代表 后台服务地址
    - / 代表前端路由地址
    - /uploaded/ 静态资源访问地址
    - 使用同一个域名，不存在跨域问题