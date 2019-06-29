# git

配置
    配置的级别逐级覆盖，项目配置覆盖全局配置，全局配置覆盖系统配置，建议像用户名和邮箱采用项目级别配置，个人习惯如别名、ui调整等采用全局配置，和系统环境有很大关系的配置采用系统级别配置

    系统级配置：/etc/.gitconfig（自己编译的在安装目录里）
    git config --system user.name 姓名
    git config --system user.email 邮箱

    全局配置：~/.gitconfig
    git config --global user.name 姓名
    git config --global user.email 邮箱

    项目级别配置：.git/config
    git config user.name 姓名
    git config user.email 邮箱

    删除配置
    git config --unset user.name

    查看配置
    git config user.name
    git config --list

    设置别名
    git config --global alias.st staus    相当于git st == git status

    其他习惯配置
    git config --global color.ui    true        在git命令输出中开启颜色

    忽略文件权限设置
    git config core.filemode false

文件忽略
    已在版本库的文件不能忽略

    项目级共享忽略（每个目录都可以存在这个文件，用来控制所有子目录文件）
    修改.gitigonre

    本地独享忽略
    修改.git/inf/exclude

    指定一个全局的忽略文件
    git config --global core.excludesfile filename

    查看忽略
    git status --ignored -s

    忽略已经跟踪的文件
    git update-index --assume-unchanged /path/to/file

    恢复忽略跟踪文件
    git update-index --no-assume-unchanged /path/to/file

常用命令

    版本初始化
    git init 目录

    在当前目录创建版本库
    git init

    添加文件到暂存区
    git add .

    提交暂存区的变更到版本库
    git commit -m "提交说明"

    直接提交工作区的变更到版本库(强烈不建议使用)
    git commit    -a -m "提交说明"

    修补最后一次提交
    git commit --amend --allow-empty --reset-author
        --amend 对刚刚的提交进行修补
        --allow-empty    允许空提交
        --reset-author    重置author提交者的id，否者只会影响commit的id 

    重置暂存区不改变工作区
    git reset HEAD

    重置暂存区和工作区
    git reset --hard HEAD

    重置到指定提交
    git reset --hard 9e8a761

    重置暂存区指定目录
    git reset [-q] [<commit>] [--] <paths>……

    重置指定文件
    git reset HEAD filename

    重置版本库和暂存区到上一次提交
    git reset HEAD^

    挽救错误重置（查看版本库变更记录）
    git reflog show master

    基于暂存区取消工作区的修改（工作区内容覆盖不能找回）
    git checkout -- .

    基于暂存区取消工作区指定文件的变更（工作区内容覆盖不能找回）
    git checkout -- filename

工作进度管理

    保存当前未提交修改到工作进入
    git stash 

    查看保存的工作进度
    git stash list

    从最近保存的进度进行恢复，恢复后删除工作经度
    git stash pop

    从最近保存的进度进行恢复，不删除工作经度
    git stash apply
    git stash apply [<sstash>]

    删除工作进度
    git stash drop [<sstash>]

    删除所有工作进度
    git stash clear

    基于进度创建分支
    git stash branch <branchname> <stash>

状态和日志相关命令

    查看当前状态
    git status

    查看精简的状态
    git status -s


    查看日志每次提交文件变更统计
    git log --stat

    精简的日志
    git log    --oneline

    查看提交的文件清单
    git log  --name-only

    显示提交新增、修改和删除的文件清单
    git log --name-status

    显示ASCII图形表示的分支合并历史
    git log --graph

    显示每次提交的内容差异
    git log -p

    显示最近的两次提交
    git log -2

    自定义日志格式
    git log --pretty=format:"%h - %an, %ar : %s"
    选项     说明
    %H        提交对象（commit）的完整哈希字串
    %h        提交对象的简短哈希字串
    %T        树对象（tree）的完整哈希字串
    %t        树对象的简短哈希字串
    %P        父对象（parent）的完整哈希字串
    %p        父对象的简短哈希字串
    %an        作者（author）的名字
    %ae        作者的电子邮件地址
    %ad        作者修订日期（可以用 -date= 选项定制格式）
    %ar        作者修订日期，按多久以前的方式显示
    %cn        提交者(committer)的名字
    %ce        提交者的电子邮件地址
    %cd        提交日期
    %cr        提交日期，按多久以前的方式显示
    %s        提交说明

    git log其他参数
    -(n)    仅显示最近的 n 条提交
    --since, --after 仅显示指定时间之后的提交。
    --until, --before 仅显示指定时间之前的提交。
    --author 仅显示指定作者相关的提交。
    --committer 仅显示指定提交者相关的提交。    

差异比较

    比较暂存区和工作区的差异
    git diff 

    比较暂存区和版本库的差异
    git diff HEAD

    比较工作区和版本库的差异
    git diff --cached
    git diff --staged

git版本库工作区项目命令

    显示版本库.git目录所在位置
    git rev-parse --git-dir

    显示工作去根目录
    git rev-parse --show-toplevel

    显示相对于工作区根目录的相对目录
    git rev-parse --show-prefix

    显示从当前目录cd到工作区根目录的深度
    git rev-parse --show-cdup

合并忽略策略

	git config --global merge.ours.driver true
    echo "database.xml merge=ours" >> .gitattributes    // 忽略database.xml文件合并