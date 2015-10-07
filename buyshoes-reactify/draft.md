{
  fuck i want to go home. I can't get anything done here.

  this is so terrible. I start every morning on the wrong foot. every single meal wrecks havoc on my digestive tract.

  noise. pollution. how does anyone ever get anything done?? or do they even get anything done.

  wtf.

  god how do i pull through.

}

# Outline

React.createElement("h1",{},"hello world")

+ React Components
  + React.Element are values. Maybe it's kinda fun to build them by hand.
  + The most important idea is probably `shoes.map((shoe) => <Product product={shoe}>)`
    + If you get that, you are fine...

+ Introduce JSX. How it's just plain JavaScript.
  + <script> to React.
  + Play around in browser.
  + Shadow DOM and Reconciliation.

+ modularity. common js
  + first approximation: separate <script> tags.
  + second step: commonjs & browserify bundling.

+ React Component
  + One step up. Add logic to views.

  + <Checkout>, <Products>, <Cart>, <SiteTitle>

+ Store. Read-Only data.

day 3 - componentize
  + data.
  + interface with DOM. perfect-scroll.
  + state. filter for hearted products.
    + pass a heart-filter setter down to the sidebar as prop.
      + pass the heart filter down to products as prop.
    + up-down cross-component communication is admittedly clumsy. will refactor with flux.

+ done.

+ flux

+ immutability, history, and undo



+ Useful ES6 Features - Part 1
  + `let` everywhere
  + `=>` everywhere
  + the object function notation: `render() { ... }`
  + `let {a,b,c} = obj`
