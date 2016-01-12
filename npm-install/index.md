# npm install

npm install 把包都安装到那里去了？怎么执行 npm 安装的命令行工具？

<video src="intro.mp4" controls="true" preload="none"></video>

## 知识依赖

+ 知道 PATH 是什么。
+ 知道怎么改 PATH，添加查找路径。
+ 知道怎么配置终端初始化文件 ~/.bashrc 或者 ~/.zshrc。

不清楚的话请看：[PATH 环境变量的原理](../cast-path-environmental-variable)

## 扫盲知识点

+ npm 项目结构。
+ 安装 package.json 里的依赖
  + 不要用 sudo。
+ 本地安装的 npm 包在那里？
  + node_modules 嵌套结构。
  + require 用法。
+ npm 把可执行文件安装在那里？
  + 怎么执行？


# NPM 项目结构

<video src="npm-project-structure.mp4" controls="true" preload="none"></video>

# 本地安装

+ `npm install` 安装 package.json 里的依赖
  + 不要用 sudo

## 安装项目依赖

<video src="npm-install-package.mp4" controls="true" preload="none"></video>



## 引用 npm 依赖

<video src="require-lodash.mp4" controls="true" preload="none"></video>

```
> require("postcss")
'/Users/howard/casts/npm-install/node_modules/postcss/lib/postcss.js'
> require.resolve("lodash")
'/Users/howard/casts/npm-install/node_modules/lodash/index.js'
```

## 添加新的依赖

add-new-dependencies.mp4

+ `npm install lodash --save` 把新的依赖添加到 package.json。
+ --save-dev 添加开发依赖。

# 可执行文件安装在那里？

<video src="node_modules-bin.mp4" controls="true" preload="none"></video>

+ `export PATH=$PATH:./node_modules/.bin`
+ `./node_modules/.bin` 是一个相对路径。每个 npm 项目都有自己的 .bin 目录。

# 全局安装的包在那里？

全局安装的姿势：

+ `npm install -g webpack`

但不要用吧~
