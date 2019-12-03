## 遇到的问题

#### too_many_open_files

看什么进程打开最多文件句柄 
```bash
lsof -n|awk '{print $2}'|sort|uniq -c|sort -nr|more
``` 

查看系统打开文件句柄数
   ```bash
lsof | awk 'NR>1 {++S[$2]} END { for(a in S) {print a,"\t",S[a]}}'|sort -n -k 2|tail -n 1
```

解决:
 vim .bash_profile， 添加 
```bash
ulimit -S -n 2048 # or whatever number you choose
``` 
打开.zshrc文件 ，并添加source .bash_profile
重新加载 source .zshrc