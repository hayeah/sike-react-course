# Layout With Flexbox

The original CSS [visual rendering model](http://www.w3.org/TR/WD-CSS2/cover.html#toc) was drafted in 1998, at a time when most pages were documents, and making apps that runs in the browser was a radical idea.

The web had evolved multiple times since then, yet we still use the same dated layout engine from 1998. CSS2 is proven technology, there is a huge body of community knowledge, but it's a huge pain.

For example, if you want to vertically center some content, there isn't one way to do it, but many different ways that only work in special circumstances:

![](css-vertical-centering.jpg)

<Cn>

# Flexbox 布局

最初的 CSS [可视化渲染模型](http://www.w3.org/TR/WD-CSS2/cover.html#toc) 起草于 1998 年。当时页面大多只是文件类型，制作能在浏览器中运行的应用还是一个相当激进的想法。

多年以来 Web 有了多次的进化，但是我们仍然用着来自 1998 年陈旧的的布局引擎。虽然说 CSS2 是被一个通用的技术，社区也有巨大的知识沉淀，但开发起来实在是个巨大的痛苦。

举个例子，如果你想垂直居中些内容，CSS2 并没有一种通用的方法来做这个布局。不同情况有特有的解决方式：

![](css-vertical-centering.jpg)

</Cn>

To become a competent frontend developer, you have to learn all kinds of weird, ugly, unnatural tricks to handle different layout needs.

Flexbox, in contrast, was designed specifically for creating UI for the modern web. It is not simple, and just like any other complex layout engine, you'd sometime find it behaving in surprising ways. Almost always, though, there's a straightforward explanation why.

It does take practice to use flexbox well. But compared to traditional layout with CSS, Flexbox is so much easier!

<Cn>

要成为一名称职的前端开发者，你需要去学习各种各样古怪、丑陋、不自然的小技巧来处理不同的布局需求。

相比之下，Flexbox 是专门为现代 Web UI 设计的布局机制。它并不简单。和其他复杂的布局系统一样，使用的时候你可能会遇到出乎意料的布局结果。但是几乎所有情况下，“奇怪” 的布局行为总有一个简单的解释。

Flexbox 还是需要多上手才能用好。但是相比于传统 CSS 布局方法，Flexbox 简单多了！

</Cn>

### Our Mission

In this lesson we'll use flexbox to implement the basic layout of our page:

![](ilove-react-layout-only.jpg)

<Cn>

### 我们的任务

在本课程我们将用 flexbox 实现页面的基本布局：

![](ilove-react-layout-only.jpg)

</Cn>

# Design Spec

If you have Sketch, you can download the original Sketch file:

[ilovereact-plain.sketch](ilovereact-plain.sketch)

If you don't have Sketch, you can download the annotated design:

![](annotated-layout.jpg)(annotated-layout.jpg)

<Cn>

# 设计规范

如果你有 Sketch，你可以下载原始的 Sketch 文件：

[ilovereact-plain.sketch](ilovereact-plain.sketch)

如果你没有 Sketch，你可以下载带有注解的设计效果图：

![](annotated-layout.jpg)(annotated-layout.jpg)

</Cn>

# Download Design Assets

Download all the design assets from the repo:

[hayeah/iLoveReact-assets](https://github.com/hayeah/iLoveReact-assets)

Add these images to the `img` directory in your project.

<Cn>

# 下载设计资源

从这个仓库下载所有的设计资源：

[hayeah/iLoveReact-assets](https://github.com/hayeah/iLoveReact-assets)

把这些图片添加到你项目的 `img` 目录。

</Cn>

# Introduction To Flex

[A Complete Guide to Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/) has a good summary of all the CSS properties related to flexbox. It's a bit too much information to get started with.

We'll start with just three properties: `flex-direction`, `align-items`, `justify-content`.

<Cn>

# Flex 介绍

[Flexbox 完整指南](https://css-tricks.com/snippets/css/a-guide-to-flexbox/) 里有所有和 Flexbox 的相关 CSS 属性的总结。这篇文章的讲解对于初次接触 Flexbox 来说信息量有点过大，很难一次消化。

因此，我们先选这三个 Flexbox 属性开始介绍：`flex-direction`，`align-items`，`justify-content`。

</Cn>

+ `flex-direction` - whether child items are arranged horizontally or vertically.

  ![](flex-direction.jpg)

+ Centering children both horizontally and vertically in a container:

  ![](flex-centering.jpg)

+ `align-items` and `justify-content` - centering the children in the parent container, or put them against the edge.

  ![](flex-align-justify.jpg)

<Cn>

+ `flex-direction` - 控制子元素是水平排列还是垂直排列。

  ![](flex-direction.jpg)

+ 在容器里同时水平和垂直居中元素：

  ![](flex-centering.jpg)

+ `align-items` 和 `justify-content` - 在父容器中居中元素，或者靠边。

  ![](flex-align-justify.jpg)

</Cn>


Pay special attention to the last example. Observe how `flex-direction` affects the behaviour of align-items and justify-content.

Imagine that flex-direction is an arrow pointing in the layout direction.

+ `justify-content` - controls where the items should be on the arrow.
  + This is the "main-axis" of the flex container.
+ `align-items` - controls where the arrow should be in the container.
  + This is the "cross-axis" of the flex container.

`align-items` and `justify-content` are very easy to mix up.

<Cn>

看看最后这个例子。注意，`flex-direction` 属性影响了 align-items 和 justify-content 的效果。把 flex-direction 想象为一个箭头，指向元素布局的方向。

+ `justify-content` - 控制着元素应该放在箭头上的什么位置。
  + 这是 flex 容器的 “主轴”。
+ `align-items` - 控制着箭头本身应该放在容器的什么位置。
  + 这是 flex 容器的 “横轴”。

`align-items` 和 `justify-content` 着两个属性很容易混淆。

</Cn>

So `align-items: center` could mean horizontal centering or vertical centering depending on what the flex-direction is.

### Align Self

`align-self` can give a different `align-item` value to a particular item in a flex container.

![](flex-align-self.jpg)

The CSS is like:

```css
.container {
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.red {
  /* This is the only item that's `flex-start` */
  align-self: flex-start;
}
```



<Cn>

`align-self` 可以对 Flex 容器中某个指定的元素的赋与不同的 `align-item` 值。

![](flex-align-self.jpg)

CSS 大概这样：

```css
.container {
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.red {
  /* 指定该元素为 `flex-start`  */
  align-self: flex-start;
}
```

</Cn>

# Page Style

Let's first define the typography and background color of the page:

<Cn>

# 页面样式

让我们首先定义页面的文字样式和背景颜色：

</Cn>

```css
body {
  background-color: #1F1E34;
  color: #FFF;
  font-family: "Avenir Next",
      "HelveticaNeue-Light", "Helvetica Neue Light", "Helvetica Neue", Helvetica,
      Arial, "Lucida Grande", sans-serif;
  font-weight: 100;
}

h1 {
  font-size: 64px;
  font-weight: 100;
}

h2 {
  font-size: 48px;
  font-weight: 100;
}

p {
  font-size: 24px;
}

a {
  font-weight: 400;
  color: #FFF;
}

a:hover {
  font-weight: 400;
  color: #DADADA;
  text-decoration: none;
}
```

<Cn>

```css
body {
  background-color: #1F1E34;
  color: #FFF;
  font-family: "Avenir Next",
      "HelveticaNeue-Light", "Helvetica Neue Light", "Helvetica Neue", Helvetica,
      Arial, "Lucida Grande", sans-serif;
  font-weight: 100;
}

h1 {
  font-size: 64px;
  font-weight: 100;
}

h2 {
  font-size: 48px;
  font-weight: 100;
}

p {
  font-size: 24px;
}

a {
  font-weight: 400;
  color: #FFF;
}

a:hover {
  font-weight: 400;
  color: #DADADA;
  text-decoration: none;
}
```

</Cn>

# Page Layout

### Exercise: Full Page Sections

There are four sections in this web page (`<div class="section">`). You should make each section as tall and as wide as the screen.

To make them easier to see while you are debugging, you could temporarily set their background to red:

```css
.section {
  background-color: rgba(255,0,0,0.3);
  border: 2px solid #FFF;
}
```

Hint: The `%` unit is the most conventional way to do this. You could also try the newer vw/vh unit.

+ [Make div 100% height of browser window](http://stackoverflow.com/questions/1575141/make-div-100-height-of-browser-window)
+ [Can I Use: Viewport Units vw/vh](http://caniuse.com/#feat=viewport-units)

The sections should look like:

<video src="fullpage-sections.mp4" controls></video>

<Cn>

# 页面布局

### 练习：页面的章节

这个网页有四个章节 (`<div class="section">`)。你应该让每个章节的宽高和屏幕一样。

在调试时，为了让它们可见度更高，你可以暂时把它们的背景设为红色半透明：

```css
.section {
  background-color: rgba(255,0,0,0.3);
  border: 2px solid #FFF;
}
```

提示：`%` 单位是最普遍使用的全屏方案。你也可以试试新的 vw/vh 单位。

+ [Make div 100% height of browser window](http://stackoverflow.com/questions/1575141/make-div-100-height-of-browser-window)
+ [Can I Use: Viewport Units vw/vh](http://caniuse.com/#feat=viewport-units)

这些章节应该看起来这样：

<video src="fullpage-sections.mp4" controls></video>

</Cn>

### Exercise: Title And Logo

Add the following to the first section:

<Cn>

### 练习：标题和 Logo

把下面的代码添加到第一个区域块：

</Cn>

```html
<img class="react-logo" src="img/react-logo.svg"/>
<h1>Build Native Apps With React</h1>
```

<Cn>

```html
<img class="react-logo" src="img/react-logo.svg"/>
<h1>Build Native Apps With React</h1>
```

</Cn>

Your result:

![](title-and-logo.jpg)

<Cn>

你的结果：

![](title-and-logo.jpg)

</Cn>

### Exercise: Navigation Links

Since items in a single flex container can only flow in one direction, you need to nest flex containers if some items are vertically arranged, and some other items are horizontally arranged:

![](flex-nesting.jpg)

<Cn>

### 练习：导航链接

因为元素在一个 flex 容器内只能按照一个方向来排列，如果你需要一些元素垂直排列，一些其他的元素水平排列，那你就要嵌套 flex 容器。

![](flex-nesting.jpg)

</Cn>

+ Add a horizontal flex container to hold the navigation links.
+ Add padding to space the links apart.

Add these internal links:

<Cn>

+ 添加一个水平 flex 容器容纳导航链接。
+ 添加 padding 为导航链接加上间隔。

</Cn>

```html
<a href="#native">
  Native iOS
</a>

<a href="#touch">
  Touch Handling
</a>

<a href="#async">
  Asynchronous
</a>

<a href="#flex">
  Flex &amp; Styling
</a>
```

<Cn>

```html
<a href="#native">
  Native iOS
</a>

<a href="#touch">
  Touch Handling
</a>

<a href="#async">
  Asynchronous
</a>

<a href="#flex">
  Flex &amp; Styling
</a>
```

</Cn>

Your result:

![](navlinks.jpg)

<Cn>

你的结果：

![](navlinks.jpg)

</Cn>

# Flex-Grow And Stretch

Our next goal is to divide the page into two equal parts:

![](left-right-partitions-no-content.jpg)

The obvious way to accomplish this is to set the width to 50%, and height to 100%. For this exercise, though, we are going to use flexbox. First, let's add the following html to the second section:

<Cn>

# Flex-Grow 如何拉长元素

我们下一个目标是把页面分为两个相等的部分：

![](left-right-partitions-no-content.jpg)

完成这项工作最直观的方式是把宽度设置为 50%，高度设置为 100%。但是为了练习，我们要用 flexbox 来做。首先，让我们把下面的 html 添加到第二个区域块：

</Cn>


```html
<div class="iphone-demo">
  iphone demo
</div>

<div class="feature-description">
  awesome feature description
</div>
```

```css
.iphone-demo {
  background-color: rgba(255,0,0,0.3);
}

.feature-description {
  background-color: rgba(0,255,0,0.3);
}
```

<Cn>

```html
<div class="iphone-demo">
  iphone demo
</div>

<div class="feature-description">
  awesome feature description
</div>
```

```css
.iphone-demo {
  background-color: rgba(255,0,0,0.3);
}

.feature-description {
  background-color: rgba(0,255,0,0.3);
}
```

</Cn>

Initially, these two containers are just big enough to contain their text content:

![](flex-grow-none.jpg)

If you remove their content, they'd collapse to a 0x0 box (try it!).

<Cn>

一开始，这两个容器刚好足够大包围它们的文字内容：

![](flex-grow-none.jpg)

如果移除了它们的内容，这些容器会折叠成一个 0x0 的盒子（尝试下！）。

</Cn>

There are two properties you can set to make a flexbox bigger than its content:

+ `align-self: stretch` - stretch the element along the cross-axis.
+ `flex-grow: 1` - stretch the element along the main-axis.

<Cn>

你可以设置下面这两个属性使 flexbox 比它的内容更大：

+ `align-self: stretch` - 沿着横轴的方向拉伸元素。
+ `flex-grow: 1` - 沿着主轴的方向拉伸元素。

</Cn>

Their behaviour also depends on the flex direction:

![](flex-and-stretch.jpg)

<Cn>

它们的具体行为也受 flex direction 影响：

![](flex-and-stretch.jpg)

</Cn>

Why is `flex-grow` a number? If `flex-grow` is 0, that element doesn't grow. Otherwise, the number is the proportion an element should stretch to fill the available space.



![](flex-grow-factor.jpg)

<Cn>

为什么 `flex-grow` 是个数字呢？如果 `flex-grow` 为 0，那个元素就不会扩大。所有不为 0 的元素，按比例来分配空间，拉长覆盖所有可用的空白区域。

![](flex-grow-factor.jpg)

</Cn>

As a concrete example, we can make the first container take up 1/3 of the space, and the second container take up 2/3 of the space:

<Cn>

具体举个例子，我们可以让第一个容器占有 1/3 的空间，让第二个容器占有 2/3 的空间：

</Cn>

```css
.iphone-demo {
  flex-grow: 1;
  background-color: rgba(255,0,0,0.3);
}

.feature-description {
  flex-grow: 2;
  background-color: rgba(0,255,0,0.3);
}
```

![](flew-grow-1-and-2.jpg)

<Cn>

```css
.iphone-demo {
  flex-grow: 1;
  background-color: rgba(255,0,0,0.3);
}

.feature-description {
  flex-grow: 2;
  background-color: rgba(0,255,0,0.3);
}
```

![](flew-grow-1-and-2.jpg)

</Cn>


### Exercise: Left Right Layout

Use `flex-grow` and `align-self` to divide the second section in two.

You result:

![](left-right-partitions-no-content.jpg)

<Cn>

### 练习：左右布局

使用 `flex-grow` 和 `align-self` 把第二部分分成两半。

你的结果：

![](left-right-partitions-no-content.jpg)

</Cn>

Note: Make sure these containers are empty. If the amount of content in one container is more than the other, then one container would be bigger than the other.

In the screenshot below, the right container is wider than the left container, because the text is a bit longer in the right.

![](flex-basis-auto.jpg)

We'll fix this problem in the next exercise.

<Cn>

注：你得先把着两个容器清空。如果一个容器内容比另外一个容器多，那内容多的容器就会比较宽。

下面这个截图右边的容器比左边宽了一点点，因为右边的文字较多。

![](flex-basis-auto.jpg)

我们会在下个练习修复这个问题。

</Cn>

# Flex Basis

The size of a flexbox is determined by two factors:

1. How much content there is in the flexbox.
2. How much free space there is in the parent container. If `flex-grow` is non-zero, grow to fill the space.

If a flexbox is allowed to grow, it would grow as much as possible to fit the content. Let's add more content to the right container:

```html
<div class="feature-description">
  <h2>Native Experience</h2>
  <p>
    Takes advantage of native iOS components to give your app a consistent look and feel with the rest of the platform ecosystem, and keeps the quality bar high.
  </p>
</div>
```

<Cn>

# Flex Basis 属性

Flexbox 容器的高宽由取决于两个因素：

1. Flexbox 中有多少的内容。
2. 父容器中有多少空闲空间。如果 `flex-grow` 非零，应占满所有空闲空间。

如果你允许一个 flexbox 容器成长 (flex-grow 非零), 那它会无限制地加长，以保证有足够的空间适配内容。我们来试试加上一个很长的字串：

```html
<div class="feature-description">
  <h2>Native Experience</h2>
  <p>
    Takes advantage of native iOS components to give your app a consistent look and feel with the rest of the platform ecosystem, and keeps the quality bar high.
  </p>
</div>
```

</Cn>

It expands to fit the content in one line, squeezing out the left container:

<video src="flex-basis-auto-greedy.mp4" controls loop></video>

<Cn>

容器拉得很长，包围一整行的内容，结果就把左边的容器挤没有了：

<video src="flex-basis-auto-greedy.mp4" controls loop></video>

</Cn>

The key to understand this behaviour is that the amount of "free space" is calculated after growing the containers to fit the content:

![](flex-basis-auto-free-spzce.jpg)

<Cn>

理解这种行为的关键是 “空闲空间” 有多少是在容器适配内容之后才计算的：

![](flex-basis-auto-free-spzce.jpg)

</Cn>

Then the free space is divided according to `flew-grow` factors. This explains why the containers are not the same width:

![](flex-basis-auto.jpg)

<Cn>

然后计算出来的空闲空间会根据 `flew-grow` 因素来瓜分。这就解释了为什么之前两个容器的宽度会不一样：

![](flex-basis-auto.jpg)

</Cn>

We can use the `flex-basis` property to override the size of a flexbox when calculating free space. If we set `flex-basis: 0` for both children, it is as though their widths are zero when their parent calculates the free space. The free space (the entire width of the parent) is then divided between the two children:

![](flex-basis-zero.jpg)

<Cn>

我们可以使用 `flex-basis` 来指定计算空闲空间时 flexbox 应该是什么尺寸。如果我们给两个子元素设置了 `flex-basis: 0`，当它们的父容器计算空闲空间时，它们的宽度就好像为零一样。之后，空闲空间（整个父容器的宽度）才会被这两个子元素瓜分：

![](flex-basis-zero.jpg)

</Cn>

The `flex-basis` property is like the min-width of flexbox. It determines how much space the flexbox reserves for itself. The default `flex-basis: auto` means "reserve as much space as needed to fit the content". And `flex-basis: 50px` means reserve 50px, but grow if there's more free space.

Question: Setting `flex-basis: 50%` also makes the two containers equal. Why? How is it different from `flex-basis: 0`?

<Cn>

`flex-basis` 属性就好像是 flexbox 的最小值。它决定了 flexbox 给自己保留多少空间。默认的 `flex-basis: auto` 意为 “保留尽可能多的空间来适配内容”。`flex-basis: 50px` 意为保留 50 像素，但是如果有更多空闲空间就会扩大。

</Cn>

### Exercise: Adjust Flex-Basis

Add to section 2:

```html
<div class="iphone-demo">
  <img src="img/iphone-frame.svg"/>
</div>

<div class="feature-description">
  <h2>Native Experience</h2>
  <p>
    Takes advantage of native iOS components to give your app
    a consistent look and feel with the rest of the platform ecosystem,
    and keeps the quality bar high.
  </p>
</div>
```

<Cn>

### 练习：适应 Flex-Basis

添加到第二区域块：

```html
<div class="iphone-demo">
  <img src="img/iphone-frame.svg"/>
</div>

<div class="feature-description">
  <h2>Native Experience</h2>
  <p>
    Takes advantage of native iOS components to give your app
    a consistent look and feel with the rest of the platform ecosystem,
    and keeps the quality bar high.
  </p>
</div>
```

</Cn>

Add to section 3:

```html
<div class="iphone-demo">
  <img src="img/iphone-frame.svg"/>
</div>

<div class="feature-description">
  <h2>Touches &amp; Gestures</h2>
  <p>
    React Native implements a powerful responder system to
    negotiate touches in complex view hierarchies. It allows you
    to build complex UI, and handle user interactions precisely.
  </p>
</div>
```

Your result:

![](feature-layout-done.jpg)

<Cn>

添加到第三区域块：

```html
<div class="iphone-demo">
  <img src="img/iphone-frame.svg"/>
</div>

<div class="feature-description">
  <h2>Touches &amp; Gestures</h2>
  <p>
    React Native implements a powerful responder system to
    negotiate touches in complex view hierarchies. It allows you
    to build complex UI, and handle user interactions precisely.
  </p>
</div>
```

你的结果：

![](feature-layout-done.jpg)

</Cn>

# Android Is Coming

### Exercise: Add "Android Is Here"

In the last section, add:

```html
<h1>Android Is Here</h1>
```

Result:

![](android-is-here.jpg)

<Cn>

# Android 即将到来

### 练习：添加 “Android Is Here”

在最后的区域块，添加：

```html
<h1>Android Is Here</h1>
```

结果：

![](android-is-here.jpg)

</Cn>

# Flex: 1

In ReactNative you'd often seen the mysterious setting `flex: 1` to grow a flexbox. `flex` is a shorthand that sets `flex-grow`, `flex-shrink`, and `flex-basis` at the same time. The default is:

```css
flex: 0 1 auto;
/*
flex-grow: 0;
flex-shrink: 1;
flex-basis: auto;
*/
```

<Cn>

# Flex: 1

在使用 ReactNative 时你会经常看到一个神秘的设定 `flex: 1`，用了来扩大一个 flexbox。`flex` 是一个简写，同时设置 `flex-grow`，`flex-shrink` 和 `flex-basis` 三个属性。它们的默认值为：

```css
flex: 0 1 auto;
/*
flex-grow: 0;
flex-shrink: 1;
flex-basis: auto;
*/
```

</Cn>

And `flex: 1` means `flex: 1 1 auto`, or written in full:

```css
flex-grow: 1;
flex-shrink: 1;
flex-basis: auto;
```

<Cn>

`flex: 1` 意为 `flex: 1 1 auto`。全写出来如下：

```css
flex-grow: 1;
flex-shrink: 1;
flex-basis: auto;
```

</Cn>

# Summary

There are many properties and many possible settings. A good reference to refresh your memory:

+ [CSS Tricks - A Complete Guide to Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/)

If you run into a strange layout problem, it's often quicker to read the spec to understand the exact behaviour:

+ [CSS Flexible Box Layout Module Level 1](http://www.w3.org/TR/css-flexbox-1)

<Cn>

# 总结

Flexbox 有很多属性和设定。这个参考资料可以刷新你的记忆：

+ [CSS Tricks - A Complete Guide to Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/)

如果你碰到了一个奇怪的布局问题，阅读规范常常可以更快地了解到底为什么:

+ [CSS Flexible Box Layout Module Level 1](http://www.w3.org/TR/css-flexbox-1)

</Cn>

Control which axis (horizontal or vertical) items are arranged:

+ `flex-direction: row`, `flex-direction: column`

<Cn>

控制使用哪个轴（水平或者垂直）来排列元素的方法：

+ `flex-direction: row`, `flex-direction: column`

</Cn>

Control where in the container items are (against one of the edges, or centered):

+ `align-items`, `align-self`, `justify-content`

<Cn>

控制元素在容器的什么位置（对着一边，或者居中）：

+ `align-items`, `align-self`, `justify-content`

</Cn>

Control how items grow or shrink：

+ `flex-grow`, `flex-basis`, `flex-shrink`, `align-self: stretch`

<Cn>

控制元素如何扩大或收缩：

+ `flex-grow`, `flex-basis`, `flex-shrink`, `align-self: stretch`

</Cn>

Shorthand:

+ `flex: 1` means `flex-grow: 1; flex-shrink: 1; flex-basis: auto;`

<Cn>

简写：

+ `flex: 1` 意为 `flex-grow: 1; flex-shrink: 1; flex-basis: auto;`

</Cn>
