# Buyshoes

+ I think it'd be nice if the lesson could retrace flux's origin. sort of "hey, i could've invented flux too!"



# lesson strategy

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






