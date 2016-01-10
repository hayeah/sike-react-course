# Setup A Frontend Project

We are going to setup a frontend project from scratch. Although it is boring to setup a project, it's a good way to learn the common tools and structure of a frontend project.

There are too many different project build tools out there. Even worse, a new build tool becomes popular every other week! It's impossible to keep up.

All these build tools, though, share same basic ideas. We'll show you these basic ideas, so that no matter which build tool you eventually use (Grunt, Gulp, Webpack), you'd be able learn them quickly.

<cn>

# 初始化一个前端项目

我们将要从零开始建立一个前端项目。

虽然建立一个项目的过程很乏味，但这是一个学习基础的好方法。你可以学到前端项目的基本结构，也能接触到常见的工具。

市面上有太多的前端项目构建工具。然而非常不幸地，每周貌似都有个新的工具流行起来。要跟上这个步伐是不可能的。

但这些构建工具都有相同的基本思路。我们会教你你这些基本思路，所以不论你最终使用哪种前段构建工具（Grunt，Gulp，Webpack），你将能够快速地上手。

</cn>

For this project, we will:

+ Use NPM to install open-source packages.
+ Break CSS files into modules.
+ Use BrowserSync for live-editing.
+ Use Makefile to run project tasks.

Although NPM is the package management tool for NodeJS, it's also great for frontend development.

This week we won't need to use too much JavaScript. Next week, we'll learn more advanced ways to organize JavaScript:

+ Instead of one big JS file, write smaller CommonJS modules.
+ Use `browserify` to combine JavaScript modules into a bundle.
+ Use Babel to convert ES6 to ES5, so you can run your JS everywhere.

Before you start, make sure that your NodeJS version is recent enough.

<cn>

在这个项目我们会：

+ 使用 NPM 安装开源包
+ 把 CSS 模块化
+ 使用 BrowserSync 实现实时编辑
+ 使用 Makefile 运行项目任务

尽管 NPM 是 NodeJS 的包管理工具，它也非常适合前端开发。

这一周我们对 JavaScript 的需求不多。下一周我们才会开始学习更进阶的 JavaScript 组织方式：

+ 编写更小的 CommonJS 模块而不是一个大的 JS 文件
+ 使用 `webpack` 把 JavaScript 模块组合成一个 bundle。
+ 使用 Babel 把 ES6 转换为 ES5，这样你可以在各处运行你的 JS 了。

在你开始之前，请确保你的 NodeJS 版本是最新的。

</cn>

```
$ node -v
v4.2.1
```
Versions 2, 3, 4, or above should be ok.

<cn>

```
$ node -v
v4.2.1
```
版本 2， 3， 4， 或更高都可以.

</cn>

# Create A New Project

[Mission: Use NPM to create a new project](../project-init)

+ Use `npm init` to create `ilovereact` project.
+ Create a new git repo.
+ Use .gitignore to avoid adding junk into the repo.
+ Push the git repo to GitHub.
+ Use GitHub pages to host static websites.

Resources:

+ [github/gitignore](https://github.com/github/gitignore)
+ [MarkDown Syntax](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)

<cn>

# 创建一个新项目

[任务：使用 NPM 创建一个新项目](../project-init/?lang=cn)

+ 使用 `npm init` 创建 `ilovereact` 项目。
+ 创建一个新的 git 仓库。
+ 使用 .gitignore 避免把垃圾文件添加进仓库。
+ 推送 git 仓库到 GitHub 上。
+ 使用 GitHub pages 来托管静态网站。

资源：

+ [github/gitignore](https://github.com/github/gitignore)
+ [MarkDown Syntax](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)

</cn>

# Live Edit With BrowserSync

[Mission: Get BrowserSync running](../live-edit)

+ Use `npm install` to install open source packages (e.g. BrowserSync).
+ Run executables installed in `node_modules`.
+ Add `./node_modules/.bin` to PATH
+ Run BrowserSync for live-editing.
+ Create `Makefile`. Add the `server` task to run BrowserSync.

Resources:

+ [Using GNU Make as a Front-end Development Build Tool](http://www.sitepoint.com/using-gnu-make-front-end-development-build-tool/)

<cn>

# 使用 BrowserSync 实时编辑

[任务：让 BrowserSync 运行起来](../live-edit/?lang=cn)

+ 使用 `npm install` 安装开源包（比如 BrowserSync）。
+ 运行安装在 `node_modules` 内的可执行文件。
+ 把 Add `./node_modules/.bin` 添加到 PATH 中。
+ 运行 BrowserSync 实现实时编辑
+ 创建 `Makefile`。添加 `server` 任务以运行 BrowserSync。

</cn>

<cn>
### NPM 镜像

因为国内特别的网路环境，梯子不给力请用淘宝 NPM 镜像。

临时使用:

```
npm --registry https://registry.npm.taobao.org install express
```

持久使用:

```
npm config set registry https://registry.npm.taobao.org

// 配置后可通过下面方式来验证是否成功
npm config get registry
```
</cn>

# The Project CSS Base

[Mission: Configure the project CSS base](../css-base)

Before we start working on the project, we should add some CSS to solve common cross-browser problems:

+ Use [PostCSS](https://github.com/postcss/postcss) to add features to standard CSS.
+ Use [autoprefixer](https://github.com/postcss/autoprefixer) to automatically add vendor prefixes to properties.
+ Include [normalizer.css](http://necolas.github.io/normalize.css) to minimize browser inconsistencies.

Furthermore, to make CSS layout easier, we'll adopt the ReactNative's flexbox settings:

+ Use the [ReactNative flexbox settings](https://github.com/facebook/css-layout#default-values) globally.

We'll talk about what these settings mean in the next lesson. Don't worry if you don't understand them yet.

Resources:

+ Check [Can I Use](http://caniuse.com) for browser support.

<cn>

# 项目 CSS 基础

[任务：配置项目 CSS 基础](../css-base/?lang=cn)

在这个项目开工之前，我们应该添加一些 CSS 来解决常见的跨浏览器问题：

+ 使用 [PostCSS](https://github.com/postcss/postcss) 为标准 CSS 添加一些新功能。
+ 使用 [autoprefixer](https://github.com/postcss/autoprefixer) 自动为属性加上浏览器前缀，提高兼容性。
+ 引入 [normalizer.css](http://necolas.github.io/normalize.css) 来减少浏览器之间的不一致。

此外，为了让 CSS 布局更简单，我们会采用 ReactNative 的 flexbox 设置：

+ 在全局使用 [ReactNative flexbox settings](https://github.com/facebook/css-layout#default-values)。

我们会在下一课中谈到这些设置的意义。如果你还没理解这些设置，别担心。

资源：

+ 查看某个功能的浏览器支持 [Can I Use](http://caniuse.com)。

</cn>
