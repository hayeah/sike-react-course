NOTE: assets must be relative path for GitHub pages to work.

Forget inline, inline-block, inline, float. We'll make EVERYTHING a flexbox. Let's pretend you don't know anything about CSS layout. How do we start from first principle? This is a heretical introduction to CSS layout.

hmm. how do i do this?

The important properties are

The wtf property you'd see a lot in ReactNative:

+ flex: 1

+ align-items
+ align-self
+ justify-content
+ flex-direction: row | column

# Download Assets

# HTML

# splash background

+ set width & height to 100% to fill the screen.
+ set background image

# react logo and title

+ content that has "intrinsic size": img, text, video.

+ element big enough to wrap content.
  + if no content, size is 0x0.
+ elements are lined up one after another along the vertical axis.

```
flex-direction: row;
// horizontally centered
align-items: center;
```

# nav

changing the axis

```
flex-direction: row;
// horizontally centered
justify-content: center;
```

# iphone demo

40% 60% layout

```
// left container
flex-grow: 0.4

// right container
flex-grow: 0.6
```

# android slide

+ parent needs to have "position: relative"
+ absolute positioning
+ top 100%, left: 100%.
+ `translate(50%,50%)`
  + unlike negative margin, we don't need to know the size of the image
+ `rotate(-45deg)`

# draft

+ you basically have all these rectangles nesting in each other.

+ content: img, video and text
+ one element follow another along the "main-axis".
+ div, h1, a
  + big enough to wrap content.
  + if no content, size is 0x0
+ if width is given, that's the width of the element.
+ if height is given, that's the height of the element.

+ if there's less content than can fit the container, justify-content/align-items are used to determine where to put the content.
+ if there's more content than can fit the container, things poke out.
  + width/height might be overriden by the flex-grow and flex-shrink properties. ignore for now.

+ padding is space inside the element
+ margin is space outside the element

# absolute position

goal: position the android robot

+ container does not expand to wrap the positioned element.
+ stacking-order comes in front of other elements.
+ each positioned element (relative/absolute) establishes a coordinate space. The origin (0,0) is the closest