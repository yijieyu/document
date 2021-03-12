允许Linux上的目录/var/www能被多个用户进行读写，

用户www-data（组www-data）

用户dev（组dev）

将用户dev添加进用户组www-data
设置目录为组www-data可写可读
sudo usermod -a -G www-data dev
sudo chgrp -R www-data /var/www
sudo chmod -R g+w /var/www