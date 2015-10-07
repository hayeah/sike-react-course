# Reactify The Buy Shoes Page



In this lesson we are going to rewrite the `buyshoes` page in React. We'll cover the following:

+ Compile ES6 & JSX with Babel.
+ Virtual DOM with JSX.
+ Writing React components.
+ Breaking an app into CommonJS modules.
+ Bundling modules into a single file with Browserify.

# Web Or Mobile? That's The Question

But why do we even want to rewrite a perfectly good HTML page in React? Use the server-client architecture, the browser makes a request, the server responds with the updated page, then the browser does a full-page refresh.

Done. No JavaScript at all.

On the

People expects page refresh.


Tricks to reduce latency time:

full-page refresh
  + the page flickers. reload HTML, CSS, JS.

```
document.location.reload()
```

pjax/turbolink
  + minimize page flickering. only reload the main content.

```
document.querySelector("#content").innerHTML = newContentHTML;
```

ajax
  + reload a tiny part of the content.


```js
updateTodo(todo1,function(updatedTodoHTML) {
  document.querySelector(".todo[todoID=1]").innerHTML = updatedTodoHTML;
});
```

```js
createTodo(newTodo,function(updatedTodoListHTML) {
  document.querySelector("#todolist").innerHTML = updatedTodoListHTML;
  // remmeber to update todo counter

  var $todos = document.querySelectorAll(".todo");

  var $todoCounter = document.querySelector(".todoCounter");
  $todoCounter.innerHTML = $todos.length;

});
```


optimistic update gets complicated. at this point, you are kinda like, fuck, i give up.

```js
// this is the todo we are updating
var $todo = document.querySelector(".todo[data-todo-id=1]")
var oldTodoHTML = $todo.innerHTML;

// Client needs to know how to render todo...
todoHTML = renderTodo(todo);

updateTodo(todo,function(error) {
  if(error) {
    // revert on failure
    $todo.innerHTML = oldTodoHTML;
    alert("update error!");
  }
});
```


```js
// global variable that stores all the todos...
var todos = ...;


function createTodo(todo) {
  todos.push(todos);

  var $todolist = document.querySelector(".todolist");

  // Client needs to know how to render todo...
  var newTodoHTML = renderTodo(newTodo);

  var $newTodo = blah blah
  $todoList.appendChild($newTodo);

  var $todoCounter = document.querySelector(".todoCounter");
  $todoCounter.innerHTML = todos.length;

  createTodo(todo,function(error,newTodo) {
    if(error) {
      // revert on failure
      todos.pop();

      // update UI
      $todolist.removeChild($newTodo);
      $todoCounter.innerHTML = todos.length;

      alert("create error!");
    }
  });
}
```

So it's getting complicated. You need frontend MVC to help you.

Why do we want to add so much client-side complexity

+ Web-client is one of many. There is one shared API, but no "server-side rendering".


### Clone Project

Use the starter kit.

# Install React And Babel


# innerHTML

Use a template language and replace innerHTML.

Why not innerHTML all the way? Indeed some people do that!
  3 granularities of server-side rendering: full-page refresh. pjax. ajax partial.

Why taking the server-side rendering model to the client-side?
  + offline usage.
  + instantaneous response. (latency compensation. update the UI immediately).
    + DOM is hard!
  + "is it ok to feel like a web-app". converse: "do i need it to feel like a mobile app"



# JSX & Virtual DOM



```
var name = "Howard";

var view = (
  <div className="hello">
    Hello, <b className="hello__name">{name}</b>
  </div>
);

React.render(view,document.body);
```

# DOM Reconciliation

+ vs body.innerHTML

```js
sayHello(name) {
  // ...
}
```

# ES6+Babel

How to compile React JSX to JS. Put everything in app for now.

Install babel:

```
npm install babel@5.8.23 --save-dev
```


```js
var App = React.createClass({
  render: function() {
    return (
      <div className="site">
        <h1>Buy Some Shoes!!!</h1>
      </div>
    );
  },
});

window.onload = () => {
  React.render(<App/>,document.querySelector("#root"));
}


```

compiles to

```js
var App = React.createClass({
  // hmm. magically added. interesting...
  displayName: "App",

  render: function render() {
    return React.createElement(
      "div",
      { id: "content" },
      React.createElement(
        "h1",
        null,
        "Buy Some Shoes!"
      )
    );
  }
});

React.render(React.createElement(App, null), document.querySelector("#content"));
```

(it's annoying that we can't render directly to body because of DOM injection by browser-sync)

# Build app.jsx

js/app.jsx -> build/app.js

js/app.jsx -> build/app.js

Create a new Makefile task.

```
mkdir -p build
babel --watch js/app.jsx  -o build/app.js
```

```
<script type="text/javascript" src="build/app.js"></script>
```

# link react

```html
<script type="text/javascript" src="node_modules/react/dist/react.js"></script>
```

+ Mounting React component.
+ Use <script> for now...

Install react (runtime support):

```
npm install react@0.13.3
```

Play with React manually in browser.

renderToString & render.



# browserify & commonjs

build/app.js -> bundle/app.js

sourcemap