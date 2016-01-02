# Create React Native Project Quickly

Creating a new ReactNative project can take a long time, because each time the latest ReactNative is re-installed from npm.

We can speed up the process by not downloading ReactNative each time we want to create a new project.

# Install ReactNative Dependencies

Run the following commands to see if your software are up to date.

NodeJS should be higher than 4.0.

```
$ node -v
v4.2.1
```

> `nvm install v4.2`. See: https://github.com/creationix/nvm#usage

```
$ watchman version
{
    "version": "4.1.0"
}
```
> `brew install watchman`

```
$ react-native -v
0.1.7
```

> `npm install -g react-native-cli`

For details see:

http://facebook.github.io/react-native/docs/getting-started.html#requirements

# Create ReactNative Project (Normal)

```
react-native init MyApp
```

It will install the latest ReactNative version.

# Create ReactNative Project (Fast)

Download the ReactNative package (中国镜像):

```
APP_NAME=MyApp
mkdir $APP_NAME && cd $APP_NAME
wget http://7fvhy7.com1.z0.glb.clouddn.com/rn-0.17.0.tgz
mkdir -p node_modules/react-native
tar -zxf rn-0.17.0.tgz -C node_modules/react-native
```

Generate the ReactNative project:

```
$ node -e "require('react-native/cli').init('./','$APP_NAME')"
Setting up new React Native app in ./
To run your app on iOS:
   Open /Users/howard/tmp/MyApp/ios/MyApp.xcodeproj in Xcode
   Hit the Run button
To run your app on Android:
   Have an Android emulator running (quickest way to get started), or a device connected
   cd /Users/howard/tmp/MyApp
   react-native run-android
```

You should have these files in your project:

```
$ ls -a
.                .gitignore       index.android.js node_modules
..               .watchmanconfig  index.ios.js     package.json
.flowconfig      android          ios              rn-0.17.0.tgz
```

Initialize NPM project:

```
$ npm init --yes
{
  "name": "MyApp",
  "version": "1.0.0",
  "main": "index.android.js",
  "dependencies": {
    "react-native": "^0.17.0"
  },
  "devDependencies": {},
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "keywords": [],
  "author": "",
  "license": "ISC",
  "description": ""
}
```

You should now be able to run the packager:

```
$ react-native start
 ┌────────────────────────────────────────────────────────────────────────────┐
 │  Running packager on port 8081.                                            │
 │                                                                            │
 │  Keep this packager running while developing on any JS projects. Feel      │
 │  free to close this tab and run your own packager instance if you          │
 │  prefer.                                                                   │
 │                                                                            │
 │  https://github.com/facebook/react-native                                  │
 │                                                                            │
 └────────────────────────────────────────────────────────────────────────────┘
Looking for JS files in
   /Users/howard/tmp/MyApp

[2:46:29 PM] <START> Building Dependency Graph
[2:46:29 PM] <START> Crawling File System
[2:46:29 PM] <START> Loading bundles layout
[2:46:29 PM] <END>   Loading bundles layout (1ms)
```