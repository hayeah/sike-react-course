# Setup A Frontend Project

We are going to setup a frontend project from scratch. Although it is boring to setup a project, it's a good way to learn the common tools and structure of a frontend project.

There are too many different project build tools out there. Even worse, a new build tool becomes popular every other week! It's impossible to keep up.

All these build tools, though, share same basic ideas. We'll show you these basic ideas, so that no matter which build tool you eventually use (Grunt, Gulp, Webpack), you'd be able learn them quickly.



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



```
$ node -v
v4.2.1
```
Versions 2, 3, 4, or above should be ok.



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



# Live Edit With BrowserSync

[Mission: Get BrowserSync running](../live-edit)

+ Use `npm install` to install open source packages (e.g. BrowserSync).
+ Run executables installed in `node_modules`.
+ Add `./node_modules/.bin` to PATH
+ Run BrowserSync for live-editing.
+ Create `Makefile`. Add the `server` task to run BrowserSync.

Resources:

+ [Using GNU Make as a Front-end Development Build Tool](http://www.sitepoint.com/using-gnu-make-front-end-development-build-tool/)





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


