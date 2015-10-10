# Modular JavaScript

JavaScript is broken by default. `let` fixes common bugs caused by `var`. `=>` fixes common bugs caused by `this`. Having no way to break JavaScript into modules is another big problem that needs fixing.

Up to now, we've been using `<script>` to load the JavaScript files we need. It's actually a terrible way to do things.

First, each `<script>` tag is one extra HTTP request, which slows down the page load.

Then you have to make sure that you load a library before you use it in your own code. If you use a jQuery plugin, you'd have to load the jQuery plugin after jQuery, but before your own app.

Finally, because all scripts share the same global namespace, you have to make sure that these scripts don't accidentally define the same variable.

A module system should solve three problems:

+ Dependencies. Each module should be declare what modules it depends on. The module system will then ensure that all dependencies are loaded in the correct order.
+ Namespacing. Functions, classes, and variables in a file shouldn't pollute the global namespace.
+ Bundling. Combine small modular files into one big bundle, so the browser only need to make one HTTP request to download it.

# CommonJS

We'll focus our attention mostly on CommonJS, which is the module system adopted by NodeJS and NPM. Because of NodeJS' popularity, CommonJS is the most popular module system, and enjoys the best tool support.

CommonJS doesn't introduce any new syntax. It adds the function `require` and the object `module`. ES6 module's design is heavily influenced by CommonJS, and almost exactly the same except for the special syntax `import` and `export`.

The CommonJS API:

+ `require` - A function used to load module dependencies.
  + Equivalent to ES6 `import`.
+ `module.exports` - The module's exported content.
  + Equivalent to ES6 `export`.

In a typical project, CommonJS and ES6 are interchangeable.

# Using CommonJS Module

In CommonJS, each file you create is a module, and you can load these files with the `require` function. There's no special syntax creating or loading modules. Just plain JavaScript.

Since CommonJS is built into NodeJS, we'll use the `node` command to experiment with it. Let's start with a simple module that provides mathematical constants.

Create the file `constants.js`:

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

+ The file has its own toplevel scope. Normally the variables would be global, but here they are local to the module.
  + Before you'd have to wrap your code in a closure. Now you don't have to.
+ There is the special object `module.exports`. When you `require` this module, this object is the return value.

Now let's try to load this module with `require`. First open a node shell:

```
$ node
```

You can eval JavaScript with the node shell. Call `require` with the path of the module to load it:

```
> require("./constants")
{ pi: 3.14159, e: 2.71828 }
```

The value of `module.exports` is returned. If you try to use `pi`, you get an error:

```
> pi
ReferenceError: pi is not defined
```

This is because `require` doesn't modify the current scope. It only returns the value of the loaded module. You have to assign the result of require to a variable so you can use it:

```js
> var pi = require("./constants").pi;
> pi
3.14159
```

With ES6's destructuring assignments, you can concisely import the values more concisely, like this:

```js
let {pi,e} = require("./constants");
var pie = pi + e;
```

ES6's own import syntax looks almost the same:

```js
import {pi,e} from "./constants";
var pie = pi + e;
```

By default, Babel compiles import to CommonJS's require:

```js
var _constants = require("./constants");
var pie = _constants.pi + _constants.e;
```

# Proejct Structure

+ explain that relative path is relative to the current file.
+ use an example project... probably.

# Requiring NPM Package

```js
let React = require("react");
let {pi,e} = require("./constants");
let view = (
  <div>
    Hello CommonJS!
  </div>
);

console.log(React.renderToString(view));
```

node doesn't understand JSX or ES6. We need to compile the file with babel first.

```
babel hello.jsx -o hello.js
node hello.js
```

How does require know where to load React?

```
# this is the file specified in package.json's `main` property.
> require.resolve("react")
./node_modules/react/react.js
```

# Bundling With Browserify

For NodeJS, it's not necessary to bundle the modules into a single file, since it's cheap for NodeJS to load the needed modules from disk one file at a time. This could translate to dozens or hundreds of HTTP requests in the browser.

For the browser we'll use [browserify](http://browserify.org/) to package all the modules into a single file.

Install it:

```
npm install browserify@11.2.0 --save-dev
```

Verify that it's installed:

```
$ browserify --version
11.2.0
```

Let's try bundling with browserify. First, create the `pie.js` file:

```js
var constants =  require("./constants");
var pi = constants.pi;
var e = constants.e;
var pie = pi + e;
```

The browserify command can process `pie.js` to find all the dependencies recursively. The bundled file looks like:

```
$ browserify pie.js
(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
// constants.js
var pi = 3.14159;
var e = 2.71828;

var secretAnswer = 42;

module.exports = {
  pi: pi,
  e: e,
};
},{}],2:[function(require,module,exports){
var constants =  require("./constants");
var pi = constants.pi;
var e = constants.e;
var pie = pi + e;
},{"./constants":1}]},{},[2]);
```

### Exercise: Use Browserify To Include Dependencies

We are using `<script>` tags to include React and PerfectScroll.

# Watchify



