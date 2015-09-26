# ScrollMagic 滚动效果

我们会通过这一系列的视频来学习如何使用 ScrollMagic 来实现 ILoveReact 的滚动效果。

<video src="http://7xn15n.media1.z0.glb.clouddn.com/scrollmagic-leadin.mp4" controls preload="none"></video>

# ScrollMagic 案例

<video src="http://7xn15n.media1.z0.glb.clouddn.com/scrollmagic-rant.mp4" controls preload="none"></video>

+ 好例子. http://www.pixate.com/
  + 有垂直滚动韵律，方便用户定位当前的位置。
  + 滚动一个屏到下一个屏有圆滑的连贯性，、

+ 坏例子：http://appliancetecltd.com/
  + 失去了垂直的空间感，用户不好定位当前的位置。
  + 滚动页面像走迷宫。

+ 性能问题。
  + https://mixpanel.com/inapp-notifications

+ 返璞归真，不做炫酷滚动效果。
  + 方便快速扫描页面信息，容易直接跳到自己要的地方。
  + http://early-access.notion.so/

终而言之 ScrollMagic 着东西玩玩就好。认真你就输了。

# ScrollMagic 使用机制

<video src="http://7xn15n.media1.z0.glb.clouddn.com/scrollmagic-trigger.mp4" controls preload="none"></video>

触发点 (trigger) 和动画 “长度” (duration)

+ `duration: 0` 会按照动画本身设定的时间去执行动画效果，和滚动轴没有绑定。
+ 非 0 的 `duration` 表示一个滚动距离，把动画的时间轴和滚动轴绑定。
+ `duration: 300` 是滚动 300 素数后完成动画。
+ `duration: 100%` 是滚动屏高 100% 后完成动画。

<video src="http://7xn15n.media1.z0.glb.clouddn.com/controller-and-scene.mp4" controls preload="none"></video>

控制器 (Controller) & 场景 （Scene）:

+ 控制器等于是整个页面的滚动轴，在不同垂直区域会有场景。
+ 滚动到了场景会执行该场景指定的动画效果。
+ 动画效果可以用 GreenSock，Velocity.js，或者是 CSS 来驱动。
+ 若要把动画效果和滚动轴绑定，GreenSock 是唯一的选择。

<video src="http://7xn15n.media1.z0.glb.clouddn.com/pinning.mp4" controls preload="none"></video>

固定 （pinning):

+ 一个场景可以在其滚动区域固定某个元素，停止它的滚动。

# 安装 ScrollMagic

```
npm install scrollmagic@2.0.5 --save
```

# 滚动淡出背景

实现思路:

<video src="http://7xn15n.media1.z0.glb.clouddn.com/background-fading-plan.mp4" controls preload="none"></video>

请自己先试着实现再看示范:

# 滚动飘移 iPhone 效果

实现思路:

<video src="http://7xn15n.media1.z0.glb.clouddn.com/iphone-movement-implementation-plan.mp4" controls preload="none"></video>

请自己先试着实现再看示范:

<video src="http://7xn15n.media1.z0.glb.clouddn.com/iphone-movement-implementation.mp4" controls preload="none"></video>

### 固定 iPhone

<video src="http://7xn15n.media1.z0.glb.clouddn.com/pin-iphone.mp4" controls preload="none"></video>

# 响应式调整

<video src="http://7xn15n.media1.z0.glb.clouddn.com/responsive-demo.mp4" controls preload="none"></video>

+ 没有什么系统性的解决方案。有洞补洞的体力活。
+ 移动端通常需要自己的布局。我们这里就不管了。
