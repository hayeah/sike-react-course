# Layout With Flexbox

The original CSS [visual rendering model](http://www.w3.org/TR/WD-CSS2/cover.html#toc) was drafted in 1998, at a time when most pages were documents, and making apps that runs in the browser was a radical idea. The web had evolved multiple times since then, yet we still use the same dated layout engine from 1998. CSS2 is proven technology, there is a huge body of community knowledge, but it's a huge pain.

For example, if you want to vertically center some content, there isn't one way to do it, but many different ways that only work in special circumstances:

![](css-vertical-centering.jpg)

<cn>

# Flexbox 布局

最初的 CSS [可视化渲染模型](http://www.w3.org/TR/WD-CSS2/cover.html#toc) 起草与 1998 年，那时候大部分页面都是文档，而且制作运行在浏览器中的应用是一个激进的想法。Web 从那之后已经进化多次了，但是我们仍旧用着同样陈旧的来自 1998 年的布局引擎。CSS 2 是被公认的技术，社区知识内容有很多（bad translation），这是巨大的痛苦。

举个例子，如果你想垂直居中某些内容，没有一种确定的方法来实现它，但有很多种在特殊情况下工作的方法：

![](css-vertical-centering.jpg)

</cn>

To become a competent frontend developer, you have to learn all kinds of weird, ugly, unnatural tricks to handle different layout needs.

Flexbox, in contrast, was designed specifically for creating UI for the modern web. It is not simple, and just like any other complex layout engine, you'd sometime find it behaving in surprising ways. Almost always, though, there's a straightforward explanation why.

It does take practice to use flexbox well. But compared to traditional layout with CSS, flexbox is so much easier!

<cn>

为了成为一名称职的前端开发者，你需要去学习各种各样古怪、丑陋、不自然的敲门来处理不同的布局需求。

Flexbox，相反，是被特别设计为现代 web 创造 UI 的。它并不简单，而且就像其他复杂的布局系统一样，你有时候会发现它的行为令人惊讶。但是几乎所有情况下，总有一个简单的解释为什么。

需要实践才能把 flexbox 用好。但是相比于传统 CSS 布局方法，flexbox 简单多了！

</cn>

### Our Mission

In this lesson we'll use flexbox to implement the basic layout of our web page:

![](ilove-react-layout-only.jpg)

<cn>

### 我们的任务

在本次课程中我们将用 flexbox 实现 web 页面的基本布局：

![](ilove-react-layout-only.jpg)

</cn>

# Design Spec

If you have Sketch, you can download the original Sketch file:

[ilovereact-plain.sketch](ilovereact-plain.sketch)

If you don't have Sketch, you can download the annotated design:

[![](annotated-layout.jpg)](annotated-layout.jpg)

<cn>

# 设计规范

如果你有 Sketch，你可以下载原始的 Sketch 文件：

[ilovereact-plain.sketch](ilovereact-plain.sketch)

如果你没有 Sketch，你可以下载带有注解的设计：

[![](annotated-layout.jpg)](annotated-layout.jpg)

</cn>

# Download Design Assets

You can download all the design assets from the repo:

[hayeah/iLoveReact-assets](https://github.com/hayeah/iLoveReact-assets)

Add these images to the `img` directory in your project.

<cn>

# 下载设计资源

你可以从这个仓库下载所有的设计资源：

[hayeah/iLoveReact-assets](https://github.com/hayeah/iLoveReact-assets)

把这些图片添加到你项目的 `img` 目录。

</cn>

# Introduction To Flex

[A Complete Guide to Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/) has a good summary of all the CSS properties related to flexbox. It's a bit too much information to get started with.

We'll start with just three properties: `flex-direction`, `align-items`, `justify-content`.

<cn>

# Flex 介绍

[Flexbox 完整指南](https://css-tricks.com/snippets/css/a-guide-to-flexbox/) 里有所有有关 flexbox 的 CSS 属性的总结。这里面的信息稍微有点多（bad translation）。

我们将从三个属性开始：`flex-direction`，`align-items`，`justify-content`。

</cn>

+ `flex-direction` - whether child items are arranged horizontally or vertically.

  ![](flex-direction.jpg)

+ `align-items` and `justify-content` - controls how the children should be arranged in the parent container.

  ![](flex-align-justify.jpg)


+ To center items both horizontally and vertically in a container:

  ![](flex-centering.jpg)
  
<cn>

+ `flex-direction` - 子元素是水平排列还是垂直排列。

  ![](flex-direction.jpg)

+ `align-items` 和 `justify-content` - 控制着子元素如何在父容器中排列。

  ![](flex-align-justify.jpg)


+ 把元素在容器里同时水平和垂直居中：

  ![](flex-centering.jpg)

</cn>


The alignment properties `align-items` and `justify-content` are very easy to mix up.

+ `justify-content` - controls where the items should be on the arrow.
  + This is the "main-axis" of the flex container.
+ `align-items` - controls where the arrow should be in the container.
  + This is the "cross-axis" of the flex container.

<cn>

对齐属性 `align-items` 和 `justify-content` 很容易混在一起使用。

+ `justify-content` - 控制着元素应该放在箭头的什么位置。
  + 这是 flex 容器的“主轴”。
+ `align-items` - 控制着箭头应该放在容器的什么位置。
  + 这是 flex 容器的“横轴”。

</cn>

So `align-items: center` could mean horizontal centering or vertical centering depending on what the flex-direction is. One trick to help us remember is by understanding the `align-self` property.

`align-self` can give a different `align-item` value to a particular item in a flex container.

<cn>

因此 `align-items: center` 意为根据 flex-direction 的值，决定水平居中或者垂直居中。一个帮助我们记住这个方法是理解 `align-self` 属性。

`align-self` 可以对 flex 容器中得一个特殊的元素的 `align-item` 赋不同的值。

</cn>

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

<cn>

```css
.container {
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.red {
  /* 只有一个元素那就是 `flex-start` */
  align-self: flex-start;
}
```

</cn>

What should `align-self` do? `align-self` should not move the item along the flex-direction (the main-axis), otherwise the order of the `.red` element would change:

![](align-self-bad-design.jpg)

<cn>

`align-self` 应该做了些什么？`align-self`不应该把 item 沿着 flex-direction（主轴）移动，否则 `.red` 元素的顺序会改变：

![](align-self-bad-design.jpg)

</cn>

`align-self` should move the item cross-axially so the horizontal order of the row stays the same:

![](flex-align-self.jpg)

Following this reasoning, it's easy to remember which direction `align-self` is. And `align-items` must be in the same direction as `align-self`.

<cn>

`align-self` 应该把元素按照横轴轴的方向移动，因此行的水平顺序保持相同：

![](flex-align-self.jpg)

根据这个原因，很简单就能基础 `align-self` 的方向。而且 `align-items` 和 `align-self` 方向相同。

</cn>

# Page Style

Let's first define the typography and background color of the page:

<cn>

# 页面样式

让我们首先定义页面的文字样式和背景颜色：

</cn>

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

<cn>

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

</cn>

# Page Layout

### Exercise: Full Page Sections

There are four sections in this web page. You should make each section as tall and as wide as the screen.

To make them easier to see, you could temporarily set their background to red:

<cn>

# 页面布局

### 练习：页面的部分

这个网页有四个部分。你应该让每个章节的宽高和屏幕一样。

为了让它们可见度更高，你可以暂时把它们的背景设为红色：

</cn>

```css
.section {
  background-color: rgba(255,0,0,0.3);
  border: 2px solid #FFF;
}
```

<cn>

```css
.section {
  background-color: rgba(255,0,0,0.3);
  border: 2px solid #FFF;
}
```

</cn>

The sections should look like:

<video src="fullpage-sections.mp4" controls></video>

<cn>

这些部分应该看起来像：

<video src="fullpage-sections.mp4" controls></video>

</cn>

### Exercise: Title And Logo

Add the following to the first section:

<cn>

### 练习：标题和 Logo

把下面的代码添加到第一个部分：

</cn>

```html
<img class="react-logo" src="img/react-logo.svg"/>
<h1>Build Native Apps With React</h1>
```

<cn>

```html
<img class="react-logo" src="img/react-logo.svg"/>
<h1>Build Native Apps With React</h1>
```

</cn>

Your result:

![](title-and-logo.jpg)

<cn>

你的结果：

![](title-and-logo.jpg)

</cn>

### Exercise: Navigation Links

Since items in a single flex container can only flow in one direction, you need to nest flex containers if some items are vertically arranged, and some other items are horizontally arranged:

![](flex-nesting.jpg)

<cn>

### 练习：导航链接

因为元素在单一的 flex 容器只能按照一个方向排列，如果一些元素垂直排列，一些其他的元素水平排列，你需要嵌套使用 flex 容器。

![](flex-nesting.jpg)

</cn>

+ Add a horizontal flex container to hold the navigation links.
+ Add padding to space them apart.

Add these internal links:

<cn>

+ 添加一个水平 flex 容器容纳导航链接。
+ 添加 padding 让它们分开。

</cn>

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

<cn>

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

</cn>

Your result:

![](navlinks.jpg)

<cn>

你的结果：

![](navlinks.jpg)

</cn>

# Flex-Grow And Stretch

{ya fuck. flex-grow is unwieldly. it grows to unbounded length if there's too much text content. so stupid.}

Our next goal is to divide the page into two equal parts:

![](left-right-partitions.jpg)

The obvious way to accomplish this is to set the width to 50%, and height to 100%. For this exercise, though, we are going to use flexbox. First, let's add the following html to the second section:

<cn>

# Flex-Grow 和延伸

{呀，我靠。flex-grow 笨死了。如果内容太多，它会变得无限长。太蠢了。}

我们下一个目标是把页面分为两个相等的部分：

![](left-right-partitions.jpg)

完成这项工作的直接方式是把宽度设置为 50%，高度设置为 100%。但是在这个练习中，我们将要使用 flexbox。首先，让我们把下面的 html 添加到第二部分：

</cn>


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

<cn>

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

</cn>

Initially, these two containers are just big enough to contain their text content:

![](flex-grow-none.jpg)

If you remove their content, they'd collapse to a 0x0 box (try it!).

<cn>

一开始，这两个容器就很大，足够容纳它们的文字内容：

![](flex-grow-none.jpg)

如果你想要移除它们的内容，他们会折叠成一个 0x0 的盒子（尝试下！）。

</cn>

There are two properties you can set to make a flexbox bigger than its content:

+ `align-self: stretch` - stretch the element along the cross-axis.
+ `flex-grow: 1` - stretch the element along the main-axis.

<cn>

你可以设置两个属性使 flexbox 比它的内容更大：

+ `align-self: stretch` - 沿着横轴的方向拉伸元素。
+ `flex-grow: 1` - 沿着主轴的方向拉伸元素。

</cn>

Like `align-items` and `justify-content`, their behaviour depends on the flex direction:

![](flex-and-stretch.jpg)

<cn>

正如 `align-items` 和 `justify-content`，它们的行为取决于 flex direction：

![](flex-and-stretch.jpg)

</cn>

Why is `flex-grow` a number? If `flex-grow` is 0, that element doesn't grow. Otherwise, the number is the proportion an element should stretch to fill the available space.

![](flex-grow-factor.jpg)

<cn>

为什么 `flex-grow` 是个数字？如果 `flex-grow` 为 0，那个元素就不会扩大。否则，这个数字就是元素应该拉伸覆盖可用空间的比例。

![](flex-grow-factor.jpg)

</cn>

As a concrete example, we can make the first container take up 1/3 of the space, and the second container take up 2/3 of the space:

<cn>

就具体的例子而言，我们可以让第一个容器占有 1/3 的空间，让第二个容器占有 2/3 的空间：

</cn>

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

<cn>

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

</cn>


### Exercise: Left Right Layout

Use `flex-grow` and `align-self` to divide the second section in two.

You result:

![](left-right-partitions-no-content.jpg)

<cn>

### 练习：左右布局

使用 `flex-grow` 和 `align-self` 把第二部分分成两半。

你的结果：

![](left-right-partitions-no-content.jpg)

</cn>

Note: Make sure these containers are empty. If the amount of content in one container is more than the other, then one container would be bigger than the other.

![](flex-basis-auto.jpg)

We'll fix this problem in the next exercise.

<cn>

注：确保这些容器是空得。如果一个容器内容的数量超过了其他容器，这个容器就会比其他的大。

![](flex-basis-auto.jpg)

我们讲在下个课程中修复这个问题。

</cn>

# Flex Basis

The size of a flexbox is determined by two factors:

1. How much content there is in the flexbox.
2. How much free space there is in the parent container. If `flex-grow` is non-zero, grow to fill the space.

If a flexbox is allowed to grow, it would grow as much as possible to fit the content. Let's add more content to the right container:

<cn>

# Flex 基础

Flexbox 的大小由两个因素决定：

1. Flexbox 中的内容有多少。
2. 父容器中有多少空闲空间。如果 `flex-grow` 为非零值，它会扩大填满这些空间。

</cn>

```html
<div class="feature-description">
  <h2>Native Experience</h2>
  <p>
    Takes advantage of native iOS components to give your app a consistent look and feel with the rest of the platform ecosystem, and keeps the quality bar high.
  </p>
</div>
```

<cn>

```html
<div class="feature-description">
  <h2>Native Experience</h2>
  <p>
    Takes advantage of native iOS components to give your app a consistent look and feel with the rest of the platform ecosystem, and keeps the quality bar high.
  </p>
</div>
```

</cn>

It expands to fit the content in one line, squeezing out the left container:

<video src="flex-basis-auto-greedy.mp4" controls loop></video>

<cn>

它扩张来适配一行的内容，从左容器中挤出去了：

<video src="flex-basis-auto-greedy.mp4" controls loop></video>

</cn>

The key to understand this behaviour is that the amount of "free space" is calculated after growing the containers to fit the content:

![](flex-basis-auto-free-spzce.jpg)

<cn>

理解这种行为的关键是“空闲空间”的数量是在容器适配内容扩张之后才计算的：

![](flex-basis-auto-free-spzce.jpg)

</cn>

Then the free space is divided according to `flew-grow` factors. This explains why the containers are not the same width:

![](flex-basis-auto.jpg)

<cn>

然后空闲空间感觉 `flew-grow` 因素划分。这就解释了为什么容器宽度不同：

![](flex-basis-auto.jpg)

</cn>

We can use the `flex-basis` property to override the size of a flexbox when calculating free space. If we set `flex-basis: 0` for both children, it is as though their widths are zero when their parent calculates the free space. The free space (the entire width of the parent) is then divided between the two children:

![](flex-basis-zero.jpg)

<cn>

当计算空闲空间时，我们可以使用 `flex-basis` 属性覆盖 flexbox 的大小。如果我们给子元素设置了 `flex-basis: 0`，当它们的父容器计算空闲空间时，它们的宽度就好像为零一样。空闲空间（整个父容器的宽度）后来被这两个子元素分割了：

![](flex-basis-zero.jpg)

</cn>

The `flex-basis` property is like the min-width of flexbox. It determines how much space the flexbox reserves for itself. The default `flex-basis: auto` means "reserve as much space as needed to fit the content". And `flex-basis: 50px` means reserve 50px, but grow if there's more free space.

Question: Setting `flex-basis: 50%` also makes the two containers equal. Why? How is it different from `flex-basis: 0`?

<cn>

`flex-basis` 属性就像是 flexbox 的最小值。它决定了 flexbox 给自己保留多少空间。默认的 `flex-basis: auto` 意为“保留尽可能多的空间来适配内容”。`flex-basis: 50px` 意为保留 50 像素，但是如果有更多空闲空间就会扩大。

</cn>

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

<cn>

### 练习：适应 Flex-Basis

添加到第二部分：

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

</cn>

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

<cn>

添加到第三部分：

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

</cn>

# Android Is Coming

### Exercise: Add "Android Is Here"

In the last section, add:

```html
<h1>Android Is Here</h1>
```

Result:

![](android-is-here.jpg)

<cn>

# Android 即将到来

### 练习：添加“Android Is Here”

在最后的部分，添加：

```html
<h1>Android Is Here</h1>
```

结果：

![](android-is-here.jpg)

</cn>

# Flex: 1

In ReactNative you'd often seen the mysterious setting `flex: 1` to grow a flexbox. `flex` is a shorthand that sets `flex-grow`, `flex-shrink`, and `flex-basis` at the same time. The default is:

```css
flex: 0 1 auto;
/*
flex-grow: 0;
flex-shrink: 1;
flex-shrink: auto;
*/
```

<cn>

# Flex: 1

在 ReactNative 里你会经常看到神秘的设定 `flex: 1` 来扩大一个 flexbox。`flex` 是同时设置 `flex-grow`，`flex-shrink` 和 `flex-basis` 的简写。默认为：

```css
flex: 0 1 auto;
/*
flex-grow: 0;
flex-shrink: 1;
flex-shrink: auto;
*/
```

</cn>

And `flex: 1` means `flex: 1 1 auto`, or written in full:

```css
flex-grow: 1;
flex-shrink: 1;
flex-basis: auto;
```

<cn>

`flex: 1` 意为 `flex: 1 1 auto`，或者写为完整形式：

```css
flex-grow: 1;
flex-shrink: 1;
flex-basis: auto;
```

</cn>

# Summary

There are many properties and many possible settings. A good reference to refresh your memory:

+ [CSS Tricks - A Complete Guide to Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/)

If you run into a strange layout problem, it's often quicker to read the spec to understand the exact behaviour:

+ [CSS Flexible Box Layout Module Level 1](http://www.w3.org/TR/css-flexbox-1)

<cn>

# 总结

有很多属性和可能的设定。刷新你记忆的不错参考：

+ [CSS Tricks - A Complete Guide to Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/)

如果你碰到了一个奇怪的布局问题，阅读规范，理解准确的行为通常来得更快：

+ [CSS Flexible Box Layout Module Level 1](http://www.w3.org/TR/css-flexbox-1)

</cn>

Control which axis (horizontal or vertical) items are arranged:

+ `flex-direction: row`, `flex-direction: column`

<cn>

控制使用哪个轴（水平或者垂直）来排列元素：

+ `flex-direction: row`, `flex-direction: column`

</cn>

Control where in the container items are (against one of the edges, or centered):

+ `align-items`, `align-self`, `justify-content`

<cn>

控制元素在容器的什么位置（对着一条边，或者聚众）：

+ `align-items`, `align-self`, `justify-content`

</cn>

Control how items grow or shrink：

+ `flex-grow`, `flex-basis`, `flex-shrink`, `align-self: stretch`

<cn>

控制元素如何扩大或收缩：

+ `flex-grow`, `flex-basis`, `flex-shrink`, `align-self: stretch`

</cn>

Shorthand:

+ `flex: 1` means `flex-grow: 1; flex-shrink: 1; flex-basis: auto;`

<cn>

简写：

+ `flex: 1` 意为 `flex-grow: 1; flex-shrink: 1; flex-basis: auto;`

</cn>