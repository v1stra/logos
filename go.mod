module logos

go 1.21.5

// update the path (../n0thing) to point to go-clr (bypasses some static detections)
replace example.com/nothingtoseehere => ../n0thing

require (
	example.com/nothingtoseehere v0.0.0-20200805161622-a9381fbe4fcd // indirect
	github.com/abiosoft/ishell/v2 v2.0.2 // indirect
	github.com/abiosoft/readline v0.0.0-20180607040430-155bce2042db // indirect
	github.com/billgraziano/dpapi v0.5.0 // indirect
	github.com/fatih/color v1.12.0 // indirect
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.3.3 // indirect
)
