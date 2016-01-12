# The PATH Environmental Variable

### Questions:

找不到命令？安装了还是无法运行命令？

+ What are executables?
+ How do you create a new executable?
+ How can you execute an executable?
+ Where does Unix look for executables when you run a command?

### 问题:

+ 什么是可执行文件？
+ 怎么运行可执行文件？
+ 什么创建新的可执行文件？
  + 可执行权限 (x) 有吗？
    + 用 ls -la 查看
    + 用 `chown u+x <file>`  修改
+ 运行命令时 Unix 怎么找到可执行文件？
  + PATH 环境变量配置好了吗？

### Outline:

+ All executables are files.
  + When you enter a command, Unix tries to find a matching executable file.
  + which ping
  + less -R /sbin/ping

  + `which sudo`
  + `which git`

+ Create a new shell script called: `hello`
  + exec in current directory

```
while true;
  do echo hello world
  sleep 1
done
```

+ The executable searching process
  + echo $PATH

+ Appending new paths to PATH
  + Different languages put their executables in different directories.

+ Changing PATH only affects the current shell.
  + Environmental variable inheritance

+ Make all new shell processes inherit a different PATH value.
  + man zsh
  + Configuring .bashrc/.zshrc
    + `.` in front means its a hidden file.
  + source ~/.zshrc
    + debug with echo

+ Why add other to the end of the PATH, not in front.