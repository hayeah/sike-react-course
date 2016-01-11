# Create A New Project With NPM

First, create a project directory called `ilovereact`:

```
$ mkdir ilovereact
$ cd ilovereact
```

<cn>

# 用 NPM 创建一个新的项目

首先，创建一个名为 `ilovereact` 的目录：

```
$ mkdir ilovereact
$ cd ilovereact
```

</cn>

Create the project:

<cn>

创建项目：

</cn>

```
$ npm init
This utility will walk you through creating a package.json file.
It only covers the most common items, and tries to guess sensible defaults.

See `npm help json` for definitive documentation on these fields
and exactly what they do.

Use `npm install <pkg> --save` afterwards to install a package and
save it as a dependency in the package.json file.

Press ^C at any time to quit.
name: (ilovereact)
version: (1.0.0)
description:
entry point: (index.js)
test command:
git repository:
keywords:
author:
license: (ISC)
About to write to /Users/howard/w/react/ilovereact/package.json:

{
  "name": "ilovereact",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC"
}


Is this ok? (yes) yes
```

<cn>

```
$ npm init
This utility will walk you through creating a package.json file.
It only covers the most common items, and tries to guess sensible defaults.

See `npm help json` for definitive documentation on these fields
and exactly what they do.

Use `npm install <pkg> --save` afterwards to install a package and
save it as a dependency in the package.json file.

Press ^C at any time to quit.
name: (ilovereact)
version: (1.0.0)
description:
entry point: (index.js)
test command:
git repository:
keywords:
author:
license: (ISC)
About to write to /Users/howard/w/react/ilovereact/package.json:

{
  "name": "ilovereact",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC"
}


Is this ok? (yes) yes
```

</cn>

Having run the above command, you should see `package.json` like this:

<cn>

运行上述命令后，你应该可以看到 `package.json` 大致是这样：

</cn>

```
$ cat package.json
{
  "name": "ilovereact",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC"
}
```

<cn>

```
$ cat package.json
{
  "name": "ilovereact",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC"
}
```

</cn>

# Initialize Git Repository

We don't want junks like log or installed packages to bloat our git repository. So let's create a `.gitignore` file to prevent anyone from adding unnecessary files into the repo.

We can get one suitable for a NodeJS project from gitignore.io:

```
curl https://www.gitignore.io/api/node > .gitignore
```

<cn>

# 初始化 Git 仓库

我们不希望把日志或者依赖包之类的垃圾被添加到 git 仓库里，导致膨胀。我们使用 `.gitignore` 文件来阻止任何人向仓库添加不必要的文件。

我们可以从 gitignore.io 下载一个适合 NodeJS 项目的文件：

```
curl https://raw.githubusercontent.com/github/gitignore/master/Node.gitignore > .gitignore
```

</cn>

Then we create the git repo:

```
$ git init
Initialized empty Git repository in ilovereact/.git/
$ git add .
$ git commit -m "Project init"
[master (root-commit) d7a71e7] Project init
 1 file changed, 11 insertions(+)
 create mode 100644 package.json
```

<cn>

然后我们创建 git 仓库：

```
$ git init
Initialized empty Git repository in ilovereact/.git/
$ git add .
$ git commit -m "Project init"
[master (root-commit) d7a71e7] Project init
 1 file changed, 11 insertions(+)
 create mode 100644 package.json
```

</cn>

You can see the commit you've just created:

```
$ git show HEAD
commit d7a71e7d7b8b08d3c09a1d146625502b1f45a3e7
Author: Howard Yeh <howard@metacircus.com>
Date:   Tue Sep 15 11:43:50 2015 +0800

    Project init

diff --git a/package.json b/package.json
new file mode 100644
index 0000000..e2d56c6
--- /dev/null
+++ b/package.json
@@ -0,0 +1,11 @@
+{
+  "name": "ilovereact",
+  "version": "1.0.0",
+  "description": "",
+  "main": "index.js",
+  "scripts": {
+    "test": "echo \"Error: no test specified\" && exit 1"
+  },
+  "author": "",
+  "license": "ISC"
+}
```

<cn>

你可以看到你刚创建的提交：

```
$ git show HEAD
commit d7a71e7d7b8b08d3c09a1d146625502b1f45a3e7
Author: Howard Yeh <howard@metacircus.com>
Date:   Tue Sep 15 11:43:50 2015 +0800

    Project init

diff --git a/package.json b/package.json
new file mode 100644
index 0000000..e2d56c6
--- /dev/null
+++ b/package.json
@@ -0,0 +1,11 @@
+{
+  "name": "ilovereact",
+  "version": "1.0.0",
+  "description": "",
+  "main": "index.js",
+  "scripts": {
+    "test": "echo \"Error: no test specified\" && exit 1"
+  },
+  "author": "",
+  "license": "ISC"
+}
```

</cn>

# HTML Boilerplate

Let's create `index.html`. Rather than starting from scratch, it's faster to tweak [HTML5-boilerplate](https://github.com/h5bp/html5-boilerplate/blob/master/src/index.html) to suit your needs.

For this project you could use a template like this:

<cn>

# HTML 样板文件

让我们来创建 `index.html`。 与其从零开始，按需求来调整 [HTML5 常用模板页面](https://github.com/h5bp/html5-boilerplate/blob/master/src/index.html) 会更快。

对于这个项目，你可以使用类似这样的模板：

</cn>

```html
<!-- Uses html5 syntax -->
<!doctype html>

<html lang="en">
<head>
    <meta charset="utf-8">

    <!-- Forces IE to follow modern standards -->
    <meta http-equiv="x-ua-compatible" content="ie=edge">

    <title></title>
    <meta name="description" content="">

    <!-- Disable zooming on mobile. Useful for responsive design. -->
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- <link rel="apple-touch-icon" href="apple-touch-icon.png"> -->

    <!-- <link rel="stylesheet" href="css/app.css"> -->

</head>
<body>


<!-- <script src="js/app.js"></script> -->

</body>
</html>
```

<cn>

```html
<!-- 使用 html5 语法 -->
<!doctype html>

<html lang="en">
<head>
    <meta charset="utf-8">

    <!-- 强制 IE 遵循现代标准 -->
    <meta http-equiv="x-ua-compatible" content="ie=edge">

    <title></title>
    <meta name="description" content="">

    <!-- 在移动端禁用缩放。对响应式设计有用。 -->
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- <link rel="apple-touch-icon" href="apple-touch-icon.png"> -->

    <!-- <link rel="stylesheet" href="css/app.css"> -->

</head>
<body>


<!-- <script src="js/app.js"></script> -->

</body>
</html>
```

</cn>

### Exercise: I Love React

Add an h1 element to `index.html`

```html
<h1>I Love React</h1>
```

<cn>

### 练习：我爱 React

把 h1 元素添加到 `index.html`

```html
<h1>I Love React</h1>
```

</cn>

# Publish To GitHub

Let's push this project to GitHub so everyone can see it. My username is `hayeah`. You'll need to use your own username for the examples below.

First, go on GitHub to [create a new repository](https://github.com/new). Call it `sikeio-ilovereact`. Then add the remote repo:

<cn>

# 发布到 GitHub

让我们推送这个项目到 GitHub 上，这样每个人都能看到它。我的用户名是 `hayeah`。下面的命令你需要改用自己的 GitHub 用户名。

首先，到 GitHub 上去 [创建一个新项目](https://github.com/new), 命名为 `sikeio-ilovereact`. 接下来，添加远程仓库：

</cn>

```
$ git remote add origin git@github.com:hayeah/sikeio-ilovereact.git
$ git push origin master -u
Counting objects: 13, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (12/12), done.
Writing objects: 100% (13/13), 1.69 KiB | 0 bytes/s, done.
Total 13 (delta 2), reused 0 (delta 0)
To git@github.com:hayeah/sikeio-ilovereact.git
 * [new branch]      master -> master
Branch master set up to track remote branch master from origin.
```

<cn>

```
$ git remote add origin git@github.com:hayeah/sikeio-ilovereact.git
$ git push origin master -u
Counting objects: 13, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (12/12), done.
Writing objects: 100% (13/13), 1.69 KiB | 0 bytes/s, done.
Total 13 (delta 2), reused 0 (delta 0)
To git@github.com:hayeah/sikeio-ilovereact.git
 * [new branch]      master -> master
Branch master set up to track remote branch master from origin.
```

</cn>

And you can use GitHub Pages to host this web-page. All you need to do is to push to the branch `gh-pages`:

<cn>

你还可以使用 GitHub Pages 来托管这个网页。你只需要把页面推送到 `gh-pages` 这个分支：

</cn>

```
$ git push origin master:gh-pages
Total 0 (delta 0), reused 0 (delta 0)
To git@github.com:hayeah/sikeio-ilovereact.git
 * [new branch]      master -> gh-pages
```

<cn>

```
$ git push origin master:gh-pages
Total 0 (delta 0), reused 0 (delta 0)
To git@github.com:hayeah/sikeio-ilovereact.git
 * [new branch]      master -> gh-pages
```

</cn>

If it's successful, you should be able to see the page at: http://hayeah.github.io/sikeio-ilovereact/

<cn>

成功的话，你应该可以打开这个页面：http://hayeah.github.io/sikeio-ilovereact/

</cn>