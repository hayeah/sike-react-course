


We could write add these two features in a very straightforward way：

```
// hardwired actions
function addCartItem(productId) {
  CartStore.addCartItem(productId);

  LoggingService.log({
    time: new Date(),
    action: {type: "addCartItem", productId: productId}})

  UndoStore.saveHistory();
}

function removeCartItem(productId) {
  CartStore.removeCartItem(productId);

  LoggingService.log({
    time: new Date(),
    action: {type: "removeCartItem", productId: productId}})

  UndoStore.saveHistory();
}
```

But there are two problems:

1. `LoggingService.log` and `UndoStore.saveHistory` are repeated over all the actions.
2. It's not easy to enable and disable LoggingService or UndoStore.

But why add another extra layer of indirection between the actions and the stores? It solves two problems:

+ The dispatcher makes it easy to add a service that needs to intercept every single action (cross cutting concerns like logging and monitoring).
+ The dispatcher enables "Plug and play" architecture. You can add or remove parts of the app without changing any code.

```
// pubsub actions
function addCartItem(productId) {
  dispatcher.dispatcher({type: "addCartItem", productId: productId});
}

function removeCartItem(productId) {
  dispatcher.dispatcher({type: "removeCartItem", productId: productId});
}


// Plug and play services:

// CartStore
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


# Flux Frameworks Guide

+ AltJS - most similar to the full Flux. One dispatcher.
+ Redux - Action are messages. One store, which is transformed by a stateless reduce function to change the state. Connect the store to the top component.
+ Baobab - One store, immutable data structure. Connect the store and the view with cursor. Cursor is a pointer to an arbitrary location in the data store.


The full Flux architecture might be more work,

You don't have to wait until your codebase is huge to adopt action-store pubsub. The complete Flux architecture might be slightly more work,

We'll see how action-stire pubsub makes it easy to intercept all actions.


don't wait until your app is 10k LoC. Flux is suitable for small apps too!