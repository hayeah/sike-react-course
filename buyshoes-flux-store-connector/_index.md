# Connecting Stores And Views With Ease

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

All the code we'll have to write is only about 50~60 lines, but these work at a higher level of abstractions.

+ We'll create a component that accepts a function as its body.
+ We'll define a function that returns a component definition.
+ Finally, we'll define a function that returns a function that takes a component definition and returns another component definition. Confused? Me too.

In other worrds, we'll have some fun with React metaprogramming!

# Separation of Concerns

We are not doing this refactoring just to write less code.

The real reason is so we can go back to writing React as though the app is completely static. We should create our components as though there isn't any Flux stores to connect to.

A component should be "dumb". It takes props as input, and renders the UI as output:

```js
class FooView extends React.Component {
  render() {
    let {fooData} = this.props;
    return <div>Foo data is: {fooData}</div>;
  }
}
```

A dumb component like this doesn't depend on anything except its props. It is more reusable than a component that's connected to the store, for the simple reason that anybody can pass props to it.

+ The props could come from a store.
+ The props could come from a unit test.
+ The props could come from a parent component.

Your app would be much easier to understand if all of your components are "dumb" like this.

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

This convention separates two concerns into two components:

1. `FooView` renders the actual view. It doesn't care where its props are from.
2. `FooViewStoreWrapper` passes the latest store data to `FooView`. It doesn't know what the view would render.

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


# The ConnectedStore Component

Our first step is to create a reusable wrapper component that could read data from a store, and pass the data down to a view. Let's call this wrapper component `ConnectedStore`.

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

It updates the time and tick counter every second:

[TimerView Demo](http://codepen.io/hayeah/pen/OyQQJV?editors=001)

<video src="TimerView.mp4" controls autoplay></video>

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

Now `TimerView` is a "dumb components". It doesn't know anything about the TimeStore.

How do we use the `ConnectedStore` to pass data to TimerView? By convention, we'll make the store's reader names and the view's property names to be the same:

+ `TimeStore.currentTime()` -> `this.props.currentTime`
+ `TimeStore.currentTick()` -> `this.props.currentTick`

Then we could use the wrapper component like this:

```js
<ConnectedStore store={TimerStore} propNames={["currentTime","currentTick"]}>
  {props => <TimerView {...props}/>}
</ConnectedStore>
```

+ `store` - Should connect to TimerStore by calling `addChangeListener`.
+ `propNames` - Should read these data from the store.
+ `props => ...` - Everytime the store emits "change", read data from the store, and call this function to render the view.

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

### Exercise: Implement the ConnectedStore component

Do this exercise by forking [TimerView - ConnectedStore Exercise](http://codepen.io/hayeah/pen/ZbrrgV?editors=001).

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

Your result:

<video src="TimerView.mp4" controls autoplay></video>

### Exercise: Refactor Cart component to use ConnectedStore

First, create `js/components/ConnectedStore`, and put the wrapper components you've implemented there.

Then modify the Cart components so it no longers depends on the `CartStore`. It should get `cartItems` directly from `this.props.cartItems`:

```
class Cart extends React.Componente {
  render() {
    let {cartItems} = this.props;
    // ...
  }
}
```

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

# Liked Products

So far the only store we have is the `CartStore`. Now we are going to create the `LikeStore` so we can like products.

You will first implement the feature directly in the `Products` component, then refactor it with `ConnectedStore`.

[Download the heart icon](heart-liked.svg)

![](heart-liked.svg)

### Exercise: Implement The Like Feature

+ Create the `LikeStore`.
+ `Products` should connect to LikeStore and ProductStore.


Your result:

<video src="like-items.mp4" controls></video>

### Exercise: Turn Products component into a dumb component

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

This is ugly, which we'll fix it next!

# The "ConnectedComponent" Component Factory

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

Now we want to simplify this one step further. We can use a function to construct the data component automatically:

```js
class Foo extends React.Component {
  // ...
}

module.exports = MakeConnectedComponent(Foo,FooStore,"fooData1","fooData2");
```

### Exercise: Refactor "Cart.js" with the "MakeConnectedComponent" function

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


Question: The argument "ViewComponent" must be uppercase. Why? What if it's lowercase?

Hint: A class in JavaScript is implemented with functions. It can capture variables and arguments in its scope.

```
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

Hint: Instead of using `class`, you could also use a [stateless function component](https://facebook.github.io/react/blog/2015/10/07/react-v0.14.html#stateless-functional-components) to implement this.

Hint: Don't overthink, it's easy.

### Exercise: MakeConnectedComponent should be nestable.

Rewrite the `Products` with `MakeConnectedComponent`. Connect to two stores like this:

```js
module.exports =
  MakeConnectedComponent(
    MakeConnectedComponent(Products,CartStore,"cartItems"),
    LikeStore,"likeItems");
```

Let's break this out to make it easier to explain what's going on.

```js
innerComponent = MakeConnectedComponent(Products,CartStore,"cartItems");
outterComponent = MakeConnectedComponent(innerComponent,LikeStore,"likeItems");
module.exports = outterComponent;
```


1. The outterComponent is connected to LikeStore, and passes `likeItems` to innerComponent
2. The innerComponent is connected to CartStore, it receivies `likeItems` from outterComponent, and `cartItems` from the store.

MakeConnectedComponent should correctly merge the props it receives from outside with the props it receives from the store.

With JSX, you can merge two sets of properties like this:

```
<Foo {...propsA} {...propsB}/>
```

# The "Connect" Decorator

Finally, we'll define a JavaScript decorator to connect a component to multiple stores. The syntax looks like this:

```js
// Note: There cannot be ';' after decorators
@connect(LikeStore,"likeItems")
@connect(CarStore,"cartItems")
class ConnectedProducts extends Products {}

module.exports = ConnectedProducts;
```

We can define the decorator as a function:

```
function connect(store,...cartItems) {
  return (klass) => {

    // calculate klassReplacement

    return klassReplacement;
  };
}
```

This function receives the decorator arguments, and returns a function. The returned function replaces the decorated class with something else.

To see how this works, we could compile the ConnectedProducts class with babel. Put the class definition in `connected-products.js`, and compile:

```
// Stage 0 enables all experimental ES7 features, including the decorator.
babel --stage 0 connected-products.js
```

The output:

```
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

The key is how `ConnectedProducts` is replaced twice. The double function invokation may be confusing:

```js
connect(CarStore, "cartItems")(ConnectedProducts)
```

We can rewrite it to make it little bit clearer:

```
var calculateClassReplacement;

// `connect` returns a function that calculates the replacement for ConnectedProducts
calculateClassReplacement = connect(CarStore, "cartItems");
ConnectedProducts = calculateClassReplacement(ConnectedProducts);

calculateClassReplacement = connect(LikeStore, "likeItems");
ConnectedProducts = calculateClassReplacement(ConnectedProducts);

return ConnectedProducts;
```

To make the decorator syntax works for webpack, enable the `stage=0` option:

```
webpack ... --module-bind "js=babel?stage=0"
```

Note: Remember to restart webpack.

### Exercise: Implement the `@connect` decorator

Create the decorator function in `js/components/connect.js`.

Then refactor `Products.js` again:

```js
@connect(CartStore,"cartItems")
@connect(LikeStore,"likeItems")
class ConnectedProducts extends Products {};

module.exports = ConnectedProducts;
```

You could also write it without using the decorator syntax:

```js
class ConnectedProducts extends Products {};
ConnectedProducts = connect(LikeStore,"likeItems")(ConnectedProducts);
ConnectedProducts = connect(CartStore,"cartItems")(ConnectedProducts);
module.exports = ConnectedProducts;
```

# Liked Products Filtering

Clicking the heart in the left sidebar should toggle between showing all products or showing only liked products.

### Exercise: Filter products by whether it's liked

Create `ProductStore.js`.

```js
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

Your result:

<video src="filter-liked-items.mp4" controls></video>

For now, hardwire the `_products` variable with all the available products.

# Summary

Instead of hardwiring stores into view components, we used 3 different techniques to separate the view and the data. By keeping the view components dumb, we can mroe easily reuse them in different situations.


