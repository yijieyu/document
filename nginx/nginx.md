# Nginx 配置

### try_files
`以 try_files $uri $uri/ /index.php; 为例。 首先查找 $uri 地址 ，然后$url/ $url/ 查找的$url/文件夹 ， 最后找 /index.php`
``` 当用户请求http://servers.blog.ustc.edu.cn/example  时，这里的 $uri就是 /example。try_files 会到硬盘里尝试找这个文件。如果存在名为 /$root/example（其中 $root 是 WordPress 的安装目录）的文件，就直接把这个文件的内容发送给用户。显然，目录中没有叫 example 的文件。然后就看 $uri/，增加了一个 /，也就是看有没有名为 /$root/example/ 的目录。又找不到，就会 fallback 到 try_files 的最后一个选项 /index.php，发起一个内部 “子请求”，也就是相当于 nginx 发起一个 HTTP 请求到 http://servers.blog.ustc.edu.cn/index.php。这个请求会被 location ~ \.php$ { ... } catch 住，也就是进入 FastCGI 的处理程序。而具体的 URI 及参数是在 REQUEST_URI 中传递给 FastCGI 和 WordPress 程序的，因此不受 URI 变化的影响。```
