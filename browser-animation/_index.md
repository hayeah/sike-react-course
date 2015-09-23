

http://www.html5rocks.com/en/tutorials/speed/high-performance-animations/#toc-composite-properties

+ translateZ(0) adds a layer that GPU can accelerate. memory cost.
  + `backface-visibility: hidden` to fix flickering problem on webkit.

+ CSS animation limitations: http://greensock.com/css-performance
  + css animation thread and js thread might have synchronization issues.
  + js animation turns out to be quite fast.
    + allow timeline. different easing curves for different property animations. different timing. greater control.

+ CSS animation is more declarative and built in the browser. not necessarily more performant tho.
