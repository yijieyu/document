# 常用命令

- ffprobe -v quiet -print_format json -show_format -show_streams #{source} 查看音频视频文件信息并输出为json格式
- ffmpeg -i #{source} -acodec copy -vcodec copy -f mp4 test.mp4 视频转码格式,如：ts转mp4
- ffmpeg -i 1.mp4 -vcodec copy -acodec copy -vbsf h264_mp4toannexb 1.ts ts转mp4




sudo docker run --rm -v "$PWD":/tmp/ftp -w /tmp/ftp golang:1.12 go build -mod vendor -tags "jsoniter" -v -o ftp
