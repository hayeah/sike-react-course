# Action To Store PubSub With Dispatcher

We've seen how Flux uses EventEmitter to connect stores and views, so if a piece of data changes, multiple views would update:

![](store-views-pubsub.jpg)

In this lesson we'll see add pubsub between actions and stores, so one action can trigger multiple stores to update:

![](action-stores-pubsub.jpg)

Usually there is a single dispatcher for your app, and every action goes through it.

# So Much Loose Coupling!

![](doge-loose-coupling.jpg)

Why so much pubsub?

+ It's "plug and play" architecture, so we can add and remove modules without changing any code.
+ Modules are better isolated. The only way to communicate with a module is by sending it an event. In other words, [Message Passing](https://en.wikipedia.org/wiki/Message_passing).
+ Modules can control exactly how they could be triggered, by choosing which events to listen to.

Although you have to write more lines of code for each module, the app as a whole becomes simpler. B

ut for a small to medium sized app, the benefits of action-store pubsub could be more theoretical than real. The supposed benefits of pubsub may sound more like:

+ blah blah blah blah~

In this lesson we'll add two additional features:

+ Logging service. To log all actions a user have taken.
+ Undo service. To be able to undo add to cart and remove from cart actions.

Even though our codebase is small, we'll see how action-store pubsub makes it easy to add these two features.

# The Central Dispatcher

The dispatcher is essentially an EventEmitter:

```js
import EventEmitter from "events";

const EVENT_NAME = "action";

module.exports = class Dispatcher {
  constructor() {
    this.emitter = new EventEmitter();
  }

  register(handler) {
    this.emitter.addListener(EVENT_NAME, handler);
  }

  unregister(handler) {
    this.emitter.removeListener(EVENT_NAME,handler);
  }

  dispatch(action) {
    this.emitter.emit(EVENT_NAME,action);
  }
}
```

By convention, an action is an object with the `type` property. And to pass arguments, you add additional properties to the object:

```js
function addCartItem(productId) {
  let actionObject = {
    // Type of the action.
    type: "addCartItem",

    // Arguments
    productId: productId,
  }
  dispatcher.dispatch(actionObject);
}
```

Then dispatcher passes action objects to stores. The store may look at the action object's type, and decide whether to handle the action or to ignore it. The CartStore looks like:

```js
// CartStore.js
// Listens to "action" events.
AppDispatcher.register((action) => {
  let handler = handlers[action.type];
  // Ignores the action if the store doesn't have a handler for it.
  handler && handler(action);
})

let handlers = {
  // Private Writer API. Don't export these functions!
  addCartItem(action) {
    let {productId} = action;
    _cartItems[productId] = {
      id: productId,
      quantity: 1,
    };

    emitChange();
  },
}
```

Previously we allowed views to call the store's writer API, now the only way to trigger store update is by passing action events through dispatchers.

Let's see how this extra layer of indirection changes the structure of our app. Consider the `addCartItem` and `removeCartItem` actions. Without using the dispatcher, we'd wire in three different services directly into the action:

```js
// hardwired actions
function addCartItem(productId) {
  CartStore.addCartItem(productId);

  LoggingService.log(...)

  UndoStore.saveHistory(...);
}

function removeCartItem(productId) {
  CartStore.removeCartItem(productId);

  LoggingService.log(...)

  UndoStore.saveHistory(...);
}
```

With the dispatcher, the two actions only emit events:

```js
function addCartItem(productId) {
  dispatcher.dispatcher({type: "addCartItem", productId: productId});
}

function removeCartItem(productId) {
  dispatcher.dispatcher({type: "removeCartItem", productId: productId});
}
```

These actions don't need to specify what exactly need to be done. Individual modules can choose to handle the action events however they want:

```js
// plug-and-play modules:

// CartStore.js
dispatcher.register(action => {
  // ...
});

// LoggingService.js
dispatcher.register(action => {
  // ...
});

// UndoStore.js
dispatcher.register(action => {
  // ...
});
```

# Modify Search Suggestions To Use Dispatcher

[Search Suggestions With Dispatcher Demo](http://codepen.io/hayeah/pen/qOoqov?editors=001)

First create a dispatcher for the app:

```
let dispatcher = new Dispatcher();
```

The `updateSearchQuery` function dispatches the `updateSearchQuery` action. It also calls `receiveSuggestions` when the RemoteAPI returns with data:

```
function updateSearchQuery(query) {
  dispatcher.dispatch({type: "updateSearchQuery", query: query});

  RemoteAPI.fetchSuggestions(query,(suggestions) => {
    receiveSuggestions(suggestions);
  });
}
```

The `receiveSuggestions` action no longer calls the store method directly:

```js
function receiveSuggestions(suggestions) {
  dispatcher.dispatch({type: "receiveSuggestions", suggestions: suggestions});
  // suggestionsStore.setSuggestions(suggestions);
}
```

The suggestionsStore listens to action events from the dispatcher, and its writer methods is no longer visible to the outside:

```js
let suggestionsStore = (() => {
  let _suggestions = [];

  let emitter = new EventEmitter();

  // The writer API is now private.
  function setSuggestions({suggestions}) {
    _suggestions = suggestions;
    emitter.emit("change");
  }

  // The only way to call the writer method is by listening to action.
  dispatcher.register((action) => {
     if(action.type === "receiveSuggestions") {
       setSuggestions(action);
     }
  });

  // Only the Reader API is exported.
  return {
    getSuggestions() {
      return _suggestions;
    },

    addChangeListener(callback) {
      emitter.addListener("change",callback);
    },
  };
})();
```

The views stay the same as before.

# Logging Service

Let's create a logging service so we can monitor all the action events that are being emitted.

### Exercise: Implement Logging Service

The logging service itself is super simple:

```js
let AppDispatcher = require("./AppDispatcher");

module.exports = function enableLogging() {
  AppDispatcher.register((action) => {
    console.log(JSON.stringify({
      timestamp: new Date(),
      action
    },undefined,2));
  })
}
```

Please modify the following three actions to use the dispatcher:

+ addCartItem
+ removeCartItem
+ updateCartItemQuantity

The shopping cart functionality should be the same as before. You'll need to make quite a few changes:

+ `actions.js` -  Dispatch action objects to the dispatcher.
+ `AppDispatcher.js` -  An instance of Dispatcher, shared by the whole app.
+ Modify the views to call functions exported by `actions.js`
+ `LoggingService.js` is the logging service as defined above.
  + Call `enableLogging` to start the logging service.
+ Modify the `CartStore.js` writer API to be private.

Your result:

<video src="buyshoes-logging.mp4" controls></video>

# Specify Update Order With waitFor

While the dispatcher is essentially an event-emitter, you sometimes want to make sure that a store gets to process a message before another store. Facebook's Flux Dispatcher adds the `waitFor` method to ensure exactly that.

Install it with npm.

```
npm install flux@2.1.1
```

The API of our DIY dispatcher is exactly the same. You can replace the DIY dispatcher with the dispatcher from Flux:

```
const {Dispatcher} = require("flux");
```

Your code should work as before. Now let's see in what order the subscribers receive an event. Create the file `test-dispatch-order.js`,

```
let dispatcher = new Dispatcher();

dispatcher.register((action) => {
  console.log("A", action);
});

dispatcher.register((action) => {
  console.log("C", action);
});

dispatcher.register((action) => {
  console.log("B", action);
});


dispatcher.dispatch({type: "test"});
```

Run with babel:

```
babel-node test-dispatch-order.js
```

You should see the output:

```
A { type: 'test' }
C { type: 'test' }
B { type: 'test' }
```

Try changing the order of the subscribers to get this output:

```
C { type: 'test' }
A { type: 'test' }
B { type: 'test' }
```

We can use `waitFor` to enforce an order. Here C waits for B, and B waits for C:

```
let tokenC = dispatcher.register((action) => {
  dispatcher.waitFor([tokenB]);
  console.log("C", action);
});

let tokenA = dispatcher.register((action) => {
  console.log("A", action);
});

let tokenB = dispatcher.register((action) => {
  dispatcher.waitFor([tokenA]);
  console.log("B", action);
});
```

So we always get the same output:

```
A { type: 'test' }
B { type: 'test' }
C { type: 'test' }
```

Without `waitFor`, the order of action handling depends on which modules are loaded first. Using `waitFor` is a much better solution than trying to get a module to load before others.

In general, stores should be designed so it doesn't matter which one get to run the action first. You should use `waitFor` only for special circumstances.

# Undo Service

We'll add an undo button beside the "shopping cart" title:

![](undo-button.jpg)

+ If you've just added an item, undo should remove that item.
+ If you've just removed an item, undo should add that item.

### Exercise: Implement Shopping Cart Undo

The HTML structure looks like:

```html
<div className="cart">
  <h3 className="cart__title">Shopping Cart</h3>
  <div className="cart__content">
    ...
  </div>

  <h3 className="cart__undo"><a onClick={this.undo}>undo</a></h3>
</div>
```

And the CSS for `cart__undo`:

```css
.cart__undo {
  position: absolute;
  top: 0; right: 0;
  padding: 15px 0;
  justify-content: center;
  z-index: 20;
}
```

The `UndoStore` has a history array that stores snapshots of CartStore:

```js
// UndoStore.js

// Array of cartItems
let history = [
];
```

The idea to implement this feature is simple. Everytime `UndoStore` receives the `addCartItem` and `removeCartItem` actions, it should copy `CartStore.cartItems`, and store it in the history. Use the [cloneDeep](https://lodash.com/docs#cloneDeep) utility function to copy cartItems.

The `undoShoppingCart` action is like this:

```
function undoShoppingCart() {
  let carItems = UndoStore.lastHistoryItem();
  dispatcher.emit({type: "undoShoppingCart", cartItems: cartItems})
}
```

This action should trigger two changes:

+ CartStore should restore its internal data to the snapshot.
+ UndoStore's history items should decrease by 1.

If there's no more history to undo, the "undo" button should be hidden.

Hint: `CartStore` and `UndoStore` will need to require each other, but circular dependency is a bit broken for the CommonJS module syntax. For these two modules, don't use CommonJS export syntax:

```js
module.exports = {
  ...
}
```

Replace the above with ES6 module syntax:

```js
export default {
  ...
}
```

Use `import` to use these modules:

```
// In UndoStore.js
import CartStore from "./CartStore"

// In CartStore.js
import UndoStore from "./UndoStore"
```

Your result:

<video src="undo.mp4" controls></video>