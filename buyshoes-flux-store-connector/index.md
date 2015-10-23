# Liked Products

+ be able to hide and show liked products
+ `getLikedProducts` reader method.
+ embed the products data in ProductStore for now. Later we'll see how to load products using remote api.

# Coupon

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


But connecting a view to a store is a bit verbose:

```js
class FooView extends React.Component {
  componentDidMount() {
    FooStore.addChangeListener(this.forceUpdate.bind(this));
  }

  render() {
    let fooData = FooStore.getFooData();
  }
}
```

In the next lesson we'll create a [JavaScript decorator](https://github.com/wycats/javascript-decorators) to make view-store connection more elegant:

```js
@connectTo(FooStore,"fooData");
class FooView extends React.Component {
  render() {
    let {fooData} = this.props;
  }
}
```

