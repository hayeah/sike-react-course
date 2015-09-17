# Create A New Project With NPM

First, create a project directory called `ilovereact`:

```
$ mkdir ilovereact
$ cd ilovereact
```

Create the project:

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

Having run the above command, you should see `package.json` like this:

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

# Initialize Git Repository

We don't want junks like log or installed packages to bloat our git repository. So let's use a `.gitignore` file to prevent anyone from adding unnecessary files into the repo.

You can find common project-specific `.gitignore` files in the repo [github/gitignore](https://github.com/github/gitignore). Some examples are:

+ [ObjectiveC .gitignore](https://github.com/github/gitignore/blob/master/Objective-C.gitignore)
+ [Rails .gitignore](https://github.com/github/gitignore/blob/master/Rails.gitignore)
+ [Node .gitignore](https://github.com/github/gitignore/blob/master/Node.gitignore)

We'll use NodeJS specific: `.gitignore`. Download it to your project directory:

```
curl https://github.com/github/gitignore/blob/master/Node.gitignore > .gitignore
```

Then we can create the repo:

```
$ git init
Initialized empty Git repository in ilovereact/.git/
$ git add *
$ git commit -m "Project init"
[master (root-commit) d7a71e7] Project init
 1 file changed, 11 insertions(+)
 create mode 100644 package.json
```

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

# HTML Boilerplate

Let's create `index.html`. Rather than starting from scratch, it's faster to tweak [HTML5-boilerplate](https://github.com/h5bp/html5-boilerplate/blob/master/src/index.html) to suit your needs.

For this project you could use something like this:


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

### Exercise: I Love React

Add an h1 element to `index.html`

```html
<h1>I Love React</h1>
```

# Publish To GitHub

Let's push this project to GitHub so everyone can see it.

My username is `hayeah`. You'll need to change it to yourown.

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

And you can use GitHub Pages to host this web-page. All you need to do is to push to the branch `gh-pages`:

```
$ git push origin master:gh-pages
Total 0 (delta 0), reused 0 (delta 0)
To git@github.com:hayeah/sikeio-ilovereact.git
 * [new branch]      master -> gh-pages
```

Then you should be able to see the page at: http://hayeah.github.io/sikeio-ilovereact/