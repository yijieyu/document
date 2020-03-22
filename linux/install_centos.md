## 安装centos

#### 官网下载 https://centos.org/
    找到镜像列表 也可以到（http://mirrors.aliyun.com/centos/7.7.1908/isos/x86_64/） 这里下载

## 配置USB启动盘（https://blog.seisman.info/linux-usb-installer/）
   
#### mac制作
   - 1.将安装文件转换成 dmg 格式：
   ```bash
cd ~/Download
hdiutil convert -format UDRW -o centos.dmg CentOS-7-x86_64-LiveGNOME-1708.iso
```
   - 2.插入 U 盘，确定U盘 盘符 执行如下命令：
   ```bash
diskutil list
```
   通常最后一个盘符是 U 盘，请根据磁盘大小确认哪一个是 U 盘。如果觉得没有把握，可以使用 磁盘工具来辅助判断。如果找错盘符，后续可能会导致其他盘数据丢失。
   - 3.卸载 U 盘：
      
      使用如下命令卸载 U 盘，但不要拔出。也不要使用图形界面推出 U 盘。
   ```bash
diskutil unmountDisk /dev/disk2  # 根据自己的实际情况修改 /dev/disk2
Unmount of all volumes on disk2 was successful
```
   - 4.将镜像写入 U 盘
   ```bash
sudo dd if=./centos.dmg of=/dev/rdisk2 bs=1m
```
   这里的盘符千万别写错。如果正常，这里需要等一小会，完成后就可以在 Finder 推出 U 盘了。
   
   
#### Windows制作

Windows 下制作 USB 安装镜像可以使用 [Universal USB installer](#https://www.pendrivelinux.com/universal-usb-installer-easy-as-1-2-3/)， 点击鼠标执行即可。


#### ts转mp4
```bash
/usr/local/bin/ffmpeg -i /Users/fuyi/Downloads/video/上传动漫.ts -acodec copy -vcodec copy -f mp4 /Users/fuyi/Downloads/video/test.mp4
```

#### 获取视频信息
```bash
ffprobe -v quiet -print_format json -show_format -show_streams videoPath
```