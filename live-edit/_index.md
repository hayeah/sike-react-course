# Install BrowserSync

[BrowserSync](http://www.browsersync.io) is a live editing tool. When you change a file, the browser automatically refreshes. It is super helpful if you have 2 screens (you should!) so you can put the browser in one screen, and your editor in another.

Let's use npm to install BrowserSync 2.9.3:

<cn>

# 安装 BrowserSync

[BrowserSync](http://www.browsersync.io) 是一个实时编辑工具。当你修改一个文件时，浏览器会自动刷新。有两个屏幕（买买买！）实在很爽，这样你可以把浏览器放在一个屏幕上，把编辑器放在另一个上。

让我们用 npm 安装 BrowserSync 2.9.3：

</cn>

```
$ npm install browser-sync@2.9.3 --save -d
npm info it worked if it ends with ok
npm info using npm@2.13.2
npm info using node@v2.5.0
npm info attempt registry request try #1 at 12:36:38
npm http request GET https://registry.npm.taobao.org/browser-sync
npm http 304 https://registry.npm.taobao.org/browser-sync
npm info install browser-sync@2.9.3 into /Users/howard/w/react
npm info installOne browser-sync@2.9.3
```

<cn>

```
$ npm install browser-sync@2.9.3 --save -d
npm info it worked if it ends with ok
npm info using npm@2.13.2
npm info using node@v2.5.0
npm info attempt registry request try #1 at 12:36:38
npm http request GET https://registry.npm.taobao.org/browser-sync
npm http 304 https://registry.npm.taobao.org/browser-sync
npm info install browser-sync@2.9.3 into /Users/howard/w/react
npm info installOne browser-sync@2.9.3
```

</cn>

+ `--save` adds browser-sync to `package.json`.
+ `-d` to make npm output more logging information.

After the command is finished, you should see that the `browser-sync` package is installed in your local project directory:

<cn>

+ `--save` 把 browser-sync 添加到 `package.json`。
+ `-d` 使 npm 输出更多的日志信息。

在这个命令完成后，你应该可以看到 `browser-sync` 包被安装在你本地项目目录中：

</cn>

```
$ find node_modules
node_modules/browser-sync/node_modules/ua-parser-js/test/test.js
node_modules/browser-sync/node_modules/ua-parser-js/ua-parser-js.jquery.json
node_modules/browser-sync/node_modules/ucfirst
node_modules/browser-sync/node_modules/ucfirst/.npmignore
node_modules/browser-sync/node_modules/ucfirst/.travis.yml
node_modules/browser-sync/node_modules/ucfirst/index.js
node_modules/browser-sync/node_modules/ucfirst/package.json
node_modules/browser-sync/node_modules/ucfirst/README.md
node_modules/browser-sync/node_modules/ucfirst/test.js
node_modules/browser-sync/package.json
node_modules/browser-sync/README.md
```

<cn>

```
$ find node_modules
node_modules/browser-sync/node_modules/ua-parser-js/test/test.js
node_modules/browser-sync/node_modules/ua-parser-js/ua-parser-js.jquery.json
node_modules/browser-sync/node_modules/ucfirst
node_modules/browser-sync/node_modules/ucfirst/.npmignore
node_modules/browser-sync/node_modules/ucfirst/.travis.yml
node_modules/browser-sync/node_modules/ucfirst/index.js
node_modules/browser-sync/node_modules/ucfirst/package.json
node_modules/browser-sync/node_modules/ucfirst/README.md
node_modules/browser-sync/node_modules/ucfirst/test.js
node_modules/browser-sync/package.json
node_modules/browser-sync/README.md
```

</cn>

Run `git diff`, you can see that the `--save` option added browser-sync to the project dependencies:

<cn>

运行 `git diff`，你能够看到 `--save` 选项把 browser-sync 添加到了项目依赖中：

</cn>

```
$ git diff
diff --git a/package.json b/package.json
index e2d56c6..97eb580 100644
--- a/package.json
+++ b/package.json
@@ -7,5 +7,8 @@
     "test": "echo \"Error: no test specified\" && exit 1"
   },
   "author": "",
-  "license": "ISC"
+  "license": "ISC",
+  "dependencies": {
+    "browser-sync": "^2.9.3"
+  }
 }
```

<cn>

```
$ git diff
diff --git a/package.json b/package.json
index e2d56c6..97eb580 100644
--- a/package.json
+++ b/package.json
@@ -7,5 +7,8 @@
     "test": "echo \"Error: no test specified\" && exit 1"
   },
   "author": "",
-  "license": "ISC"
+  "license": "ISC",
+  "dependencies": {
+    "browser-sync": "^2.9.3"
+  }
 }
```

</cn>

BrowserSync is one single package, but it also depends on MANY other packages, and these packages also have their own dependencies. We can see the dependency tree by running the `npm ls` command:

<cn>

BrowserSync 是一个单独的包，但是它也依赖于 <u>很多</u> 其他的包，而且这些包也有它们自己的依赖。我们可以通过运行 `npm ls` 看到依赖树：

</cn>

```
$ npm ls
ilovereact@1.0.0
└─┬ browser-sync@2.9.3
  ├─┬ anymatch@1.3.0
  │ ├── arrify@1.0.0
  │ └─┬ micromatch@2.2.0
  │   ├─┬ arr-diff@1.1.0
  │   │ ├── arr-flatten@1.0.1
  │   │ └── array-slice@0.2.3
...
  │ │ │ └── xmlhttprequest@1.5.0
  │ │ ├─┬ has-binary@0.1.6
  │ │ │ └── isarray@0.0.1
  │ │ ├── indexof@0.0.1
  │ │ ├── object-component@0.0.3
  │ │ ├─┬ parseuri@0.0.2
  │ │ │ └─┬ better-assert@1.0.2
  │ │ │   └── callsite@1.0.0
  │ │ └── to-array@0.1.3
  │ └─┬ socket.io-parser@2.2.4
  │   ├── benchmark@1.0.0
  │   ├── component-emitter@1.1.2
  │   ├── debug@0.7.4
  │   ├── isarray@0.0.1
  │   └── json3@3.2.6
  ├── ua-parser-js@0.7.9
  └── ucfirst@0.0.1
```

<cn>

```
$ npm ls
ilovereact@1.0.0
└─┬ browser-sync@2.9.3
  ├─┬ anymatch@1.3.0
  │ ├── arrify@1.0.0
  │ └─┬ micromatch@2.2.0
  │   ├─┬ arr-diff@1.1.0
  │   │ ├── arr-flatten@1.0.1
  │   │ └── array-slice@0.2.3
...
  │ │ │ └── xmlhttprequest@1.5.0
  │ │ ├─┬ has-binary@0.1.6
  │ │ │ └── isarray@0.0.1
  │ │ ├── indexof@0.0.1
  │ │ ├── object-component@0.0.3
  │ │ ├─┬ parseuri@0.0.2
  │ │ │ └─┬ better-assert@1.0.2
  │ │ │   └── callsite@1.0.0
  │ │ └── to-array@0.1.3
  │ └─┬ socket.io-parser@2.2.4
  │   ├── benchmark@1.0.0
  │   ├── component-emitter@1.1.2
  │   ├── debug@0.7.4
  │   ├── isarray@0.0.1
  │   └── json3@3.2.6
  ├── ua-parser-js@0.7.9
  └── ucfirst@0.0.1
```

</cn>

If an NPM package is a command-line tool, the executable file is installed at the `node_modules/.bin` directory.

```
$ ls node_modules/.bin
browser-sync
```

<cn>

假如某个 NPM 包是一个命令行工具，可执行文件会被安装在 `node_modules/.bin` 这个目录。

```
$ ls node_modules/.bin
browser-sync
```

</cn>

We can run the executable with `--help` option to see the help message:

```
$ ./node_modules/.bin/browser-sync --help
Live CSS Reload & Browser Syncing

  Usage:
  ---------

      $ browser-sync <command> [options]

  Commands:
  ---------

      init    Creates a default config file
      start   Start Browser Sync
      reload  Send a reload event over HTTP protocol
...
```

<cn>

执行命令时我们可以加上 `--help` 选项来查看帮助信息：

```
$ ./node_modules/.bin/browser-sync --help
Live CSS Reload & Browser Syncing

  Usage:
  ---------

      $ browser-sync <command> [options]

  Commands:
  ---------

      init    Creates a default config file
      start   Start Browser Sync
      reload  Send a reload event over HTTP protocol
...
```

</cn>

### The PATH Environment Variable

To be able to type `browser-sync` instead of the full path `./node_modules/.bin/browser-sync`, we need to change the PATH environment variable so that the system can find `browser-sync`.

First, let's take a look at the current value of `PATH`:

```
$ echo $PATH
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games
```

<cn>

### PATH 环境变量

为了可以用 `browser-sync` 来执行命令，而不是完整路径 `./node_modules/.bin/browser-sync`，我们需要修改 PATH 环境变量。这样系统才能找到 `browser-sync` 这个命令。

首先，让我们看看 `PATH` 的当前值：

```
$ echo $PATH
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games
```

</cn>

If you type the `foo` command, the system searches `foo` the following directories one after another:

```
/usr/local/sbin
/usr/local/bin
/usr/sbin
/usr/bin
/sbin
/bin
/usr/games
```

<cn>

如果你输入 `foo` 命令，系统会按顺序在以下目录里查找 `foo`：

```
/usr/local/sbin
/usr/local/bin
/usr/sbin
/usr/bin
/sbin
/bin
/usr/games
```

</cn>

To make the system search for `browser-sync` in `./node_modules/.bin`, we need to add that directory to the PATH variable:

```
$ export PATH=$PATH:./node_modules/.bin
```

<cn>

为了让系统在 `./node_modules/.bin` 中查找 `browser-sync`，我们需要把那个目录添加到 PATH 环境变量中：

```
$ export PATH=$PATH:./node_modules/.bin
```

</cn>

Now, executing `browser-sync` should work:

```
$ browser-sync --version
2.8.2
```

<cn>

现在，执行 `browser-sync` 应该可以工作了：

```
$ browser-sync --version
2.8.2
```

</cn>

To make the PATH permanent, add to your shell's startup file:

```
# ~/.bashrc or ~/.zshrc
export PATH=$PATH:./node_modules/.bin
```

<cn>

把 PATH 设定持久化，添加到你的命令行的启动文件中：

```
# ~/.bashrc or ~/.zshrc
export PATH=$PATH:./node_modules/.bin
```

</cn>

# Live-Edit With BrowserSync

Run the BrowserSync server:

<cn>

# 使用 BrowserSync 实时编辑

运行 BrowserSync 服务器：

</cn>

```
$ browser-sync start --server --files=index.html
[BS] Access URLs:
 --------------------------------------
       Local: http://localhost:3002
    External: http://192.168.5.106:3002
 --------------------------------------
          UI: http://localhost:3003
 UI External: http://192.168.5.106:3003
 --------------------------------------
[BS] Serving files from: ./
[BS] Watching files...
[BS] File changed: index.html
```

<cn>

```
$ browser-sync start --server --files=index.html
[BS] Access URLs:
 --------------------------------------
       Local: http://localhost:3002
    External: http://192.168.5.106:3002
 --------------------------------------
          UI: http://localhost:3003
 UI External: http://192.168.5.106:3003
 --------------------------------------
[BS] Serving files from: ./
[BS] Watching files...
[BS] File changed: index.html
```

</cn>

+ `--files=index.html` monitors `index.html` for changes, and refreshes the browser automatically.

Use the browser to open `http://localhost:3002`, and you can start editing live!

Try changing `Hello World` to `Make Everything With React!`.

`Ctrl-C` to shutdown the server when you are done.

<cn>

+ `--files=index.html` 监听 `index.html`，一旦改变了会自动刷新浏览器。

使用浏览器打开 `http://localhost:3002`，你可以开始实时编辑了！

尝试把 `Hello World` 改为 `Make Everything With React!`。

使用完毕后用 `Ctrl-C` 来关闭服务器。

</cn>

# Makefile

For small projects, a Makefile is much simpler to get started with than Grunt/Gulp/Webpack. There are tasks you need to run frequently, like compiling JavaScript, bundling, or running BrowserSync. The commands for the tasks you need to run can be long and complicated. Put these tasks in the Makefile can simplify your life.

<cn>

# Makefile

对于小项目来说，Makefile 比 Grunt/Gulp/Webpack 容易上手多了。你会有些经常需要运行的任务，比如编译 JavaScript，打包，或者运行 BrowserSync。这些任务的命令可能更长更复杂，把它们放在 Makefile 里会很方便。

</cn>

### Exercise: Create the `Makefile`.

Create a file called `Makefile` at the root of your project directory. We define the task `server` like this:

```
.PHONY: server
server:
  browser-sync start --server --files=index.html
```

Note that the space in front of the command MUST BE A TAB!

```
.PHONY: server
server:
<tab>browser-sync start --server --files=index.html
```

Now you can use make to run  browser-sync:

```
make server
```

<cn>

### 练习：创建 `Makefile`。

在项目的根目录创建一个命名为 `Makefile` 的文件。我们可以这样定义 `server` 这个任务：

```
.PHONY: server
server:
  browser-sync start --server --files=index.html
```

注意，命令之前的空白<u>必须是 TAB</u>

```
.PHONY: server
server:
<tab>browser-sync start --server --files=index.html
```

现在你可以用 make 运行 browser-sync：

```
make server
```

Note: If you use spaces instead of tab, you'd get this error,

```
make server
Makefile:3: *** missing separator.  Stop.
```

### Phony Target

</cn>

Usually a Makefile "rule" would create files as output. `.PHONY: server` is saying that `server` is a task, and doesn't create any file.

The Makefile tutorials out there are mostly for C/C++ projects. See: [Using GNU Make as a Front-end Development Build Tool](http://www.sitepoint.com/using-gnu-make-front-end-development-build-tool/) for a good intro.

<cn>

一个 Makefile “规则” 通常会创建作为输出的文件。`.PHONY: server` 是说 `server` 是一个任务，并不会输出文件。

市面上的 Makefile 教程几乎都是关于 C/C++ 项目的。查看一个不错的介绍：[Using GNU Make as a Front-end Development Build Tool](http://www.sitepoint.com/using-gnu-make-front-end-development-build-tool/)
</cn>
