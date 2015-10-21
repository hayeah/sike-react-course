let loginController = {
  login(name) {
    console.log("hello",name);
  },

  logout(name) {
    console.log("goodbye",name);
  }
}

loginController.login("howard");
loginController.logout("howard");


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


window.appDispatcher = new Dispatcher();

let CartStore = (function() {
  let _cartItems = {};

  function handleAddCartItem(action) {
    let {productId} = action;

    _cartItems[productId] = {
      id: product
    };
  }

  appDispatcher.register(action => {
    if(action.type == "addCartItem") {
      handleAddCartItem(action);
    }
  });

  return {
    getCartItems() {
      return _cartItems;
    }
  };
})();

// Action Creator
function addCartItem(productId) {

}