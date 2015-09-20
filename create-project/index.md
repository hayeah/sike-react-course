We are going to setup a frontend project from scratch. It's tedious to setup a project. But if you are new to frontend development, it's a good way to learn the basics before you start using a more complicated build tool.

There are too many different project build tools out there. Even worse, a new build tool becomes popular every other week! It's impossible to keep up.

All these build tools, though, share same basic ideas. We'll show you these basic ideas, so that no matter which build tool you eventually use (Grunt, Gulp, Webpack), you'd be able learn them quickly.

<cn>

我们将要从零开始建立一个前端项目。建立一个项目的过程很乏味。但是如果你刚接触前端开发，这是一个在你开始使用更复杂的工具前，学习基础的好方法。

市面上有很多不同的项目构建工具。更不幸的是，每隔一周，一个新的构建工具就会变得流行起来。跟上这个步伐是不可能的。

尽管如此，这些构建工具使用了想用的基本思路。我们会展示给你这些基本思路，因此不论你最终使用哪种前段构建工具（Grunt，Gulp，Webpack），你将能够快速地学习它们。

</cn>

For this project, we will:

+ Use NPM to install open-source packages.
+ Break CSS files into modules.
+ Use BrowserSync for live-editing.
+ Use Makefile to run project tasks.

Although NPM is the package management tool for NodeJS, it's also great for frontend development. Later, we'll learn how to organize JavaScript:

+ Instead of one big JS file, write smaller CommonJS modules.
+ Use `browserify` to combine JavaScript modules into a bundle.
+ Use Babel to convert ES6 to ES5, so you can run your JS everywhere.

Before you start, make sure that your NodeJS version is recent enough.

<cn>

对于这个项目，我们将：

+ 使用 NPM 安装开源包
+ 把 CSS 模块化
+ 使用 BrowserSync 实现实时编辑
+ 使用 Makefile 运行项目任务

尽管 NPM 是 NodeJS 的包管理工具，它也非常适合前端开发。以后，我们将学习如何组织 JavaScript：

+ 编写更小的 CommonJS 模块而不是一个大的 JS 文件
+ 使用 `browserify` 把 JavaScript 模块组合成一个 bundle。
+ 使用 Babel 把 ES6 转换为 ES5，这样你可以在各处运行你的 JS 了。

在你开始之前，确保你的 NodeJS 版本是最新的。

</cn>

```
$ node -v
v2.5.0
```
Versions 2, 3, 4, or above should be ok.

<cn>

```
$ node -v
v2.5.0
```
版本 2， 3， 4， 或者更高都可以.

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

[任务：使用 NPM 创建一个新项目](init)

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

[Mission: Get BrowserSync running](live-edit)

+ Use `npm install` to install open source packages (e.g. BrowserSync).
+ Run executables installed in `node_modules`.
+ Add `./node_modules/.bin` to PATH
+ Run BrowserSync for live-editing.
+ Create `Makefile`. Add the `server` task to run BrowserSync.

Resources:

+ [Using GNU Make as a Front-end Development Build Tool](http://www.sitepoint.com/using-gnu-make-front-end-development-build-tool/)

<cn>

# 使用 BrowserSync 实时编辑

[任务：让 BrowserSync 运行起来](live-edit)

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

[Mission: Configure the project CSS base](css-base)

Before we start working on the project, we should add some CSS to solve common cross-browser problems:

+ Use [PostCSS](https://github.com/postcss/postcss) to add features to standard CSS.
+ Use [autoprefixer](https://github.com/postcss/autoprefixer) to add vendor prefixes automatically.
+ Include [normalizer.css](http://necolas.github.io/normalize.css) to fix browser inconsistencies.

Furthermore, to make CSS layout easier, we'll adopt the ReactNative's flexbox settings:

+ Use the [ReactNative flexbox settings](https://github.com/facebook/css-layout#default-values) globally.

We'll talk about what these settings mean in the next lesson. Don't worry if you don't understand them yet.

Resources:

+ Check [Can I Use](http://caniuse.com) for browser compatibility.

<cn>

# 项目 CSS 基础

[任务：配置项目 CSS 基础](css-base)

在我们开始继续在这个项目工作前，我们应该添加一些 CSS 来解决常见的跨浏览器问题：

+ 使用 [PostCSS](https://github.com/postcss/postcss) 向标准 CSS 添加特性。
+ 使用 [autoprefixer](https://github.com/postcss/autoprefixer) 添加浏览器引擎前缀。
+ 引入 [normalizer.css](http://necolas.github.io/normalize.css) 来修复浏览器不一致的问题。

此外，为了让 CSS 布局更简单，我们会采用 ReactNative 的 flexbox 设置：

+ 在全局使用 [ReactNative flexbox settings](https://github.com/facebook/css-layout#default-values)。

我们会在下一课中谈到这些设置的意义。如果你还没理解它们，不用担心。

资源：

+ 为浏览器兼容性查看 [Can I Use](http://caniuse.com)。

</cn>