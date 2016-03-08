# Open Source Markdown Extensions

markdown-ast
  + special markdown treatment to give more structure to a markdown document.

markdown-react
  + react-based markdown renderer.
  + a Component plugin mechanism
  + render should receive global data, and Component should be able to access global data.

# Sike Specific

sike-markdown-react
  + sike specific custom components. Used to render pages.
  + `<Exercise>`
  + `<Hint>`
  + `<Video>`
    + Use this to switch CDN later.

sike-web
  + put everything together, be able to render course content.
  + student be able to "checkin".
  + a simple progress tracker.
  + automatic email reminder to come back and do homework.



```
# Build your first page

Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.

<Exercise>
# Install NodeJS

Go to nodejs.org and download it and blah blah blah.
</Exercise>

This is the design spec for the website we'll build

<ZoomableImage src="designs-spec.jpg"/>

<CN>
# 构建你的第一个页面

罗伦益普孙 has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.

chinese translation lorem ipsum

<Exercise>
# 安装 NodeJS

前往 nodejs.org 下载安装文件然后这样那样...
</Exercise>

这是这个页面的设计规范：

<ZoomableImage src="designs-spec.jpg"/>
</CN>
```




# Redesign

+ Use HTML tag as heredoc, to extend markdown.
  + Use uppercase `<Tag>` as heredoc.

When you see `<Foo ...>`, naively search for `</Foo>`, capturing everything in between as heredoc. Also, parse the attributes as string values. Write a simple version first, don't worry about edge cases (e.g. escaping quotation in attributes).

```
<Foo a="1" b="c">
* foo
* bar
</Foo>
```

The node type is:


```
{
  type: "heredoc"
  tag: "ZoomableImage",
  props: {a: "1", b="c"},
  content: "",
}
```

+ Use AST walker to transform and postprocess HTML heredocs.

```
function walk(node,fn) {
  fn(node);

  if(node.children) {
    // recursive walk ...
    for child in children
      walk(child,fn)
  }
}

function transform(node) {
  if(node.type == "foobar") {
    // write the node
    return {
      ...node,
    }
  } else {
    return node;
  }
}

walk(tree,transform);
```


+ Use React to render whatever shit we need.


# TODO

+ edit "why react"


+ Breaking an app into CommonJS modules.
+ Bundling modules into a single file with Browserify.




NOTE: assets must be relative path for GitHub pages to work.

+ add the actual "good" HTML boilerplate to the `init` mission.

# Extra

http://material-ui.com uses the JS StyleSheet approach, also quite interesting.


