# Getting Started With Flux

Our shopping cart is just a static page. In this lesson, we'll use Flux to make it possible to add and remove shopping cart items.

If you are new to Flux, the best way to get started is to imitate the example apps Facebook provided:

+ [Todo Flux App](https://facebook.github.io/flux/docs/todo-list.html#content)
+ [Chat Flux App](https://github.com/facebook/flux/tree/master/examples/flux-chat/js)

Looking at these two apps, you might find Flux to be somewhat verbose and over-engineered. To better understand why Flux is designed that way, we will pretend that Flux doesn't exist and implement our own Flux framework from scratch.

At first, our code would follow a more conventional coding style, similar to MVC. As we implement our shopping cart application, though, we'll discover new needs. By refactoring our code to satisfy these needs, we would eveolve our app to be closer and closer to the real Flux.

The Flux architecture has four parts:

![](flux-simple-f8-diagram-1300w.png)

In this lesson we'll focus on Store->View. In other words, when we change the application data, the views should update.
<cn>
# 开始使用 Flux

我们的购物车目前只是一个静态页面。在本课中，我们将使用 Flux 来让它能够添加和删除购物车内的商品。
如果你没有接触过 Flux，那么最好的开始方法是模仿 Facebook 提供的应用例子：

+ [Todo Flux App](https://facebook.github.io/flux/docs/todo-list.html#content)
+ [Chat Flux App](https://github.com/facebook/flux/tree/master/examples/flux-chat/js)

观察这两个应用，你也许会发现 Flux 有点啰嗦且过度设计了。为了更好的理解为什么 Flux 要这样设计，我们将假设 Flux 并不存在，从零开始实现我们自己的 Flux 框架。

一开始我们的代码将遵循一个更加常见的代码风格，类似于 MVC。然而，在我们实现购物车应用时，我们会发现新的需求。通过重构代码来满足需求的方式，我们将把应用进化到和真正的 Flux 越来越接近。

Flux 的结构由以下四个部分组成：

![](flux-simple-f8-diagram-1300w.png)

在本课中，我们将把注意力集中在 Store->View 上。换句话说，当我们修改应用的数据时，view 将会改变。
</cn>
# Lesson StarterKit

We have provided a reference implementation for the `buyshoes` page.

```
git clone https://github.com/hayeah/sikeio-buyshoes-startkit.git sikeio-buyshoes-react
cd sikeio-buyshoes-react
git checkout flux-start
```

You could use your own if you want.

<cn>
# 课程的启动包

我们提供了一个 `buyshoes` 页面的参考实现。

```
git clone https://github.com/hayeah/sikeio-buyshoes-startkit.git sikeio-buyshoes-react
cd sikeio-buyshoes-react
git checkout flux-start
```

你也可以使用自己的代码。
</cn>
# Flux Is Events Driven

The procedural programming model is what we are most used to. We tell the program to do one command after another:

```js
main() {
  doX()
  doY()
  doZ()
}
```

Another way to structure a program is with the event-driven architecture. The computer waits for something to happen, then it responds to the event. Suppose we are building a home automation system:

+ When X happens, do Y.
+ When the windows open, turn off the air conditioner.
+ When "howard" wakes up, make coffee.
+ When "howard" start working, turn off music.

This is called [inversion of control](https://en.wikipedia.org/wiki/Inversion_of_control), because you no longer tell the program when to do something. All that the program does is responding to what happens in the world.

User interface is another good fit for event-driven architecture, since all it does is responding to user interactions:

+ When user does X, application should do Y.
+ When user clicks "add to cart" button, add product to the shopping cart.
+ When user clicks on "trash can" button, remove product from the shopping cart.

MVC is procedural. The controller is the brain, telling different parts of the system to do as it commands.

Flux is an events-driven publish/subscribe system (pubsub). When a user interacts with the application, a message is broadcasted throughout the entire system. Any component of the system can decide whether it should respond or ignore an event. There is no centralized control.

The key advantage of publish/subscribe is that it's super extensible. Because the system is decentralized, you can easily add new features without modifying any existing code.
<cn>
# Flux 是事件驱动的

面向过程的编程模型是我们最习惯的。我们告诉程序依次执行命令：

```js
main() {
  doX()
  doY()
  doZ()
}
```

另一种构建程序的方式是通过事件驱动的架构。计算机等待某种事情的发生，然后相应这个事件。假设我们在构建一个家庭自动化系统：

+ 当 X 发生时，执行 Y；
+ 当窗户打开时，关闭空调；
+ 当 "howard" 醒来时，煮咖啡；
+ 当 "howard" 开始工作时，关闭音乐。

这叫做 [控制反转](https://zh.wikipedia.org/wiki/%E6%8E%A7%E5%88%B6%E5%8F%8D%E8%BD%AC)，因为你不需要告诉程序什么时候做什么事情。程序做的所有事情只是响应世界上发生的事情。

用户界面是另一个适合事件驱动架构的场景，因为它做的所有事情都是响应用户交互：

+ 当用户做 X 动作时，应用应该执行 Y；
+ 当用户点击 "add to cart" 按钮时，向购物车添加商品；
+ 当用户点击 "trash can" 按钮时，将商品从购物车中去掉。

MVC 是过程化的。控制器是大脑，告诉系统的其他部分去执行它的指令。

Flux 是一个事件驱动的 发布/订阅 系统（pubsub）。当用户和应用交互时，有一条消息将广播到整个系统中。任何系统中的组件都可以决定应该接受还是忽略该事件。没有中心控制。

发布/订阅系统的最大好处是它非常适合扩展。因为系统是去中心化的，你可以轻松的添加新的功能，而不需要修改任何已有的代码。
</cn>
# Keeping PubSub Simple With Flux

But because there's no clear logical flow, a pubsub system can be horribly confusing. In a complicated pubsub system, an event could trigger more events, like spreading a gossip (or contagious disease):

![](rumour-spreading.jpg)

It can be extremely difficult to trace the chain of events back to the original source. Unless we impose discipline in how we structure a pubsub system, it could easily become an impossible to understand mess.

Flux restricts its pubsub system by enforcing the four-steps "unidirectional flow" convention:

![](flux-simple-f8-diagram-1300w.png)

+ The only way to initiate an update cycle is by creating an action.
+ Stores can ONLY receive events from the dispatcher.
+ Views can ONLY receive events from stores.

With these restrictions in place, the Flux pub/sub system looks like:

![](unidirectional-flow.jpg)
<cn>
# 应用 Flux 保持简单的发布订阅

但是因为没有一个明确的逻辑流，一个发布订阅系统可能会很令人困惑。在一个复杂的发布订阅系统中，一个事件可以触发更多的事件，就像散布流言一样（或者是传染病）：

![](rumour-spreading.jpg)

沿着事件链条追溯到最初的源头是非常困难的。除非我们强加一些构建发布订阅系统的的规则，否则它将很容易变成一团不可理解的乱麻。

Flux 通过强制使用四步"单向流"的约定，来约束它的发布订阅系统：

![](flux-simple-f8-diagram-1300w.png)

+ 初始化一个更新循环的唯一方法是创建一个 action；
+ Store **只能** 从 dispatcher 处接受事件；
+ View **只能** 从 store 处接受事件。

在这样的约束下，Flux 的发布订阅系统看上去是这样的：

![](unidirectional-flow.jpg)

</cn>
### No Cascading Allowed

A store should never never NEVER cause other stores to update ("no cascading"). Let's see an example of how cascading might be useful, so we can appreciate what exactly it is we need to avoid.

Suppose I want to `like` a product, the MVC code might look like:

```js
let likedProducts = {
  products: [],

  addProduct(product) {
    this.products.push(product);

    // Update the associated view
    updateLikedProductsView();
  }
}

let product = {
  setLiked() {
    this.liked = true;
    likedProducts.addProduct(this);

    // Update the associated view
    updateProductView();
  }
}

product.setLiked();
```

1. Call `product.setLiked()`
2. `product` updates its associated view.
3. `product` adds itself to `likedProducts`.
4. `likedProducts` updates its associated view.

This seems like perfectly reasonable code, but you are not allowed to do this with Flux. When we build our shopping cart, please remember the following 3 important rules:

1. Store should never update other stores.
2. Store should never update other stores.
3. Store should never update other stores.
<cn>
### 不允许层叠

一个 store 永远永远永远不会更新其它的 store（“不要层叠”）。让我们来看一个可能用到层叠的例子，这样我们能领会我们要避免怎么做。

假设我希望可以 `喜欢` 一个商品，MVC 的代码看上去是这样子的：

```js
let likedProducts = {
  products: [],

  addProduct(product) {
    this.products.push(product);

    // 更新关联的 view
    updateLikedProductsView();
  }
}

let product = {
  setLiked() {
    this.liked = true;
    likedProducts.addProduct(this);

    // 更新关联的 view
    updateProductView();
  }
}

product.setLiked();
```

1. 调用 `product.setLiked()`
2. `product` 更新它关联的 view
3. `product` 将它自己添加到 `likedProducts`
4. `likeProducts` 更新它关联的 view

这看上去是极为合理的代码，但是 Flux 不允许你这么干。当我们构建购物车时，请记住以下3条重要的规则：

1. Store 永远不会更新其它的 store
2. Store 永远不会更新其它的 store
3. Store 永远不会更新其它的 store
</cn>
# The Store Pattern

A store is simply a set of functions to read and write data in the module. It's like a mini database.

Facebook's original Flux is more a pattern than a framework.  Rather than creating the `Store` class, we'll use plain JavaScript:

```js
// From NodeJS standard library
const EventEmitter = require("events");
function emitChange() {
  emitter.emit("change");
}

// The store's private data. Only accessible inside the module.
let _data = [];
// There can be other pieces of data you want the store to control.
let _foo = new Foo();

let _bar = new Bar();

let emitter = new EventEmitter();

module.exports = {
  // Reader API
  getData() {
    return _data;
  },

  // Writer API
  modifyData(data) {
    _data = data;

    // Every writer method should emit the "change" event.
    emitChange();
  },

  // Views can subscribe to "change" events.
  addChangeListener(callback) {
    emitter.addListener("change",callback);
  },

  removeChangeListener(callback) {
    emitter.removeListener("change",callback);
  },
}
```

+ There can be many reader methods.
+ There can be many writer methods.
+ All store should have the `addChangeListener` and `removeChangeListener` methods.
+ Never allow the outside world to modify the store's internal data directly.
+ The views would get data from a store by calling its reader methods.


Rather than having one store for the whole app, you usually create many stores to serve different needs.

This is the first version of the pattern, which we will refactor as we make progress.
<cn>
# Store 模式

一个 store 仅仅就是一组对模块内数据进行读写的函数。它像一个迷你的数据库。

Facebook 最初的 Flux 更像是一个模式而不是框架。我们用原生的 JavaScript 来代替创建 `Store` 类：

```js
// 从 NodeJS 标准库中得到
const EventEmitter = require("events");
function emitChange() {
  emitter.emit("change");
}

// store 的私有数据。只能在模块内可取得。
let _data = [];
// 可能有其他你希望 store 控制的数据
let _foo = new Foo();

let _bar = new Bar();

let emitter = new EventEmitter();

module.exports = {
  // 读方法 API
  getData() {
    return _data;
  },

  // 写方法 API
  modifyData(data) {
    _data = data;

    // 每个写方法都应该发出一个 "change" 事件
    emitChange();
  },

  // View 可以订阅 "change" 事件
  addChangeListener(callback) {
    emitter.addListener("change",callback);
  },

  removeChangeListener(callback) {
    emitter.removeListener("change",callback);
  },
}
```

+ 有很多读方法
+ 有很多写方法
+ 所有的 store 都应该有 `addChangeListener` 和 `removeChangeListener` 方法
+ 永远不允许外部世界直接修改 store 内部的数据
+ view 通过调用 store 的读方法来获得其数据

通常你会创建多个 store 来满足不同的需求，而不是在整个应用中只用一个 store。

这是 store 模式的第一个版本，我们会在之后的进程中重构它。
</cn>
### The NodeJS "events" Module

[EventEmitter](https://nodejs.org/api/events.html) is from NodeJS' standard library. Webpack would automatically find the `events` module from NodeJS, so you don't need to install the "events" package with npm.

The EventEmitter API is similar to how you might respond to a click event:

```js
function handleClick(event) {
  console.log("target",event.target);
  console.log("type",event.type);
}
button.addEventListener("click",handleClick);
```

We can use EventEmitter to create a fake button:

```js
const EventEmitter = require("events");

let fakeButton = new EventEmitter()

function handleClick(event) {
  console.log("target",event.target);
  console.log("type",event.type);
}

fakeButton.addListener("click",handleClick);

// Fake a click event
let event = {target: fakeButton, type: "click"};
fakeButton.emit("click");
```
<cn>
### NodeJS 的 "events" 模块

[EventEmitter](https://nodejs.org/api/events.html) 是 NodeJS 的一个标准库。Webpack 会从 NodeJS 中自动找到 `events` 模块，所以你不需要用 npm 安装 "events" 包。

EventEmitter 的 API 和你（译者注：用 DOM）响应一个点击事件是一样的：

```js
function handleClick(event) {
  console.log("target",event.target);
  console.log("type",event.type);
}
button.addEventListener("click",handleClick);
```

我们可以使用 EventEmitter 来创建一个假按钮：

```js
const EventEmitter = require("events");

let fakeButton = new EventEmitter()

function handleClick(event) {
  console.log("target",event.target);
  console.log("type",event.type);
}

fakeButton.addListener("click",handleClick);

// 虚构一个 click 事件
let event = {target: fakeButton, type: "click"};
fakeButton.emit("click");
```
</cn>
# Implement Search Suggestions

Now let's see how Flux works in practice. We'll build a search input box with auto suggestions:

<video src="search-suggestions.mp4" controls></video>

We'll build this feature twice.

1. First time, using a centralized controller (MVC).
2. Second time, using EventEmitter (Flux).

The actual functionality is the same. These two examples only differ by how the pieces of code are glued together.
<cn>
# 实现搜索建议

现在让我们来看下 Flux 在实践中如何工作。我们将构建一个带有自动建议的搜索输入框：

<video src="search-suggestions.mp4" controls></video>

我们将实现这个功能两次。

1. 第一次，用一个集中的 controller 实现（MVC）
2. 第二次，用 EventEmitter（Flux）

（两个版本）的实际功能是一样的。这两个例子的区别只是在于各部分代码之间是如何连接的。
</cn>
# Search Suggestions With Controller

[Search Suggestions MVC - Codepen Demo](http://codepen.io/hayeah/pen/pjdpPP?editors=001)

The MVC version of the "search suggestions" uses a centralized controller to glue the parts together:

![](mvc-search-suggestions.jpg)

First, `SearchInputView` calls `updateSearchInput` whenever its input changes:

```js
class SearchInputView extends React.Component {
  onChange(e) {
    let value = e.target.value;
    suggestionController.updateSearchInput(value);
  }

  render() {
    return <input onChange={this.onChange.bind(this)} placeholder="enter country name"/>;
  }
};
```

The controller uses the RemoteAPI to get the matching suggestions:

```js
let suggestionController = {
  updateSearchInput(queryString) {
    RemoteAPI.fetchSuggestions(queryString,(suggestions) => {
      this.receivedSuggestions(suggestions);
    });
  },
}
```

When the RemoteAPI returns with the result, the controller updates the store and the view:

```js
let suggestionController = {
  updateSuggestionsDisplayView() {
    // Forces the React component to re-render.
    this.suggestionsDisplayView.forceUpdate();
  },

  receivedSuggestions(suggestions) {
    suggestionsStore.setSuggestions(suggestions);
    this.updateSuggestionsDisplayView();
  },
}
```
<cn>
# 用 Controller 实现搜索建议

[搜索建议 MVC - Codepen Demo](http://codepen.io/hayeah/pen/pjdpPP?editors=001)

MVC 版本的 "搜索建议" 使用了一个集中的 controller 来将各个部分粘结在一起：

![](mvc-search-suggestions.jpg)

首先，`SearchInputView` 在输入框变化时调用 `updateSearchInput` 方法：

```js
class SearchInputView extends React.Component {
  onChange(e) {
    let value = e.target.value;
    suggestionController.updateSearchInput(value);
  }

  render() {
    return <input onChange={this.onChange.bind(this)} placeholder="enter country name"/>;
  }
};
```

controller 使用 RemoteAPI 来得到对应的搜索建议：

```js
let suggestionController = {
  updateSearchInput(queryString) {
    RemoteAPI.fetchSuggestions(queryString,(suggestions) => {
      this.receivedSuggestions(suggestions);
    });
  },
}
```

当 RemoteAPI 返回结果时，controller 更新 store 和 view：

```js
let suggestionController = {
  updateSuggestionsDisplayView() {
    // 强制 React 组件重新渲染
    this.suggestionsDisplayView.forceUpdate();
  },

  receivedSuggestions(suggestions) {
    suggestionsStore.setSuggestions(suggestions);
    this.updateSuggestionsDisplayView();
  },
}
```
</cn>
The `suggestionStore` is just a getter/setter API for an array of strings:

```js
// Normally the store is in a module. Here we use closure.
let suggestionsStore = (() => {

  let _suggestions = [];

  // We could replace `return` with module
  // module.exports = { ... }

  return {
    // Reader API
    getSuggestions() {
      return _suggestions;
    },

    // Writer API
    setSuggestions(suggestions) {
      _suggestions = suggestions;
    }
  };
})();
```

The `SuggestionsDisplayView` registers itself with the controller when it's mounted:

```js
class SuggestionsDisplayView extends React.Component {
  componentDidMount() {
    suggestionController.setSuggestionsDisplayView(this);
  }
};
```

When the controller tells the `SuggestionsDisplayView` to update, it reads the latest data from `suggestionsStore`:

```js
class SuggestionsDisplayView extends React.Component {
  render() {
    let suggestions = suggestionsStore.getSuggestions();
    ...
  }
};
```
<cn>
`suggestionStore` 只是一个供字符串数组使用的读写 API：

```js
// 通常 store 是在模块中的。在这里我们使用闭包
let suggestionsStore = (() => {

  let _suggestions = [];

  // 我们可以用模块代替 `return`
  // module.exports = { ... }

  return {
    // 读方法 API
    getSuggestions() {
      return _suggestions;
    },

    // 写方法 API
    setSuggestions(suggestions) {
      _suggestions = suggestions;
    }
  };
})();
```

`SuggestionsDisplayView` 在加载后将自身注册到 controller 中：

```js
class SuggestionsDisplayView extends React.Component {
  componentDidMount() {
    suggestionController.setSuggestionsDisplayView(this);
  }
};
```

当 controller 通知 `SuggestionsDisplayView` 更新时，它从 `suggestionsStore` 中读取最新的数据：

```js
class SuggestionsDisplayView extends React.Component {
  render() {
    let suggestions = suggestionsStore.getSuggestions();
    ...
  }
};
```
</cn>
# Search Suggestions With EventEmitter

[Search Suggestions - EventEmitter Codepen Demo](http://codepen.io/hayeah/pen/NGwXXb?editors=001)

Now we remove the controller, and glue code together with pubsub instead:

![](events-search-suggestions.jpg)

Note: In "real" Flux, actions would use pubsub to cause stores to update. For now we'll keep it simple by allowing actions to call stores directly.

First the input box triggers the action `updateSearchQuery`:

```js
class SearchInputView extends React.Component {
  onChange(e) {
    let value = e.target.value;
    updateSearchQuery(value);
  }

  render() {
    return (
      <p>
        Country Name: <br/>
        <input onChange={this.onChange.bind(this)} placeholder="enter country name"/>
      </p>
    );
  }
};
```

The action `updateSearchQuery` fetches the data from server:

```js
// Action
function updateSearchQuery(query) {
  RemoteAPI.fetchSuggestions(query,(suggestions) => {
    receiveSuggestions(suggestions);
  });
}
```

When the suggestions return from the server, the action `receiveSuggestions` is triggered. This action would cause the store to update:

```js
function receiveSuggestions(suggestions) {
  suggestionsStore.setSuggestions(suggestions);
}
```
<cn>
# 应用 EventEmitter 实现搜索建议

[搜索建议 - EventEmitter Codepen Demo](http://codepen.io/hayeah/pen/NGwXXb?editors=001)

现在我们将 controller 去掉，然后用发布订阅模式来将代码连结起来：

![](events-search-suggestions.jpg)

注意：在 “真正的” Flux 中，action 会利用发布订阅使 store 更新。目前我们先简单地让 action 直接调用 store。

首先输入框触发 action `updateSearchQuery`：

```js
class SearchInputView extends React.Component {
  onChange(e) {
    let value = e.target.value;
    updateSearchQuery(value);
  }

  render() {
    return (
      <p>
        Country Name: <br/>
        <input onChange={this.onChange.bind(this)} placeholder="enter country name"/>
      </p>
    );
  }
};
```

`updateSearchQuery` 这个 action 从服务器中读取数据：

```js
// Action
function updateSearchQuery(query) {
  RemoteAPI.fetchSuggestions(query,(suggestions) => {
    receiveSuggestions(suggestions);
  });
}
```

当服务器返回建议（数据）时，`receiveSuggestions` 这个 action 被触发。这个 action 将让 store 更新：

```js
function receiveSuggestions(suggestions) {
  suggestionsStore.setSuggestions(suggestions);
}
```
</cn>
The store is mostly the same as before. The only difference is that the setter method `setSuggestions` now emits the `change` event:

```js
let suggestionsStore = (() => {
  let _suggestions = [];

  let emitter = new EventEmitter();

  return {
    getSuggestions() {
      return _suggestions;
    },

    setSuggestions(suggestions) {
      _suggestions = suggestions;
      emitter.emit("change");
    },

    addChangeListener(callback) {
      emitter.addListener("change",callback);
    },
  };
})();
```

Finally, the `SuggestionsDisplayView` listens to the `suggestionsStore`.

```js
class SuggestionsDisplayView extends React.Component {
  componentDidMount() {
    suggestionsStore.addChangeListener(this.forceUpdate.bind(this));
  }
};
```

The `SuggestionsDisplayView` reads the latest data from the `suggestionStore` whenever it has to renderer:

```js
class SuggestionsDisplayView extends React.Component {
  render() {
    let suggestions = suggestionsStore.getSuggestions();
    ...
  }
};
```
<cn>
store 的代码和之前几乎是一样的。唯一的区别在于写方法 `setSuggestions` 现在发出 `change` 事件：

```js
let suggestionsStore = (() => {
  let _suggestions = [];

  let emitter = new EventEmitter();

  return {
    getSuggestions() {
      return _suggestions;
    },

    setSuggestions(suggestions) {
      _suggestions = suggestions;
      emitter.emit("change");
    },

    addChangeListener(callback) {
      emitter.addListener("change",callback);
    },
  };
})();
```

最后，`SuggestionDisplayView` 监听 `suggestionStore` 的变化。

```js
class SuggestionsDisplayView extends React.Component {
  componentDidMount() {
    suggestionsStore.addChangeListener(this.forceUpdate.bind(this));
  }
};
```

当需要渲染时，`SuggestionsDisplayView` 从 `suggestionStore` 中读取最新的数据：

```js
class SuggestionsDisplayView extends React.Component {
  render() {
    let suggestions = suggestionsStore.getSuggestions();
    ...
  }
};
```
</cn>
# Extending The Search Suggestions

Let's practice Flux by adding a few additional features. You can do these exercises directly in Codepen by forking the original code:

![](codepen-fork.jpg)

The codepen editor embeds the page in an iframe, so debugging is difficult. There's a debug mode that opens your code in a new window without the iframe:

![](codepen-enter-debug-mode.jpg)

You can also take a look at the JavaScript settings. The forked codepen project should inherit the same JavaScript settings:

+ Uses Babel.
+ Included React 0.14.
+ Included EventEmitter.

![](codepen-javascript-settings.jpg)
<cn>
# 扩展搜索建议功能

让我们通过添加一些额外的功能来练习 Flux。你可以直接通过 fork 以下的源代码来在 Codepen 上直接完成练习：

![](codepen-fork.jpg)

codepen 编辑器在一个 iframe 中嵌套了页面，所以 debug 会有点困难。幸好有一个 debug 模式可以在一个新窗口里打开你的代码，不需要 iframe：

![](codepen-enter-debug-mode.jpg)

你还可以看一下 JavaScript 的设置。fork 的 codepen 项目应该继承了相同的 JavaScript 设置：

+ 使用 Babel；
+ 包含 React 0.14；
+ 包含 EventEmitter。

![](codepen-javascript-settings.jpg)
</cn>
### Exercise: Show the length of the query string.

Fork the demo: [Search Suggestions - EventEmitter Codepen Demo](http://codepen.io/hayeah/pen/NGwXXb?editors=001)

Add a the `QueryLengthView` component to the length of the search string:

```js
class QueryLengthView extends React.Component {
  componentDidMount() {
    queryStore.addChangeListener(this.forceUpdate.bind(this));
  }

  render() {
    let query = queryStore.getQuery();

    return (
      <div>Query length: {query.length}</div>
    );
  }
}

let App = () => {
  return (
    <div>
      <SearchInputView/>
      <QueryLengthView/>
      <SuggestionsDisplayView/>
    </div>
  )
}
```

You'll need to:

+ Create the `queryStore`.
+ Modify `updateSearchQuery` so it causes `queryStore` to update.

Your result:

<video src="search-query-length-display.mp4" controls></video>
<cn>
### 练习：显示查询字符串的长度

Fork 这个 demo: [Search Suggestions - EventEmitter Codepen Demo](http://codepen.io/hayeah/pen/NGwXXb?editors=001)

添加一个 `QueryLengthView` 组件对应显示搜索字符串的长度：

```js
class QueryLengthView extends React.Component {
  componentDidMount() {
    queryStore.addChangeListener(this.forceUpdate.bind(this));
  }

  render() {
    let query = queryStore.getQuery();

    return (
      <div>Query length: {query.length}</div>
    );
  }
}

let App = () => {
  return (
    <div>
      <SearchInputView/>
      <QueryLengthView/>
      <SuggestionsDisplayView/>
    </div>
  )
}
```

你需要：

+ 创建 `queryStore`。
+ 修改 `updateSearchQuery` 这样它会让 `queryStore` 更新。

效果如下：

<video src="search-query-length-display.mp4" controls></video>
</cn>
### Exercise: Sum population and area of matching countries

The countries data has the population and area of each country:

```js
 window.countriesData = [
  {
    "countryCode": "AD",
    "countryName": "Andorra",
    "population": "84000",
    "areaInSqKm": "468.0"
  },
  {
    "countryCode": "AE",
    "countryName": "United Arab Emirates",
    "population": "4975593",
    "areaInSqKm": "82880.0"
  },
  ...
]
```

Create a component to display the sums of all the countries that match the query string.

```js
class MatchingCountriesSummaryView extends React.Component {
  componentDidMount() {
    ...
  }

  render() {
    ...

    return (
      <div>
        <h3>Matching countries ({countries.length})</h3>
        <p>total area: {areaSum} </p>
        <p>total population: {populationSum} </p>
      </div>
    );
  }
}
```

Your result:

<video src="search-query-sum-area-population.mp4" controls></video>

Question: Which of a,e,i,o,u has the greatest population sum?
<cn>
### 练习：将匹配到的国家的人口和面积加起来

国家的数据包括各个国家的人口和区域：

```js
 window.countriesData = [
  {
    "countryCode": "AD",
    "countryName": "Andorra",
    "population": "84000",
    "areaInSqKm": "468.0"
  },
  {
    "countryCode": "AE",
    "countryName": "United Arab Emirates",
    "population": "4975593",
    "areaInSqKm": "82880.0"
  },
  ...
]
```

创建一个组件来显示所有匹配搜索字符串的国家的（人口和区域面积）的总和。

```js
class MatchingCountriesSummaryView extends React.Component {
  componentDidMount() {
    ...
  }

  render() {
    ...

    return (
      <div>
        <h3>Matching countries ({countries.length})</h3>
        <p>total area: {areaSum} </p>
        <p>total population: {populationSum} </p>
      </div>
    );
  }
}
```

效果如下：

<video src="search-query-sum-area-population.mp4" controls></video>

问题：a,e,i,o,u 这五个查询字符串，那个得到的总和最大？
</cn>
### Exercise: Be able to select a suggestion

When you click on a suggestion, the input box's value should be replaced. We want `SearchInputView` update whenever the `queryStore` changes.

There are now two separate paths that `SearchInputView` could be updated:

1. SearchInputView could modify itself by triggering `updateSearchQuery`.
2. SuggestionDisplayView triggering `updateSearchQuery`.

![](search-update-search-query.jpg)


The first case where SearchInputView could trigger its own update makes it a [controlled component](https://facebook.github.io/react/docs/forms.html#controlled-components). It looks like:

```js
class SearchInputView extends React.Component {

  // ...

  onChange(e) {
    let value = e.target.value;
    updateSearchQuery(value);
  }

  render() {
    let query = queryStore.getQuery();

    return (
      <p>
        Country Name: <br/>
        <input value={query} onChange={this.onChange.bind(this)} placeholder="enter country name"/>
      </p>
    );
  }
};
```

The circular update of the controlled component is a little strange. It goes like this:

1. The input element updates the store (or `this.state`) with a new value.
2. The store notifies the input component that it changed.
3. The new value of the store (or `this.state`) is now circled back to the input component, and sets the value.

This maintains the unidirectional flow.

Your result:

<video src="search-suggestion-select.mp4" controls></video>

Question: How can you limit the user's query input to a maximum length of 10 characters? Can you add a single line in `SearchInputView` to impose this limit?
<cn>
### 练习：选择一个搜索建议

当你点击一个建议时，输入框的值应该被（这个值）代替。我们希望 `SearchInputView` 在每次 `queryStore` 变化时进行更新。

有两个办法可以让 `SearchInputView` 更新：

1. SearchInputView 触发 `updateSearchQuery` 来更新自己。
2. SuggestionDisplayView 触发 `updateSearchQuery`。

![](search-update-search-query.jpg)

第一种情况中，SearchInputView 可以触发自己的 update，这让它成为一个 [controlled component](https://facebook.github.io/react/docs/forms.html#controlled-components)。它看上去是这样的：

```js
class SearchInputView extends React.Component {

  // ...

  onChange(e) {
    let value = e.target.value;
    updateSearchQuery(value);
  }

  render() {
    let query = queryStore.getQuery();

    return (
      <p>
        Country Name: <br/>
        <input value={query} onChange={this.onChange.bind(this)} placeholder="enter country name"/>
      </p>
    );
  }
};
```

controlled component 的循环更新有一点奇怪。它的流程是这样的：

1. input 元素用一个新的值更新 store（或者 `this.state`）；
2. store 告诉 input 组件它被更新了；
3. store 的新值（或者 `this.state`）现在又循环回到 input 组件，并设置它的值。

这样就维持了（数据的）单向流动。

效果如下：

<video src="search-suggestion-select.mp4" controls></video>

问题：如何才能限制用户的请求字符串不超过10个字符？你能否在 `SearchInputView` 中添加一行来加入这个限制？
</cn>
# Implement Shopping Cart With Flux

We'll start using Flux to implement the shopping cart functionality.

For now, to keep it simple, we'll allow views to call the store's writer methods directly. In other words, the "actions" are the same as the store's writer methods.

In the future we'll refactor the app so actions and stores are glued together with pub/sub instead of direct function calls.

<cn>
# 用 Flux 实现购物车

我们将用 Flux 来实现购物车的功能了。

目前，为简洁起见，我们允许 view 直接调用 store 的写方法。换句话说，"actions" 和 store 的写方法是一样的。

在将来的课程中，我们将重构应用，这样 action 和 store 会用发布订阅模式来连结，以取代直接的函数调用。
</cn>
### Exercise: Be able to add products to shopping cart

We'll put all the stores the `js/stores` directory. Create the file `js/stores/CartStore.js`:

```js
const EventEmitter = require("events");

let emitter = new EventEmitter();

function emitChange() {
  emitter.emit("change");
}

let _cartItems = {
  // "jameson-vulc": {
  //   id: "jameson-vulc",
  //   quantity: 1,
  // },
};

module.exports = {
  // Reader methods
  ...

  // Writer methods. These are the "actions".
  ...

  addChangeListener(callback) {
    emitter.addListener("change",callback)
  },

  removeChangeListener(callback) {
    emitter.removeListener("change",callback)
  },
}
```

All the stores you create would look similar to the above.


The Flux components are:

+ Action: `addCartItem(productId)`.
+ Store: `CartStore`.
+ Views: `Products` and `Cart`.

They should be glued together like this:

+ Clicking the "add to cart" button should trigger the `addCartItem` action.
+ `addCartItem` should cause the `CartStore` to update.
+ The `CartStore` should notify `Products` and `Cart` of changes by calling `emitChange`.
+ The `Products` and `Cart` components should read cart items from `CartStore`.

Instead of calling `CartStore.addCartItem`, we should call an action by its name:

```js
const CartStore = require("./CartStore");
const {addCartItem} = CartStore;

addCartItem(productId);
```

When we refactor the actions to an independent module, we could change the require, and keep everything else the same:

```js
const CartStore = require("./CartStore");
const {addCartItem} = require("./actions");

addCartItem(productId);

```

Your result:

<video src="addCartItem.mp4" controls></video>
<cn>
### 练习：向购物车添加商品

我们把所有的 store 都放置在 `js/stores` 目录下。创建文件 `js/stores/CartStore.js`：

```js
const EventEmitter = require("events");

let emitter = new EventEmitter();

function emitChange() {
  emitter.emit("change");
}

let _cartItems = {
  // "jameson-vulc": {
  //   id: "jameson-vulc",
  //   quantity: 1,
  // },
};

module.exports = {
  // 读方法
  ...

  // 写方法。这些就是 "action"
  ...

  addChangeListener(callback) {
    emitter.addListener("change",callback)
  },

  removeChangeListener(callback) {
    emitter.removeListener("change",callback)
  },
}
```

你创建的所有 store 都会和上面的看上去差不多。

Flux 组件是：

+ Action: `addCartItem(productId)`.
+ Store: `CartStore`.
+ View: `Products` and `Cart`.

它们连结在一起应该是：

+ 点击 "add the cart" 按钮会触发 `addCartItem` action；
+ `addCartItem` 会让 `CartStore` 更新；
+ `CartStore` 会通过调用 `emitChange` 方法通知 `Products` 和 `Cart` 数据的变更；
+ `Products` 和 `Cart` 组件会从 `CartStore` 中读取购物车条目。

我们用 action 的名字来调用它，而不是调用 `CartStore.addCartItem`：

```js
const CartStore = require("./CartStore");
const {addCartItem} = CartStore;

addCartItem(productId);
```

当我们将 action 重构到一个独立模块时，我们可以只修改 require，而其他部分保持不变：

```js
const CartStore = require("./CartStore");
const {addCartItem} = require("./actions");

addCartItem(productId);

```

效果如下：

<video src="addCartItem.mp4" controls></video>
</cn>
### Exercise: Be able to remove products to shopping cart

Add the `removeCartItem(productId)` action.

Use the delete operator to remove a key from object.

```js
delete object[key]
```

Your result:

<video src="removeCartItem.mp4" controls></video>
<cn>
### 练习：可以从购物车中删除商品

添加 `removeCartItem(productId)` 这个 action。

使用 delete 操作符来从一个对象中删除一个键。

```js
delete object[key]
```

效果如下：

<video src="removeCartItem.mp4" controls></video>
</cn>
### Exercise: Be able to adjust quantities

Add the `updateCartItemQuantity(productId,quantity)` action. It works like this:

```js
// increase by 1
updateCartItemQuantity(productId,quantity+1);

// decrease by 1
updateCartItemQuantity(productId,quantity-1);
```

+ In the `CartStore` make sure that quantity can never be less than 1.
+ `QuantityControl` does not need to listen to the store, since its parent is already listening.

Your result:

<video src="updateCartItemQuantity.mp4" controls></video>
<cn>
### 练习：可以调整数量

添加 `updateCartItemQuantity(productId,quantity)` 这个 action。它的用法如下：

```js
// 增加1
updateCartItemQuantity(productId,quantity+1);

// 减少1
updateCartItemQuantity(productId,quantity-1);
```

+ 在 `CartStore` 中，要保证数量不能少于1
+ `QuantityControl` 不需要监听 store，因为它的 parent 已经在监听了。

效果如下：

<video src="updateCartItemQuantity.mp4" controls></video>
</cn>
### Exercise: Update Checkout's subtotal calculation

Finally, Checkout component's subtotal should also update.

You should ensure that only 2 decimal digits are displayed (rather than something liek 59.999967):

```js
subtotal.toFixed(2);
```

Your result:

<video src="update-checkout-subtotal.mp4" controls></video>
<cn>
### 练习：更新 结算 中 小计部分 的计算值

最后，结算组件的小计金额也需要更新。

你可以保证只显示两位小数（而不是显示像 59.999967 的数）：

```js
subtotal.toFixed(2);
```

结果如下：

<video src="update-checkout-subtotal.mp4" controls></video>
</cn>
# Summary

We've used Flux to connect the views with changing data. It's worth mentioning again that the views are exactly the same as before, yet they can automatically update the UI with the latest data.

Isn't it great that we didn't have to add any code to update the DOM manually?

+ Flux is events-driven. It removes the centralized controller.
+ We've used EventEmitter to glue together stores and views.
+ The unidirectional flow from actions to stores to views.
+ A store should never update other stores.

Our Flux code is very simple, making it easy to see what exactly is going on. We'll make it a bit more elegant in the next lesson.
<cn>
# 总结

我们用 Flux 来将变化的数据和 view 连接起来。值得再次一提的是，view 和之前是一模一样的，但是现在它们可以根据最新的数据来更新 UI。

我们不需要添加任何的代码来手动更新 DOM 了，是不是很棒？

+ Flux 是事件驱动的。它去掉了集中的 controller；
+ 我们用 EventEmitter 来将 store 和 view 连结起来；
+ 从 action 到 store 到 view 的单向流；
+ 一个 store 永远不会更新其他的 store

我们的 Flux 代码非常简单，很容易能看到到底是怎么回事。我们会在下一课中让代码变得优雅一点。
</cn>
