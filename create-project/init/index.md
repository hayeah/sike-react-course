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

# Add README.md

[MarkDown](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)

# Publish To GitHub

Let's push this new project to GitHub.


