<h1 align="center"><code>dataurl</code></h1>
<p align="center">convert given file to data URL</p>

# Installation

```sh
$ go install github.com/abiriadev/dataurl@latest
```

# Usage

## Encode

```sh
$ dataurl ./image.png
data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEUAAACnej3aAAAAAXRSTlMAQObYZgAAAApJREFUCNdjYAAAAAIAAeIhvDMAAAAASUVORK5C
```

## Specify MIME types

```sh
$ dataurl ./post.md
data:text/markdown; charset=utf-8;base64,IyBIZWxsbywgd29ybGQh

$ dataurl --mime 'text/plain; charset=UTF-8' ./post.md
data:text/plain; charset=UTF-8;base64,IyBIZWxsbywgd29ybGQh
```

## Read from standard input

```sh
$ dataurl < ./post.md
data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEUAAACnej3aAAAAAXRSTlMAQObYZgAAAApJREFUCNdjYAAAAAIAAeIhvDMAAAAASUVORK5C

$ dataurl --mime 'text/plain; charset=UTF-8' ./post.md
data:text/plain; charset=UTF-8;base64,IyBIZWxsbywgd29ybGQh
```

# License

[![Licence](https://img.shields.io/github/license/abiriadev/dataurl?style=for-the-badge)](./LICENSE)
