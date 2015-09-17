We are going to setup a frontend project from scratch. Usually it's simpler to use a boilerplate, but if you are new to frontend, it's good to go through the process of setting up a project, so you know what the useful tools are.

There are too many different project build tools out there. Even worse, a new build tool becomes popular every other week! It's impossible to keep up.

All these build tools, though, share same basic ideas. We'll show you these basic ideas, so no matter what build tool you have to use (Grunt, Gulp, Webpack), you'd can learn to use them quickly.

We'll learn to:

+ Use NPM to install open-source packages.
+ Break CSS files into modules.
+ Use BrowserSync for live-editing.
+ Use Makefile to run project tasks.

Later, we'll learn how to organize JavaScript:

+ Instead of one big JS file, write smaller CommonJS modules.
+ Use `browserify` to bundle JavaScript modules.
+ Use Babel to convert ES6 to ES5, so you can run your JS everywhere.

Although NPM is the package management tool for NodeJS, it's also great for frontend development.

Before you start, make sure that your NodeJS version is recent enough.

```
$ node -v
v2.5.0
```
Versions 2, 3, 4, or above should be ok.

# Create A New Project

[Mission: Use NPM to create a new project](init)

+ Use `npm init` to create `ilovereact` project.
+ Create a new git repo.
+ Push the git repo to GitHub.

Resources:

+ [github/gitignore](https://github.com/github/gitignore)
+ [MarkDown Syntax](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)

# Live Edit With BrowserSync

[Mission: Get BrowserSync running](live-edit)

+ Use `npm install` to install open source packages (e.g. BrowserSync).
+ Run executables installed in `node_modules`.
+ Add `./node_modules/.bin` to PATH
+ Run BrowserSync for live-editing.
+ Create `Makefile`. Add a task to run BrowserSync.

Resources:

+ [Using GNU Make as a Front-end Development Build Tool](http://www.sitepoint.com/using-gnu-make-front-end-development-build-tool/) for a good intro.

<zh>
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
</zh>

# The Project CSS Base

[Mission: Configure the project CSS base](css-base)

Before we start working on the project, we should add some CSS to solve common cross-browser problems:

+ Include [normalizer.css](http://necolas.github.io/normalize.css) to fix browser inconsistencies.
+ Use [autoprefixer](https://github.com/postcss/autoprefixer) to add vendor prefixes automatically.

Furthermore, to make CSS layout easier, we'll adopt the ReactNative's flexbox settings:

+ Use the [ReactNative flexbox settings](https://github.com/facebook/css-layout#default-values) globally.

We'll talk about what these settings mean in the next lesson. Don't worry if you don't understand what they mean yet.

Resources:

+ Always check [Can I Use](http://caniuse.com) for browser compatibility.
