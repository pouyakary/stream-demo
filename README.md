
# Chunks Assembler

The `stream.js` reads the patrick image and slices it into multiple parts, then the the `stream.go` reads those parts written by the `stream.js` in the `parts.txt` (each part is a line) and then rewrites a new image based on the patrick sample as `something.jpg` in the directory. You can run this by:

```
node stream.js
go run stream.js
```