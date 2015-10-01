# Buy Shoes Details

<cn>
# Buy Shoes 页面细节
</cn>

Previously we've completed the skeletal layout for the `buyshoes` page. In this lesson, we'll fill in the details.

<cn>
上一课中，我们已经搭建起来了 `buyshoes` 页面的基本骨架结构。这节课我们将继续完善细节。
</cn>

+ Implement the complete design.
+ Responsive tweaks.

<cn>
+ 完成网页的整个设计。
+ 响应式调整。
</cn>

# Design Spec

<cn>
# 设计细节
</cn>

![](buy-shoes-spec.jpg)

<cn>
![](buy-shoes-spec.jpg)
</cn>

The original Sketch file: [design.sketch](design.sketch)

<cn>
原始的 Sketch 文件：[design.sketch](design.sketch)
</cn>

# Clone Assets

<cn>
# 克隆素材
</cn>

Download the assets from this project from:

https://github.com/hayeah/sikeio-buyshoes-assets

Copy the images into the `img` directory as needed.

<cn>
从以下网址下载你的素材文件：

https://github.com/hayeah/sikeio-buyshoes-assets

复制这些图片到 `img` 文件夹中。
</cn>

# Product Catalogue

<cn>
# 商品目录
</cn>

Let's start by building the product catalogue. Since the layout is based on percentage, it should adjust itself as we resize the window:

<video src="products-layout-percentage-adjustable.mp4" controls autoplay loop></video>

<cn>
现在开始，让我们来一起制作商品目录。既然框架是基于百分比调整的，那么整个页面就应该根据窗口大小调整。

<video src="products-layout-percentage-adjustable.mp4" controls autoplay loop></video>
</cn>

It looks reasonable as long as there's enough space. Later we'll fix the layout problems when there isn't enough space:

![](products-too-narrow.jpg)

<cn>
在窗口留有足够空间的时候它表现得很不错！接下来我们将修复窗口大小不够的情况下的问题：

![](products-too-narrow.jpg)
</cn>

## Product Image

<cn>
## 商品图片
</cn>

Let's put in the product image. The image should be as big as big as the container.

![](product__image-todo.jpg)

<cn>
这时我们把商品图片放进来，它应该和容器同等大小。

![](product__image-todo.jpg)
</cn>

The only tricky part is maintaining the image's aspect ratio. Our default flexbox settings would cause the image to stretch:

![](product-image-stretch.jpg)

<cn>
需要的唯一的技巧就是如何维持图片的宽高比。我们默认的 flexbox 设置会导致图片拉长：

![](product-image-stretch.jpg)
</cn>

Unfortunately the image is only stretched horizontally (the cross-axis). To scale the image while maintaining the aspect ratio, give the product image an explicit width:

![](product-image-aspect-scale.jpg)

<cn>
不幸的是，图片仅仅在水平方向（即 cross-axis）上被拉长了。为了让图片保持宽高比，我们给它一个固定的宽度：

![](product-image-aspect-scale.jpg)
</cn>

Since the image is set to the same width as its container, it doesn't matter whether align-items is center, flex-start, or flex-end.

Note: Why does `stretch` not maintain the aspect ratio of the image? The algorithm calculates the width following (roughly) these steps:

<cn>
现在图片已经和容器有一样的宽度了，那就无所谓 align-items 属性是 center、flex-start，亦或是 flex-end 了。

注意：为什么 `stretch` 无法保持图片的原始比例的？通过下面这个算法逻辑，你将会知道如何计算图片的宽度：
</cn>

1. Determine the intrinsic size of the original image. Use that to calculate the aspect ratio.
2. Use the width given if specified (e.g. 100%). Otherwise use the intrinsic width.
3. Scale the image's height to maintain aspect ratio.
4. Stretch the image's width to be the same as the container's width.

<cn>
1. 确定原始图片本身的大小。使用原始大小计算宽高比。
2. 如果明确给出了宽度（例如 100%）则使用给出的宽度；反之使用原始宽度。
3. 调整图片高度，维持原来的宽高比。
4. 将图片的宽度拉伸为容器宽度。
</cn>

Because stretch happens last, it ignores the image's aspect ratio. If `stretch` happens before scaling, then the aspect ratio would've been maintained.

<cn>
因为 stretch 总是最后执行，它会忽略图片的宽高比。若是 stretch 在缩放前执行，那就能够保持图片的相对比例啦！
</cn>

### Exercise: Set The Product Image

<cn>
### 练习：调整产品图片
</cn>

Start with:

```html
<div class="product">
  <div class="product__display">
    <img class="product__img" src="img/shoe1.jpg"/>
  </div>
</div> <!-- product -->
```

<cn>
一开始的代码：

```html
<div class="product">
  <div class="product__display">
    <img class="product__img" src="img/shoe1.jpg"/>
  </div>
</div> <!-- product -->
```
</cn>

Even though we are using border-box, the margins of flex items are still added outside of the boxes rather than inside. When using percentage layout, `margin` could break the layout by adding extra space so the content no longer fit. In our case, if there's margin, two shoes can't fit in the container:

![](product-broken-into-multiple-lines.jpg)

<cn>
尽管我们在使用 border-box，然而所有 flex 元素的 margin 始终是加在 box 外部而不是内部的。当我们使用百分比布局时，`margin` 可能会通过增加元素的大小而破坏整个布局，以至于元素无法正确的容纳在容器之中。在这个例子中，如果有 margin 的话，你会发现两双鞋是无法占满整个容器的：

![](product-broken-into-multiple-lines.jpg)
</cn>

To avoid this issue, use padding to add space inside the box.

<cn>
为了避免这个问题，我们使用 padding 来扩展 box 内部空间。
</cn>

Your result:

![](product-image-set.jpg)

<cn>
你的结果：

![](product-image-set.jpg)
</cn>

### Exercise: Add To Cart And Price

<cn>
### 练习：添加到购物车和价格
</cn>

Change `product__display` to:

```html
<div class="product__display">
  <img class="product__img" src="img/shoe1.jpg"/>
  <a class="product__add">
    <img class="product__add__icon" src="img/cart-icon.svg"/>
  </a>

  <div class="product__price">
    $299
  </div>
</div>
```

<cn>
调整 `product__display` 的结构至：

```html
<div class="product__display">
  <img class="product__img" src="img/shoe1.jpg"/>
  <a class="product__add">
    <img class="product__add__icon" src="img/cart-icon.svg"/>
  </a>

  <div class="product__price">
    $299
  </div>
</div>
```
</cn>

Note: Previously we've added an explicit height to `.product` so it acts as a placeholder. Now it's a good time to remove it.

<cn>
注意：之前我们给 `.product` 添加了一个固定的高度，作为商品的占位用途。现在是时候把这个属性去掉了。
</cn>

Your result:

![](product-display-done.jpg)

<cn>
你的结果：

![](product-display-done.jpg)
</cn>

### Exercise: Product Name & Heart

<cn>
### 练习：商品名称和喜欢按钮
</cn>

![](product__description.jpg)

<cn>
![](product__description.jpg)
</cn>

Change `.product` to:

```html
<div class="product">
  <div class="product__display">
    <img class="product__img" src="img/shoe1.jpg"/>
    <a class="product__add">
      <img class="product__add__icon" src="img/cart-icon.svg"/>
    </a>

    <div class="product__price">
      $299
    </div>
  </div>

  <div class="product__description">
    <div class="product__name">
      Marana E-Lite
    </div>

    <img class="product__heart" src="img/heart.svg"/>
  </div>
</div> <!-- product -->
```

<cn>
更改 `.product` 结构：

```html
<div class="product">
  <div class="product__display">
    <img class="product__img" src="img/shoe1.jpg"/>
    <a class="product__add">
      <img class="product__add__icon" src="img/cart-icon.svg"/>
    </a>

    <div class="product__price">
      $299
    </div>
  </div>

  <div class="product__description">
    <div class="product__name">
      Marana E-Lite
    </div>

    <img class="product__heart" src="img/heart.svg"/>
  </div>
</div> <!-- product -->
```
</cn>

It's a left/right layout, where one part should grow and the other should not. Do you remember how to do it?

<cn>
这是应该一个左右的分栏结构。其中一部分需要扩展空间而另一部分则不需要。还记得如何实现吗？
</cn>

Change one of the product names to be very long to test your layout is still correct.

```html
<div class="product__name">
  Marana E-Lite Marana E-Lite Marana E-Lite Marana E-Lite Marana E-Lite
</div>
```

<cn>
当商品名称很长的时候，测试你的布局是否正确。

```html
<div class="product__name">
  Marana E-Lite Marana E-Lite Marana E-Lite Marana E-Lite Marana E-Lite
</div>
```
</cn>

Your result:

![](product__description-done.jpg)

<cn>
你的结果：

![](product__description-done.jpg)
</cn>

# Shopping Cart Items

<cn>
# 购物车中的物品
</cn>

Now we layout the cart items:

![](cart-item--todo.jpg)

<cn>
现在我们来给购物车中的物品布局：

![](cart-item--todo.jpg)
</cn>

The layout strategy is like this:

![](cart-item--layout.jpg)

<cn>
布局的方式类似这样：

![](cart-item--layout.jpg)
</cn>

+ The image should occupy `33%` of its container's width, so it becomes larger or smaller depending on the size of the screen.
+ The product name and price should grow to occupy the middle part.
+ The trash can should be just big enough.

<cn>
+ 图片需要占据容器宽度的 33%，能够根据屏幕自适应大小。
+ 商品名称和价格需要占据中间的位置。
+ 移入回收站的按钮需要足够大。
</cn>

### Exercise: Cart Items

<cn>
### 练习：购物车中的物品
</cn>

HTML:

```html
<div class="cart-item">
  <div class="cart-item__top-part">
    <div class="cart-item__image">
      <img  src="img/shoe1.jpg"/>
    </div>

    <div class="cart-item__top-part__middle">
      <div class="cart-item__title">
        Marana E-Lite
      </div>

      <div class="cart-item__price">
        $299
      </div>
    </div>

    <img class="cart-item__trash" src="img/trash-icon.svg"/>
  </div> <!-- cart-item__top-part -->


  <div class="cart-item__qty">
    <div class="adjust-qty">
      <a class="adjust-qty__button">-</a>
      <div class="adjust-qty__number">2</div>
      <a class="adjust-qty__button">+</a>
    </div>
  </div>

</div> <!-- cart-item -->
```

<cn>
HTML:

```html
<div class="cart-item">
  <div class="cart-item__top-part">
    <div class="cart-item__image">
      <img  src="img/shoe1.jpg"/>
    </div>

    <div class="cart-item__top-part__middle">
      <div class="cart-item__title">
        Marana E-Lite
      </div>

      <div class="cart-item__price">
        $299
      </div>
    </div>

    <img class="cart-item__trash" src="img/trash-icon.svg"/>
  </div> <!-- cart-item__top-part -->


  <div class="cart-item__qty">
    <div class="adjust-qty">
      <a class="adjust-qty__button">-</a>
      <div class="adjust-qty__number">2</div>
      <a class="adjust-qty__button">+</a>
    </div>
  </div>

</div> <!-- cart-item -->
```
</cn>

Your result:

<video src="cart-items.mp4" controls></video>

<cn>
你的结果：

<video src="cart-items.mp4" controls></video>
</cn>

Make sure that the cart items can scroll under the "shopping cart" title:

![](cart-item-scroll-under.jpg)

<cn>
确保购物车中的物品是在购物车标题下滚动的：

![](cart-item-scroll-under.jpg)
</cn>

# Checkout

<cn>
# 结算
</cn>

![](checkout--todo.jpg)

<cn>
![](checkout--todo.jpg)
</cn>

### Exercise: Checkout Component

<cn>
### 练习：结算元素
</cn>

Nothing new here. This time, write your own HTML using the BEM convention. Start with:

```html
<div class="checkout">
  <hr class="checkout__divider"/>

  <input type="text" class="checkout__coupon-input" placeholder="coupon code"></input>

  <!-- other stuff -->

</div> <!-- checkout -->
```

<cn>
这边依旧没有什么新内容。现在需要你来用 BEM 命名约定来写 HTML 了。一开始的模板：


```html
<div class="checkout">
  <hr class="checkout__divider"/>

  <input type="text" class="checkout__coupon-input" placeholder="coupon code"></input>

  <!-- 其他的内容 -->

</div> <!-- checkout -->
```

</cn>

There are three variants of amount in the checkout component. We can use BEM's modifier convention to name them:

![](checkout__amount--variants.jpg)

<cn>
对于结算元素我们有三种不同的样式。通过 BEM 修饰符来命名他们。

![](checkout__amount--variants.jpg)
</cn>

+ `checkout__amount`
  + This is the base style.
+ `checkout__price--strikeout`
  + Adds `text-decoration: line-through` to the base style.
+ `checkout__price--saving`
  + Change `font-weight` and `color`.

<cn>
+ `checkout__amount`
  + 基础样式。
+ `checkout__price--strikeout`
  + 对基础样式添加 `text-decoration: line-through`。
+ `checkout__price--saving`
  + 更改 `font-weight` 和 `color` 属性。
</cn>

Your result:

<video src="checkout-done.mp4" controls></video>

<cn>
你的结果：

<video src="checkout-done.mp4" controls></video>
</cn>

# Responsive Tweaks

<cn>
# 响应式调整
</cn>

We are pretty much done. Now that all the major components are in place, we can tweak the layout for different screen sizes.

<cn>
我们离完成已经不远了。既然现在主要的部件已经在它们正确的位置上了，我们就能对不同大小的屏幕做一些布局上的调整了。
</cn>

So... how do we decide what media queries to add for which screen sizes? There are basically two ways:

1. Add tweaks only for specific devices: Desktop, iPad, iPhone.
2. Continually decrease the screen size until something breaks. Fix & repeat.

<cn>
唔，我们如何决定在不同屏幕大小上使用什么样的 Media 查询呢？下面介绍两种基本方法：

1. 只对特定的设备添加调整：例如 Desktop，iPad 还有 iPhone。
2. 持续性减小屏幕大小直到某些元素断行。不断重复修改调整。
</cn>

The second method may seem messy, as you might end up with one tweak for screens that are smaller than 850px, another tweak for 900px or less, and yet another at 935px. In practice though, you might group them together at 900px. At the end, there are usually just 2~3 widths you'd add media queries for. A more nuanced responsive layout might have 6.

<cn>
第二种方法看上去很烦人，你可能需要给小于等于 850px 的屏幕做一次调整，又要给小于等于 900px 的屏幕做一次调整，还要给小于等于 935px 的屏幕做另一次调整。在我们的练习中，你可以把这三种情况看做在基于 900px 屏幕的上的调整。最后一般只要增加两到三个 Media 查询就能完成不同屏幕的调整。如果需要细枝末节上的更改，可能需要 6 个左右的 Media 查询。
</cn>

So let's start with a reasonably big window size, and start to narrow it down:

<video src="responsive-checks.mp4" controls></video>

<cn>
从一个相当大的屏幕开始，逐渐把屏幕变小：

<video src="responsive-checks.mp4" controls></video>
</cn>

+ At 1000px it still looks ok.
+ At 950px it's looking cramped, but still ok.

<cn>
+ 1000px 的时候看上去还行。
+ 950px 的时候感觉到有一些限制，但依旧可以正常显示。
</cn>

At 900px it's just at the point of breaking:

![](sidebar-responsive-check-900px.jpg)

<cn>
900px 已经到了无法正常显示的边缘：

![](sidebar-responsive-check-900px.jpg)
</cn>

At 800px, it's definitely broken:

![](sidebar-responsive-check-800px.jpg)

<cn>

800px 的时候，整个布局就彻底无效了：

![](sidebar-responsive-check-800px.jpg)
</cn>

So we'd want to fix the layout between 900px and 1000px. 950px is probably a good point. Your could choose a different value if you like.

<cn>
所以我们选择在 900px 和 1000px 中间进行修复布局的工作。950px 看上去是一个不错的选择。如果你喜欢的话也可以换一个不同的值。
</cn>

### Exercise: Tweak Sidebar Layout

<cn>
### 练习：调整侧边栏布局
</cn>

At 950px, do these tweaks:

+ Increase the size of `sidebar` to `40%` of screen.
+ Decrease the `site__content` to `60%` of screen.
+ Increase the size of `product` to `100%` of container, so there is one product per row.

```css
@media (max-width: 950px) {
  ...
}
```

<cn>
在 950px 的位置，调整下面的部分：

+ 将 `sidebar` 增加到 `40%` 的屏幕大小宽度。
+ 减少 `site__content` 至 `60%` 的屏幕大小宽度。
+ 调整 `product` 大小至 `100%` 的容器大小，使得一行只能有一件商品。

```css
@media (max-width: 950px) {
  ...
}
```
</cn>

Your result:

<video src="responsive-layout--950px.mp4" controls></video>

<cn>
你的结果：

<video src="responsive-layout--950px.mp4" controls></video>
</cn>

### Exercise: Mobile Layout

<cn>
### 练习：移动端布局
</cn>

At 600px the layout breaks again:

![](sidebar-responsive-check-600px.jpg)

<cn>
600px 的时候，布局再一次被打乱：

![](sidebar-responsive-check-600px.jpg)
</cn>

We'll apply these tweaks:

+ Increase right sidebar to `80%` of screen.
+ Hide the left sidebar.
+ Make `site__main` 100%.
+ (Temporarily) Change right sidebar's opacity to 0.3 so we can see what's beneath it.

<cn>
我们来应用下面这些调整：

+ 右侧侧边栏调整为 `80%` 的屏幕宽度。
+ 隐藏左侧侧边栏。
+ 调整 `site__main` 至 100%。
+ （临时性地）将右侧侧边栏的不透明度调整为 0.3。这样我们能看到它的下面有什么内容。
</cn>

We'll use JavaScript to hide & show the right sidebar in the next exercise. In this step just adjust the layout.

<cn>
下一个练习我们将使用 JavaScript 来显示和隐藏右侧的侧边栏，这一步我们先调整它的布局。
</cn>

Your result:

![](responsive-layout-mobile.jpg)

<cn>
你的结果：

![](responsive-layout-mobile.jpg)
</cn>

### Exercise: Add The Right Sidebar Hide/Show Toggle

<cn>
### 练习：增加右侧侧边栏的显示和隐藏的切换
</cn>

We want to be able to hide & show the checkout sidebar when using mobile layout. Add a sidebar toggle button:

```html
<a class="site__right-sidebar-toggle">
  <img src="img/arrow-icon.svg"/>
</a>
```

<cn>
在移动端布局时，我们想要能够显示和隐藏右侧的侧边栏。给侧边栏加上一个切换的按钮：

```html
<a class="site__right-sidebar-toggle">
  <img src="img/arrow-icon.svg"/>
</a>
```
</cn>

We need to remember whether the checkout sidebar is shown in a global state. We could add a class to the body:

```js
var $toggle = document.querySelector(".site__right-sidebar-toggle");
$toggle.addEventListener("click",function() {
  document.body.classList.toggle("js-show-right-sidebar");
});
```

<cn>
结算的侧边栏显示状态应该是个全局的状态。我们可以给 body 部分增加一个 class：

```js
var $toggle = document.querySelector(".site__right-sidebar-toggle");
$toggle.addEventListener("click",function() {
  document.body.classList.toggle("js-show-right-sidebar");
});
```
</cn>

This way it's easy to adjusts the style of the page's components depending on this global state:

```css
.site__right-sidebar-toggle {
  /* styles for when the sidebar is hidden */
}

body.js-show-right-sidebar .site__right-sidebar-toggle {
  /* styles for when the sidebar is shown */
}
```

<cn>
基于这个全局的状态，我们能很方便的调节页面的样式：

```css
.site__right-sidebar-toggle {
  /* 侧边栏隐藏时的样式 */
}

body.js-show-right-sidebar .site__right-sidebar-toggle {
  /* 侧边栏显示时的样式 */
}
```
</cn>

A few other requirements:

+ The sidebar should always show for non-mobile layouts.
+ The toggle button should be hidden for non-mobile layouts.

<cn>
一些其他的要求：

+ 侧边栏在非移动端需要一直显示。
+ 侧边栏切换显示和隐藏的按钮在非移动端需要隐藏。
</cn>

Your result:

<video src="sidebar-toggle.mp4" controls></video>

<cn>
你的结果：

<video src="sidebar-toggle.mp4" controls></video>
</cn>

# Design Details

<cn>
# 设计细节
</cn>

Functionality-wise we are done. We'll add a few transparent overlays to create some fading effects.

<cn>
功能上我们已经完成了。我们接下来添加几个透明遮罩来制作一些渐出的效果。
</cn>

There are two base colors:

+ Checkout sidebar background - `#C4CCDA == rgba(192,207,221,1)`
+ Main background - `#F7F8FA == rgba(247,248,250,1)`

<cn>
两个基本颜色：

+ 结算侧边栏的背景 - `#C4CCDA == rgba(192,207,221,1)`
+ 整体的背景 - `#F7F8FA == rgba(247,248,250,1)`
</cn>

You'll need to use linear gradients to adjust the alpha value (the transparency) of these two colors for different overlays.

<cn>
你需要对不同的遮罩使用使用线性渐变来调整 Alpha 值（即不透明度）。
</cn>

### Exercise: Shopping Cart Gradient/Transparent Overlays

<cn>
### 练习：购物车线性渐变/透明遮罩
</cn>

![](cart-overlays.jpg)

<cn>
![](cart-overlays.jpg)
</cn>

### Exercise: Background

<cn>
### 练习：背景
</cn>

![](background-overlay.jpg)

<cn>
![](background-overlay.jpg)
</cn>

### Exercise: Transparent Gradient For Right Sidebar

<cn>
### 练习：为右侧侧边栏加上透明渐变
</cn>

![](sidebar-transparency.jpg)

<cn>
![](sidebar-transparency.jpg)
</cn>

# Summary

<cn>
# 总结
</cn>

+ Set image's width to 100% to fill the container, while preserving aspect ratio.
+ Use percentage to divide a container into parts.
+ Continuously decrease the size of the window to tweak responsive layout.

<cn>
+ 将图片的宽度设置为 100% 来填充容器，同时也能保持它的宽高比。
+ 使用百分比将容器划分为多个部分。
+ 通过不断减小窗口大小来调整响应式布局。
</cn>
