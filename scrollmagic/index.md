
# ScrollMagic

<video src="scrollmagic-leadin.mp4" controls></video>

Just for fun. Probably not even a good idea... Find other ways to impress/attract your users.

+ What is ScrollMagic?

# Examples

<video src="scrollmagic-rant.mp4" controls></video>

+ Good.
  + Scroll rhythm.
  + Continuity.
  + http://www.pixate.com/

+ Bad.
  + No sense of vertical space.
  + Complicated maze.
  + http://appliancetecltd.com/

+ Bad performance.
  + https://mixpanel.com/inapp-notifications

+ Even better, a normal scrollable page!
  + Many people (me) scroll the page REALLY FAST to skim.
  + http://early-access.notion.so/

# Mechanics

<video src="scrollmagic-trigger.mp4" controls></video>

Trigger & Duration:

+ `duration: 0` means trigger the tween, and run in "actual time".
+ Otherwise duration means scroll distance.
  + In pixels or percentage.
  + 100% is the full height of the screen.

<video src="controller-and-scene" controls></video>

Controller & Scene:

+ Scene triggers.
  + Trigger isn't necessarily the animation target.
  + Entering and exiting a scene.
+ GreenSock or Velocity.js or CSS

<video src="pinning.mp4" controls></video>

Pinning:

+ Keep an element at its current position.
+ One pin only.

# Install ScrollMagic

```
npm install scrollmagic@2.0.5 --save
```

# Background Fade

Plan:

<video src="background-fading-plan.mp4" controls></video>

Implementation:

+ Start with a paused Tween, and seek from console.
  + Add an overlay & animate opacity.

# iPhone Pinning

<video src="iphone-movement-implementation-plan.mp4" controls></video>

+ iPhone move.

<video src="iphone-movement-implementation.mp4" controls></video>

+ iPhone pin.

<video src="pin-iphone.mp4" controls></video>

### responsive tweeks

<video src="responsive-demo.mp4" controls></video>