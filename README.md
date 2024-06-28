Small project to learn golang.
Usage: 
[imgname] -f [filter] {[additional flags]}
 -f string
        to be applied Filter(s). Can be chained with ",".
        Possible Values are:
        b : blur
        c : comic
        h : heat
        e : edge
        i : invert
        s : spot

  -h    display flag help
  -i int
        Number of iterations of Applying. If not specified filter default values will be used (default -1)
  -o string
        Offset in Pixel from the middle of picture. Usage: x,y (default "0,0")
  -p string
        Name of the picture (default "[name].png")
  -r int
        Radius in pixel from the middle of image (default -1)
  -rp, int
        Radius in percent of the smallest side of the image (default -1)
  -t int
        Number of threads to be used for converting. If not specified, uses max available Cores of the System. (default -1)
