# ASCII Tree Generator
a cli tool to get your project structure in ascii tree format
## build
### Windows
`go build -o gen-ascii-tree.exe cmd/main.go`
### Unix-like
`go build -o gen-ascii-tree cmd/main.go`
## execute
### default options
`gen-ascii-tree`
### include hidden file/directory
`gen-ascii-tree --include-hidden`
### set output file
`gen-ascii-tree --output-file=output.txt`
### replace existing output file
`gen-ascii-tree --output-file=output.txt --replace`
