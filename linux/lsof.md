## lsof 操作



#### 查看打开最多句柄的PID
lsof | awk 'NR>1 {++S[$2]} END { for(a in S) {print a,"\t",S[a]}}'|sort -n -k 2|tail -n 1


#### 查看每个PID 打开的句柄
lsof -n|awk '{print $2}'|sort|uniq -c|sort -nr|more