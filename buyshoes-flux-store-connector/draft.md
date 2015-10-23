
# Coupon

+ CouponStore
+ connect to ProductStore, CouponStore, CartStore

```
let coupons = {
  "SWEET20": {
    percent: 20
  },

  "GIVEME5": {
    percent: 5
  },
}`
```




In fact, we could rewrite `TimerView` as a function:

```js
let TimerView = (props) => {
  let {currentTime,currentTick} = props;
  return (
    <div>
      <p>Time: {currentTime.toString()}</p>
      <p>Tick: {currentTick}</p>
    </div>
  );
}
```

(This is called a [stateless functional components](https://facebook.github.io/react/blog/2015/10/07/react-v0.14.html#stateless-functional-components), introduced in React 0.14.)