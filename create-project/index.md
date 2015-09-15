Usually, it's faster to start with a boiler plate project setup. But we are going to start from scratch, and configure a simple frontend project. Throughout this project, we'll incrementally add more features to our setup.

There are too many different project build tools out there. Even worse, there's a new build tool that comes out every week. It's impossible to keep up. All these build tools, though, share same basic ideas.

We'll show you these basic ideas, so no matter what build tool you use (Grunt, Gulp, Webpack, Duo.js), you'd have some intuition of how they work. Specifically, we will learn to use the following tools:

+ Use NPM to install open-source packages.
+ Break CSS files into modules.
+ Use BrowserSync for live-editing.
+ Use Makefile to run project tasks.
+ Instead of one big JS file, write smaller CommonJS modules.
+ Use `browserify` to bundle JavaScript modules.
+ Use Babel to convert ES6 to ES5, so you can run your JS everywhere.

Although NPM is the package management tool from NodeJS, it's also great for frontend development.

Before you start, make sure that your NodeJS version is recent enough.

```
$ node -v
v2.5.0
```
Versions 2, 3, 4, or above should be ok.

# Create A New Project

[Mission: Use NPM to create a new project](project-create)

+ Use `npm init` to create `ilovereact` project.
+ Create a new git repo.
+ Push the git repo to GitHub.

Resources:

+ [github/gitignore](https://github.com/github/gitignore)
+ [MarkDown Syntax](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)

# Live Edit With BrowserSync

+ Use `npm install` to install open source packages.
+ Run executables installed in `node_modules`.
+ Run BrowserSync for live-editing.
+ Add project tasks to a `Makefile`.

Add "Build Apps with React!"

# CSS

+ @include directive.

+ Include [normalizer.css](http://necolas.github.io/normalize.css).
+ autoprefix.
+ app.css -> bundle.css
