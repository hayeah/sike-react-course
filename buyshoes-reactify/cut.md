+ Useful ES6 Features
  + `let` everywhere to replace `var`.
  + `=>` everywhere to replace `function`.
  + the class syntax
    + the object function notation: `render() { ... }`
  + `let {a,b,c} = obj`
  + Promise & await


# ES6/7 Features And Fixes

You can see the new language features that Babel supports, in its list of [transformers](https://babeljs.io/docs/advanced/transformers/).

There are quite a few. But there are four categories.

+ Fix language design issues.


language fixes
  + let
  + =>
  + for ... of


codify common practices in an officially sanctioned syntax
  + class
  + Default arguments. Spread arguments.
  + module


syntatic sugar
  + Spread operator
  + Destructuring


language extensions
  + Promise & Await for async.
  + List comprehension
  + decorators




# JSX & Virtual DOM

Without further ado, let's see React in action.

```js
{
  "name": "Howard",
  "time": "Thu Oct 01 2015 21:45:17 GMT+0800 (CST)",
}
```

The result you want:

```html
<div>
  <p>Hello, <b>Howard</b></p>
  <p>Thu Oct 01 2015 21:45:17 GMT+0800 (CST)</p>
</div>
```

Template:

```html
<div>
  <p>Hello, <b>{name}</b></p>
  <p>{time}</p>
</div>
```





Is a shopping cart a web page or an app? Is it ok to wait for the page to reload everytime the user adds a new product?



Works fine for many use cases. News sites. Blogs. Media stuff.

As is, doesn't work as well for sites that feel like apps. User demand faster UI response.

Let's see what a hypothetical "upgrade path" would look like for a web-app to become more responsive.





{

I might not have to launch into an editorial here. The key idea is the equivalence to setting innerHTML.

+ how do I motivate that setting innerHTML is a good idea?
  + setting innerHTML is actually pretty fast.
  + event listeners are removed. all states are gone.
+ why move from innerHTML to virtual dom?
  + also: batched update. one reconciliation per multiple changes.


how can i illustrate the horror that is DOM?

consider the addProductToCart function.

+ client-server request-response cycle.
+ pjax.
+ ajax.
+ optimistic UI update.
  + here we have to introduce client-side rendering.
+ react. take us back to client-server programing model, but everything is at client-side now.

}

But why do we want to rewrite a perfectly good HTML page in React? Use the server-client architecture, the browser makes a request, the server responds with the updated page, then the browser does a full-page refresh.

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

