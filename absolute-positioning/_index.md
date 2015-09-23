# Absolute Positioning

<cn>
# 绝对定位
</cn>

The layout system (flexbox or traditional CSS document flow) is good at arranging the items relative to each other, one after another:

![](layout-relative.jpg)

<cn>

布局系统（flex 或者传统的 CSS 文档流）善于排列相对于彼此的元素，或者一个元素在另外一个的后面：

![](layout-relative.jpg)

</cn>

Sometimes, the items are not relative to each other, but relative to their container. You can use absolute positioning to specify where these items should be. For example, put the items at the corners:

![](layout-absolute.jpg)

<cn>

有时，元素并不彼此相关，但是和容器相关。你可以用绝对定位来指定这些元素应该在哪里。举个例子，把元素放在角上：

![](layout-absolute.jpg)

</cn>

In the same container you can have mix items that are positioned by the layout system with items that are absolutely positioned:

![](layout-mixed.jpg)

<cn>

在同一个容器里，你可以把使用布局系统定位的元素，和使用绝对定位的元素混合在一起。

![](layout-mixed.jpg)

</cn>

Absolutely positioned items are like stickers. They might cover over what's below them:

![](layout-stacking-order.jpg)

Finally, you can position items outside of the container:

![](layout-absolute-outside.jpg)

<cn>

绝对定位的元素就像贴纸一样。它们可能覆盖了下面的元素：

![](layout-stacking-order.jpg)

最后，你可以把元素放在容器外面：

![](layout-absolute-outside.jpg)

</cn>


### Your Mission

Embed the app demo in the iPhone frame:

![](iphone-demo.jpg)

<cn>

### 你的任务

把这个 app demo 嵌入到 iPhone 框架里：

![](iphone-demo.jpg)

</cn>

Add a transparent gradient overlay to the background image transitions smoothly to the background color of the page:

![](background-overlay.jpg)

Add the Android robot at the bottom right corner of the last section:

![](android-robot-peek.jpg)

<cn>

给背景图片添加一个透明渐变覆盖层，平滑地过渡到页面背景色：

![](background-overlay.jpg)

把 Android 机器人添加到最后部分的右下角：

![](android-robot-peek.jpg)

</cn>

# Static vs Relative

Usually an absolutely positioned item is relative to its container. Let's position a red box inside a green box:


<cn>

# Static 对比 Relative

一个绝对定位的元素经常与它的容器有关。让我们把一个红盒子放在一个绿盒子的内部：

</cn>

```html
<div class="green">
  <div class="red">
  </div>
</div>
```

```css
.red {
  position: absolute;

  top: 10px;
  left: 10px;
}
```

<cn>

```html
<div class="green">
  <div class="red">
  </div>
</div>
```

```css
.red {
  position: absolute;

  top: 10px;
  left: 10px;
}
```

</cn>

If the green box is `position: relative`, it works as expected:

![](relative-container.jpg)

<cn>

如果绿色盒子是 `position: relative`，它的行为符合预期：

![](relative-container.jpg)

</cn>

If the green box is `position: static`, the red box now ignores the green box as though it's not there:

![](static-container.jpg?)

(Demo: http://codepen.io/hayeah/pen/ZbOzWV?editors=110)

<cn>

如果绿盒子是 `position: static`，红盒子现在忽略掉了绿盒子，它就不在那儿了：

![](static-container.jpg?)

（演示： http://codepen.io/hayeah/pen/ZbOzWV?editors=110）

</cn>

`position: static` is the CSS default, but it's almost never useful. When using absolute positioning, you'd always have to remember to set the parent container to `position: relative`.

ReactNative changes the default to `position: relative` so we don't have to worry about this problem:

<cn>

`position: static` 是 CSS 的默认设定，但是它基本没用。当使用绝对定位的时候，你总应该记住把父容器设为 `positon: relative`。

ReactNative 把默认值改成了 `position: relative`，因此我们不需要担心这种问题：

</cn>

```css
/* ReactNative defaults */
body, div, span {
  box-sizing: border-box;

  /******************/
  position: relative;
  /******************/

  display: flex;
  flex-direction: column;
  align-items: stretch;
  flex-shrink: 0;
  align-content: flex-start;

  border: 0 solid black;
  margin: 0;
  padding: 0;
}
```

<cn>

```css
/* ReactNative 默认值 */
body, div, span {
  box-sizing: border-box;

  /******************/
  position: relative;
  /******************/

  display: flex;
  flex-direction: column;
  align-items: stretch;
  flex-shrink: 0;
  align-content: flex-start;

  border: 0 solid black;
  margin: 0;
  padding: 0;
}
```

</cn>

### Exercise: iPhone Demo

Use absolute positioning to embed the demo image inside the iPhone frame. The measurements are:

<cn>

### 练习：iPhone 演示

使用绝对定位把演示图片嵌入 iPhone 框架框架里。尺寸为：

</cn>

![](annotated-iphone-frame.jpg?)

<cn>

![](annotated-iphone-frame.jpg?)

</cn>

Add to the "Native Experience" section:

```html
<img src="img/tumblr-demo.jpg"/>
```

<cn>

添加到“Native Experience”部分：

```html
<img src="img/tumblr-demo.jpg"/>
```

</cn>

Add to the "Touch & Gestures" section:

```html
<img src="img/swype-demo.jpg"/>
```

<cn>

添加到“Touch & Gestures”部分：

```html
<img src="img/swype-demo.jpg"/>
```

</cn>

Your result:

![](embeded-iphone-demo.jpg)

<cn>

你的结果：

![](embeded-iphone-demo.jpg)

</cn>

# Percentage Positioning

For responsive design, you'd often have to use percentage (%) to position items, so where they are is relative to size of the window/screen.

It's easy to move an item to the center of a container:

![](topleft-50-50.jpg)

<cn>

# 百分比定位

对于响应式设计，你应该经常使用百分比（%）来定位元素，因此它们的位置与窗口/屏幕的大小有关：

![](topleft-50-50.jpg)

</cn>

There is a problem, though... `top, left` moves the top-left corner of the positioned item. More often than not, you'd want the center of the positioned item to be centered. The easiest way to move the center of the item is to use the transform property, and move it by 50% of the item's size:

![](topleft-50-50-translated.jpg)

<cn>

尽管有一个问题... `top, left` 把左上角移动到了被定位的元素里。你往往想把被定位的元素居中。移动元素中心的最简单方法是使用 transform 属性，按元素大小的 50% 移动它：

![](topleft-50-50-translated.jpg)

</cn>

This technique works regardless of the size of the positioned item. To center items along the edge of a container:

![](centered-positions.jpg)

<cn>

无论被定为元素有多大，这种技术都可以工作。把元素定位在容器边缘的中间：

![](centered-positions.jpg)

</cn>

It's confusing that percentage means different things for different CSS properties, but they are mostly what you'd expect:

+ `left, right` - percentage refers to the width of the container.
+ `top, bottom` - percentage refers to the height of the container.
+ `translate(x%,y%)` - percentage refers to the width & height of the transformed item.
+ `padding, margin` - percentage refers to the width of the container.
  + Useful horizontally. Useless vertically.

<cn>

百分比对于不同 CSS 属性的意义有很多，这点令人迷惑。但是这基本是你期望看到的：

+ `left, right` - 参考容器宽度的百分比。
+ `top, bottom` - 参考容器高度的百分比。
+ `translate(x%,y%)` - 参考变换元素宽(x)或高(y)的百分比。
+ `padding, margin` - 参考容器宽度的百分比。
  + 对水平方向有用，垂直方向无用。

</cn>

> Note: It is very very hard to vertically center items with traditional CSS layout techniques, mostly `margin-top` and `margin-bottom` refers to the width of the container, not the height!
>
> Also, if you can't use CSS3 `transform`, then you need to know the width & height of a positioned item in order to offset it with a negative margin...

<cn>

> 注：使用传统 CSS 布局技术来垂直居中元素非常非常难，经常是与容器宽度有关的 `margin-top` 和 `margin-bottom`，而不是与高度有关。
>
> 而且，如果你不会用 CSS3 的 `transform`，你就需要知道被定位元素的宽度和高度的具体值，并根据这些值给被定位元素设置一个负的 margin 值。通过这些操作后，最终才能使被定位元素居中。
</cn>

### Exercise: Android Robot

Position the android robot at the bottom right corner.

```html
<img id="android-robot" src="img/android-robot.svg"/>
```

+ `overflow: hidden` - the container should hide the robot if it's outside.
+ `rotate(-45deg)` - rotate the robot.

Your result:

![](android-robot-peek-done.jpg)

<cn>

### 练习：Android 机器人

把 android 机器人放在右下角.

```html
<img id="android-robot" src="img/android-robot.svg"/>
```

+ `overflow: hidden` - 如果超出边界，容器应该把机器人隐藏。
+ `rotate(-45deg)` - 旋转这个机器人。

你的结果：

![](android-robot-peek-done.jpg)

</cn>

# Fixed Positioning

`position: fixed` is similar to `position: absolute`. The only difference is that a fixed item is relative to the screen/window, not the container. The item stays at the same place even when you scroll. It's perfect for UI, navigation, and menus.

<cn>

# Fixed 定位

`position: fixed`  和 `position: absolute` 的行为相似。唯一的不同点在于 fixed 是相对于浏览器窗口而非父容器进行定位的。由于是相对于浏览器窗口定位，所以上下滚屏也不会改变元素在窗口中显示的位置。这样的性质对于制作 UI 导航栏和菜单栏尤为有用。

</cn>

### Exercise: Slider Control

Add html:

<cn>

### 练习：滚动控制

添加 html：

</cn>

```html
<div id="slider-control">
  <a href="#native" class="active">
    <div class="dot"></div>
  </a>

  <a href="#touch">
    <div class="dot"></div>
  </a>

  <a href="#async">
    <div class="dot"></div>
  </a>

  <a href="#flex">
    <div class="dot"></div>
  </a>
</div>
```

<cn>

```html
<div id="slider-control">
  <a href="#native" class="active">
    <div class="dot"></div>
  </a>

  <a href="#touch">
    <div class="dot"></div>
  </a>

  <a href="#async">
    <div class="dot"></div>
  </a>

  <a href="#flex">
    <div class="dot"></div>
  </a>
</div>
```

</cn>

Some basic stylings:

```css
#slider-control {
  padding: 3px;
}

#slider-control a {
  padding: 3px;
}

#slider-control .dot {
  width: 16px;
  height: 16px;
  border: 1px solid #fff;
  border-radius: 8px;
}

#slider-control .active .dot {
  background-color: #fff;
}
```

<cn>

一些基础的样式：

```css
#slider-control {
  padding: 3px;
}

#slider-control a {
  padding: 3px;
}

#slider-control .dot {
  width: 16px;
  height: 16px;
  border: 1px solid #fff;
  border-radius: 8px;
}

#slider-control .active .dot {
  background-color: #fff;
}
```

</cn>

Your Result:

<video src="fixed-slider-control.mp4" controls autoplay loop></video>

<cn>

你的结果：

<video src="fixed-slider-control.mp4" controls autoplay loop></video>

</cn>

# Overlay

Without this overlay the background transition looks abrupt:

![](bg-without-overlay.jpg)

We need to create a gradient overlay to smoothly transition from the background image to the background color. Like this:

![](gradient-overlay.jpg?)

<cn>

# Overlay

不用 overlay，背景变换看起来很唐突：

![](bg-without-overlay.jpg)

我们需要创建一个渐变 overlay 从背景颜色平滑过渡。就像这样：

![](gradient-overlay.jpg?)

</cn>

There are three layers that should be stacked:

1. Bottom - The background image.
2. Middle - The linear gradient overlay.
3. Top - The content (title, navigation links).

<cn>

有三个图层需要堆在一起：

1. Bottom - 背景图片。
2. Middle - 线性渐变 overlay。
3. Top - 内容（标题，导航链接）。

</cn>

The `img` and `h1` elements are `position: static`. `z-index` doesn't work for static elements, so they are always below absolutely positioned elements:

![](z-order-static.jpg)

You should set the title to `position: relative` to give it a z-index. Now the title would appear above the overlay:

![](z-order-relative.jpg)

<cn>

`img` 和 `h1` 元素都是 `position: static`。`z-index` 对 static 元素不起作用，因此它们经常在绝对定位的元素下面：

![](z-order-static.jpg)

你应该把标题设为 `position: relative` 来给它设定一个 z-index。现在标题会在 overlay 上出现了：

![](z-order-relative.jpg)

</cn>

### Exercise: Intro Section Background

Set the background to an image:

```css
#intro-section {
  background-image: url(../img/background.jpg);
  background-size: cover;
  background-position: center;
}
```

<cn>

### 练习：介绍部分的背景

把背景设为一张图片：

```css
#intro-section {
  background-image: url(../img/background.jpg);
  background-size: cover;
  background-position: center;
}
```

</cn>

Your Result:

![](bg-without-overlay.jpg)

<cn>

你的结果：

![](bg-without-overlay.jpg)

</cn>

### Exercise: Linear Gradient Overlay

The style for the overlay is:

```css
.overlay {
  background-image: linear-gradient(rgba(0,0,0,0),rgba(31,30,52,1));
}
```

<cn>

### 练习：线性渐变 Overlay

Overlay 的样式为：

```css
.overlay {
  background-image: linear-gradient(rgba(0,0,0,0),rgba(31,30,52,1));
}
```

</cn>

Make sure that the content isn't behind the overlay:

![](z-index-problem.jpg)

Your Result:

![](bg-with-overlay.jpg)

<cn>

确保内容不在 overlay 的后面：

![](z-index-problem.jpg)

你的结果：

![](bg-with-overlay.jpg)

</cn>

# Summary

`position: static` is useless, and sometimes evil. It might not be a bad idea to set globally:

```css
* {
  position: relative;
}
```

<cn>

# 总结

`position: static` 没什么用，有时还会起负面作用。把全局设为这个值并不是一个坏主意：

```css
* {
  position: relative;
}
```

</cn>

+ Container should not be `position: static`.
+ User percentage to position items responsively.
+ Use CSS transform's `translate` function to offset an element by percentage of its own size.
+ If an element is `position: static`, it doesn't have z-index.

<cn>

+ 容器不能为 `position: static`。
+ 使用百分比对元素进行响应式定位。
+ 使用 CSS transform 的 `translate` 函数根据它自己大小的百分比来偏移一个元素。
+ 如果一个元素为 `position: static`，它不会有 z-index。

</cn>



