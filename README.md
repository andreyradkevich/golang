I've made each file in `pkg` folder , as `main` , to make it possible to run it independently,
for example for checking closure , you can navigate to closure solder `cd ./pkg/closure` and
execute there `go run .` to run code and see result. For making it possible to import into `main` root package
we just can navigate to `index.go` file in `pkg/closure` folder and rename package on `closure` , and then
in root `main.go` file import this package by `import("golang/pkg/closure")`

`git clean -n -d -x` - clean cached files

`export PATH=$PATH:/usr/local/go/bin`

`go build -gcflags "-m -l"` - check which parts of codes excape to the
heap
