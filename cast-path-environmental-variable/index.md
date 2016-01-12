# PATH 环境变量的原理

找不到命令？安装了还是无法运行命令？

<video src="intro.mp4" controls="true" preload="none"></video>

### 扫盲知识点

+ 什么是可执行文件？怎么运行可执行文件？
+ 怎么创建新的可执行文件？
  + 有没有可执行权限 (x) 有吗？
    + 用 ls -la 查看
    + 用 `chown u+x <file>`  修改
+ 运行命令时 Unix 怎么找到可执行文件？
  + PATH 环境变量配置好了吗？
+ 如何在终端启动时配置 PATH？


# 什么是可执行文件

<video src="executable-files.mp4" controls="true" preload="none"></video>

所有可执行文件都是文件。输入命令后系统会去查找一个匹配该命令的可执行文件。

+ `which ping`
+ `less -R /sbin/ping`
+ `which sudo`
+ `which git`

# 怎么创建新的可执行文件？

<video src="create-a-new-executable.mp4" controls="true" preload="none"></video>

+ 创建 hello 执行文件

  ```
  while true;
    do echo hello world
    sleep 1
  done
  ```

+ 用 ls -la 查看权限。
  + `rwxr-xr-x   1 howard  staff  2660288 20 Sep 22:26 foo`
  + `r` - 读取权限
  + `w` - 写入权限
  + `x` - 执行权限
  + `howard` - 拥有者
+ 用 `chmod u+x hello` 添加执行权限。

# 运行命令时 Unix 怎么找到可执行文件？

<video src="modifying-path.mp4" controls="true" preload="none"></video>

+ `echo $PATH` 查看所有运行文件的搜索目录。
  + 系统会用命令的名字一个一个目录去查找匹配执行文件。
+ `export PATH=$PATH:/a/b/c` 添加目录到 PATH。
  + 更改 PATH 只影响当前的终端

## 环境变量继承的原理

<video src="environmental-variable-inheritance.mp4" controls="true" preload="none"></video>

+ 系统里所有进程是由父子关系串在一起
+ 改变环境变量会影响子进程，但不会影响兄弟或者父进程


# 如何在终端启动时配置 PATH？

<video src="edit-and-loading-zshrc.mp4" controls="true" preload="none"></video>

+ ~ 是 HOME 目录
+ 配置文件大都是 `.` 开头, 影藏起来了。
  + 用 `ls -a` 查看影藏文件。
+ 修改 Shell 启动文件，在启动时配置 PATH
  + bash: 修改 ~/.bashrc
    + 还要确认 ~/.bash_profile 这个文件有 `source ~/.bashrc` 这段。
  + zsh: 修改 ~/.zshrc
+ 可以手动 `source ~/.zshrc` 去运行启动文件
  + 开新终端会自动运行



# 添加路径在 PATH 最后

<video src="edit-and-loading-zshrc.mp4" controls="true" preload="none"></video>



