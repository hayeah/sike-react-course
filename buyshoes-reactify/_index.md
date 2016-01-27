# Reactify The Buy Shoes Page

<cn>
# React 化 Buy Shoes 页面
</cn>

In this lesson we are going to rewrite the `buyshoes` page in React. We'll cover the following:

+ Compile ES6 & JSX with Babel.
+ Virtual DOM with JSX.
+ Writing React components.
+ Using the `key` property to allow reordering.
+ Component lifecycle hooks.
+ Using the `ref` property to access browser DOM.

<cn>
这一课中，我们将使用 React 来重写 `buyshoes` 页面。我们将学习以下内容：

+ 使用 Babel 编译 ES6 & JSX。
+ Virtual DOM + JSX。
+ 编写 React 组件。
+ 使用 `key` 属性来重新排列、调换次序。
+ 了解组件的生命周期。
+ 使用 `ref` 属性来获取浏览器 DOM 内容。
</cn>

# Lesson StarterKit

<cn>
# 课程基础项目
</cn>

For this lesson, we have provided a reference implementation for the `buyshoes` page. It includes the HTML and css.

<cn>
这里我们将提供 `buyshoes` 页面的一个实现。这个基础项目包括了 HTML 和 CSS。
</cn>

```
git clone https://github.com/hayeah/sikeio-buyshoes-startkit.git sikeio-buyshoes-react
cd sikeio-buyshoes-react
git checkout styled
```

<cn>
```
git clone https://github.com/hayeah/sikeio-buyshoes-startkit.git sikeio-buyshoes-react
cd sikeio-buyshoes-react
git checkout styled
```
</cn>

You could use your own if you want.

<cn>
如果你喜欢也可以用自己的！
</cn>

## Project Assets

<cn>
## 项目素材
</cn>

We'll add different kinds of shoes.

Download images here: https://github.com/hayeah/sikeio-buyshoes-assets/tree/master/shoes

<cn>
我们将添加几种不同款式的鞋子。

在这里可以下载到资源：https://github.com/hayeah/sikeio-buyshoes-assets/tree/master/shoes
</cn>

# Full Page Refresh And Virtual DOM

<cn>
# 页面刷新和 Virtual DOM
</cn>

A browser refresh is the simplest way to update a page. The server renders the new page with the latest data, and the browser loads the new page. There is no need for client-side code to manipulate the DOM.

<cn>
刷新浏览器可谓是最简单的一个用来更新页面的方法了。服务器使用最新的数据重新渲染整个新的页面，然后告诉浏览器需要重新载入。客户端代码没有必要再去控制整个 DOM 的排版。
</cn>

React's Virtual DOM is like server-side rendering for the client-side. Every time user clicks a button, or enters some text that changes the data, React generates the new UI using the latest data. A render in React is essentially the same as replacing the body's `innerHTML` with the new UI:

```js
document.body.innerHTML = updatedHTML;
```

<cn>
React 中的 Virtual DOM 的作用就像是服务器端为客户端重新渲染页面一样。每当用户点击一个按钮、亦或是在文本框中输入了某些文字导致更改了数据，React 都能够使用最新的数据重新生成新的 UI 页面。在 React 中的一次渲染本质上来说等同于使用新的 UI 替换了 body 部分的 `innerHTML`。

```js
document.body.innerHTML = updatedHTML;
```

</cn>

Like a browser refresh, there is no DOM manipulation. But setting innerHTML is inefficient because it destroys the old DOM tree and rebuilds an entirely new DOM tree.

The virtual DOM makes the "full page refresh" technique efficient. React detects what had changed, and only updates the DOM elements that had changed. It also adds, removes, and reorders DOM elements as necessary.

<cn>
就像浏览器刷新，这里也没有 DOM 操作。但是改变 `innerHTML` 的内容是一件效率很低的事情，因为它破坏了旧的 DOM 树，然后再去建立了一个新的 DOM 树。

Virtual DOM 就让这种「整页刷新」的技术高效起来。React 会先检测到有哪些地方更新了，再只更新这些被更改的 DOM 元素。有必要时，React 也会增加、移除、重新排列 DOM 元素。
</cn>

With React, you get the simplicity of programming client-side as though you are always doing a full-page refresh, yet with the efficiency of direct DOM manipulation.

Let's look at an example that contrasts DOM manipulation with "full page refresh".

<cn>
有了 React，你就可以感受到客户端部分编程的简洁性——好像总是在做整个页面的刷新，却又能够高效快捷地直接操控 DOM 元素。

让我们一起来看下面这个例子，使用「整页刷新」來控制操作 DOM 元素。
</cn>

Here's a snippet that generates a greeting message:

```js
var name = "Howard";

var view = (
  '<div>' +
    '<p>Hello, <b>' + name + '</b></p>'+
    '<p>Current Time: ' + new Date() + '</p>'+
  '</div>'
);

document.body.innerHTML = view;
```

<cn>

这是生成打招呼信息的一段代码片段：

```js
var name = "Howard";

var view = (
  '<div>' +
    '<p>Hello, <b>' + name + '</b></p>'+
    '<p>Current Time: ' + new Date() + '</p>'+
  '</div>'
);

document.body.innerHTML = view;
```

</cn>

[Codepen Demo](http://codepen.io/hayeah/pen/LpWVXa?editors=001)

<cn>

[Codepen Demo](http://codepen.io/hayeah/pen/LpWVXa?editors=001)

</cn>

If we want to update the time every second, we might manipulate the DOM like this:

```js
var name = "Howard";

var view = (
  '<div>' +
    '<p>Hello, <b>' + name + '</b></p>'+
    '<p class="current-time">Current Time: ' + new Date() + '</p>'+
  '</div>'
);

document.body.innerHTML = view;

function updateTime() {
  var $currentTime = document.querySelector(".current-time");
  $currentTime.innerHTML = "Current Time: " + new Date();
}

setInterval(updateTime,1000);
```

<cn>
如果我们想要每秒来更新一下信息，可能需要对 DOM 这么操作：

```js
var name = "Howard";

var view = (
  '<div>' +
    '<p>Hello, <b>' + name + '</b></p>'+
    '<p class="current-time">Current Time: ' + new Date() + '</p>'+
  '</div>'
);

document.body.innerHTML = view;

function updateTime() {
  var $currentTime = document.querySelector(".current-time");
  $currentTime.innerHTML = "Current Time: " + new Date();
}

setInterval(updateTime,1000);
```

</cn>

[Codepen Demo](http://codepen.io/hayeah/pen/KdWpjX?editors=001)

<cn>

[Codepen Demo](http://codepen.io/hayeah/pen/KdWpjX?editors=001)

</cn>

But as soon as we start to manipulate DOM, the code gets ugly. The `view` no longer tells you everything about what is shown to the user, because now you have to know how other pieces of code could change the view.

<cn>

我们的确改变了 DOM 元素，但是代码看上去却很脏。`View` 无法告诉你用户所看到的内容，因为你也需要知道其他的代码是怎样改变 `View` 部分的内容的。

</cn>

A much simpler programming model is to regenerate the whole UI every second:

```js
var name = "Howard";

function render() {
  var view = (
    '<div>' +
      '<p>Hello, <b>' + name + '</b></p>'+
      '<p>Current Time: ' + new Date() + '</p>'+
    '</div>'
  );

  document.body.innerHTML = view;
}

setInterval(render,1000);
```

<cn>
下面有一个简单的模型来每隔一秒重新生成 UI：

```js
var name = "Howard";

function render() {
  var view = (
    '<div>' +
      '<p>Hello, <b>' + name + '</b></p>'+
      '<p>Current Time: ' + new Date() + '</p>'+
    '</div>'
  );

  document.body.innerHTML = view;
}

setInterval(render,1000);
```

</cn>

[Codepen Demo](http://codepen.io/hayeah/pen/rOyVov?editors=001)

<cn>

[Codepen Demo](http://codepen.io/hayeah/pen/rOyVov?editors=001)

</cn>

There is no DOM manipulation. The `view` tells us exactly what the UI is. While the "full page refresh" programming model is simple, it is too inefficient. Setting the innerHTML destroys the old DOM tree and replaces it with an newly built DOM tree:

<cn>
这里并没有 DOM 操作。`View` 视图直接告诉了我们 UI 长什么样子。尽管「整页刷新」的编程模型的确很简单，但是它太低效了。设置 innerHTML 破坏了整个旧的 DOM 树结构，之后再用新的 DOM 树替代：
</cn>

<video src="setting-innerHTML.mp4" controls autoplay loop></video>

<cn>

<video src="setting-innerHTML.mp4" controls autoplay loop></video>

</cn>

React avoids the expenses of rebuilding the DOM tree. The code in React looks almost the same:

```js
var name = "Howard";

function render() {
  var view = (
    <div>
      <p>Hello, <b>{name}</b></p>
      <p>Current Time: {(new Date()).toString()}</p>
    </div>
  );

  // document.body.innerHTML = view;
  React.render(view,document.body);
}

setInterval(render,1000);
```

<cn>

React 避免了重新构建 DOM 树的开销。React 中的代码也看上去很类似：

```js
var name = "Howard";

function render() {
  var view = (
    <div>
      <p>Hello, <b>{name}</b></p>
      <p>Current Time: {(new Date()).toString()}</p>
    </div>
  );

  // document.body.innerHTML = view;
  React.render(view,document.body);
}

setInterval(render,1000);
```

</cn>

But React detects that only the time had changed, and updates only that specific DOM element:

<cn>
但是 React 能识别只有时间发生了改变，然后仅仅更新对应的 DOM 元素：
</cn>

<video src="dom-reconciliation.mp4" controls autoplay loop></video>

<cn>

<video src="dom-reconciliation.mp4" controls autoplay loop></video>

</cn>

Efficient "full page refresh" is React's most important idea.

<cn>
高效的「整页刷新」正是 React 最重要的思想！
</cn>

# Babel Compiler For ES6 And JSX

<cn>
# 使用 Babel 编译 ES6 和 JSX
</cn>

Babel is the ES6 compiler that is most closely linked with React. It has great support for React right out of the box.

+ CommonJS support.
+ JSX syntax for virtual dom.
+ ES6 and ES7 features.
+ Compiles to readable ES5 that can run anywhere.

<cn>
Babel 是一个 ES6 编译器。它也和 React 息息相关。它为 React 在功能上提供了非常便捷的支持：

+ 支持 CommonJS。
+ 为 Virtual Dom 提供 JSX 格式支持。
+ ES6 和 ES7 特性。
+ 可以编译成可读、到处可运行的 ES5 文件。
</cn>

To install Babel:

```
npm install babel@5.8.23 --save-dev
```

<cn>
安装 Babel：

```
npm install babel@5.8.23 --save-dev
```
</cn>

Verify that Babel is installed:

```
$ babel --version
5.8.21 (babel-core 5.8.22)
```

<cn>

验证 Babel 已被安装：

```
$ babel --version
5.8.21 (babel-core 5.8.22)
```
</cn>

For a good summary of ES6 features see: [ES6 Features](https://github.com/lukehoban/es6features).

<cn>
ES6 特性总结请看：[ES6 Features](https://github.com/lukehoban/es6features)
</cn>

You can see some of the ES6 & ES7 features that Babel supports in its [list of transformers](http://babeljs.io/docs/learn-es2015/). It's possible to turn on/off individual features to control exactly what you'd allow your team to use.

<cn>
Babel 所支持的 ES6 和 ES7 中的特性可以在[这里](http://babeljs.io/docs/learn-es2015/)看到。你还可以通过开启或关闭某些单独的特性来控制团队所使用的功能。
</cn>

The two most important ES6 features are `let` and `=>`. You should always use them.

+ Use `let` to replace `var`. Always.
+ Use `=>` to replace anonymous functions. Always.

<cn>
有两个最重要的 ES6 中的特性就是 `let` 和 `=>` 了。你会经常使用到它们。

+ 记住永远使用 `let` 替代 `var`。
+ 记住永远使用 `=>` 替代匿名函数。

</cn>

ES6 helps you to write more succinct code. See [ES6 With React](../es6-with-react).

<cn>
ES6 能帮助你写出更佳简练的代码。可以查看 [ES6 With React](../es6-with-react)。
</cn>

# Getting Started With React Virtual DOM

<cn>
# 从 React Virtual DOM 开始
</cn>

Our first task is to get React running. Install React:

```
npm install react@0.13.3 --save-dev
```

<cn>
我们的第一项任务就是让 React 跑起来。安装 React：

```
npm install react@0.13.3 --save-dev
```
</cn>

Let's create the component `App` in a new file called `js/app.jsx`:

```js
// The App component.
let App = React.createClass({
  // The `render` method will generate the `buyshoes` site's virtual DOM.
  render() {
    return (
      <div className="site">
        <h1>Buy Some Shoes!!!</h1>
      </div>
    );
  },
});

window.onload = () => {
  // Replace innerHTML of `#root` with the App component.
  React.render(<App/>,document.querySelector("#root"));
}
```

<cn>

让我们先在 `js/app.jsx` 中创建一个新的组件 `App`：

```js
// App 组件
let App = React.createClass({
  // `render` 方法将生成 `buyshoes` 网页的 Virtual DOM。
  render() {
    return (
      <div className="site">
        <h1>Buy Some Shoes!!!</h1>
      </div>
    );
  },
});

window.onload = () => {
  // 使用 App 组件替换 `#root` 的 innerHTML。
  React.render(<App/>,document.querySelector("#root"));
}
```

</cn>

Use Babel to compile `js/app.jsx`, we can see the result in the terminal:

```js
$ babel js/app.jsx
var App = React.createClass({
  displayName: "App",

  render: function render() {
    return React.createElement(
      "div",
      { className: "site" },
      React.createElement(
        "h1",
        null,
        "Buy Some Shoes!!!"
      )
    );
  }
});

window.onload = function () {
  React.render(React.createElement(App, null), document.querySelector("#root"));
};
```

<cn>
使用 Babel 来编译 `js/app.jsx`。可以在终端中看到结果：

```js
$ babel js/app.jsx
var App = React.createClass({
  displayName: "App",

  render: function render() {
    return React.createElement(
      "div",
      { className: "site" },
      React.createElement(
        "h1",
        null,
        "Buy Some Shoes!!!"
      )
    );
  }
});

window.onload = function () {
  React.render(React.createElement(App, null), document.querySelector("#root"));
};
```

</cn>

JSX is just a simple mapping from HTML onto `React.createElement` calls. The first argument is the type of element. The second argument is a map of properties.

<cn>
JSX 是一个从 HTML 到 `React.createElement` 调用的简单映射。第一个参数是元素的类型，第二个参数是属性映射的集合。
</cn>

The `React.createElement` API:

```js
ReactElement createElement(
  string/ReactClass type,
  [object props],
  [children ...]
)
```

<cn>

`React.createElement` API:

```js
ReactElement createElement(
  string/ReactClass type,
  [object props],
  [children ...]
)
```

</cn>

### Virtual DOM Is Side-Effect Free

<cn>
### Virtual DOM 不存在副作用
</cn>

JSX may resemble other conventional template languages, like PHP:

```html
<html>
   <head>
      <title>Online PHP Script Execution</title>
   </head>

   <body>

      <?php
         echo "<h1>Hello, PHP!</h1>";
      ?>

   </body>
</html>
```

<cn>
JSX 有点像其他的传统模版语言，比如 PHP：

```html
<html>
   <head>
      <title>Online PHP Script Execution</title>
   </head>

   <body>

      <?php
         echo "<h1>Hello, PHP!</h1>";
      ?>

   </body>
</html>
```

</cn>

But in fact, JSX is a thin syntatic sugar coating for the virtual DOM API, which has a completely different programming model. The virtual DOM is a tree of ordinary JavaScript objects. There is a lot of flexibility in how we build the virtual DOM.

<cn>
但是事实上，JSX 只是包含 Virtual DOM API 的一个简单的语法糖，它包含了一个完全不一样的编程模型。Virtual DOM 可被看作为包含 JavaScript 元素的一棵树。构建 Virtual DOM 会带来很多的便捷性。
</cn>

For example, we could break render into two separate methods:

```js
// The App component.
let App = React.createClass({
  renderTitle() {
    return (
      <h1>Buy Some Shoes!!!</h1>
    );
  },

  render() {
    return (
      <div className="site">
        {this.renderTitle()}
      </div>
    );
  },
});
```

<cn>
例如，我们可以将 Render 分解为两个不同的方法：


```js
// App 组件
let App = React.createClass({
  renderTitle() {
    return (
      <h1>Buy Some Shoes!!!</h1>
    );
  },

  render() {
    return (
      <div className="site">
        {this.renderTitle()}
      </div>
    );
  },
});
```

</cn>

Or we can generate three titles using `map`:

```js
let App = React.createClass({
  renderTitle() {
    return (
      <h1>Buy Some Shoes!!!</h1>
    );
  },

  render() {
    return (
      <div className="site">
        {[1,2,3].map(this.renderTitle)}
      </div>
    );
  },
});
```

<cn>
或者是我们可以通过 `map` 方法来生成三个标题：

```js
let App = React.createClass({
  renderTitle() {
    return (
      <h1>Buy Some Shoes!!!</h1>
    );
  },

  render() {
    return (
      <div className="site">
        {[1,2,3].map(this.renderTitle)}
      </div>
    );
  },
});
```
</cn>

Or we could move the three titles outside the `.site` div:

```js
let App = React.createClass({
  renderTitle() {
    return (
      <h1>Buy Some Shoes!!!</h1>
    );
  },

  render() {
    let titles = [1,2,3].map(this.renderTitle);

    return (
      <div className="site">
        {titles}
      </div>
    );
  },
});
```

<cn>
又或者是我们可以把三个标题移到 `.site` 之外：

```js
let App = React.createClass({
  renderTitle() {
    return (
      <h1>Buy Some Shoes!!!</h1>
    );
  },

  render() {
    let titles = [1,2,3].map(this.renderTitle);

    return (
      <div className="site">
        {titles}
      </div>
    );
  },
});
```
</cn>

Using functional programming jargon, the virtual DOM API is "side-effect free". Being side-effect free gives you three important guarantees:

+ You can call render methods many times.
+ You can call render methods in any order.
+ As long as the inputs are the same (state, props), the virtual DOM output is the same.

<cn>
使用了函数式编程的思想，Virtual DOM API 就是「无副作用的」。它能给你带来以下三点重要的保障：

+ 你可以调用任意多次 Render 方法。
+ 你可以以任意的次序调用 Render 方法。
+ 只要所给的输入时相同的（例如状态、属性），Virtual DOM 的输出就是相同的。
</cn>

These guarantees gives you a lot of flexibilty in how you could refactor your code.

<cn>
这些保障带给了你重构代码的可能和便捷性！
</cn>

Conventional templating languages aren't this flexible. For example, if we reordered `document.write`, we'd get the wrong output:

```js
renderTitle() {
  document.write('<h1>Buy Some Shoes!!!</h1>');
}

render() {
  [1,2,3].forEach(this.renderTitle);

  document.write('<div className="site">');
  document.write('<div>');
}
```

<cn>
传统的模版语言久可能没这么便捷和有弹性了。例如，如果我们改变了 `document.write` 的顺序，我们可能得到错误的输出：

```js
renderTitle() {
  document.write('<h1>Buy Some Shoes!!!</h1>');
}

render() {
  [1,2,3].forEach(this.renderTitle);

  document.write('<div className="site">');
  document.write('<div>');
}
```

</cn>

Don't let the familiar HTML syntax deceive you. The essence of JSX is the `React.createElement` method.

<cn>
不要让熟悉的 HTML 语法欺骗了你！JSX 最本质的东西就是 `React.createElement` 方法。
</cn>

### Exercise: Create The App Component

<cn>
### 练习：创建 App 组件
</cn>

Create the file `js/app.jsx`:

```js
$ babel js/app.jsx
var App = React.createClass({
  displayName: "App",

  render: function render() {
    return React.createElement(
      "div",
      { className: "site" },
      React.createElement(
        "h1",
        null,
        "Buy Some Shoes!!!"
      )
    );
  }
});

window.onload = function () {
  React.render(React.createElement(App, null), document.querySelector("#root"));
};
```

<cn>
新建 `js/app.jsx` 文件：

```js
$ babel js/app.jsx
var App = React.createClass({
  displayName: "App",

  render: function render() {
    return React.createElement(
      "div",
      { className: "site" },
      React.createElement(
        "h1",
        null,
        "Buy Some Shoes!!!"
      )
    );
  }
});

window.onload = function () {
  React.render(React.createElement(App, null), document.querySelector("#root"));
};
```
</cn>

Previously we used `babel` to output the compiled result to the terminal. Now we want to save the compiled result to `build/app.js`:

```
# Make sure that the `build` directory exists.
mkdir -p build

# Compiles js/app.jsx
babel js/app.jsx --out-file build/app.js
```

<cn>
之前我们使用 `Babel` 将编译后的结果输出到终端中。先在我们希望将它保存到 `build/app.js` 里：

```
# 确保 `build` 目录存在
mkdir -p build

# 编译 js/app.jsx
babel js/app.jsx --out-file build/app.js
```

</cn>

Next, load React and `build/app.js`:

```js
<script type="text/javascript" src="node_modules/react/dist/react.js"></script>
<script type="text/javascript" src="build/app.js"></script>
```

<cn>
接下来载入 React 和 `build/app.js`：

```js
<script type="text/javascript" src="node_modules/react/dist/react.js"></script>
<script type="text/javascript" src="build/app.js"></script>
```

</cn>

Note: Rendering React onto body would remove elements added by BrowserSync. To avoid that, we render the App component onto the element `#root`.

<cn>
注意：使用 React 渲染 body 时会删除 BrowserSync 所添加的元素。为了避免这种情况，我们用 App 组件替换掉 `#root` 中的内容。
</cn>

Your result:

![](react-first-render.jpg)

<cn>

你的结果：

![](react-first-render.jpg)

</cn>

### Exercise: Get LiveReload To Work

<cn>
### 练习：让 LiveReload 起作用
</cn>

The project configuration is slightly more complicated. We need to:

1. Automatically compile `js/app.jsx`.
2. Create the `make js` task to compile jsx.
3. Modify BrowserSync to watch the compiled jsx.

<cn>
项目的配置要变的稍微的复杂些了。我们需要：

1. 自动编译 `js/app.jsx`。
2. 创建 `make js` 任务来编译 JSX。
3. 修改 BrowserSync 来监控编译后的 JSX 文件。
</cn>

Add the `--watch` option so Babel would automatically recompile `js/app.jsx` when it changed:

```
$ babel --watch js/app.jsx --out-file build/app.js
```

<cn>
在 Babel 命令后加上 `--watch` 选项可以自动的重新编译更改后的 `js/app.jsx` 文件：


```
$ babel --watch js/app.jsx --out-file build/app.js
```

</cn>

Open three terminals, and run:

1. `make css` - rebuild css.
2. `make js` - rebuild jsx.
3. `make server` - reload the browser.

<cn>
打开三个终端，分别运行：

1. `make css` － 重新编译 CSS。
2. `make js` － 重新编译 JSX。
3. `make server` － 重载浏览器。

</cn>

### Exercise: `make all`

<cn>
### 练习：`make all`
</cn>

It's annoying to have to open three terminals to run three different processes. You could add an `all` task that runs all three at the same time:

```
.PHONY: all
all:
  (make css & make js & make server & wait)
```

<cn>
每次我们都需要打开三个终端来运行三个不同的进程，这让我们很反感和困扰。你可以在 Makefile 中添加一个 `all` 的任务来让三个进程一起执行：

```
.PHONY: all
all:
  (make css & make js & make server & wait)
```

</cn>

Hit `Ctrl-C` to stop them all at the same time.

<cn>
使用 `Ctrl-C` 来停止。
</cn>

# React Components

<cn>
# React 组件
</cn>

```
npm install react@0.13.3 --save-dev
```

<cn>
```
npm install react@0.13.3 --save-dev
```
</cn>

Instead of one big render method, you should break a page into multiple components so each one is responsible for a smaller part of the page. Generally speaking there are three uses for components:

+ If a UI element is reused, then it should be a component.
  + `.cart-item` and `.product`.
+ A container that holds a list of items should be its own component.
  + `.products` is a component that holds a list of `.product` components.
+ If a a component gets too big or to complicated, you could try to break it into subparts.

<cn>
区别于一个冗长的 Render 方法，你可以将一个页面拆分成很多小的组件，每个组件又正好负责页面中的一小部分。通常来说组件有三种用途：

+ 如果一个 UI 元素被复用，那么它应该变成一个组件。
  + `.cart-item` 和 `.product`。
+ 那些容纳元素列表的容器应当变成一个组件：
  + `.products` 是一个容纳 `.product` 组件的容器，本身也应当时一个组件。
+ 如果一个组件过于庞大或者复杂，你应该把它拆分成几个小的子组件。
</cn>

You should try to keep the `render` method between 50~80 lines of code.

It's up to you to judge whether it's worth the effort to turn a fragment of HTML into its own component. Sometimes it might make more sense to create a helper method instead.

<cn>
你应该尝试着将 `render` 方法保持在 50 到 80 行左右。

你也可以决定将部分 HTML 片段转化成组件这件事是否值得去做。有时可能创建一个辅助的方法会变得有效一点。
</cn>

Consider the `.bg` element:

```html
<div class="site">
  <div class="bg">
    <div class="bg__img">
    </div>
  </div>
  <!-- other stuff -->
</div>
```

<cn>
考虑 `.bg` 元素：

```html
<div class="site">
  <div class="bg">
    <div class="bg__img">
    </div>
  </div>
  <!-- 其他东西 -->
</div>
```

</cn>

You could create a component:

```html
<div className="site">
  <Background/>
</div>
```

<cn>
你可以创建一个组件：

```html
<div className="site">
  <Background/>
</div>
```
</cn>

Or you could create a helper method:

```html
<div className="site">
  {this.renderBackground()}
</div>
```

<cn>
或者，你可以创建一个辅助的方法：

```html
<div className="site">
  {this.renderBackground()}
</div>
```
</cn>

# Reactify The App

<cn>
# React 化 App
</cn>

The React components we'll create mirror closely the BEM components we already have. Let's start with the site layout.

<cn>
我们所创建的 React 组件和之前已经存在的 BEM 元素是很类似的。让我们从网站的布局先开始。
</cn>

### Exercise: The App should define the layout

<cn>
### 练习：App 决定了布局
</cn>

Your App component should have a structure like this:

```js
<div class="site">
  <div class="bg">
    <div class="bg__img">
    </div>
  </div>

  <div class="site__main">
    <div class="site__left-sidebar">
      <!-- <SiteTitle/> -->
    </div>

    <div class="site__content">
      <!-- <Products/> -->

    </div> <!-- site__content -->
  </div> <!-- site__main -->

  <div class="site__right-sidebar">

  <!-- <Cart/> -->
  <!-- <Checkout/> -->

  </div> <!-- site__right-sidebar -->

  <a class="site__right-sidebar-toggle">
    <img src="img/arrow-icon.svg"/>
  </a>

</div> <!-- site -->
```

<cn>
你的 App 组件应该像这样类似的结构：

```js
<div class="site">
  <div class="bg">
    <div class="bg__img">
    </div>
  </div>

  <div class="site__main">
    <div class="site__left-sidebar">
      <!-- <SiteTitle/> -->
    </div>

    <div class="site__content">
      <!-- <Products/> -->

    </div> <!-- site__content -->
  </div> <!-- site__main -->

  <div class="site__right-sidebar">

  <!-- <Cart/> -->
  <!-- <Checkout/> -->

  </div> <!-- site__right-sidebar -->

  <a class="site__right-sidebar-toggle">
    <img src="img/arrow-icon.svg"/>
  </a>

</div> <!-- site -->
```
</cn>

First, translate the HTML into JSX.

Then you need to create 4 new React components: "SiteTitle", "Products", "Cart", and "Checkout". Leave these components empty for now. We'll fill in the details later.

<cn>
首先，把 HTML 转换成 JSX。

然后你应该建立四个 React 组件：「SiteTitle」，「Products」，「Cart」还有「Checkout」。现在先不管组件内部是空的情况，我们一会儿再来补充细节。
</cn>

Note: If you are lazy like me, you'd probably Google for a tool that automatically converts HTML to React. See: [HTML to JSX Compiler](https://facebook.github.io/react/html-jsx.html)

<cn>
注意：如果你和我一样懒，也许你可以 Google 来找到一个工具可以帮助你自动地将 HTML 转换成 React，可以查看 [HTML to JSX Compiler](https://facebook.github.io/react/html-jsx.html)。
</cn>

Note: In JSX, comments must be a JavaScript comment: `{/* I am a JavaScript comment */}`. Check the Babel compilation result to see how it works.

<cn>
注意：在 JSX 中，注释必须是以 JavaScript 格式存在：`{/* 我是一个 JavaScript 注释 */}`。检查一下 Babel 的编译情况来查看它是如何正常工作的。
</cn>

Your result:

![](layout-components.jpg)

<cn>
你的结果：

![](layout-components.jpg)
</cn>

# Componentize Products

<cn>
# 商品组件化
</cn>

We'll do this in two steps:

1. Display a single product.
2. Display list of different products.

<cn>
我们把这件事拆分为两步：

1. 显示单个商品。
2. 显示不同的商品列表。
</cn>

### Exercise: Create the Product component

<cn>
### 练习：创建商品组件
</cn>

Let's convert the `.product` HTML to a component:

```html
<div class="product">
  <div class="product__display">
    <div class="product__img-wrapper">
      <img class="product__img" src="img/shoe1.jpg"/>
    </div>

    <a class="product__add">
      <img class="product__add__icon" src="img/cart-icon.svg"/>
    </a>

    <div class="product__price">
      $299
    </div>
  </div>

  <div class="product__description">
    <div class="product__name">
      Marana E-Lite
    </div>

    <img class="product__heart" src="img/heart.svg"/>
  </div>
</div> <!-- product -->
```

<cn>
将 `.product` HTML 转化成一个组件：

```html
<div class="product">
  <div class="product__display">
    <div class="product__img-wrapper">
      <img class="product__img" src="img/shoe1.jpg"/>
    </div>

    <a class="product__add">
      <img class="product__add__icon" src="img/cart-icon.svg"/>
    </a>

    <div class="product__price">
      $299
    </div>
  </div>

  <div class="product__description">
    <div class="product__name">
      Marana E-Lite
    </div>

    <img class="product__heart" src="img/heart.svg"/>
  </div>
</div> <!-- product -->
```
</cn>

We want the Product component to render the product object:

```js
let product = {
  name: "Jameson Vulc",
  price: 64.99,
  imagePath: "img/shoes/jameson-vulc-brown-gum-orig.png",
  gender: "man",
};
```

<cn>
我们想要使用 Product 组件来重新渲染这个元素：

```js
let product = {
  name: "Jameson Vulc",
  price: 64.99,
  imagePath: "img/shoes/jameson-vulc-brown-gum-orig.png",
  gender: "man",
};
```
</cn>

For now, we'll hardwire the `Products` container to render a single product:

```js
let Products = React.createClass({
  render() {
    let product = {
      name: "Jameson Vulc",
      price: 64.99,
      imagePath: "img/shoes/jameson-vulc-brown-gum-orig.png",
      gender: "man",
    };

    return (
      <div className="products">
        <Product product={product}/>
      </div>
    );
  }
});
```

<cn>
现在开始，`Products` 容器中只渲染单个物品：

```js
let Products = React.createClass({
  render() {
    let product = {
      name: "Jameson Vulc",
      price: 64.99,
      imagePath: "img/shoes/jameson-vulc-brown-gum-orig.png",
      gender: "man",
    };

    return (
      <div className="products">
        <Product product={product}/>
      </div>
    );
  }
});
```

</cn>

The most important piece of code in the `Products` container is:

```js
<Product product={product}/>
```

<cn>
`Products` 中最重要的一部分代码就是：

```js
<Product product={product}/>
```

</cn>

It creates a Product component, with `product` as a property. The `Product` component looks like:

```js
let Product = React.createClass({
  render() {
    // This component requires the `product` property.
    let {name,price,imagePath} = this.props.product;

    ...
  }
});
```

<cn>
它创建了一个 Product 组件，将 `product` 作为一个属性。`Product` 组件看上去像这样：

```js
let Product = React.createClass({
  render() {
    // 这个组件需要 `product` 属性。
    let {name,price,imagePath} = this.props.product;

    ...
  }
});
```
</cn>

Your result:

![](product-component.jpg)

<cn>
你的结果：

![](product-component.jpg)
</cn>

### Exercise: Render multiple products

<cn>
### 练习：渲染多个商品
</cn>

Now we want the `Products` container to render a list of products, not just one. Add the following as a global variable to `js/app.jsx`:

```js
let products = {

  "jameson-vulc": {
    id: "jameson-vulc",
    name: "Jameson Vulc",
    price: 64.99,
    imagePath: "img/shoes/jameson-vulc-brown-gum-orig.png",
    gender: "man",
  },

  ...

  "corby-womens-2": {
    id: "corby-womens-2",
    name: "Corby Women's",
    imagePath: "img/shoes/corby-womens-2-tan-white-orig.png",
    price: 44.99,
    gender: "woman",
  },
};
```

<cn>
现在我们想让 `Products` 容器渲染商品的列表，而不是单单一个商品。将下面的代码添加到 `js/app.jsx` 中作为全局变量：

```js
let products = {

  "jameson-vulc": {
    id: "jameson-vulc",
    name: "Jameson Vulc",
    price: 64.99,
    imagePath: "img/shoes/jameson-vulc-brown-gum-orig.png",
    gender: "man",
  },

  ...

  "corby-womens-2": {
    id: "corby-womens-2",
    name: "Corby Women's",
    imagePath: "img/shoes/corby-womens-2-tan-white-orig.png",
    price: 44.99,
    gender: "woman",
  },
};
```

</cn>

Download the full list of products: [products.js](products.js)

<cn>
下载所有商品的列表：[products.js](products.js)
</cn>

Note: Remember that a JSX element is just a value. You can put it in a variable, return it from a function, or put it in an array:

```js
let children = [
  <div>1</div>,
  <div>2</div>,
  <div>3</div>,
];

return (
  <div className="container">
    {children}
  </div>
);
```

<cn>
注意：记住 JSX 元素只是一个值。你可以把它赋给一个变量，做为函数的返回值，或者把它放进数组中：

```js
let children = [
  <div>1</div>,
  <div>2</div>,
  <div>3</div>,
];

return (
  <div className="container">
    {children}
  </div>
);
```

</cn>

Your result:

![](multiple-products.jpg)

<cn>
你的结果：

![](multiple-products.jpg)
</cn>

Note: If you look at the console, React complains that `key` prop is missing.

<cn>
注意：在终端中，React 会提醒缺少 `key` 属性。
</cn>

```
Warning: Each child in an array or iterator should have a unique "key" prop. Check the render method of Products. See https://fb.me/react-warning-keys for more information.
```

<cn>
```
警告：数组和迭代器中的每一个元素都应该有一个独一无二的 `key` 属性。检查 Products 的 render 方法。可以查看 https://fb.me/react-warning-keys 了解更多。
```
</cn>

We'll fix this problem in the next section.

![](unique-key-prop-warning.jpg)

<cn>
下一节我们将修复这个问题：

![](unique-key-prop-warning.jpg)
</cn>

# Component Key

<cn>
# 组件键值
</cn>

Suppose you have 6 divs:

```html
<div>1</div>
<div>2</div>
<div>3</div>
<div>4</div>
<div>5</div>
<div>6</div>
```

<cn>
假设你有 6 个 div：

```html
<div>1</div>
<div>2</div>
<div>3</div>
<div>4</div>
<div>5</div>
<div>6</div>
```
</cn>

And you want to move the first div to the end:

```html
<div>2</div>
<div>3</div>
<div>4</div>
<div>5</div>
<div>6</div>
<div>1</div>
```

<cn>
然后你想把第一个元素移动到最后一个：

```html
<div>2</div>
<div>3</div>
<div>4</div>
<div>5</div>
<div>6</div>
<div>1</div>
```
</cn>

If you don't use the `key` prop, React would update the divs like this:

1. update div 1 to be 2
2. update div 2 to be 3
3. update div 3 to be 4
4. update div 4 to be 5
5. update div 5 to be 6
6. update div 6 to be 1

<cn>
如果我们没有 `key` 属性，React 会这样更新 div：

1. 更新 div 1 到 2
2. 更新 div 2 到 3
3. 更新 div 3 到 4
4. 更新 div 4 到 5
5. 更新 div 5 到 6
6. 更新 div 6 到 1

</cn>

The `key` prop gives a child a unique identity, so React knows that that it can reorder these children rather than updating them.

<cn>
`key` 属性给每一个子元素一个独一无二的辨识属性。这能让 React 知道我们只是改变了子元素的顺序而不是在更新它们。
</cn>

Let's consider this example in actual code. `Numbers` is component that moves the first number to the end, once every second:

```js
let Numbers = React.createClass({
   getInitialState() {
    return {
      numbers: [1,2,3,4,5,6,7,8,9,0]
    };
  },

  // Cycle numbers every second.
  componentDidMount() {
    setInterval(this.cycleNumbers,1000);
  },

  // Move the first number to the end.
  cycleNumbers() {
    let {numbers} = this.state;
    let copiedNumbers = shuffle(numbers.slice());

    let firstNumber = copiedNumbers.shift();
    copiedNumbers.push(firstNumber);

    // Change this.state.numbers. Would cause `render` to calculate the new virtual DOM.
    this.setState({numbers: copiedNumbers});
  },

  render() {
    let {numbers} = this.state;

    let children = numbers.map(n => {
      return (
        <div>{n}</div>
      );
    });

    return (
      <div>{children}</div>
    )
  }
});
```

<cn>
让我们看一下下面这个实际运用中的例子。`Numbers` 是一个每隔一秒把第一个数字移动到最后一个的组件：

```js
let Numbers = React.createClass({
   getInitialState() {
    return {
      numbers: [1,2,3,4,5,6,7,8,9,0]
    };
  },

  // 每一秒对数组中的数字循环。
  componentDidMount() {
    setInterval(this.cycleNumbers,1000);
  },

  // 移动第一个数字到最后一个。
  cycleNumbers() {
    let {numbers} = this.state;
    let copiedNumbers = shuffle(numbers.slice());

    let firstNumber = copiedNumbers.shift();
    copiedNumbers.push(firstNumber);

    // 改变 this.state.numbers。可以让 `render` 计算生成新的 Virtual DOM。
    this.setState({numbers: copiedNumbers});
  },

  render() {
    let {numbers} = this.state;

    let children = numbers.map(n => {
      return (
        <div>{n}</div>
      );
    });

    return (
      <div>{children}</div>
    )
  }
});
```
</cn>

[Codepen Demo](http://codepen.io/hayeah/full/dYvNXy)

<cn>

[Codepen Demo](http://codepen.io/hayeah/full/dYvNXy)

</cn>

But if you look in Chrome Inspector, you'd see that all the divs are being updated:

<video src="number-cycle-updates-all.mp4" controls autoplay loop></video>

<cn>
但是打开 Chrome 中的审查元素，你会看到所有的 div 都被更新了。

<video src="number-cycle-updates-all.mp4" controls autoplay loop></video>
</cn>

Instead of updating all the divs, we want them to be reordered. To enable reordering, add the `key` property to each div:

```js
let children = numbers.map(n => {
  return (
    <div key={n}>{n}</div>
  );
});
```

<cn>
我们并不想更新所有的 div，而是想要让这些数字重新排序。为了实现这点，我们给每个 div 加上 `key` 属性：


```js
let children = numbers.map(n => {
  return (
    <div key={n}>{n}</div>
  );
});
```
</cn>

[Codepen Demo](http://codepen.io/hayeah/full/KdWWPb/)

<cn>

[Codepen Demo](http://codepen.io/hayeah/full/KdWWPb/)

</cn>

<video src="numbers-cycle-movement.mp4" controls autoplay loop></video>

<cn>

<video src="numbers-cycle-movement.mp4" controls autoplay loop></video>

</cn>

Each key uniquely identifies a div element, so React is now able to reorder them. You should always provide the `key` property to an array of children.

<cn>
每个键值代表且仅代表了一个 div 元素。这样 React 就能对它们重新排序了。对于包含了子元素的数组，你总是需要提供 `key` 属性。
</cn>

### Exercise: Fix the "should have a unique key prop" error

<cn>
### 练习：修复「需要拥有独一无二的 key 属性」的错误
</cn>

# Reactify Cart

<cn>
# React 化购物车
</cn>

Now let's create the `Cart` component. It follows a similar structure as `Products`.

<cn>
现在让我们创建 `Cart` 组件。结构和 `Products` 类似。
</cn>

### Exercise: Implement the Cart component

<cn>
### 练习：实现购物车组件
</cn>

Again, let's fake the data (for now) by declaring a global variable:

```js
let cartItems = {
  "jameson-vulc": {
    id: "jameson-vulc",
    quantity: 1,
  },

  "scout-womens-6": {
    id: "scout-womens-6",
    quantity: 2,
  },
};
```

<cn>

同样地，我们暂时先在全局制造一下假数据：

```js
let cartItems = {
  "jameson-vulc": {
    id: "jameson-vulc",
    quantity: 1,
  },

  "scout-womens-6": {
    id: "scout-womens-6",
    quantity: 2,
  },
};
```

</cn>

Remember to add a "x 2" if the quantity is greater than 1.

<cn>
如果商品数量大于 1 时，记得添加「x 2」。
</cn>

Your result:

![](cart-reactified.jpg)

<cn>
你的结果：

![](cart-reactified.jpg)
</cn>

# Reactify Checkout

<cn>
# React 化结算部分
</cn>

Finally, let's create the `Checkout` component. For now, we'll just calculate the subtotal of all the cart items. Later we'll implement coupon verification and discount.

<cn>
最后我们来创建 `Checkout` 组件。现在我们要计算在购物车中的商品的总价。之后我们会实现验证优惠券和折扣优惠。
</cn>

### Exercise: Implement the Checkout component

<cn>
### 练习：实现结算组件
</cn>

Your result:

![](checkout-no-coupon.jpg)

<cn>
你的结果：

![](checkout-no-coupon.jpg)
</cn>

# Product Quantity Adjustment

<cn>
# 商品数量调整
</cn>

If a product is already in the shopping cart, it should display the quantity control instead of the add-to-cart button:

![](products-with-quantity-control.jpg)

<cn>
如果一个商品已经在购物车中了，那么在商品的下面应该显示调整数量的按钮而不是添加到购物车的按钮：

![](products-with-quantity-control.jpg)
</cn>

### Exercise: Create QuantityControl Component

<cn>
### 练习：创建数量调整组件
</cn>

Start with this HTML:

```html
<div className="adjust-qty">
  <a className="adjust-qty__button">-</a>
  <div className="adjust-qty__number">{quantity}</div>
  <a className="adjust-qty__button">+</a>
</div>
```

<cn>
从这个 HTML 开始：


```html
<div className="adjust-qty">
  <a className="adjust-qty__button">-</a>
  <div className="adjust-qty__number">{quantity}</div>
  <a className="adjust-qty__button">+</a>
</div>
```

</cn>

The `QuantityControl` component should accept a cartItem object as property:

```html
<QuantityControl item={item}/>
```

<cn>
数量控制的组件会将购物车中的一个物品做为其属性：

```html
<QuantityControl item={item}/>
```

</cn>

The component looks different depending on where it is used. Implement the `variant` property to determine which style to use:

```html
<QuantityControl item={item} variant="gray"/>
```

<cn>
在不同的地方，这个组件看上去也不太相同。实现 `variant` 属性来决定使用哪种样式：

```html
<QuantityControl item={item} variant="gray"/>
```

</cn>

The "gray" variant is already styled with the `.adjust-qty--gray` BEM variant.

<cn>
「gray」已经通过 BEM 中的 `.adjust-qty--gray` 样式所定义了。
</cn>

# Component Lifecycle

<cn>
# 组件的生命周期
</cn>

The [component lifecycle](https://facebook.github.io/react/docs/working-with-the-browser.html#refs-and-finddomnode) hooks are useful when you need to interface with code that's outside of React views. For example:

+ Subscribe to events handler.
+ Set a timer.
+ Access to the actual browser DOM.

<cn>
当你需要和 React 视图外部的代码进行交互时，[组件的生命周期](https://facebook.github.io/react/docs/working-with-the-browser.html#refs-and-finddomnode) 将会非常的有用。例如：

+ 订阅并且响应处理时间。
+ 设置一个定时器。
+ 获得浏览器中实际的 DOM。
</cn>

Use the following hooks for setup and teardown:

+ `componentDidMount` - Any setup code.
+ `componentDidUpdate` - Do something whenever the view changes.
+ `componentWillUnmount` - Any cleanup code.

<cn>
使用下面的钩子（Hooks）能够方便我们在开始和结束阶段调用：

+ `componentDidMount` － 加载监听时的代码。
+ `componentDidUpdate` － 视图变化时的更新代码。
+ `componentWillUnmount` － 结束监听时的清理代码。
</cn>

## Accessing Browser DOM

<cn>
## 获得浏览器 DOM
</cn>

Many useful JavaScript plugins are not written with React. You often need access to the actual browser DOM in order to initialize these JavaScript plugins.

<cn>
很多有用的 JavaScript 插件并不是用 React 写的。所以你经常需要获得浏览器中实际的 DOM 来初始化这些 JavaScript 插件。
</cn>

You'd usually add DOM related JavaScript in the `componentDidMount` hook. When this method is called, you can be sure that the browser DOM element exists.

<cn>
你通常会在 `componentDidMount` 这个钩子方法中添加 DOM 所对应的 JavaScript。当这个方法被调用时，你要确保浏览器 DOM 元素已经存在。
</cn>

One way to enable PerfectScroll for a content container is to add an id:

```js
let ComponentA = React.renderClass({
  componentDidMount() {
    let $content = document.querySelector("#content");
    Ps.initialize($content);
  },

  render() {
    return (
      <div id="content">
        a
        lot
        of
        content
      </div>
    );
  },
});
```

<cn>
有一种能让 PerfectScroll 对一个容器有效的方法就是把它加到 id 中：

```js
let ComponentA = React.renderClass({
  componentDidMount() {
    let $content = document.querySelector("#content");
    Ps.initialize($content);
  },

  render() {
    return (
      <div id="content">
        a
        lot
        of
        content
      </div>
    );
  },
});
```

</cn>

There is a big problem though. Once the id `content` is used in ComponentA, we can't use it anywhere else. In a big project with many components and many team members, the team need to adopt a naming convention to avoid name collisions. BEM is one way to solve this problem:

```js
let ComponentA = React.renderClass({
  componentDidMount() {
    let $content = document.querySelector(".js-componentA__content");
    Ps.initialize($content);
  },

  render() {
    return (
      <div className="js-componentA__content">
        a
        lot
        of
        content
      </div>
    );
  },
});
```

<cn>
还有一个比较大的问题。一旦某个 id（例如 `content`）被用在了某个组件上（例如 `ComponentA`)，那么我们就没办法在其他地方使用这个 id 了。在一个很大的项目中，我们可能会有很多的组件，也会有很多一起干活的项目成员，整个团队需要使用一种命名约定来避免这种命名的冲突。BEM 就是一种用来解决这个问题的方法：

```js
let ComponentA = React.renderClass({
  componentDidMount() {
    let $content = document.querySelector(".js-componentA__content");
    Ps.initialize($content);
  },

  render() {
    return (
      <div className="js-componentA__content">
        a
        lot
        of
        content
      </div>
    );
  },
});
```

</cn>

But if there are two of same components, like this:

```js
let App = React.renderClass({
  render() {
    <div>
      <ComponentA/>
      <ComponentA/>
    </div>
  }
});
```

<cn>

但是如果有两个相同的组件，像这样：

```js
let App = React.renderClass({
  render() {
    <div>
      <ComponentA/>
      <ComponentA/>
    </div>
  }
});
```

</cn>

Then ComponentA's `componentDidMount` method would be called twice, once for each instance of componentA.

React's component lifecycle is designed for attaching JavaScript to one DOM at a time. The special `ref` property is like the `id` property, but only visible within a component.

<cn>
那么 ComponentA 的 `componentDidMount` 方法将会被调用两次。对于每个 ComponentA 的实例都会被调用一次。

React 中的组件生命周期就是为了将 JavaScript 一次性地插入到 DOM 中。这个特殊的 `ref` 属性就和 `id` 属性一样，但是只对特定的单个组件有效。
</cn>

In general, where you might've used `id`, you can replace that with `ref`:


```js
let ComponentA = React.renderClass({
  componentDidMount() {
    let $content = React.findDOMNode(this.refs.content);
    Ps.initialize($content);
  },

  render() {
    return (
      <div ref="content">
        a
        lot
        of
        content
      </div>
    );
  }
});
```

<cn>
通常，你可以使用 `ref` 来替代需要使用到 `id` 的地方：

```js
let ComponentA = React.renderClass({
  componentDidMount() {
    let $content = React.findDOMNode(this.refs.content);
    Ps.initialize($content);
  },

  render() {
    return (
      <div ref="content">
        a
        lot
        of
        content
      </div>
    );
  }
});
```

</cn>

Since `ref` is local within a componentA, other components can also use the 'content' reference name:

```js
let AnotherComponent = React.renderClass({
  render() {
    return (
      <div ref="content">
        a different component with different content.
      </div>
    );
  }
});
```

<cn>
既然 `ref` 只局部作用于某个特定的 ComponentA，那么在其他的组件中也可以使用 `content` 这个 `ref` 名字。

```js
let AnotherComponent = React.renderClass({
  render() {
    return (
      <div ref="content">
        a different component with different content.
      </div>
    );
  }
});
```

</cn>

See: [Refs and findDOMNode()](https://facebook.github.io/react/docs/working-with-the-browser.html#refs-and-finddomnode)

<cn>

阅读：[Refs and findDOMNode()](https://facebook.github.io/react/docs/working-with-the-browser.html#refs-and-finddomnode)

</cn>

### Exercise: Enable PerfectScroll

<cn>
### 练习：使用 PerfectScroll
</cn>

Add more cart items so the shopping cart overflows:

```js
let cartItems = {
  "jameson-vulc": {
    id: "jameson-vulc",
    quantity: 1,
  },

  "marana-x-hook-ups": {
    id: "marana-x-hook-ups",
    quantity: 2,
  },

  "scout-womens-6": {
    id: "scout-womens-6",
    quantity: 2,
  },

  "scout-womens-coco-ho-5": {
    id: "scout-womens-coco-ho-5",
    quantity: 1,
  },

  "jameson-2-womens-8": {
    id: "jameson-2-womens-8",
    quantity: 1,
  },
};
```

<cn>
向购物车里增加点商品能让其溢出：

```js
let cartItems = {
  "jameson-vulc": {
    id: "jameson-vulc",
    quantity: 1,
  },

  "marana-x-hook-ups": {
    id: "marana-x-hook-ups",
    quantity: 2,
  },

  "scout-womens-6": {
    id: "scout-womens-6",
    quantity: 2,
  },

  "scout-womens-coco-ho-5": {
    id: "scout-womens-coco-ho-5",
    quantity: 1,
  },

  "jameson-2-womens-8": {
    id: "jameson-2-womens-8",
    quantity: 1,
  },
};
```

</cn>

Add a `ref` property to initialize the PerfectScroll plugin.

<cn>
初始化 PerfectScroll 时，增加一个 `ref` 属性。
</cn>

# Summary

<cn>
# 总结
</cn>

We've translated the `buyshoes` page to React. The page now uses real data, but it's still not that much different from a static page. We can't add or remove shopping cart items yet.

<cn>
我们已经将 `buyshoes` 页面 React 化了。这个页面现在用着真实的数据，但是它和静态页面并无太大差异。我们还无法从购物车内添加或者移除商品。
</cn>

But in fact, the "static page" is already 80% of the final product. When we use the flux architecture to implement adding & removing shopping cart items, you'll be surprised by how little we have to modify the existing views.

<cn>
但是事实上，这个「静态页面」离最后的完成还只剩下 20% 了！当我们接下来使用 flux 来实现增加和移除购物车中的商品时，你会惊喜地发现：这时我们只需改变极少的代码了。
</cn>

Until then, here's a quick summary of the important ideas you should know from this lesson:

+ JSX is a syntactic sugar for `React.createElement`.
  + Can put virtual DOM created JSX in variables, arrays, or return from function.
+ Virtual DOM construction should be side-effects free.
  + Given the same `this.state` and `this.props`, `render` should calculate the same virtual DOM.
+ Parent component passes data to child components using the `this.props` property.
  + Use ES6 destructuring to make your code more concise.
+ Use `key` to give child components a unique identity so React can reorder them.
+ Use `componentDidMount` to do setup after browser DOM is ready.
+ Replace `id` with `ref`.

<cn>
最后我们来快速地总结一下这节课的主要要点：

+ JSX 是一个辅助于 `React.createElement` 的语法糖。
  + JSX 能够将虚拟的 DOM 作为变量、数组元素或者函数的返回值。
+ Virtual DOM 的构建应该是无副作用的。
  + 给定相同的 `this.state` 和 `this.props`，`render` 方法能渲染出相同的 Virtual DOM 树。
+ 父组件通过 `this.props` 属性对子组件传值。
  + 使用 ES6 中的析构函数让代码更佳简洁。
+ 使用 `key` 给子组件一个独一无二的辨识属性，这能让 React 知道要对它们重新排序。
+ 在 DOM 加载完毕后，使用 `componentDidMount` 方法来做初始化。
+ 使用 `ref` 替代 `id`。
</cn>

