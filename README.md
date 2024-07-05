# Image Filter CLI

This is a small project created to learn Golang by building a command-line tool that applies various filters to png images.

## Usage

```sh
[imgname] -f [filter] {[additional flags]}
```

### Filters

The `-f` flag is used to specify filters to be applied. Filters can be chained together using a comma (`,`) as a separator. Available filters are:

-   `b` : blur
-   `c` : comic
-   `h` : heat
-   `e` : edge
-   `i` : invert
-   `s` : spot
-   `rs` : rowSort

### Flags

-   `-h` : Display flag help.
-   `-i int` : Number of iterations for applying the filter. If not specified, default values for the filter will be used (default: -1).
-   `-o string` : Offset in pixels from the middle of the picture. Usage: `x,y` (default: "0,0").
-   `-p string` : Name of the picture (default: "[name].png").
-   `-r int` : Radius in pixels from the middle of the image (default: -1).
-   `-rp int` : Radius in percent of the smallest side of the image (default: -1).
-   `-t int` : Number of threads to be used for converting. If not specified, uses the maximum available cores of the system (default: -1).

## Examples

Apply a blur filter with default settings:

```sh
[imgname] -f b
```

Apply a heat filter with a radius of 50 pixels:

```sh
[imgname] -f h -r 50
```

Apply an edge filter with 3 iterations and an offset of 10,20 pixels:

```sh
[imgname] -f e -i 3 -o 10,20
```

## License

This project is licensed under the MIT License.
