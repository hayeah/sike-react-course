# Install BrowserSync

[BrowserSync](http://www.browsersync.io) is a live editing tool. When you change a file, the browser automatically refreshes. It is super helpful if you have 2 screens (you should!) so you can put the browser in one screen, and your editor in another.

Let's use npm to install BrowserSync 2.9.3:



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



+ `--save` adds browser-sync to `package.json`.
+ `-d` to make npm output more logging information.

After the command is finished, you should see that the `browser-sync` package is installed in your local project directory:



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



Run `git diff`, you can see that the `--save` option added browser-sync to the project dependencies:



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



BrowserSync is one single package, but it also depends on MANY other packages, and these packages also have their own dependencies. We can see the dependency tree by running the `npm ls` command:



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



If an NPM package is a command-line tool, the executable file is installed at the `node_modules/.bin` directory.

```
$ ls node_modules/.bin
browser-sync
```



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



### The PATH Environment Variable

To be able to type `browser-sync` instead of the full path `./node_modules/.bin/browser-sync`, we need to change the PATH environment variable so that the system can find `browser-sync`.

First, let's take a look at the current value of `PATH`:

```
$ echo $PATH
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games
```



If you type the `foo` command, the system searches for `foo` in the following order:

```
/usr/local/sbin
/usr/local/bin
/usr/sbin
/usr/bin
/sbin
/bin
/usr/games
```



To make the system search for `browser-sync` in `./node_modules/.bin`, we need to add that directory to the PATH variable:

```
$ export PATH=$PATH:./node_modules/.bin
```



Now, executing `browser-sync` should work:

```
$ browser-sync --version
2.8.2
```



To make this permanent, add to your shell's startup file:

```
# ~/.bashrc or ~/.zshrc
export PATH=$PATH:./node_modules/.bin
```



# Live-Edit With BrowserSync

Run the BrowserSync server:



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



+ `--files=index.html` specifies that if the file `index.html` changes, refresh the browser automatically.

Use the browser to open `http://localhost:3002`, and you can start editing live!

Try changing `Hello World` to `Make Everything With React!`.

`Ctrl-C` to shutdown the server when you are done.



# Makefile

For small projects, a Makefile is much simpler to get started with than Grunt/Gulp/Webpack. There are tasks you need to run frequently, like compiling JavaScript, bundling, or running BrowserSync. The command you need to run can get longer and more complicated. These tasks are perfect inside a Makefile.



### Exercise: Create the `Makefile`.



```
# Makefile
.PHONY: server
server:
  # WARNING: The indentation MUST be a tab. Spaces won't work.
  browser-sync start --server --files=index.html
```



Now you can run browser-sync this way:

```
make server
```



Usually a Makefile "rule" would create files as output. `.PHONY: server` is saying that `server` is a task, and doesn't create any file.

The Makefile tutorials out there are mostly for C/C++ projects. See: [Using GNU Make as a Front-end Development Build Tool](http://www.sitepoint.com/using-gnu-make-front-end-development-build-tool/) for a good intro.


