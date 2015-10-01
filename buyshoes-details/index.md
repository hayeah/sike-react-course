# Buy Shoes Details



Previously we've completed the skeletal layout for the `buyshoes` page. In this lesson, we'll fill in the details.



+ Implement the complete design.
+ Responsive tweaks.



# Design Spec



![](buy-shoes-spec.jpg)



The original Sketch file: [design.sketch](design.sketch)



# Clone Assets



Download the assets from this project from:

https://github.com/hayeah/sikeio-buyshoes-assets

Copy the images into the `img` directory as needed.



# Product Catalogue



Let's start by building the product catalogue. Since the layout is based on percentage, it should adjust itself as we resize the window:

<video src="products-layout-percentage-adjustable.mp4" controls autoplay loop></video>



It looks reasonable as long as there's enough space. Later we'll fix the layout problems when there isn't enough space:

![](products-too-narrow.jpg)



## Product Image



Let's put in the product image. The image should be as big as big as the container.

![](product__image-todo.jpg)



The only tricky part is maintaining the image's aspect ratio. Our default flexbox settings would cause the image to stretch:

![](product-image-stretch.jpg)



Unfortunately the image is only stretched horizontally (the cross-axis). To scale the image while maintaining the aspect ratio, give the product image an explicit width:

![](product-image-aspect-scale.jpg)



Since the image is set to the same width as its container, it doesn't matter whether align-items is center, flex-start, or flex-end.

Note: Why does `stretch` not maintain the aspect ratio of the image? The algorithm calculates the width following (roughly) these steps:



1. Determine the intrinsic size of the original image. Use that to calculate the aspect ratio.
2. Use the width given if specified (e.g. 100%). Otherwise use the intrinsic width.
3. Scale the image's height to maintain aspect ratio.
4. Stretch the image's width to be the same as the container's width.



Because stretch happens last, it ignores the image's aspect ratio. If `stretch` happens before scaling, then the aspect ratio would've been maintained.



### Exercise: Set The Product Image



Start with:

```html
<div class="product">
  <div class="product__display">
    <img class="product__img" src="img/shoe1.jpg"/>
  </div>
</div> <!-- product -->
```



Even though we are using border-box, the margins of flex items are still added outside of the boxes rather than inside. When using percentage layout, `margin` could break the layout by adding extra space so the content no longer fit. In our case, if there's margin, two shoes can't fit in the container:

![](product-broken-into-multiple-lines.jpg)



To avoid this issue, use padding to add space inside the box.



Your result:

![](product-image-set.jpg)



### Exercise: Add To Cart And Price



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



Note: Previously we've added an explicit height to `.product` so it acts as a placeholder. Now it's a good time to remove it.



Your result:

![](product-display-done.jpg)



### Exercise: Product Name & Heart



![](product__description.jpg)



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



It's a left/right layout, where one part should grow and the other should not. Do you remember how to do it?



Change one of the product names to be very long to test your layout is still correct.

```html
<div class="product__name">
  Marana E-Lite Marana E-Lite Marana E-Lite Marana E-Lite Marana E-Lite
</div>
```



Your result:

![](product__description-done.jpg)



# Shopping Cart Items



Now we layout the cart items:

![](cart-item--todo.jpg)



The layout strategy is like this:

![](cart-item--layout.jpg)



+ The image should occupy `33%` of its container's width, so it becomes larger or smaller depending on the size of the screen.
+ The product name and price should grow to occupy the middle part.
+ The trash can should be just big enough.



### Exercise: Cart Items



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



Your result:

<video src="cart-items.mp4" controls></video>



Make sure that the cart items can scroll under the "shopping cart" title:

![](cart-item-scroll-under.jpg)



# Checkout



![](checkout--todo.jpg)



### Exercise: Checkout Component



Nothing new here. This time, write your own HTML using the BEM convention. Start with:

```html
<div class="checkout">
  <hr class="checkout__divider"/>

  <input type="text" class="checkout__coupon-input" placeholder="coupon code"></input>

  <!-- other stuff -->

</div> <!-- checkout -->
```



There are three variants of amount in the checkout component. We can use BEM's modifier convention to name them:

![](checkout__amount--variants.jpg)



+ `checkout__amount`
  + This is the base style.
+ `checkout__price--strikeout`
  + Adds `text-decoration: line-through` to the base style.
+ `checkout__price--saving`
  + Change `font-weight` and `color`.



Your result:

<video src="checkout-done.mp4" controls></video>



# Responsive Tweaks



We are pretty much done. Now that all the major components are in place, we can tweak the layout for different screen sizes.



So... how do we decide what media queries to add for which screen sizes? There are basically two ways:

1. Add tweaks only for specific devices: Desktop, iPad, iPhone.
2. Continually decrease the screen size until something breaks. Fix & repeat.



The second method may seem messy, as you might end up with one tweak for screens that are smaller than 850px, another tweak for 900px or less, and yet another at 935px. In practice though, you might group them together at 900px. At the end, there are usually just 2~3 widths you'd add media queries for. A more nuanced responsive layout might have 6.



So let's start with a reasonably big window size, and start to narrow it down:

<video src="responsive-checks.mp4" controls></video>



+ At 1000px it still looks ok.
+ At 950px it's looking cramped, but still ok.



At 900px it's just at the point of breaking:

![](sidebar-responsive-check-900px.jpg)



At 800px, it's definitely broken:

![](sidebar-responsive-check-800px.jpg)



So we'd want to fix the layout between 900px and 1000px. 950px is probably a good point. Your could choose a different value if you like.



### Exercise: Tweak Sidebar Layout



At 950px, do these tweaks:

+ Increase the size of `sidebar` to `40%` of screen.
+ Decrease the `site__content` to `60%` of screen.
+ Increase the size of `product` to `100%` of container, so there is one product per row.

```css
@media (max-width: 950px) {
  ...
}
```



Your result:

<video src="responsive-layout--950px.mp4" controls></video>



### Exercise: Mobile Layout



At 600px the layout breaks again:

![](sidebar-responsive-check-600px.jpg)



We'll apply these tweaks:

+ Increase right sidebar to `80%` of screen.
+ Hide the left sidebar.
+ Make `site__main` 100%.
+ (Temporarily) Change right sidebar's opacity to 0.3 so we can see what's beneath it.



We'll use JavaScript to hide & show the right sidebar in the next exercise. In this step just adjust the layout.



Your result:

![](responsive-layout-mobile.jpg)



### Exercise: Add The Right Sidebar Hide/Show Toggle



We want to be able to hide & show the checkout sidebar when using mobile layout. Add a sidebar toggle button:

```html
<a class="site__right-sidebar-toggle">
  <img src="img/arrow-icon.svg"/>
</a>
```



We need to remember whether the checkout sidebar is shown in a global state. We could add a class to the body:

```js
var $toggle = document.querySelector(".site__right-sidebar-toggle");
$toggle.addEventListener("click",function() {
  document.body.classList.toggle("js-show-right-sidebar");
});
```



This way it's easy to adjusts the style of the page's components depending on this global state:

```css
.site__right-sidebar-toggle {
  /* styles for when the sidebar is hidden */
}

body.js-show-right-sidebar .site__right-sidebar-toggle {
  /* styles for when the sidebar is shown */
}
```



A few other requirements:

+ The sidebar should always show for non-mobile layouts.
+ The toggle button should be hidden for non-mobile layouts.



Your result:

<video src="sidebar-toggle.mp4" controls></video>



# Design Details



Functionality-wise we are done. We'll add a few transparent overlays to create some fading effects.



There are two base colors:

+ Checkout sidebar background - `#C4CCDA == rgba(192,207,221,1)`
+ Main background - `#F7F8FA == rgba(247,248,250,1)`



You'll need to use linear gradients to adjust the alpha value (the transparency) of these two colors for different overlays.



### Exercise: Shopping Cart Gradient/Transparent Overlays



![](cart-overlays.jpg)



### Exercise: Background



![](background-overlay.jpg)



### Exercise: Transparent Gradient For Right Sidebar



![](sidebar-transparency.jpg)



# Summary



+ Set image's width to 100% to fill the container, while preserving aspect ratio.
+ Use percentage to divide a container into parts.
+ Continuously decrease the size of the window to tweak responsive layout.


