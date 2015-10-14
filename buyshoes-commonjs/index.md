# Modular JavaScript With CommonJS

Note: I now recommend [Webpack](http://webpack.github.io) over browserify/watchify. Webpack is a lot faster if your project has a lot of dependencies. The ideas are the same. It's easy to switch from Browserify to Webpack, so don't worry.

JavaScript is broken by default. `let` fixes common bugs caused by `var`, and `=>` fixes common bugs caused by `this`.

Up to now, we've been using `<script>` to load the JavaScript files we need. This is yet another big problem that needs fixing. The problems are:

+ Each `<script>` tag is one extra HTTP request. Slows down page load.
+ No explicit dependency tree between files.
+ You need to specify the correct loading order.
+ All script shares the same namespace.

In this lesson we'll use CommonJS break up `app.js` into smaller modules.

# CommonJS Introduction

CommonJS is the module system adopted by NodeJS/NPM. It's probably the most popular module system, and enjoys the best tools support. Don't confuse CommonJS with NPM, though.

+ CommonJS - Breaks a project into smaller modules.
+ NPM - Create and install packages, which may or may not use CommonJS.

CommonJS doesn't introduce any new syntax to JavaScript. The CommonJS API adds just two things:

+ `require` - A function used to load a module.
+ `module.exports` - The module's exported content.

Like ES6, you need to compile your CommonJS project into normal JavaScript that the browser can understand.

### CommonJS Interop With ES6 Modules

ES6 module's design is heavily influenced by CommonJS. Because they are so similar, Babel allows you to use either interchangeably.

```js
// CommonJS
let foo = require("foo");

// ES6
import foo from "foo";
```

Babel will compile the ES6 `import` syntax to `require`.

# Using CommonJS Module

Since CommonJS is built into NodeJS, we can use the `node` interpreter to experiment with it. Let's start with a simple module that provides mathematical constants.

In CommonJS, each file you create is a module, there's no special syntax. Create the file `constants.js`:

```js
// constants.js
var pi = 3.14159;
var e = 2.71828;

var secretAnswer = 42;

module.exports = {
  pi: pi,
  e: e,
};
```

This file is different from a file loaded with `<script>` in two ways:

+ The file has its own namespace. No need to wrap your code in a closure.
+ There is the special object `module.exports`.

CommonJS modules don't have names; they are just files. To load them, you use the `require` function, and specify which file you want to load. When the module is loaded, the `require` function returns `module.exports` as the result.

Let try to `require` the module `constants.js`. The value should be:

```js
{
  pi: pi,
  e: e,
}
```

### Exercise: Load a module with require

First open a node shell:

```
// babel-node supports ES6
$ babel-node
```

Call `require` with the path to the module file:

```
> require("./constants")
{ pi: 3.14159, e: 2.71828 }
```

The value of `module.exports` is returned. If you try to use `pi`, you get an error:

```
> pi
ReferenceError: pi is not defined
```

This is because `require` doesn't modify the current scope. It only returns the value of the loaded module.

Create a new local variable `pi`:

```js
> let pi = require("./constants").pi;
> pi
3.14159
```

Using ES6's destructuring we can create variables more concisely:

```js
let {pi,e} = require("./constants");
var pie = pi + e;
```

### Exercise: ES6 import syntax

The `import` is similar to CommonJS:

```js
> import "./constants";
{ pi: 3.14159, e: 2.71828 }
```

And to create multiple variables at the same time:

```js
import {pi,e} from "./constants";
var pie = pi + e;
```

Use `babel` to compile the above code, you should see:

```js
var _constants = require("./constants");
var pie = _constants.pi + _constants.e;
```

### Exercise: Add a new export value

Add a new number to `constants.js`:

```js
// The golden ratio
let phi = 1.61803;
```

From `babel-node`:

```
> import "./constants";
{ pi: 3.14159, e: 2.71828, phi: 1.61803 }
```

Note: Remember to restart `babel-node`, or else you wan't see the new module value.

Question: If you `require` a module 3 times, how many times is the file evaluated?

# Bundling With Webpack

[Webpack](http://webpack.github.io/) is a tool that turns a CommonJS project into normal JavaScript that the browser can understand.

There are other older/mature/popular tools like [Grunt](http://gruntjs.com/), [Gulp](http://gulpjs.com/), and [Browserify](https://github.com/substack/node-browserify), is it safe to use a relatively new tool like Webpack? Maybe next week another build tool would become popular. If your project is already using an existing tool, it's likely not worth the effort to convert to Webpack.

Webpack is a complicated tool, with lots of [features](http://webpack.github.io/docs/) and [configuration options](http://webpack.github.io/docs/configuration.html). We'll avoid using the more advanced stuff, and focus on the core functionalities that all build tools would have:

+ Collects all CommonJS modules into a single file.
+ Provides fake `require` in the browser.
+ Ensures that a module is evaluated only once, and in the right order.
+ Converts ES6/JSX to ES5 (normal JavaScript).

Because any future build tools should have these features, there is less risk of being locked into Webpack. Indeed, it takes almost no work to convert between using Browserify and Webpack!

Install:

```
npm install webpack@1.12.2 --save-dev
```

To be able to convert ES6/JSX to ES5, we'd also need to install the Webpack Babel plugin:

```
npm install babel-loader --save-dev
```

### Exercise: Bundling pie.js

Let's try create a bundle with Webpack. First, create the `pie.js` file:

```js
let {pi,e} =  require("./constants");
console.log("pie =",pi + e);
```

Evaluating `pie.js` with NodeJS should print out its value:

```
$ babel-node pie.js
pie = 5.85987
```

Now let's make this work for the browser.

```
# webpack [entry-file] [bundle-file]
$ webpack pie.js pie-bundle.js --module-bind "js=babel"
webpack pie.js pie-bundle.js
Hash: c499792d1a74823ee7da
Version: webpack 1.12.2
Time: 65ms
        Asset     Size  Chunks             Chunk Names
pie-bundle.js  1.69 kB       0  [emitted]  main
   [0] ./pie.js 67 bytes {0} [built]
   [1] ./constants.js 145 bytes {0} [built]
```

+ The `entry-file` - The entry of the project. Put `window.onload` here.
+ The `bundle-file` - The bundled file.
+ `--module-bind` - All files with the `.js` extension should be compiled with Babel.

The bundled project is like:

```js
/******/ (function(modules) {

/* webpack loader. omitted */

/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ function(module, exports, __webpack_require__) {

  "use strict";

  var _require = __webpack_require__(1);

  var pi = _require.pi;
  var e = _require.e;

  console.log("pie =", pi + e);

/***/ },
/* 1 */
/***/ function(module, exports) {

  // constants.js
  "use strict";

  var pi = 3.14159;
  var e = 2.71828;

  module.exports = {
    pi: pi,
    e: e
  };

/***/ }
/******/ ]);
```

Run the bundled code in the browser to verify that it works!

Notice how the modules are wrapped in a function to ensure a new scope:

```
function(module,exports,__webpack_require__){
  // module code.
}
```

Also, the `require` function is replaced with `__webpack_require__`.

## Webpack Bootstrap

Reading the `webpackBootstrap` code is a good way to understand exactly how CommonJS works.

Modules are closures:

```js
// (function(modules) { ... })([modules])
let modules = [
/* 0 */
/***/ function(module, exports, __webpack_require__) {

  "use strict";

  var _require = __webpack_require__(1);

  var pi = _require.pi;
  var e = _require.e;

  console.log("pie =", pi + e);

/***/ },
/* 1 */
/***/ function(module, exports) {

  // constants.js
  "use strict";

  var pi = 3.14159;
  var e = 2.71828;

  module.exports = {
    pi: pi,
    e: e
  };

/***/ }
/******/ ]
```

The definition for `require` is like this:

```js
var installedModules = {};

// The require function
function __webpack_require__(moduleId) {

  // Check if module is in cache
  if(installedModules[moduleId])
    return installedModules[moduleId].exports;

  // Create a new module (and put it into the cache)
  var module = installedModules[moduleId] = {
    exports: {},
    id: moduleId,
    loaded: false
  };

  // Execute the module function
  modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);

  // Return the exports of the module
  return module.exports;
}
```

+ It caches the module in `installedModules`, so each module is executed just once.
+ It returns the value of `module.exports` at the end.

The most interesting line is the module execution:

```js
modules[moduleId].call(
  // `this` is module.exports
  module.exports,
  // Make `module` available to the module code
  module,
  // Make `exports` available to the module code
  module.exports,
  // For recursive require.
  __webpack_require__
);
```

The only difference is that `__webpack_require__`  uses webpack's internal module id, which is the position of the module in the `modules` array.

And what is the "entry file"? It's the module that is automatically evaluated when the bundle is loaded:

```js
return __webpack_require__(0);
```

# Bundling BuyShoes Dependencies

Let's bundle PerfectScrollbar and React with `app.jsx`.

### Exercise: Bundling PerfectScrollbar and React

Remove from `index.html`:

```html
<script type="text/javascript" src="node_modules/perfect-scrollbar/dist/js/perfect-scrollbar.js"></script>
<script type="text/javascript" src="node_modules/react/dist/react.js"></script>
```

In `app.jsx` add:

```js
const Ps = require("../node_modules/perfect-scrollbar/index");
const React = require("../node_modules/react/react");
```

Use Webpack to create the bundle to `build/app.js`. Bundling now takes somewhat longer because React is pretty big. Add the `--progress` option to the `webpack` command to see how many modules webpack had bundled.

Note: The require paths are relative to the module file. Depending on where a file is, the relative path to `node_modules` is different:

+ `app.jsx` - require("./node_modules/...")
+ `a/app.jsx` - require("../node_modules/...")
+ `a/b/app.jsx` - require("../../node_modules/...")

# Require By Package Name

We can also use the package name to `require` React and PerfectScrollbar:

```js
const Ps = require("perfect-scrollbar");
const React = require("react");
```

Usually `require` loads a file by its path. If it's a package name, NodeJS uses the `require.resolve` function to find which file to load. See which file `require("react")` would load:

```
$ node
> require.resolve("react")
./node_modules/react/react.js
```

### Exercise: Use package name to bundle React and PerfectScrollbar

The result should be the same as before.

# Live-Edit

Webpack can automatically rebundle the project when you make changes. Just add the `--watch` option:

```
$ webpack --watch --progress ...
```

Because Webpack caches all the modules in memory, it needs to recompile only the module that had changed. In one of my projects, browserify+watchify takes 3~4 seconds to rebundle, but Webpack can do it in ~300ms.

### Exercise: Modify Makefile for live-edit.

When you edit a file,

1. Webpack should compile and bundle the project.
2. BrowserSync should see that bundle file had changed, and reload the browser.

Change `make js` to make this happen.

# Modularize BuyShoes

Our goal is to reduce `app.jsx` to just this:

```js
// When the window is loaded, render the App component.
const App = require("./components/App");

window.onload = () => {
  React.render(<App/>,document.querySelector("#root"));
}
```

### Exercise: Modularize Fake Data

Create the file `js/data.js`:

```js
module.exports = {
  cartItems: ...
  products: ...
}
```

Then import data into `js/app.jsx`.

### Exercise: Modularize SiteTitle

Put the `SiteTitle` component into its own module. We'll put all components in the `js/components` directory.

First, create the directory `js/components`.

Then create the file `js/components/SiteTitle.js`:

```js
const React = require("react");
let SiteTitle = React.createClass({
  render() {
    return (
      <div className="title">
        <h2>Buy Me Shoes</h2>
        <img className="title__heart" src="img/heart.svg" />
      </div>
    );
  }
});

module.exports = SiteTitle;
```

Modify `app.jsx` to import this component.

Note: The extension must be `.js`, not `.jsx`.

### Exercise: Modularize Everything Else

Turn all the components into modules.

It'd be easier to start with a simple App, then migrate the components one by one. Start with the components commented out:

```html
<div className="site">
  <div className="bg">
    <div className="bg__img">
    </div>
  </div>

  <div className="site__main">
    <div className="site__left-sidebar">
      <SiteTitle/>
    </div>
    <div className="site__content">
      {/* <Products/> */}
    </div> {/* site__content */}
  </div> {/* site__main */}
  <div className="site__right-sidebar">
    {/* <Cart/> */}
    {/* <Checkout/> */}
  </div> {/* site__right-sidebar */}
  <a className="site__right-sidebar-toggle">
    <img src="img/arrow-icon.svg" />
  </a>
</div>
```

+ `App` in `js/components/App.js`
+ `Cart` in `js/components/Cart.js`
+ `Products` in `js/components/Products.js`
+ etc.

# Source Map For Debugging

The bundled `build/app.js` is a huge file, making it hard to debug.

![](no-sourcemap.jpg)

Thankfully, we can ask Webpack to generate [source map](http://www.html5rocks.com/en/tutorials/developertools/sourcemaps/), so Chrome can correlate between the JavaScript that runs in the browser, and the original source files you've written.

Add the `-d` option to the `webpack` command to enable the "development mode", which generates source map for your project.

See: [Webpack CLI - development shortcut -d](https://webpack.github.io/docs/cli.html#development-shortcut-d)

With source map enabled, Chrome can now shows you the original source files:

![](with-sourcemap.jpg)

# Minified JavaScript

For production, you'd want to:

1. Make the bundle smaller. Removing comments and whitespace, etc.
2. Obfuscate the source so it's harder for other people to borrow/steal it.

[Uglify](https://github.com/mishoo/UglifyJS2) is the most popular tool for minifying JavaScript. Webpack make it super easy. Just add the `-p` option to enable production mode, and you'd get a final output like:

```js
!function(e){function t(o){if(n[o])return n[o].exports;var r=n[o]={exports:{},id:o,loaded:!1};return e[o].call(r.exports,r,r.exports,t),r.loaded=!0,r.exports}var n={};return t.m=e,t.c=n,t.p="",t(0)}
```

Comparing the `-p` (production) vs `-d` (development), the file decreased from 710k to 188k:

```
-rw-r--r--  1 howard  staff   188K 12 Oct 19:57 bundle/app.js
-rw-r--r--  1 howard  staff   710K 12 Oct 19:49 build/app.js
```

### Exercise: Create the minjs task

Add the `minjs` task to Makefile. It should create `bundle/app.js`, which is the minified version of `build/app.js`.

# Summary

We've seen how we can break a big file into modules.

+ Every file is a CommonJS module.
+ CommonJS adds `require` and `module.exports`.
+ ES6 modules adds the `import` and `export` syntax.
+ Load a package by calling `require` with the path to a file, or with a package name.
+ The `require` path is relative to the requiring file.
+ Use Webpack to bundle a CommonJS for the browser.
