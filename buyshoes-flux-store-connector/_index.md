# Connecting Stores And Views

<cn>
# 连接 Store 和 View
</cn>

We've turned the shopping cart into a Flux app. But connecting views to stores is clumsy and verbose:

```js
class FooView extends React.Component {
  componentDidMount() {
    FooStore.addChangeListener(this.forceUpdate.bind(this));
  }

  render() {
    let fooData = FooStore.getFooData();
  }
}
```

<cn>
我们已经将购物车转化成了一个 Flux 应用，但是连接 Store 和 View 是一件笨拙而又繁琐的事情：

```js
class FooView extends React.Component {
  componentDidMount() {
    FooStore.addChangeListener(this.forceUpdate.bind(this));
  }

  render() {
    let fooData = FooStore.getFooData();
  }
}
```
</cn>

In this lesson we'll create a [JavaScript decorator](https://github.com/wycats/javascript-decorators) to connect views and stores:

```js
@connect(FooStore,"fooData");
class FooView extends React.Component {
  render() {
    let {fooData} = this.props;
    return <div>Foo data is: {fooData}</div>;
  }
}
```

<cn>
这节课我们来做一个 [JavaScript Decorator](https://github.com/wycats/javascript-decorators) 来连接 View 和 Store：

```js
@connect(FooStore,"fooData");
class FooView extends React.Component {
  render() {
    let {fooData} = this.props;
    return <div>Foo data is: {fooData}</div>;
  }
}
```
</cn>

All the code we'll have to write is only about 50~60 lines, but these work at a higher level of abstractions.

+ We'll create a wrapper component that accepts a function as its body.
+ We'll define a function that returns a component definition.
+ Finally, we'll define a function that returns a function that takes a component definition and returns another component definition. Confused? Me too.

<cn>
我们将要写的代码可能只有 50 到 60 行，但是这些代码是高度抽象的：

+ 我们将会创建一个 wrapper 组件，使其能将任意一个函数作为它的内容。
+ 我们将定义一个函数用于返回一个组件的定义。
+ 最后，我们还会定义一个函数，用于返回一个返回值是另一个组件的定义、并自身带有一个组件的定义的函数。有点绕，不过不要紧！
</cn>

In other words, we'll have some fun with React metaprogramming!

<cn>
换句话说，我们将从 React 元编程中获得很多乐趣！
</cn>

# Separation of Concerns

<cn>
# 化解疑虑
</cn>

We are not doing this refactoring just to write less code.

The real reason is so we can go back to writing React as though the app is completely static. We should create our components as though there isn't any Flux stores to connect to.

<cn>
我们重构只是为了写更少的代码。

真正的原因是我们可以将整个应用看作是静态的，然后再来写 React 代码。我们可以创建单独的组件就像是 Flux Store 并还没有对接上来一样。
</cn>

A component should be "dumb". It takes props as input, and renders the UI as output:

```js
class FooView extends React.Component {
  render() {
    let {fooData} = this.props;
    return <div>Foo data is: {fooData}</div>;
  }
}
```

<cn>
组件应该是「笨拙」的。它只会将参数作为输入，然后将渲染出的 UI 作为输出：

```js
class FooView extends React.Component {
  render() {
    let {fooData} = this.props;
    return <div>Foo data is: {fooData}</div>;
  }
}
```

</cn>

A dumb component like this doesn't depend on anything except its props. It is more reusable than a component that's connected to the store, for the simple reason that anybody can pass props to it.

+ The props could come from a store.
+ The props could come from a unit test.
+ The props could come from a parent component.

<cn>
这么「笨拙」的组件看上去只依赖于输入的参数。和 Store 已经连接的组件相比，它有更好的复用性，因为可以传递任何的参数：

+ 参数可以从 Store 中来。
+ 参数可以从单元测试中来。
+ 参数可以从其继承的父元素组件中来。

</cn>

Your app would be much easier to understand if all of your components are "dumb" like this.

<cn>
如果在你的应用中所有的组件都是如此「笨拙」，那么你的应用也会更容易被理解。
</cn>

To connect stores to views we could create wrapper components:

```js
class FooViewStoreWrapper extends React.Component {
  componentDidMount() {
    FooStore.addChangeListener(this.forceUpdate.bind(this));
  }

  render() {
    let fooData = FooStore.getFooData();
    return <FooView fooData={fooData}/>
  }
}
```

<cn>
连接 Store 和 View，我们需要创建一个容器元素：

```js
class FooViewStoreWrapper extends React.Component {
  componentDidMount() {
    FooStore.addChangeListener(this.forceUpdate.bind(this));
  }

  render() {
    let fooData = FooStore.getFooData();
    return <FooView fooData={fooData}/>
  }
}
```

</cn>

This convention separates two concerns into two components:

1. `FooView` renders the actual view. It doesn't care where its props are from.
2. `FooViewStoreWrapper` passes the latest store data to `FooView`. It doesn't know what the view would render.

<cn>
这种约定将我们的两个疑虑划分到了所对应的两个组件中：

1. `FooView` 渲染实际的师徒。它并不关心参数从哪里来。
2. `FooViewStoreWrapper` 将 Store 的数据传递给 `FooView`。它并不知道实际会渲染出什么 View。
</cn>

This separation of responsibilities is our goal. Of course, writing the store wrappers manually is incredibly tedious, and that's why we'd want to create the `@connect` decorator:

```js
@connect(FooStore,"fooData");
class FooView extends React.Component {
  render() {
    let {fooData} = this.props;
    return <div>Foo data is: {fooData}</div>;
  }
}
```

<cn>
这种职责分离就是我们设计的奥义。当然，手动写这种 Store 的容器是一件乏味的活，所以我们想要创造 `@connect` 装饰器（Decorator）：

```js
@connect(FooStore,"fooData");
class FooView extends React.Component {
  render() {
    let {fooData} = this.props;
    return <div>Foo data is: {fooData}</div>;
  }
}
```
</cn>


# The ConnectedStore Component

<cn>
# ConnectedStore 组件
</cn>

Our first step is to create a reusable wrapper component that could read data from a store, and pass the data down to a view. Let's call this wrapper component `ConnectedStore`.

<cn>
第一步就是要创建一个可以复用的 wrapper 组件，能够从一个 Store 中获取数据，还能将数据传递给一个 View。我们暂且叫它 `ConnectedStore`。
</cn>

The `ConnectedStore` component can be used to refactor a view that reads data from a store. The `TimerView` reads `currentTime` and the `currentTick` from the `TimerStore`:


```js
class TimerView extends React.Component {
  componentDidMount() {
    TimerStore.addChangeListener(this.forceUpdate.bind(this));
  }

  render() {
    let currentTime = TimerStore.currentTime();
    let currentTick = TimerStore.currentTick();

    return (
      <div>
        <p>Time: {currentTime.toString()}</p>
        <p>Tick: {currentTick}</p>
      </div>
    );
  }
}
```

<cn>

`ConnectedStore` 用于重构一个需要从 Store 中读取数据的 View。例如 `TimerView` 从 `TimerStore` 中读取了 `currentTime` 和 `currentTick` 属性。 

```js
class TimerView extends React.Component {
  componentDidMount() {
    TimerStore.addChangeListener(this.forceUpdate.bind(this));
  }

  render() {
    let currentTime = TimerStore.currentTime();
    let currentTick = TimerStore.currentTick();

    return (
      <div>
        <p>Time: {currentTime.toString()}</p>
        <p>Tick: {currentTick}</p>
      </div>
    );
  }
}
```

</cn>

It updates the time and tick counter every second:

[TimerView Demo](http://codepen.io/hayeah/pen/OyQQJV?editors=001)

<video src="TimerView.mp4" controls autoplay></video>

<cn>
这个视图将每秒更新时间和计数器：

[TimerView Demo](http://codepen.io/hayeah/pen/OyQQJV?editors=001)

<video src="TimerView.mp4" controls autoplay></video>
</cn>

We don't want TimerView to read data from TimerStore anymore, so let's make it read data from `this.props`:

```js
class TimerView extends React.Component {
  render() {
    let {currentTime,currentTick} = this.props;
    return (
      <div>
        <p>Time: {currentTime.toString()}</p>
        <p>Tick: {currentTick}</p>
      </div>
    );
  }
}
```

<cn>
我们并不想让 TimerView 从 TimerStore 读取数据了，将它改为从 `this.props` 中读取数据：

```js
class TimerView extends React.Component {
  render() {
    let {currentTime,currentTick} = this.props;
    return (
      <div>
        <p>Time: {currentTime.toString()}</p>
        <p>Tick: {currentTick}</p>
      </div>
    );
  }
}
```

</cn>

Now `TimerView` is a "dumb components". It doesn't know anything about the TimeStore.

<cn>
现在的 `TimerView` 就是一个「笨拙的组件」。它对 TimeStore 的内容一无所知。
</cn>

How do we use the `ConnectedStore` to pass data to TimerView? By convention, we'll make the store's reader names and the view's property names to be the same:

+ `TimeStore.currentTime()` -> `this.props.currentTime`
+ `TimeStore.currentTick()` -> `this.props.currentTick`

<cn>
那么我们如何利用 `ConnectedStore` 来传递数据给 TimerView 呢？通过约定，我们会让 Store 的方法名和 View 的属性名保持一致：

+ `TimeStore.currentTime()` -> `this.props.currentTime`
+ `TimeStore.currentTick()` -> `this.props.currentTick`

</cn>

Then we could use the wrapper component like this:

```js
<ConnectedStore store={TimerStore} propNames={["currentTime","currentTick"]}>
  {props => <TimerView {...props}/>}
</ConnectedStore>
```

<cn>
然后我们就能这么使用 wrapper 组件了：

```js
<ConnectedStore store={TimerStore} propNames={["currentTime","currentTick"]}>
  {props => <TimerView {...props}/>}
</ConnectedStore>
```

</cn>

+ `store` - Should connect to TimerStore by calling `addChangeListener`.
+ `propNames` - Should read these data from the store.
+ `props => ...` - Everytime the store emits "change", read data from the store, and call this function to render the view.

<cn>

+ `store` - 通过调用 `addChangeListener` 连接 TimerStore。
+ `propNames` - 从 Store 中读取这些数据。
+ `props => ...` - 当 Store 改变时，读取数据并调用这个函数来重新渲染 View。

</cn>

If the syntax looks weird to you, see the plain JavaScript output:

```js
React.createElement(
  ConnectedStore,
  { store: FooStore, propNames: ["dataA", "dataB"] },
  function (props) {
    return React.createElement(Foo, props);
  }
);
```

<cn>
如果上面的格式看上去很古怪，看一下 JavaScript：

```js
React.createElement(
  ConnectedStore,
  { store: FooStore, propNames: ["dataA", "dataB"] },
  function (props) {
    return React.createElement(Foo, props);
  }
);
```

</cn>

So the body is a function, which can be accessed as `this.props.children` in the wrapper component:

```js
// This component should re-render whenever its store emits the "change" event.
class ConnectedStore extends React.Component {

  ...

  render() {
    // The `children` property is a function.
    let contentRenderFunction = this.props.children;

    // 1. Read all the data from store by calling reader methods dynamically.
    // 2. Pass the data to `contentRenderFunction`.

    ...

    return contentRenderFunctions(storeProps);
  }
}
```

<cn>
所以内容就是一个函数，通过  wrapper 组件中的 `this.props.children` 来连接：


```js
// 当 Store 更改时，组件应该重新渲染。
class ConnectedStore extends React.Component {

  ...

  render() {
    // `children` 属性是一个函数。
    let contentRenderFunction = this.props.children;

    // 1. 从 Store 中读取并调用对应的函数。
    // 2. 给 `contentRenderFunction` 传递数据。

    ...

    return contentRenderFunctions(storeProps);
  }
}
```
</cn>

### Exercise: Implement the ConnectedStore component

<cn>
### 练习：实现 ConnectedStore 组件
</cn>

Do this exercise by forking [TimerView - ConnectedStore Exercise](http://codepen.io/hayeah/pen/ZbrrgV?editors=001).

<cn>
做这个练习可以 Fork [TimerView - ConnectedStore Exercise](http://codepen.io/hayeah/pen/ZbrrgV?editors=001) 代码。
</cn>

The App is currently static:

```js
let App = () => {
  let propValues = {
    currentTime: new Date(),
    currentTick: 0
  };

  return (
    <TimerView {...propValues}/>
  )
}
```

<cn>
整个应用现在是静止不动的：

```js
let App = () => {
  let propValues = {
    currentTime: new Date(),
    currentTick: 0
  };

  return (
    <TimerView {...propValues}/>
  )
}
```

</cn>

Replace it with `ConnectedStore`:

```js
let App = () => {
  return (
    <ConnectedStore store={TimerStore} propNames={["currentTime","currentTick"]}>
      {propValues => <TimerView {...propValues}/>}
    </ConnectedStore>
  )
}
```

<cn>
使用 `ConnectedStore` 来替换：

```js
let App = () => {
  return (
    <ConnectedStore store={TimerStore} propNames={["currentTime","currentTick"]}>
      {propValues => <TimerView {...propValues}/>}
    </ConnectedStore>
  )
}
```

</cn>

Your result:

<video src="TimerView.mp4" controls autoplay></video>

<cn>
你的结果：

<video src="TimerView.mp4" controls autoplay></video>
</cn>

### Exercise: Refactor Cart component to use ConnectedStore

<cn>
### 练习：使用 ConnectedStore 重构 Cart 组件
</cn>

First, create `js/components/ConnectedStore`, and put the wrapper components you've implemented there.

Then modify the Cart components so it no longers depends on the `CartStore`. It should get `cartItems` directly from `this.props.cartItems`:

```js
class Cart extends React.Componente {
  render() {
    let {cartItems} = this.props;
    // ...
  }
}
```

<cn>
首先请新建文件：`js/components/ConnectedStore`。然后将你实现的 wrapper 组件放进去。

然后修改你的 Cart 组件，使之不再和 `CartStore` 有任何关系。它能直接从 `this.props.cartItems` 中获取 `cartItems`：

```js
class Cart extends React.Componente {
  render() {
    let {cartItems} = this.props;
    // ...
  }
}
```

</cn>

Modify `CartStore` to follow the `ConnectedStore` naming convention:

```js
module.exports = {
  // Reader methods
  getCartItems() {
    return _cartItems;
  },

  // An alias method for getCartItems.
  cartItems() {
    return _cartItems;
  },
}
```

<cn>
修改 `CartStore`，遵循 `ConnectedStore` 中的命名约定：

```js
module.exports = {
  // Read 方法
  getCartItems() {
    return _cartItems;
  },

  // getCartItems 方法的别名
  cartItems() {
    return _cartItems;
  },
}
```
</cn>

Instead of exporting the Cart component, export the `ConnectedCart` component:

```js
class ConnectedCart extends React.Component {
  render() {
    return (
      <ConnectedStore ...>
        ...
      </ConnectedStore>
    );
  }
}

module.exports = ConnectedCart;
```

<cn>
导出 `ConnectedCart` 组件而不是 Cart 组件：

```js
class ConnectedCart extends React.Component {
  render() {
    return (
      <ConnectedStore ...>
        ...
      </ConnectedStore>
    );
  }
}

module.exports = ConnectedCart;
```
</cn>

# Liked Products

<cn>
# 喜欢的商品
</cn>

So far the only store we have is the `CartStore`. Now we are going to create the `LikeStore` so we can like products.

<cn>
到现在为止我们只有 `CartStore`。现在我们来创建 `LikeStore` 来存放我们喜欢的商品。
</cn>

You will first implement the feature directly in the `Products` component, then refactor it with `ConnectedStore`.

[Download the heart icon](heart-liked.svg)

![](heart-liked.svg)

<cn>
你应当先在 `Products` 组件中实现这个功能，再使用 `ConnectedStore` 将它重构。

[下载 Heart Icon](heart-liked.svg)

![](heart-liked.svg)
</cn>

### Exercise: Implement The Like Feature

<cn>
### 练习：实现喜欢的功能
</cn>

+ Create the `LikeStore`.
+ `Products` should connect to LikeStore and ProductStore.

<cn>

+ 创建 `LikeStore`。
+ `Products` 应当和 LikeStore 还有 ProductStore 相关联。
</cn>

Your result:

<video src="like-items.mp4" controls></video>

<cn>
你的结果：

<video src="like-items.mp4" controls></video>
</cn>

### Exercise: Turn Products component into a dumb component

<cn>
### 练习：将 Product 组件转化成「笨拙」的组件
</cn>

Separate the view and store concerns as we've done previously:

```js
// The dumb view component.
class Products extends React.Component {
  render() {
    let {cartItems,likeItems} = this.props;
    ...
  }
}

// The store data component.
class ConnectedProducts extends React.Component {
  render() {
    return (
      ...
    );
  }
}

module.exports = ConnectedProducts;
```

<cn>
正如前面所提到的，我们将 Store 和 View 职责上的忧虑分开了：

```js
// View 组件
class Products extends React.Component {
  render() {
    let {cartItems,likeItems} = this.props;
    ...
  }
}

// Store 组件
class ConnectedProducts extends React.Component {
  render() {
    return (
      ...
    );
  }
}

module.exports = ConnectedProducts;
```

</cn>

We need to connect both LikeStore and CartStore, but `ConnectedStore` can connect to one store at a time. To connect to both stores, nest two `ConnectedStore` components like this:

```js
<ConnectedStore store={store2}>
  {propsOfStore1 => {
    return (
      <ConnectedStore store={store2}>
        {propsOfStore2 => {
          ...
        }}
      </ConnectedStore>
    )
  }}
</ConnectedStore ...>
```

<cn>
我们需要一个同时能够连接 LikeStore 和 CartStore 的组件，但是 `ConnectedStore` 一次只能连接一个组件，不是吗？为了连接两个 Store，将两个 `ConnectedStore` 组件堆叠起来，就像这样：

```js
<ConnectedStore store={store2}>
  {propsOfStore1 => {
    return (
      <ConnectedStore store={store2}>
        {propsOfStore2 => {
          ...
        }}
      </ConnectedStore>
    )
  }}
</ConnectedStore ...>
```

</cn>

This is ugly, which we'll fix it next!

<cn>
这很丑！接下来我们会修复这个问题。
</cn>

# The "ConnectedComponent" Component Factory

<cn>
# ConnectedComponent 组件工厂
</cn>

Although we've built the `ConnectedStore` wrapper component, we still build the data component by hand:

```js
class Foo extends React.Component {

}

class ConnectedFoo extends React.Component {
  render() {
    return (
      <ConnectedStore store={FooStore} propNames={["fooData1","fooData2"]}>
        {fooProps => <Foo {...fooProps}/>}
      </ConnectedStore>
    )
  }
}

module.exports = ConnectedFoo;
```

<cn>
尽管我们已经做好了 `ConnectedStore` 这个 wrapper 组件，我们仍然需要手写代码：

```js
class Foo extends React.Component {

}

class ConnectedFoo extends React.Component {
  render() {
    return (
      <ConnectedStore store={FooStore} propNames={["fooData1","fooData2"]}>
        {fooProps => <Foo {...fooProps}/>}
      </ConnectedStore>
    )
  }
}

module.exports = ConnectedFoo;
```
</cn>

Now we want to simplify this one step further. We can use a function to construct the data component automatically:

```js
class Foo extends React.Component {
  // ...
}

module.exports = MakeConnectedComponent(Foo,FooStore,"fooData1","fooData2");
```

<cn>
现在我们想进一步简化这个步骤。我们需要一个函数来自动生成所需要的数据组件：

```js
class Foo extends React.Component {
  // ...
}

module.exports = MakeConnectedComponent(Foo,FooStore,"fooData1","fooData2");
```

</cn>

### Exercise: Refactor "Cart.js" with the "MakeConnectedComponent" function

<cn>
### 练习：使用 MakeConnectedComponent 来重构 Cart.js
</cn>

Create `js/components/MakeConnectedComponent.js`. It should export a function that returns the definition of a connected component.

```js
function MakeConnectedComponent(ViewComponent,store,...propNames) {
  // Note: The argument "ViewComponent" must be uppercase. Why?

  // TODO: Define ConnectedViewComponent

  // Return the component
  return ConnectedViewComponent;
}

module.exports = MakeConnectedComponent;
```

<cn>
新建 `js/components/MakeConnectedComponent.js`。我们需要导出一个能够返回一个连接组件定义的函数：

```js
function MakeConnectedComponent(ViewComponent,store,...propNames) {
  // 注意：ViewComponent 这个参数必须大写，为什么？

  // TODO：定义 ConnectedViewComponent

  // 返回 Component
  return ConnectedViewComponent;
}

module.exports = MakeConnectedComponent;
```
</cn>

Use this function to remove the `ConnectedCart` data component:

```js
/*

// not needed anymore
class ConnectedCart extends React.Component {
  ...
}

module.exports = ConnectedCart;
*/

module.exports = MakeConnectedComponent(Cart,CartStore,"cartItems");
```

<cn>
使用这个函数来去掉 `ConnectedCart` 的数据组件：

```js
/*

// 不再需要了
class ConnectedCart extends React.Component {
  ...
}

module.exports = ConnectedCart;
*/

module.exports = MakeConnectedComponent(Cart,CartStore,"cartItems");
```
</cn>


Question: The argument "ViewComponent" must be uppercase. Why? What if it's lowercase?

<cn>
问题：ViewComponent 这个参数必须大写，为什么？如果是小写的会发生什么问题？
</cn>

Hint: A class in JavaScript is implemented with functions. It can capture variables and arguments in its scope.

```js
(function() {
  let foo = 10;

  class Foo {
    // This returns the value of `foo`.
    getFoo() {
      return foo;
    }
  }

  return Foo;
})();
```

<cn>
提示：在 JavaScript 中，类（Class）由函数构成。类在其作用域中能够捕获变量和参数。

```js
(function() {
  let foo = 10;

  class Foo {
    // 返回 `foo` 的值。
    getFoo() {
      return foo;
    }
  }

  return Foo;
})();
```

</cn>

Hint: Instead of using `class`, you could also use a [stateless function component](https://facebook.github.io/react/blog/2015/10/07/react-v0.14.html#stateless-functional-components) to implement this.

<cn>
提示：不使用 `class`，你可以使用 [stateless function component](https://facebook.github.io/react/blog/2015/10/07/react-v0.14.html#stateless-functional-components) 来实现。
</cn>

Hint: Don't overthink, it's easy.

<cn>
提示：不！要！想！太！多！其实很容易。
</cn>

### Exercise: MakeConnectedComponent should be nestable.

<cn>
### 练习：MakeConnectedComponent 应该是可嵌套的
</cn>

Rewrite the `Products` with `MakeConnectedComponent`. Connect to two stores like this:

```js
module.exports =
  MakeConnectedComponent(
    MakeConnectedComponent(Products,CartStore,"cartItems"),
    LikeStore,"likeItems");
```

<cn>
使用 `MakeConnectedComponent` 重写 `Products`。像这样连接两个 Store：

```js
module.exports =
  MakeConnectedComponent(
    MakeConnectedComponent(Products,CartStore,"cartItems"),
    LikeStore,"likeItems");
```

</cn>

Let's break this out to make it easier to explain what's going on.

```js
innerComponent = MakeConnectedComponent(Products,CartStore,"cartItems");
outterComponent = MakeConnectedComponent(innerComponent,LikeStore,"likeItems");
module.exports = outterComponent;
```

<cn>
我们将这一句分离出来，看一下究竟发生了什么：

```js
innerComponent = MakeConnectedComponent(Products,CartStore,"cartItems");
outterComponent = MakeConnectedComponent(innerComponent,LikeStore,"likeItems");
module.exports = outterComponent;
```

</cn>


1. The outterComponent is connected to LikeStore, and passes `likeItems` to innerComponent
2. The innerComponent is connected to CartStore, it receives `likeItems` from outterComponent, and `cartItems` from the store.

<cn>
1. outterComponent 连接了 LikeStore，为 innerComponent 传递了 `likeItems` 参数。
2. innerComponent 连接了 CartStore，接收从 outterComponent 传进来的 `likeItems` 参数，和 Store 中的 `cartItems` 参数。
</cn>

MakeConnectedComponent should correctly merge the props it receives from outside with the props it receives from the store.

<cn>
MakeConnectedComponent 能够正确地合并从外界传进来的参数和从 Store 传递过来的参数。
</cn>

With JSX, you can merge two sets of properties like this:

```js
<Foo {...propsA} {...propsB}/>
```

<cn>
使用 JSX 你可以这样合并：

```js
<Foo {...propsA} {...propsB}/>
```

</cn>

# The "Connect" Decorator

<cn>
# Connect 装饰器
</cn>

Finally, we'll define a JavaScript decorator to connect a component to multiple stores. The syntax looks like this:

```js
// Note: There cannot be ';' after decorators
@connect(LikeStore,"likeItems")
@connect(CarStore,"cartItems")
class ConnectedProducts extends Products {}

module.exports = ConnectedProducts;
```

<cn>
最后我们将定义一个 JavaScript 装饰器来让一个组件连接多个 Store。格式看上去像这样：

```js
// 注意：行末没有「；」
@connect(LikeStore,"likeItems")
@connect(CarStore,"cartItems")
class ConnectedProducts extends Products {}

module.exports = ConnectedProducts;
```
</cn>

We can define the decorator as a function:

```js
function connect(store,...cartItems) {
  return (klass) => {

    // calculate klassReplacement

    return klassReplacement;
  };
}
```

<cn>
我们可以定义这样一个装饰器方法：

```js
function connect(store,...cartItems) {
  return (klass) => {

    // 计算 klassReplacement

    return klassReplacement;
  };
}
```

</cn>

This function receives the decorator arguments, and returns a function. The returned function replaces the decorated class with something else.

<cn>
这个方法将一个装饰器方法作为一个参数并且返回另一个函数。返回的函数会替代之前的装饰器方法。
</cn>

To see how this works, we could compile the ConnectedProducts class with babel. Put the class definition in `connected-products.js`, and compile:

```
// Stage 0 enables all experimental ES7 features, including the decorator.
babel --stage 0 connected-products.js
```

<cn>
要看这是怎么工作的，我们需要使用 babel 编译一下 ConnectedProducts。将类的定义放入 `connected-products.js` 然后编译：

```
// Stage 0 会开启 ES7 中所有的实验性特性，包括了装饰器。
babel --stage 0 connected-products.js
```

</cn>

The output:

```js
var ConnectedProducts = (function (_Products) {
  _inherits(ConnectedProducts, _Products);

  function ConnectedProducts() {
    _classCallCheck(this, _ConnectedProducts);

    _get(Object.getPrototypeOf(_ConnectedProducts.prototype), "constructor", this).apply(this, arguments);
  }

  var _ConnectedProducts = ConnectedProducts;

  /*
  The original class is replaced by the decorator, twice.
  */
  ConnectedProducts = connect(CarStore, "cartItems")(ConnectedProducts) || ConnectedProducts;
  ConnectedProducts = connect(LikeStore, "likeItems")(ConnectedProducts) || ConnectedProducts;


  return ConnectedProducts;
})(Products);
```

<cn>
输出：


```js
var ConnectedProducts = (function (_Products) {
  _inherits(ConnectedProducts, _Products);

  function ConnectedProducts() {
    _classCallCheck(this, _ConnectedProducts);

    _get(Object.getPrototypeOf(_ConnectedProducts.prototype), "constructor", this).apply(this, arguments);
  }

  var _ConnectedProducts = ConnectedProducts;

  /*
  注意到被替换了两次。
  */
  ConnectedProducts = connect(CarStore, "cartItems")(ConnectedProducts) || ConnectedProducts;
  ConnectedProducts = connect(LikeStore, "likeItems")(ConnectedProducts) || ConnectedProducts;


  return ConnectedProducts;
})(Products);
```
</cn>

The key is how `ConnectedProducts` is replaced twice. The double function invocation may be confusing:

```js
connect(CarStore, "cartItems")(ConnectedProducts)
```

<cn>
重点是 `ConnectedProducts` 如何被替换了两次。这很让人困扰：

```js
connect(CarStore, "cartItems")(ConnectedProducts)
```
</cn>

We can rewrite it to make it little bit clearer:

```js
var calculateClassReplacement;

// `connect` returns a function that calculates the replacement for ConnectedProducts
calculateClassReplacement = connect(CarStore, "cartItems");
ConnectedProducts = calculateClassReplacement(ConnectedProducts);

calculateClassReplacement = connect(LikeStore, "likeItems");
ConnectedProducts = calculateClassReplacement(ConnectedProducts);

return ConnectedProducts;
```

<cn>
我们改写一下让函数表现得更清楚一点：

```js
var calculateClassReplacement;

// `connect` 返回一个函数，用于计算替代 ConnectedProducts 的替代物
calculateClassReplacement = connect(CarStore, "cartItems");
ConnectedProducts = calculateClassReplacement(ConnectedProducts);

calculateClassReplacement = connect(LikeStore, "likeItems");
ConnectedProducts = calculateClassReplacement(ConnectedProducts);

return ConnectedProducts;
```

</cn>

To make the decorator syntax works for webpack, enable the `stage=0` option:

```
webpack ... --module-bind "js=babel?stage=0"
```

<cn>
为了在 Webpack 下也能使用装饰器格式，启用 `stage=0` 的选项：

```
webpack ... --module-bind "js=babel?stage=0"
```

</cn>

Note: Remember to restart webpack.

<cn>
注意：记得重启 Webpack。
</cn>

### Exercise: Implement the `@connect` decorator

<cn>
### 练习：实现 `@connect` 装饰器
</cn>

Create the `connect` decorator function in `js/components/connect.js`.

<cn>
在 `js/components/connect.js` 中创建 `connect` 装饰器方法：
</cn>

Then refactor `Products.js` again to use the decorator:

```js
@connect(CartStore,"cartItems")
@connect(LikeStore,"likeItems")
class ConnectedProducts extends Products {};

module.exports = ConnectedProducts;
```

<cn>
使用装饰器重构 `Products.js`：

```js
@connect(CartStore,"cartItems")
@connect(LikeStore,"likeItems")
class ConnectedProducts extends Products {};

module.exports = ConnectedProducts;
```


</cn>

You could also write it without using the decorator syntax:

```js
class ConnectedProducts extends Products {};
ConnectedProducts = connect(LikeStore,"likeItems")(ConnectedProducts);
ConnectedProducts = connect(CartStore,"cartItems")(ConnectedProducts);
module.exports = ConnectedProducts;
```

<cn>
你也可以不使用装饰器格式来重写：

```js
class ConnectedProducts extends Products {};
ConnectedProducts = connect(LikeStore,"likeItems")(ConnectedProducts);
ConnectedProducts = connect(CartStore,"cartItems")(ConnectedProducts);
module.exports = ConnectedProducts;
```
</cn>

# Liked Products Filtering

<cn>
# 筛选出喜欢的商品
</cn>

Clicking the heart in the left sidebar should toggle between showing all products or showing only liked products.

<cn>
点击左边侧边栏的心可以在显示所有商品和只显示喜欢的商品中切换。
</cn>

### Exercise: Filter products by whether it's liked

<cn>
### 练习：筛选出喜欢的产品
</cn>

Create `ProductStore.js`.

```js
// For now, hardwire the `_products` variable with all the available products.
let _products = ...;

let _showOnlyLike = false;

module.exports = {
  // Readers
  products() {
    // Return all products
  },


  filteredProducts() {
    // Return all products or only liked products, depending on _showOnlyLike
  },

  // Actions
  toggleShowOnlyLike() {

  },
}
```

<cn>
新建 `ProductStore.js`。

```js
// 现在先暂时硬编码，将 `_products` 和所有可用的商品关联起来。
let _products = ...;

let _showOnlyLike = false;

module.exports = {
  // Reader 函数
  products() {
    // 返回所有的商品
  },


  filteredProducts() {
    // 根据 _showOnlyLike 筛选出过滤后的商品
  },

  // 行为
  toggleShowOnlyLike() {

  },
}
```
</cn>

The Products component should connect to 3 stores:

```js
@connect(CartStore,"cartItems")
@connect(LikeStore,"likeItems")
@connect(ProductStore,"filteredProducts")
class ConnectedProducts extends Products {};

module.exports = ConnectedProducts;
```

<cn>
Products 组件和三个 Store 相关联：

```js
@connect(CartStore,"cartItems")
@connect(LikeStore,"likeItems")
@connect(ProductStore,"filteredProducts")
class ConnectedProducts extends Products {};

module.exports = ConnectedProducts;
```
</cn>

Your result:

<video src="filter-liked-items.mp4" controls></video>

<cn>
你的结果：

<video src="filter-liked-items.mp4" controls></video>

</cn>

# Summary

<cn>
# 总结
</cn>

Instead of hardwiring stores into view components, we used 3 different techniques to separate the view and the data. By keeping the view components dumb, we can more easily reuse them in different situations.

<cn>
不同于硬编码将 Store 和 View 组件连接起来，我们使用了三种不同的技术来分离视图和数据的职责。通过使用这种「笨拙」的视图组件，我们能够在不同情况下，更加容易地复用这些组件。
</cn>



