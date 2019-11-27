# 常规流程
### 第一次
- git clone https://github.com/ -> ```克隆项目```
- git checkout develop -> ```切换到develop``` 
- git pull origin develop --rebase  -> ```拉取最新代码``` 
- git checkout -b fuyi/feature/xxx -> ```然后编写代码```
- git add . 或者 文件路径  -> ```文件添加到索引区```
- git commit -m "描述"  -> ```增加提交到本地仓库```
- git pull origin develop --rebase  -> ```再次拉取最新代码，其他人提交了最新代码```
- git push origin 当前分支 -> ```推送到远程仓库```
- 最后在web页面向负责人提交pr。 可以自己合并

### 以后每次开发都切换到develop 然后 拉取最新代码， 如果切换分支是新建分支，则如下操作
- git checkout develop 
- git pull origin develop --rebase -> ```每次拉取最新代码``` 
- git checkout -b fuyi/feature/xxx -> ```切换到一个新的分支。然后开始编写代码```
- git add . 或者 文件路径  -> ```编写完成 文件添加到索引区```
- git commit -m "描述"  -> ```增加提交到本地仓库```
- git pull origin develop --rebase  -> ```再次拉取最新代码，其他人提交了最新代码```
- git push origin 当前分支 -> ```推送到远程仓库```
- 最后在web页面向负责人提交pr。 可以自己合并

### 以后每次开发都切换到develop 然后 拉取最新代码， 如果已建分支，则如下操作
- git checkout develop 
- git pull origin develop --rebase -> ```每次拉取最新代码```
- git checkout fuyi/feature/xxx -> ```切换到一个已有的分支。```
- git pull origin develop --rebase -> ```然后拉取最新代码，理由：已经存在的分支，
                                    已经落后develop很多了，所以必须拉取代码。然后开始编写代码```
- git add . 或者 文件路径  -> ```编写完成 文件添加到索引区```
- git commit -m "描述"  -> ```增加提交到本地仓库```
- git pull origin develop --rebase  -> ```再次拉取最新代码，其他人提交了最新代码```
- git push origin 当前分支 -> ```推送到远程仓库```
- 最后在web页面向负责人提交pr。 可以自己合并



# 基于develop分支开发
- git checkout develop 
- git pull origin develop --rebase -> ```每次拉取代码```
- git add . 或者 文件路径  -> ```编写完成 文件添加到索引区```
- git commit -m "描述"  -> ```增加提交到本地仓库```
- git push origin develop


# 解决冲突
- 冲突出现点 ：
1：本地仓库做了代码修改后，然后develop也做了代码修改。 出现这个问题就是 开发时没有拉取最新的代码导致的
2：使用了stash区。先做了代码修改，然后执行了如下操作。 git stash 代码放入暂存区; git pull origin develop; 
git stash pop 从暂存区释放代码。 这种也会出现冲突。 最后就是不要使用暂存区。 需要暂存时，直接commit到本地

解决方式：

1：git status 查看文件状态 是否出现 both modified  这个标志
2: 使用编辑器，修改冲突文件，修正代码
3 git add .
4 git rebase --continue -> 出现  Applying: 这个标志代表解决冲突了
4 git push origin 分支 



# 简单开发流程 
### 第一次
- git clone https://github.com/
- git checkout develop -> ```切换到develop ```
- git checkout -b fuyi/feature/xxx -> ```然后编写代码```
- git add . 或者 文件路径  -> ```文件添加到索引区```
- git commit -m "描述"  -> ```增加提交到本地仓库```
- git pull origin develop --rebase  -> ```再次拉取最新代码，其他人提交了最新代码```
- git push origin 当前分支 -> ```推送到远程仓库```
- 最后在web页面向负责人提交pr。 可以自己合并

### 以后编写代码，就在第一次创建的分支操作
- git pull origin develop --rebase -> ```拉取最新代码，理由：已经存在的分支，
                                    已经落后develop很多了，所以必须拉取代码。然后开始编写代码```
- git add . 或者 文件路径  -> ```编写完成 文件添加到索引区```
- git commit -m "描述"    -> ```增加提交到本地仓库```
- git pull origin develop --rebase  -> ```再次拉取最新代码，其他人提交了最新代码```
- git push origin 当前分支 -> ```推送到远程仓库```
- 最后在web页面向负责人提交pr。 可以自己合并

















