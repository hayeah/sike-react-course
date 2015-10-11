# Modular JavaScript With CommonJS

JavaScript is broken by default. `let` fixes common bugs caused by `var`, and `=>` fixes common bugs caused by `this`.

Up to now, we've been using `<script>` to load the JavaScript files we need. This is yet another big problem that needs fixing. The problems are:

+ Each each `<script>` tag is one extra HTTP request. Slows down page load.
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

# Bundling With Browserify

[Browserify](http://browserify.org/) is a tool that turns a CommonJS project into normal JavaScript that the browser can understand. Browserify does the followings:

+ Collects all dependencies into a single file.
+ Provides fake `require` in the browser.
+ Ensures that a module is evaluated only once, and in the right order.

Install:

```
npm install browserify@11.2.0 babelify@6.3.0 --save-dev
```

By default browserify only understands ES5 (normal JavaScript). We've also installed `babelify` to extend browserify with the ability to process ES6 files.

### Exercise: Bundling pie.js

Let's try create a bundle with browserify. First, create the `pie.js` file:

```js
let {pi,e} =  require("./constants");
console.log("pie =",pi + e);
```

Evaluating `pie.js` should print out its value:

```
$ babel-node pie.js
pie = 5.85987
```

Create the bundle with browserify:

```
browserify --transform babelify pie.js --outfile pie-bundled.js
```

+ `--transform babelify` - Use Babel to compile files before bundling.
+ `--outfile` - Writes result to the specified file.

The result is:

```js
(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
// constants.js
"use strict";

var pi = 3.14159;
var e = 2.71828;

var secretAnswer = 42;

module.exports = {
  pi: pi,
  e: e,
};

},{}],2:[function(require,module,exports){
"use strict";

var _require = require("./constants");

var pi = _require.pi;
var e = _require.e;

console.log("pie =", pi + e);

},{"./constants":1}]},{},[2]);
```

Run the above code in the browser to verify that it works!

Notice how the modules are wrapped in a function to ensure a new scope:

```
function(require,module,exports){
  // module code.
}
```

The CommonJS API `require` and `module` are passed into the module code as arguments.

## DIY Require

At the top of the bundle is a chunk of scary looking code:

```js
function e(t,n,r){function s(o,u) { ... }}
```

This defines how `require` loads modules. What require does is actually quite simple. Here is a simplified version of `require`:

```js
"use strict";

// The bundled modules
var modules = {
  "./parent": function(module,require) {
    module.exports = 'parent+' + require('./child');
  },

  "./child": function(module,require) {
    module.exports = 'child';
  },
};

// Use a cache to ensure that a module is evaluated just once.
var cache = {};

function require(path) {

  // Return cached module if already loaded.
  if(cache.hasOwnProperty(path)) {
    return cache[path];
  }

  // This defines module.export
  var module = {
    exports: {}
  };


  // Evaluate the module.
  var modfn = modules[path];
  modfn(module,require);

  // Cache the module value.
  cache[path] = module.exports;

  // Return the exported objects.
  return module.exports;
}

console.log(require("./parent")); // => parent+child
```

The actual code is here: [prelude.js](https://github.com/substack/browser-pack/blob/aadeabea66feac48193d27d233daf1c85209357e/prelude.js).

To learn more see: [How Browserify Works](http://benclinkinbeard.com/posts/how-browserify-works/).

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

Use browserify to compile the bundle to `build/app.js`.

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

# Live-Edit Using Watchify

Browserify doesn't have a `--watch` flag that automatically rebundles if a file change. A separate tool [watchify](https://github.com/substack/watchify) does that.

Install:

```
npm install watchify@3.4.0 --save-dev
```

To use it, replace the browserify command with `watchify`. Everything else is the same:

```
watchify --transform babelify pie.js --outfile pie-bundled.js
```

### Exercise: Modify Makefile

+ Change `make js` to use watchify.

BrowserSync should still work.

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

The bundled `build/app.js` is too big, making it hard to debug.

![](no-sourcemap.jpg)

Thankfully, we can ask browserify to generate [source map](http://www.html5rocks.com/en/tutorials/developertools/sourcemaps/), so Chrome can correlate between the JavaScript that runs in the browser, and the original source files you've written.

Add the `--debug` option to browserify to generate Source Map. The command looks like:

```
browserify --debug --transform babelify pie.js --outfile pie-bundled.js
```

With source map enabled, Chrome can now shows you the original source files. `Cmd-P` to quickly find a file:

<video src="cmd-p-find-file.mp4" controls></video>

You can even set breakpoints:

![](with-sourcemap.jpg)

# Summary

We've seen how we can break a big file into modules.

+ Every file is a CommonJS module.
+ CommonJS adds `require` and `module.exports`.
+ ES6 modules adds the `import` and `export` syntax.
+ Load a package by calling `require` with the path to a file, or with a package name.
+ The `require` path is relative to the requiring file.
+ Use browserify to bundle a CommonJS for the browser.

