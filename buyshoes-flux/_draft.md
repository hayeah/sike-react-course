

# Monkey See Flux Monkey Do Flux

There are four steps to an update cycle:

![](flux-simple-f8-diagram-1300w.png)

When the user clicks on the "Add To Cart" button.

Action -> Dispatcher

1. The `add to cart button` calls the `addToCart` action creator.
1. The action creator sends an action to the dispatcher.

Dispatcher -> Store

1. The store receives action from the dispatcher, and process the action.
1. The store modifies its internal state.

Store -> View

1. The store emits the "change" event.
1. The view receives the "change" event.
1. The view gets the newest data from store, and rerenders.

### Code

Setup.

```
window.appDispatcher = new Dispatcher();

let CartStore = (function() {
  let _cartItems = {};

  function addCartItem(productId) {
    _cartItems[productId] = {
      id: product
    };
  }

  return {
    getCartItems() {
      return _cartItems;
    }
  };
})();
```

To implement `addToCart`, we'll need to connect the Dispatcher with the Store, and the Store with the View:

1. Create a dispatcher for the app.
1. Create the CartItemStore.
1. View should listen to CartItemStore's "change" event.
1. The CartItemStore should subscribe to dispatcher to receive actions.

When the user clicks on the add to cart button, we need to send an action into the dispatcher:

1. Define the "Action Creator" `addToCart`.
1. The action creator `addToCart` creates the an "action" object, and send it to the dispatcher.

The dispatcher sends the action to the `CartItemStore`:

1. The CartItemStore receives the `addToCart` action, and modifies its internal data.
1. The CartItemStore changes its internal state, and emit a "change" event.

Finally, the store notifies the view to update:

1. The `Cart` component recevies the `change` event from CartItemStore. Rerenders.
1. The `Products` components recevies the `change` event from CartItemStore. Rerenders.




Looking at these apps, though, you might find them confusing and unnecessarily verbose.

You might suspect that Flux solves a problem you don't have.

We won't use any of these existing frameworks. Rather,


we'll grow our own Flux from nothing. As we develop the shopping cart application we'll encounter new needs, and these needs would inspire us to add new features to our Flux framework.

By evolving the framework in lockstep with our application, we can better understand what problems different parts of Flux are designed to solve.




# Intro To Flux



+ [Facebook's original Flux](https://facebook.github.io/flux)
+ [alt.js](http://alt.js.org/)
+ [reflux](https://github.com/spoike/refluxjs)
+ [redux](https://github.com/gaearon/redux)
+ etc.

Flux is a pattern, and all the Flux frameworks mentioned above share the same basic ideas.

Of these, Facebook's original Flux is the most minimal. The core component, the [Dispatcher](https://github.com/facebook/flux/blob/e0a52d0dada7b78e7197026dc8cf35d84991f0e8/src/Dispatcher.js#L109-L240), is only 150 lines of code. The rest of Flux, Action and Store, are written with plain JavaScript, following the Flux convention.

Since there's no framework API for Action and Store, people using Flux usually get started by imitating the





Rather than jumping into an existing framework, we'll implement our own Flux.

Flux is radically different from .

By implementing your own Flux, we hope that you can quickly learn whichever Flux framework you eventually choose to use.


Flux is consisted of four types of objects. You've probably seen the Flux "unidirectional-flow" graph:

![](flux-simple.png)

+ Action Creator & Action
+ Dispatcher
+ Stores
+ Views




There are many Flux implementations available, all share the same basic ideas. Some of the most popular Flux implementations are:

+ [Facebook's original Flux](https://facebook.github.io/flux)
+ [alt.js](http://alt.js.org/)
+ [reflux](https://github.com/spoike/refluxjs)
+ [redux](https://github.com/gaearon/redux)
+ etc.


no framework API for



imitate the Flux example apps Facebook provided.


Rather than choosing an existing Flux framework  we are going to implement our own "flux".

# No Need For MVC

+ smart components. can have logic and state.
+ no need for models. just use raw data. components are our "view helpers". derived data from props.


question: what's the minimal amount of code we need to manage global application state?

80% of controller code is gone.

+ Action
+ Dispatcher
+ Store
+ View



Why Flux is like that.

Step by step. Needs driven.

Flux is a pattern. There is a chicken-and-egg learning problem. It's hard to learn Flux without a framework, and it's hard to learn a framework without having learned flux already.



# Buyshoes




+ User clicks the add to cart button.
+ The button triggers the `addToCart` action.
+ The `addToCart` action changes the application state by adding a new product to cart.
+ All the views should update to show that there is a new shopping cart item.


React component has view, logic, and state.

{
  god this is so hard to explain.

  react is a view-controller hybrid. why does it need an extra "application framework" to control the component tree?

  application state vs UI state.

  component isolation. awkward to pass state up and down the component tree.

}



Application state vs UI state.
  + UI state is something like hidden or show.
  + application state could be shared by many components.

Manage application state. MVC or Flux.

why not MVC?

+ Flux VS MVC seems like a red-herring.
+ It's more like we don't want models. We don't want controllers. Plain JS + PubSub.


It's hard to understand what flux is about because...

1. It seems unecessarily indirect.
2. Abstract.
3. Not really an API, but how code is glued together.




```
Flux is the application architecture that Facebook uses for building client-side web applications. It complements React's composable view components by utilizing a unidirectional data flow. It's more of a pattern rather than a formal framework, and you can start using Flux immediately without a lot of new code.
```

not MVC vs FLUX. Both are haaaard to understand.

procedural programming vs pubsub.

concretely: replace function calls with events.


our strategy

+ invent flux from first principle.
+ iterate/refactor by needs.




+ the code is so abstract, that it's kinda meaningless.
+ the official "flux" hardly has anything. it's more about how you use it.
+ without understanding the "pattern", it's also hard to understand a formal framework like redux or alt.





```
Flux applications have three major parts: the dispatcher, the stores, and the views (React components). These should not be confused with Model-View-Controller.

Controllers do exist in a Flux application, but they are controller-views — views often found at the top of the hierarchy that retrieve data from the stores and pass this data down to their children.

Additionally, action creators — dispatcher helper methods — are used to support a semantic API that describes all changes that are possible in the application. It can be useful to think of them as a fourth part of the Flux update cycle.

Flux eschews MVC in favor of a unidirectional data flow. When a user interacts with a React view, the view propagates an action through a central dispatcher, to the various stores that hold the application's data and business logic, which updates all of the views that are affected. This works especially well with React's declarative programming style, which allows the store to send updates without specifying how to transition views between states.

We originally set out to deal correctly with derived data: for example, we wanted to show an unread count for message threads while another view showed a list of threads, with the unread ones highlighted. This was difficult to handle with MVC — marking a single thread as read would update the thread model, and then also need to update the unread count model. These dependencies and cascading updates often occur in a large MVC application, leading to a tangled weave of data flow and unpredictable results.

Control is inverted with stores: the stores accept updates and reconcile them as appropriate, rather than depending on something external to update its data in a consistent way. Nothing outside the store has any insight into how it manages the data for its domain, helping to keep a clear separation of concerns. Stores have no direct setter methods like setAsRead(), but instead have only a single way of getting new data into their self-contained world — the callback they register with the dispatcher.
```


# this.prop vs this.state vs Flux





+ TODO: would be interesting to do something like "a tour of flux/alt/redux/baobab"



# pub/sub metaphor

+ broadcast sytem
  + caller doesn't know who's going to respond
  + receiver can choose to ignore a message if not relevant
  + easier to grow. just hook up a new component to the broadcast system, and you can start to respond.

+ funnily enough, there is no realworld metaphor for synchronous programming. it's comically inefficient.
  + there is a handoff of control. sort of like collecting signatures.

how about automating a house as a metaphor? you could use a centralize the program in a controller, or you can use a broadcast system.

"lockup"
"iamhome", "howard"


# EventEmitter

```
const EventEmitter = require("events");
let emitter = new EventEmitter();

emitter.on("login",(name) => {
  console.log("hello",name);
});

emitter.on("logout",(name) => {
  console.log("goodbye",name);
});

emitter.emit("login","howard");
emitter.emit("logout","howard");

```

add new functionality without modifying existing code:

```
let loggedInUsers = {
};

emitter.on("login",(name) => {
  loggedInUsers[name] = {
    loggedInAt: new Data(),
    name: name,
  }
});

emitter.on("logout",(name) => {
  delete loggedInUsers[name];
});
```


# EventEmitter Inversion Of Control

calling methods directly vs indirectly via events/callback.

```js
let controller = {
  searchInputUpdated() {
    updateView1();
    updateView2();
    updateView3();
  }
}

searchInput.onChange = controller.searchInputUpdated;
```

pubsub style

```
const SEARCH_INPUT_CHANGED = "SEARCH_INPUT_CHANGED";

let emitter = new EventEmitter();

searchInput.onChange = () => {
  emitter.emit(SEARCH_INPUT_CHANGED);
};

emitter.on(SEARCH_INPUT_CHANGED,updateView1);
emitter.on(SEARCH_INPUT_CHANGED,updateView2);
emitter.on(SEARCH_INPUT_CHANGED,updateView3);
```

+ loose coupling and pub/sub
  + many teams working independently.

+ loose coupling. pubsub.
  + inversion of control
+ single direction flow, originating from action.
  + anti 2-way-binding?
+ no cascading. action should not trigger more actions. store should not update stores.

Can I think of a simple example that refactor to flux?

how about an autocomplete input box?


+ MVC "direct style"
  http://codepen.io/hayeah/pen/pjdpPP?editors=001

+ PubSub "indirect style"
  http://codepen.io/hayeah/pen/NGwXXb?editors=001

+ loose coupling: extending functionality, not changing existing code.


```
searchInput.onChange = function onQueryChange(queryString) {
  autoCompleteController.updateQuery(queryString);
}

class AutoCompleteController {
  updateQuery(queryString) {
    RemoteSuggestionsAPI.getSuggestions(queryString).then((suggestions) => {
      this.receivedSuggestions(suggestions);
    });
  }

  receivedSuggestions(suggestions) {
    suggestionStore.setSuggestions(suggestions);
    suggestionChooserView.updateView();
  }
}

class SuggestionsStore {
  getSuggestions() {
    return this.suggestions;
  }

  setSuggestions(suggestions) {
    this.suggestions = [];
  }
}

class SuggestionsChooserView {
  updateView() {
    let suggestions = suggestionStore.getSuggestions();
  }
}
```

```
searchInput.onChange = function onQueryChange(queryString) {
  emitter.emit("UpdateSearchQuery",queryString);
}

emitter.on("UpdateSearchQuery",(queryString) => {
  RemoteSuggestionsAPI.getSuggestions(queryString).then((suggestions) => {
    emitter.emit("UpdateSuggestions",suggestions);
  });
});

class SuggestionsStore {
  getSuggestions() {
    return this.suggestions;
    emitter.emit("SuggestionsStoreChanged");
  }

  setSuggestions(suggestions) {
    this.suggestions = suggestions;
  }
}

class SuggestionsChooserView {
  updateView() {
    let suggestions = suggestionStore.getSuggestions();
    // blah blah blah
  }
}

emitter.on("SuggestionsStoreChanged",SuggestionsChooserView.updateView);
```

+ store. "semantic action". abstract data type.

+ event emitter

+ wrapper component

+ higher-order component
+ decorator




+ I think it'd be nice if the lesson could retrace flux's origin. sort of "hey, i could've invented flux too!"

Flux Part 1: Store and View Pub/Sub

+ higher-order component
+ "semantic action". abstract data type
+ event emitter

+ remote sync api. baked into the action creators.

Flux Part 2: Action & Dispatcher

+ refactor `saveCart` into action -> dispatcher.
+ add undo service.


ya i think that's good.

# dev note

+ upgrade react to 0.14.
  + install react-dom
  + ReactDOM.render
  + React.findDOMNode deprecated. use ref directly.
    + `$foo` is a good convention.


+ CartStore
  + NodeJS event emitter.
  + wraps a raw variable. access data via API functions.
  + can add and remove item.
  + manually bind to `change` event.

+ react-store binding
  + wrapper component
  + make a separate CartContainer. make cart itself pure.

  + abstract CartContainer into

  + higher order component
    + the lightweight component is nice. `(props) => { ... }`
    + connect(Cart,CartStore,"cartItems")

  + decorator https://github.com/wycats/javascript-decorators
    + it's cool~ but probably not worth the effort. it's pretty simple once `connect` is implemented.

+ chaining should work
  connect(connect(Cart,CartStore,"cartItems"),ProductStore,"foo","bar");

+ likes?
  + enable decorator "js=babel?stage=0"

+ todo: need a way to make Store a little bit more abstract.
  + the way `connect` is designed doesn't work well with ES6 module. there's no default object.

+ adjust quantity
  + check for min and max quantity
  + could keep components simple, and push logic to the store?

+ calculate checkout
  + CouponStore

+ how would filtering work? ProductFilterStore?
  + what i want is a single `getFilteredProduct` method.
  + i think it's easier to build a `<FilteredProduct>` component, that connects to multiple stores.

  or a utility function that derive data from two data sources.

  ```
  combine({
    stores: {store1, store2},
    foobar() {
      return store1.foo() + store2.bar();
    },
  }}
  ```

  ya. i don't think you want to a derived store, b/c it's better for the stores to be flat.

+ Store interface.

hmmm. I think it's nice that the store is just a bag of functions... but it's hard to share the boilerplate.

+ ProductStore
  + filteredProducts
  + setFilter({like: true, name: "substring"})

+ mock remote API

getCartItems
syncCartItems

hmm. change each action to sync with apidraft seems stupid... it's obviously a cross-cutting concern.

yet it's better that stores are "dumb". they shouldn't trigger further actions.

oy vey. a few choices

1. smart store, dumb action. do loading/saving in store.
2. smart action, dumb store. do loading/saving in action.
3. loading in store. saving in action.

where to put async ajax?

1. component.
2. in a store. listen to dispatcher.
  + caveat: action triggering action tend to be confusing. in general avoid.
3. in action creators.

errorHandling & retry

SyncService - listening to actions and dispatch AJAX sync requests. No binding with view.

API.js - async API with callback.

UIStore
  + spinner
  + notificationView

+ dispatachify
  + watch out for circular dependency when store and action creators both require dispatcher.
  + actions can require store, but store should never require action.

hmm.

+ undo & redo

npm install keymirror --save

action->store

hmm. the current store thing is hard to refactor into dispatcher.

bag of function is nice. switch import from store to actions, and no change to view code.


```
@connected(CartStore,"cartItems")
@connected(LikeStore,"likedItems")
export class ProductsContainer extends Products {};
```


```
<Connect store={CartStore} props={["a","b",{"c","renamedC"}]}>
  {(props) => {
    <Cart {...props}}>
  }}
</Connect>

<Connect store={CartStore} props={["a","b",{"c","renamedC"}]} view={Cart}/>
```

```
@connect(CartStore,"cartItems")
@connect(ProductStore,"products",{"foo","renamedFoo"})
class CartContainer extends Cart {
  ...
}
```


# lesson strategy

ok. so i want to "grow" towards flux.

+ local data update
+ mocked API
+ action->store pub/sub. behavioural monitoring & logging.
+ immutable & undo API


# flux structure

one dispatcher, multiple stores.

CartStore
  getCartItems

  # actions

  addItem
  removeItem

  increaseQuantity(productId)
  decreaseQuantity(productId)

ProductStore
  getProducts

  # actions
  receiveProducts

LikeStore
  getLikes

  # actions

  likeProduct

# page load init

receiveProducts(products)
receiveCartItems
receiveLikes

# async remote API

holy crap. it's not easy to do optimistic update.

+ timeout: no idea whether it succeeded or not
+ explicit error: rollback? wha?
+ explicit success: last write win?

so... i think we need to flip it around. "client" has the truth. server is just a sync storage. no rollback.

```
// Async Remote API
function saveCartItem(productId,quantity) -> Promise<{true}>
function removeCartItem(productId)
```

addItem, increaseQuantity, decreaseQuantity all trigger sync?

actually, you know what? couldn't we use pub/sub to implement the data sync service?

```
// debounce.
// auto-retry on error. display error message.
api.saveCartItems(items);
api.getCartItems(items);
```

# undo/redo API

Could use command and inverse commands,

```
addItem -> removeItem
removeItem -> addItem
increaseItem -> decreaseItem
decreaseItem -> increaseItem
```

OR transition to using immutable data structure. Save a snapshot of the shopping cart each time.

Poor man's version: copy.

Here `waitFor` is useful. CartStore wait for UndoStore to record history before updating.






















important ideas are:

+ semantic action. indirect manipulation of state. possibly pub/sub
+ store is "domain state". single source of truth.
  + no cascading. snapshot view while processing an action. state transition T(state,action).
+ pub/sub of store changes to view.

I think it's good to introduce the classic flux architecture.

+ Realistically speaking, BuyShoes doesn't need all that.
+ Can I think of an interesting example for the Action->Store pub/sub?

It probably makes sense to break into two parts, where 1 is store->view pub/sub, the other is action->store pub/sub.


dispatcher adds two things:

1. do not cascade.
  + testing is easier
  + easier to understand.
2. assert order of message delivery.


I think the granularity of flux varies. It could be very fine, or rather coarse.

+ fine - each model. each variable is a store.
+ granular - a "domain" or subsystem is a store.

If too fine, can seen overwrought.



immutable


sugar, data structure, store->view pub/sub, action->store pub/sub, integration with view

flux
  store data structure: any
  state transition: any
  store->view pub/sub: yes.
  action->store pub/sub: yes
  multiple dispatchers
  getter: store getter functions

alt
  same as flux, with sugar.

redux
  data structure: immutable
  state transition: stateless/pure reduce
  store->view pub/sub: yes
  action->store pub/sub: no
  no dispatcher
  getter: props from root component

baobab
  data structure: immutable
  state transition: stateless/pure setter
  store->view pub/sub: yes
  action->store pub/sub: no
  no dispatcher
  getter: cursor or getter functions


+ remote/async API integration in "smart actions".
  + optimistic update
+ pub/sub action->store.
  + concern 1: undo
  + concern 2: user behavioural monitoring / logging


Also... normalizing data.




http://stackoverflow.com/a/25648726


part 1: store->view pubsub

```
let carItems = {[productId]: {quantity, productId}};

getCartItems()
```

```
// cartAction
addCartItem(prductId) {
  // direct manipulation + event emitter
}
```

part 2: action->store pubsub

```
function addCartItem(productId) {
  AppDispatcher.dispatch({type:"addCartItem",productId})
}

// cartAction
_addCartItem(prductId) {
  // direct manipulation + event emitter
}

dispatcher.subscribe((action) => {
  if(action.type == "addCartItem") {
    _addCartItem(action.productId);
  }
});
```









+ waitFor is really only useful to coordinate the update orders of multiple stores.

+


ya. the original flux is totally overwrought.


+ action -> dispatcher -> store -> view
  + each stage is a many to many relationship.

+ the idea that event emitter shouldn't immediately update. should wait for coaelce.
  + hmm. not using setImmediate?
    + actually, this is a baobab feature. looking at the events library out there, most doesn't use setImmediate.
  + what's an event library that coaleces callbacks?

i still don't really get waitFor... how else could it be?

+ why not write a function that update two stores?
+ why not let a store listen on another store?
+ why not one store?

What is this "cascading update" boogeyman?


+ We originally set out to deal correctly with derived data

+ MVC — marking a single thread as read would update the thread model, and then also need to update the unread count model.

+ Control is inverted with stores: the stores accept updates and reconcile them as appropriate, rather than depending on something external to update its data in a consistent way.
  + why not functions? ADT

+ Stores have no direct setter methods like setAsRead(), but instead have only a single way of getting new data into their self-contained world: the callback they register with the dispatcher.
  + why not functions? ADT

here's another piece of the puzzle:

+ action creators — dispatcher helper methods — are used to support a semantic API that describes all changes that are possible in the application.
  + why not functions? ADT

I see. So the key to understand this is "semantic API". It's pub-sub at both ends. A dispatcher accepts a whole suite of user actions (semantic actions). This is a product oriented way to think about it. User can DO something, but there's no explicit code to glue user actions to data manipulation.

For example, you might have the original subsystem of stores to handle user actions and update UI. Another team might want to build another subsystem for tracing and debugging. with pubsub they don't need to know about each other.

The "dispatcher" is really just event-emitter + waitFor. If there's no need to enforce update order, you could've used event-emitter. The dispatcher mechanism is almost like AOP.

+ the use for `waitFor` is relatively clear. If thread list and thread unread count are handled by different stores, you'd want thread-list to update before unread count does.
  + which means both stores have to listen to the "open-thread" action.


so... what is the cascading problem? why are they hell-bent on stores not being able to update each other?

+ When updates can only change data within a single round, the system as a whole becomes more predictable.

hmmm... let me just come up with a cascading scenario... multiple boards trello.

update something that updates something else.. i think they just mean at the model level. what is it like?

I mean all their presentations they give the same example:

+ thread list
+ notifications count

They talk about how the unread count had always been buggy. How hard could that possibly be? So weird.

Dispatcher enfroces this in code. While an action is being dispatched, you can't dispatch another.

+ A dispatch causes an application-wide transformation of application state.
  + This is a moment in time, creating a snapshot of change. This is easy to reason about.
  + We cannot dispatch while dispatching and preserve this simplicity. Thus, any corollary change must be driven by the original action.

This is a good explanation of flux: http://stackoverflow.com/a/27267083

ok... I think i get it now...

+ the pub/sub mechanism is for decoupling
+ "no cascading" maintains the simple programming model. one action one change. no chasing the rabbit hole.
+ store is not orm. it's a collection of domain models.
+ action is the data that get passed around. message passing basically.

# redux

hmm. dispatcher is tied to the store. then why the hell do you even use pubsub?
  + i think it's an odd outcome of on reducers are structured. messages travel through the reducer tree.

+ just one store.
+ could have many reducers.
 + oh i see. every action would cause the state to be rebuilt. [combineReducers](http://rackt.github.io/redux/docs/api/combineReducers.html) is a way to "shard" the reducers by keys.

+ ya. recording `actions` as simple messages is pretty cool. can record every changes that ever happen to the UI.

+ i think it's fair to say that redux is flux without the dispatcher pub/sub. one store.
  + the reducer composition technique removes the need for immutable datastructure library.

Dan Abramov writes [The Case For Flux](https://medium.com/@dan_abramov/the-case-for-flux-379b7d1982c6). It argues well for why there should be a single source of data. But I still don't see why the action->store->view update cycle should be flat.



# normalizr

more trouble than it's worth...

https://github.com/gaearon/normalizr

interesting. turns nested JSON structure into normalized "tables", sort of.

# path to flux

+ pub/sub store->view
  + this is easy to understand
+ pub/sub action->store
  + aspect programming. decoupled systems. don't need to change the API
+ no cascading
  + this is kinda a refactoring step. lift all nested store updates (by stores) to the toplevel, so they listen in on actions instead.

+ waitFor
  + ok. I can finally think of a non-artificial example.

### waitFor and multiple stores

Say you want a generic notification count.

There can be many subsystems that produce notifications:

1. new message from friends
2. ads. "new coupon for you!"
3. system notifications. "please refresh browser to update"

the notification widget needs to provide previews to these events, in the order that they are received.

the actions are:

+ system-notification-new
+ system-notification-read

+ message-new
+ message-read

MessageStore subscribes to message-*

NotificationStore subscrbies to message-* and system-notification-*

The MessageAction class doesn't need to change when NotificationStore decides to plugin.

+ waitFor is necessary here because NotificationStore
  + this is NotificationStore needs to access data from MessageStore.
  + waitFor is for read-dependencies only.






