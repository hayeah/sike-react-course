# Using ES6 With React

There are quite a few ES6 features.

The most useful features you should start using are:

+ `let`
+ `=>`
+ Destructuring assignment. `let {a,b,c} = obj`
+ `class` syntax.
+ [Enhanced object literals](https://github.com/lukehoban/es6features#enhanced-object-literals) for methods.


You should always use `let` and `=>` to help avoid common bugs.

### React.createClass API

```js
let Hello = React.createClass({
  getInitialState() {
    return {
      x: 0,
      y: 0,
    }
  },

  render() {
    let {a,b} = this.props;
    let {x,y} = this.props;

    return (
      <div>{a} {b} {x} {y}</div>
    );
  },

});
```

### Class Syntax For Components

```js
class Hello extends React.component {
  constructor(props) {
    super(props);
    this.state = {
      x: 0,
      y: 0,
    }
  }

  render() {
    let {a,b} = this.props;
    let {x,y} = this.props;

    return (
      <div>{a} {b} {x} {y}</div>
    );
  }
}
```


### Let

See [MDN: Let](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/let).

Use `let` instead of `var` to declare variables. Always! Everywhere!

```js
let links = document.querySelectorAll(".links");

for(let i = 0; i < links.length; i++) {
  // Fixes the classic "addEventListner in a loop" problem.
  let link = links[i];
  link.addEventListener("click", (event) => {
    // do something with `link`
  });
}
```

Which compiles to,

```js
var links = document.querySelectorAll(".links");

var _loop = function (i) {
  // Fixes the classic "addEventListner in a loop" problem.
  var link = links[i];
  link.addEventListener("click", function (event) {
    // do something with `link`
  });
};

for (var i = 0; i < links.length; i++) {
  _loop(i);
}
```

### Arrow Function =>

See [MDN: Arrow Functions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/Arrow_functions).

Use `=>` instead of `function` for anonymous functions. Always! Everywhere!



A common mistake is forgetting to bind `this` for a callback:

```js
let person = {
  name: "Howard",

  sayHi() {
    console.log("Hello, I am ", + this.name);
  },

  sayHiRepeatedly() {
    // BUG! `this` will become 'undefined'
    setInterval(this.sayHi,1000);
  }
}
```

Fix with `=>`:

```js
let person = {
  name: "Howard",

  sayHi() {
    console.log("Hello, I am ", + this.name);
  },

  sayHiRepeatedly() {
    setInterval(() => {
      this.sayHi();
    },1000);
  }
}
```

While compiles to:

```js
var person = {
  name: "Howard",

  sayHi: function sayHi() {
    console.log("Hello, I am ", +this.name);
  },

  sayHiRepeatedly: function sayHiRepeatedly() {
    // FIX: ensures that _this stays the same when calling sayHi()
    var _this = this;

    setInterval(function () {
      _this.sayHi();
    }, 1000);
  }
};
```

### Destructuring

This feature makes your life much easier when programming React. The render method of React component usually get information from `props` and `state`. A typical render method might look like:

```js
render() {
  return (
    <div>
      <div> a: {this.props.a} </div>
      <div> b: {this.props.b} </div>
      <div> x: {this.state.x} </div>
      <div> y: {this.state.y} </div>
    </div>
  );
}
```

The JSX looks messy with all those references to `props` and `state`. You can clean it up by declaring the variables at the very beginning of the render method:

```js
render() {
  let {a,b} = this.props
  let {x,y} = this.state;

  return (
    <div>
      <div> a: {a} </div>
      <div> b: {b} </div>
      <div> x: {x} </div>
      <div> y: {y} </div>
    </div>
  );
}
```
