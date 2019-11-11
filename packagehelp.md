# GO Creating Package HELP
$ go help install
usage: go install [-i] [build flags] [packages]

### Install compiles and installs the packages named by the import paths.

### Executables are installed in the directory named by the GOBIN environment
variable, which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH
environment variable is not set. Executables in $GOROOT
are installed in $GOROOT/bin or $GOTOOLDIR instead of $GOBIN.

### When module-aware mode is disabled, other packages are installed in the
directory $GOPATH/pkg/$GOOS_$GOARCH. When module-aware mode is enabled,
other packages are built and cached but not installed.

### The -i flag installs the dependencies of the named packages as well.

For more about the build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build, go get, go clean.

# The Go Environment
$ go env
### set GO111MODULE=
### set GOARCH=amd64
### set GOBIN=
### set GOCACHE=C:\Users\xxxx\AppData\Local\go-build
### set GOENV=C:\Users\xxx\AppData\Roaming\go\env
### set GOEXE=.exe
### set GOFLAGS=
### set GOHOSTARCH=amd64
### set GOHOSTOS=windows
### set GONOPROXY=
### set GONOSUMDB=
### set GOOS=windows
### set GOPATH=C:\Users\xxx\go
### set GOPRIVATE=
### set GOPROXY=https://proxy.golang.org,direct
### set GOROOT=c:\go
### set GOSUMDB=sum.golang.org
### set GOTMPDIR=
set GOTOOLDIR=c:\go\pkg\tool\windows_amd64
### set GCCGO=gccgo
### set AR=ar
### set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
### set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\xxx\AppData\Local\Temp\go-build762227030=/tmp/go-build -gno-record-gcc-switches
