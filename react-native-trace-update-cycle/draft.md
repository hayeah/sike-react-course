# install

+ official npm
+ http://react-native.cn/docs/android-setup.html#content
+ or, template project. rename the app.

+ watchman + react-native-cli

http://7fvhy7.com1.z0.glb.clouddn.com/App-0.16.0-memory-leak-fix.zip

# large layout
First step, define layout. 5:1.5 ratio.

```
var App = React.createClass({
  render: function() {
    return (
      <View style={styles.container}>
        <View style={styles.timerFaceLayout}>
        </View>

        <View style={styles.controlsLayout}>
        </View>
      </View>
    );
  }
});

var styles = StyleSheet.create({
  container: {
    flex: 1,
    alignSelf: 'stretch',
  },

  timerFaceLayout: {
    flex: 5,
    backgroundColor: 'rgba(255,0,0,0.3)',
  },

  controlsLayout: {
    flex: 1.5,
    backgroundColor: 'rgba(0,255,0,0.3)',
  },
});
```

+ rotation should work

# clock-face component

+ getting the layout right is pretty easy.
+ something weird with line-height s.t. labels aren't aligned when 'flex-end' is used. nudge manually with marginBottom.

# add background

+ background
+ generate @2x, @3x PNG. Seems to work for Android as well.
+ remember to set root container's backgroundColor to transparent.

```
<MeasureContainer>
  {layout => this.renderContent(layout)}
</MeasureContainer>
```

+ unfortunately layout doesn't seem to resize image on rotate.
  + fix this bug by creating the MeasureContainer helper component. Wrap the root component with it.
  + make it an exercise at the end...

# controls

+ tintColorno
+ the left control should use top:0,bottom:0 + centering content. absolutely positioned.

# active timer

+ Timer.js
+ zero padding

# touchable

+ TouchableOpacity

# Adjust status bar to light

StatusBarIOS.setStyle("light-content");

# todo

+ status bar tweaks
  + how do i turn android status bar to transparent?

+ memory leak?
  + 0.16 is fucked. https://github.com/facebook/react-native/issues/4381
  + fix: https://github.com/facebook/react-native/commit/ea96a7edb8a17fbcc57db99217783302bb790056
  + latest green: https://travis-ci.org/facebook/react-native/builds/94258375
    + https://github.com/facebook/react-native/commit/b828ae4200c6e2650d1ae608982e945cfb1470ec
    + this is actually 0.12 -.-
  + https://github.com/hayeah/react-native/tree/v0.16-fix-memory
    + my own branch with the memory fix.
    + it seems that with this fix there is still memory leak, albeit much slower.
+ scheduleAnimationFrame also causes memory leak. yay!
  + ok. it seems fine. it's just that lots of garbage is being generated, and GC occurs somewhat slower.g


+ and 0.14 doesn't render the app correctly.
  + no tintColor
  + no borderRadius

+ version market share
  + http://www.umindex.com/devices/android_os
  + http://developer.android.com/about/dashboards/index.html

