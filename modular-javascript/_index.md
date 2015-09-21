# JavaScript Bling Bling

In this lesson we are going to use JavaScript to add animation effects. There are we'll implement the followings:

1. Bobbing react logo.
 + use transform.
2. Shaking android robot.
3. Scroll effect.
  + `scrollTo(id)` in its own module.

To begin with we'll just put everything in a single JavaScript file. After we are done, we'll see how we can break it into separate modules, and use build tools to combine these modules into the final file.

We'll learn

+ Animation in browser.
+ How CommonJS module works.
+ Bundling with browserify.

# Performance




# Install GreenSock

```
npm install gsap@1.18.0 --save
```

codepen demo: http://codepen.io/hayeah/pen/VvjMZY

# turn on performance monitor

https://developer.chrome.com/devtools/docs/rendering-settings

the performance tool really is quite interesting. not sure how much I should say tho.


ya ok whatever.


# Monitor Animation Performance

The goal is 60fps. The bottom-line should be 30fps.



# CSS Animation VS JS Animation

A CSS animation library like [animate.css](https://daneden.github.io/animate.css)






hmmm... how do i do this?

it's probably more fun to build the animations and effects before modularizing them. it's also a nice "before & after".

there are actually just 3 effects...




what's the important thing to know?
  + it really isn't super applicable to ReactNative...
  + even the GPU thing doesn't carry over.


I guess the biggest standout feature of GreenSock is that you can control the animation timeline by seeking. map time to scroll distance, that's how scrollmagic works.

# Animiate React Logo

TweenMax

Plugin architecture. Each option is a plugin.

+ use the canonical `css` property first. then mention the shorthand.


Should be:

```
{css: {...}}
```

+ easing curve.

# Animate Robot

TimeLineMax

# Animate Scroll

Need to add the scroll plugin.

+ updating the slider indicator seems like a pain in the ass.
  + probably hook into windows' on scroll handler.

# CSS Performance Tips



if there's no problem, there's no problem. start with the easiest implementation.

1. Animate properties that GPU can accelerate.
  + Apply null transform or will-change.
2. Avoid triggering layout.
  + if you do trigger layout, complex DOM is going to be expensive.
3. Avoid layout thrashing. Make all the changes at the same time so the browser re-calculate layout once.
  + An optimized animation library would handle this. GreenSock
5. Avoid synchronization problem. (see: http://greensock.com/css-performance)
  + Could be a problem if mix CSS and JS animation.



CPU writes a pixel at a time.

GPU writes many pixels at a time.
  + compositing is parallel.



layout -> repaint -> composite


+ <script> is terrible way to manage dependencies.
+ sharing one global namespace, also terrible.
+ npm install. the "main" file.

+ traditionally: closure.

(function(module) {

})(typeof module === "undefined" ? global : module)

# commonjs

# browserify

# Watchify + BrowserSync

+ can you do additive/delta animation with greensock?
  + you can't! ya that's too bad.



# how noconflict works.

# re-export greensock so it doesn't pollute global.

# support <script> and commonjs

http://csstriggers.com/

use requestAnimationFrame to avoid layout thrashing: http://wilsonpage.co.uk/preventing-layout-thrashing/

# very enlightening article of css vs js animation

js animation: http://davidwalsh.name/css-js-animation

>  My recommendation is to use GSAP when you require precise control over timing (e.g. remapping, pause/resume/seek), motion (e.g. bezier curve paths), or complex grouping/sequencing. These features are crucial for game development and certain niche applications, but are less common in web app UI’s.

making js animation fast:

1) optimizing DOM interaction and memory consumption to avoid stuttering,
2) leveraging the principles of requestAnimationFrame under the hood and
3) forcing hardware acceleration (leveraging the power of the GPU to improve animation performance).

> Because transitions aren't natively controlled by JavaScript (they are merely triggered by JavaScript), the browser does not know how to optimize transitions in sync with the JavaScript code that manipulates them.