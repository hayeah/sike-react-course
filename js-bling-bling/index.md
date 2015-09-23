# JavaScript Bling Bling

In this lesson we are going to use JavaScript to add animation effects.

### A Floating React Logo

<video src="react-logo-yoyo-float.mp4" controls autoplay loop></video>

### Shaking Android Robot

<video src="android-shakeit.mp4" controls autoplay loop></video>

### Animated Scrolling

<video src="animated-scroll.mp4" controls autoplay loop></video>

# JavaScript Animation VS CSS Animation

Many frontend developers would tell you that you should avoid JavaScript animation and use CSS animation to get better performance (less CPU time) and smoother effects (higher framerate).

However, the reason that JavaScript animation is slow is usually because the library you use isn't optimized for animation. An optimized JavaScript animation engine (e.g. [GreenSock](https://www.greensock.com) or [Velocity.js](http://velocityjs.org/)) has comparable performance to CSS animation. In some cases JS animation could be faster than CSS animation.

Let's try the [JavaScript Animation Speed Test](https://www.greensock.com/js/speed.html), comparing different libraries to see their performance difference. For my test, I've set the number of particles to 500. You could try more or fewer particles. The demo looks like:

<video src="particles-flyout-demo.mp4" controls loop autoplay></video>

The test results:

+ jQuery: ~10fps.
+ Zepto: ~16fps.
+ TweenJS: ~28fps.
+ GreenSock (GSAP): 50~60fps.

You could get 5~6x performance boost just by using another JavaScript library! For a comparison of JavaScript animation and CSS animation performance, see:

+ [CSS animations performance: the untold story](http://greensock.com/css-performance)

So don't worry about performance issues, it's not JavaScript's fault. We will use GreenSock for this project. Later when we implement scrolling effects, GreenSock makes it super easy to play the animation back and forth:

<video src="scrollmagic-scrubbing.mp4" controls loop></video>

Note: See [Animate.css](https://daneden.github.io/animate.css/) for a simple to use CSS animation library.

# How JavaScript Animation Works

Let's see how we can implement a JavaScript animation that swings an element horizontally:

<video src="sinewave-animation.mp4"></video>

We can use a sine wave to define its position in time:

![](sine-animation-curve.jpg)

Then we use `setTimeout` to schedule a `draw` function, updating the element 60 times a second (or once every 16ms):

```js
var deg360 = 2*Math.PI;
var $box = document.getElementById("box");

// 60 ticks a second. Use it as a counter to calculate the current time.
var tick = 0;
function draw() {
  var second = tick / 60;

  // calculate the current position
  var x = Math.sin(second * deg360) * 100;
  $box.style.left = x + "px";

  tick++;

  // set a timer to draw again in 16ms (60fps)
  setTimeout(draw,1000/60);
}


setTimeout(draw,1000/60);
```

See: [Codepen Demo](http://codepen.io/hayeah/pen/XmKYxr?editors=011)

However, the browser (or iOS for ReactNative) refreshes the screen at a fixed rate, but a `setTimeout` timer may call `draw` at an unpredictable time. Suppose `draw` takes 10ms, it could sometime finish after the the screen refreshes:

![](setTimeout-not-sync.jpg)

To make sure that `draw` always have enough time to complete, we can use `requestAnimationFrame` to schedule a call to `draw` at the same interval as screen refreshes. The timeline looks like this:

![](rAF-sync.jpg)

Invokations of `draw` are now in perfect sync with screen refreshes, they *always* finish before the screen actually refresh.

Rewriting the animation loop with `requestAnimationFrame`:


```js
var deg360 = 2*Math.PI;
var $box = document.getElementById("box");

var start = null;
function draw(time) {
  // `time` is current time in millisecond
  if(!start) {
    start = time;
  }

  var second = (time - start) / 1000;

  // One sine cycle every second.
  var x = Math.sin(second * deg360) * 100;
  $box.style.left = x + "px";

  // Redraw in sync with browser redraw.
  requestAnimationFrame(draw);
}


requestAnimationFrame(draw);
```

See: [Codepen Demo](http://codepen.io/hayeah/pen/QjExJZ?editors=011)

Note: ReactNative also supports `requestAnimationFrame`, which is built with [CADisplayLink](http://www.bigspaceship.com/ios-animation-intervals/).

Note: [Layout thrashing](http://wilsonpage.co.uk/preventing-layout-thrashing/) is another reason why naive JavaScript animation is slow. You should know what it is, but don't worry about it. An optimized JavaScript animation engine would avoid layout thrashing.

# GreenSock

GreenSock is an awesome animation library with a stupid name. Install it:

```
# (GSAP - GreenSock Animation Platform)
npm install gsap@1.18.0 --save
```

For now, we'll use `<script>` tag to load the library. Later we'll learn how to use it as a CommonJS module. The GSAP library is installed at `node_modules/gsap/src/uncompressed/TweenMax.js`. Add to `index.html`:

```js
<script type="text/javascript" src="node_modules/gsap/src/uncompressed/TweenMax.js"></script>
```

From the developer's tool, you should see that TweenMax was loaded, and that it added the global JavaScript object `TweenMax` and `TimelineMax`:

![](TweenMax-loaded.jpg?)

### TweenMax API

You can use TweenMax to animate CSS properties. The three most important methods are `to`, `from` and `fromTo`.We'll use a centered element called `#box` to illustrate these methods.

+ `TweenMax.to(object,duration,options)` - animate properties from stylesheet CSS values to your values.

```js
// Animation the `#box` element for 2 seconds.
TweenMax.to("#box",2,{
  css: {
    // animate multiple CSS properties at the same time
    left: "200px",
    opacity: 0,
  },
});
```

This is useful for animating an element out:

<video src="TweenMaxTo.mp4" autoplay controls loop></video>

[Codepen Demo](http://codepen.io/hayeah/pen/xwOmEj)

+ `TweenMax.from(object,duration,options)` - animate properties from your values to stylesheet CSS values.

```js
// Animation the `#box` element for 2 seconds.
TweenMax.from("#box",2,{
  css: {
    // animate multiple CSS properties at the same time
    left: "-200px",
    opacity: 0,
  },
});
```

This is useful for animating an element in:

<video src="TweenMaxFrom.mp4" autoplay controls loop></video>

[Codepen Demo](http://codepen.io/hayeah/full/NGreMb/)

+ `TweenMax.fromTo(object,duration,optionsFrom,optionsTo)` - animate properties from your starting values your final values.

```js
TweenMax.fromTo("#box",1, {
    // from
    css: {
      left: "-200px",
    }
  },{
    // to
    css: {
      left: "200px",
    },

    // option to repeat animation forever
    repeat: -1,

    // option to reverse the animation and rerun
    yoyo: true,
  }
);
```

<video src="TweenMaxFromToYoyo.mp4"></video>

[Codepen Demo](http://codepen.io/hayeah/full/LpZMBa)

### Easing Functions

There are three ease types:

+ `easeIn` - slow in the beginning, the speeds up until the end.
+ `easeOut` - fast in the beginning, then slows down near the end.
+ `easeInOut` - slow in the beginning, speeds up in the middle, then slows down again.

The [Ease Visualizer](http://greensock.com/ease-visualizer) is a great tool to experiment with the different easing functions that are included in GreenSock.

<video src="gsap-visualizer.mp4" controls></video>

You might notice that the yoyo animation is a bit jerky at the start of the animation (near the left). For a looping animation, easeInOut is a better easing type.

```js
TweenMax.fromTo("#box",1, {
    // from
    css: {
      left: "-200px",
    }
  },{
    // to
    css: {
      left: "200px",
    },

    // option to repeat animation forever
    repeat: -1,

    // option to reverse the animation and rerun
    yoyo: true,

    // change easing type
    ease: Power2.easeInOut,
  }
);
```

<video src="TweenYoyoEaseInOut.mp4" controls></video>

[Codepen Demo](http://codepen.io/hayeah/pen/meEvVE)

### Exercise: Animate the React Logo

+ Create `js/app.js`. Use a `<script>` tag to load it.
+ Modify `make server` to refresh when `js/app.js` changes.
+ Write the `animateLogo` function.
+ Choose a suitable easing function and animation duration that you like.

```js
// Start animating when the page is ready.
window.onload = function() {
  animateLogo();
};
```

Your result:

<video src="react-logo-yoyo-float.mp4" controls></video>

# Rendering Monitor

You can use the Chrome developer tool to see how well your animation is performing.

<video src="rendering-fps-monitor.mp4" controls></video>

You can see that we are running 60fps, as promised.

The "Show paint rectangles" shows you some green rectangles. These are areas where the browser are repainting. We can see three areas where elements are being repainted:

+ The Logo is repainting.
+ When you scroll, the srollbar is repainting.
+ The fixed "slider control" is repainting when we scroll.

Repainting is expensive because the browser is using the CPU to recreate a bitmap for that rectangle. In other words, it's not "hardware accelerated" by the GPU.

For more info about the rendering tab see [DevTools - Rendering Settings](https://developer.chrome.com/devtools/docs/rendering-settings).

# GPU Acceleration

You can think of a web page as a bunch of rectangles. The layout and drawing are done by the CPU:

1. CPU calculate layout of the rectangles. Where are the rectangles? How big are they?
2. CPU render a rectangle as bitmap.

Then if possible, rectangles are sent to the GPU for better performance:

3. CPU uploads the bitmap to GPU as texture.
4. Send instruction to GPU to manipulate the bitmap. Translate/scale/rotation, transparency, etc.

So how is the GPU faster than the GPU? Suppose we want to combine a red bitmap with a green bitmap, the CPU has to do it one pixel at a time:

<video src="CPU-composite.mp4" controls autoplay loop></video>

The GPU can combine all the pixels in parallel:

<video src="GPU-composite.mp4" controls autoplay loop></video>

To enable GPU acceleration, use the following four properties for animation:

![](cheap-operations.jpg)

Basically only CSS3 transform can be accelerated by the GPU. Any of the box model properties (top, left, width, height, padding, margin, border...) would trigger relayout and repaint.

So changing our TweenMax code to use CSS transform:

```js
TweenMax.fromTo("#box",1, {

  css: {
    // Uses CSS3 transform
    x: "-200px",
  }
},{

  css: {
    x: "200px",
  },


  repeat: -1,
  yoyo: true,
  ease: Power2.easeInOut,
}
```

[Codepen Demo](http://codepen.io/hayeah/full/YyGzva/)

TweenMax animates the `transform3d` property to enable GPU acceleration. Using the rendering monitor, you can see that the "green rectangles" are gone, there is no more browser repaint. Futhermore, turn on "show composited layer border", you'd see that the element has an orange border, meaning that it's a 3D layer, accelerated by the GPU:

<video src="show-composited-layer.mp4" controls></video>

In summary, there are three possible costs when changing CSS properties:

1. Relayout (CPU. most expensive).
2. Repaint (CPU).
3. Transform, rotate, scale, opacity (GPU, cheapest).

[CSS Triggers...](http://csstriggers.com/) has a chart of which CSS property triggers what costs. Let's take a look at border-bottom:

![](border-bottom-triggers.jpg)

+ Changing `border-color` triggers repaint. Makes sense, it has to change the color.
+ Changing `border-width` triggers relayout. Makes sense, it could change the size and position of other elements.

To learn more see: [High Performance Animations](http://csstriggers.com/)


### Exercise: Animate The Logo With GPU Acceleration

Your result:

<video src="react-logo-composited.mp4" autoplay loop controls></video>

# Animate Robot

TweenMax animates from a starting point to an end point. To build more complex animations, you can use [TimeLineMax](https://greensock.com/docs/#/HTML5/GSAP/TimelineMax/) to animate from A to B to C to D.

You can specify a different duration for each step. Here's an example:

<video src="TimelineMax.mp4" controls autoplay loop></video>

```js
var t = new TimelineMax();
t.to("#box",1,{x: 200})
  .to("#box",0.5,{rotation: "360deg"})
  .to("#box",1,{y: 100})
  .to("#box",0.5,{rotation: "-=360deg"});
```

[Codepen Demo](http://codepen.io/hayeah/full/YyGzva/)

If you need `repeat` and `yoyo`, pass them as options to the TimelineMax constructor:

```js
// yoyo the timeline animation forever
var t = new TimelineMax({yoyo: true, repeat: -1});
```

[Codepen Demo](http://codepen.io/hayeah/full/OyRVKv)

### Exercise: Shaking Android Robot

```js
function animateRobot() {
}

window.onload = function() {
  // ...
  animateRobot();
}
```

Your result:

<video src="android-shakeit.mp4" controls autoplay loop></video>

# Pro - Slider Control & Animated Scrolling

This section is optional. There are two additional features we'll implement:

1. Update the slider control to reflect the current scroll position.
2. When user click on the slider control to jump to a section, scroll there smoothly.

<video src="animated-scroll.mp4" controls loop autoplay></video>

Note: Use jQuery if you want to, but this is a good chance to practice using the DOM API.

### Exercise: Update slider control on scroll event

Whenever the the window scrolls you should use `window.scrollY` to figure out which section the window is showing.

There are four sections. You should make sure that the section ids and the the `href` property of the slider control links match up:

```html
<div id="intro-section" class="section">
</div>
<div id="native" class="section">
</div>
<div id="touch" class="section">
</div>
<div id="android" class="section">
</div>

<div id="slider-control">
  <a href="#intro-section">
    <div class="dot"></div>
  </a>

  <a href="#native">
    <div class="dot"></div>
  </a>

  <a href="#touch">
    <div class="dot"></div>
  </a>

  <a href="#android">
    <div class="dot"></div>
  </a>
</div>
```

Complete the following code fragment:

```js
function updateSliderControl() {
  // get all the slider links
  var links = document.querySelectorAll("#slider-control a")

  for(var i = 0; i < links.length; i++) {
    var link = links[i];

    // Get the section pointed to by the link
    var section = document.querySelector(...);
    var sectiontTop = ...
    var sectionBottom = ...

    // Check if window.scrollY is between the section.
    if(window.scrollY >= documentTop && window.scrollY < documentBottom) {
      link.className = "active";
    } else {
      link.className = "";
    }
  }
}

// Use the onscroll callback to update slider.
window.onscroll = function() {
  // ...
  updateSliderControl();
}

// Update the slider for the first time when the page is loaded.
window.onload = function() {
  // ...
  updateSliderControl();
};
```

+ Use querySelector and querySelectorAll to get elements by CSS selector.
+ To get the attributes of an element,
  + See: [MDN - Element.attributes](https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes)
  + See: [MDN - NamedNodeMap](https://developer.mozilla.org/en-US/docs/Web/API/NamedNodeMap)
+ To find the top of an element relative to the document, see: http://stackoverflow.com/a/21880020

Your result:

<video src="slider-control-update.mp4" controls></video>

### Exercise: Animated scrolling

When you click on the links in the slider control, the browser jumps immediately to the section targeted by the `href` property of the link. Now we want to use GreenSock to animate smoothly to the targeted section instead.

The [ScrollToPlugin](https://greensock.com/ScrollToPlugin) provides the extra functionality to animate scroll.

Let's include the scroll plugin. It must be loaded after `TweenMax.js`:

```html
<script type="text/javascript" src="node_modules/gsap/src/uncompressed/TweenMax.js"></script>
<script type="text/javascript" src="node_modules/gsap/src/uncompressed/plugins/ScrollToPlugin.js"></script>
```

Complete the code fragment:

```js
function scrollToElement(element) {
  var topOfElement = ...

  TweenMax.to(window,1,{
    scrollTo: {
      y: topOfElement,
    },

    ease: Power2.easeInOut,
  });
}

function addSmoothScrolling() {
  var links = ...

  for(...) {
    var link = links[i];

    link.addEventListener("click",function(event) {
      // `event` is the mouse click event

      // BUG WARNING! Fix with a closure or ES6 `let`.
      var href = link.blahblahblah;

      scrollToElement(...);
    });
  }
}

window.onload = function() {
  // ...
  addSmoothScrolling();
};
```

+ You need to cancel the link's default behaviour.
  + See: [MDN - Event.preventDefault()](https://developer.mozilla.org/en-US/docs/Web/API/Event/preventDefault)
+ The event listener bug is very very common.
  + See: [adding 'click' event listeners in loop](http://stackoverflow.com/questions/8909652/adding-click-event-listeners-in-loop)

Your result:

<video src="animated-scroll.mp4" controls></video>

# Summary

With an optimized animation engine, JavaScript animation can be as fast as CSS animation. Choose whichever that suit your project's needs better.

On ReactNative JavaScript animation is your only choice. Much of the performance tips we've talked about here also apply to ReactNative. The most important are:

+ Avoid triggering layout.
+ Animate properties that the GPU can accelerate.
