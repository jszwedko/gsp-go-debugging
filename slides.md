# Continous Delivery:
# Reliable software development through automation

## Jesse Szwedko
## j.szwedko@modcloth.com
## @jesse_szwedko


### Location: ModCloth, Inc. (Pittsburgh Office)
### Date: 7/31/2014
### This talk: [github.com/jszwedko/gsp-go-debugging](github.com/jszwedko/gsp-go-debugging) (slides are on gh-pages branch)

# ![modcloth](modcloth.png)

---

# Debugging
- Tracing
- Interactive debugger
- Post-mortem debugging
- Remote debugger

# Presenter Notes
- Any other debugging styles?

---

# Go Debugging
- Tracing
  - `log`
  - `logrus`
  - Other loggers
- Interactive debugger
  - `gdb`
  - LiteIDE
- Post-mortem debugging
  - `GOTRACEBACK`
- Remote debugger
  - None that I'm aware of

---

# Factorial web service

---

# Tracing
- Add log statements
- [`log`](http://golang.org/pkg/log)
  - Offers much of what you need from a logger
  - Contains a global logger to get started quickly
- [`logrus`](github.com/sirupsen/logrus)
  - Multiple output formats (text ([[l2met][http://r.32k.io/l2met-introduction]), json)
  - Multiple levels (can configure output level)
  - Hooks (e.g. sending to airbrake when an error is logged)
  - Fields (can be used for metrics or context)
- [Logs are streams](http://adam.herokuapp.com/past/2011/4/1/logs_are_streams_not_files/)
- [`GODEBUG`](http://golang.org/pkg/runtime)
  - Can be used to output information about garbage collection and the scheduler

- Logging Example (see logging/ and logrus/ subdirectories)

# Presenter Notes
- l2met can plug into services like librato
- Demonstrate logging with delays (motivate need for having identifier in each entry)
- for i in {1..10} ; do curl -i localhost:8080/factorial/$i & ; done
- Introduce bug into resource parsing
- Any other logging packages people like using?

---

# Interactive debugging

## GDB
- Build application with `go build -gcflags "-N -l"` to disable optimizations
- For tests, run `go test -gcflags "-N -l" -c` to have Go compile the test, but not run it
- Run `gdb -d $(go env GOROOT) <your binary>`

- Gotchas:
  - OSX ships with GDB 6.0 (you'll want > 7.0)
  - Wasn't able to get Go GDB extensions working with `goenv` Go versions
  - Unusable in Go 1.3
    - https://code.google.com/p/go/issues/detail?id=8256
    - https://code.google.com/p/go/issues/detail?id=7803
    - https://code.google.com/p/go/issues/detail?id=5552
  - Hit or miss about which locals are available to be printed

- [Debugging Go Code with GDB](http://golang.org/doc/gdb)
- [GDB User mangual](https://sourceware.org/gdb/current/onlinedocs/gdb/)

# Presenter Notes
- Introduce bug into Factorial function
- Compile test package and step through

---

# Interactive debugging

## GDB
- `list` Print next 10 source lines
- `break`
  - `break <file.go line>` Insert breakpoint in `file.go` at line `line`
  - `break <p.Func>` Insert breakpoint at function `Func` of package `p`
- `bt` Print backtrace
- `info args` Show function arguments
- `info locals` Show local variables
- `p <variable>` Print variable value
  - `<slc>->array[<x>]` print element `x` of `slc`
- `whatis <variable>` Print type of variable
- `continue` Continue to next break point
- `step` Run until next source line
- `next` Run until next source line in the same stack frame
- `` Execute last command again

- Go extensions (http://golang.org/src/pkg/runtime/runtime-gdb.py)
  - `info goroutines` Print goroutines (this a
  - `goroutine <n> <cmd>` Execute `cmd` in the context of goroutine `n`
  - Adds pretty printing of some non-standard types
  - `$len()` Gets len
  - `$dtype()` Get underlying type of interface value

- http://www.goinggo.net/2013/06/installing-go-gocode-gdb-and-liteide.html has a walkthrough for GDB and LiteIDE on OSX
---

# Interactive debugging

## GDB

GDB does not understand Go programs well. The stack management, threading, and runtime contain aspects that differ enough from the execution model GDB expects that they can confuse the debugger, even when the program is compiled with gccgo. As a consequence, although GDB can be useful in some situations, it is not a reliable debugger for Go programs, particularly heavily concurrent ones. Moreover, it is not a priority for the Go project to address these issues, which are difficult. In short, the instructions below should be taken only as a guide to how to use GDB when it works, not as a guarantee of success.

---

# Interactive debugging

## LiteIDE
- Graphical editor for Golang
- Integrates with GDB
- https://github.com/visualfc/liteide

### Setup
- View -> Edit Environment
  - set GOPATH and GOROOT
- Click gear icon when file with main() is open
  - Set `BUILDARGS` to `-gcflags "-N -l"`
- Set breakpoints with red orb
- Debug -> Start Debugging

- http://www.goinggo.net/2013/06/installing-go-gocode-gdb-and-liteide.html has a walkthrough for GDB and LiteIDE on OSX

# Presenter Notes
- Relatively simple to get up and running

---

# Post-mortem debugging

- `GOTRACEBACK`: Dump stack traces when Go program encounters unrecovered panic or unexpected runtime condition
  - 0: no stack traces
  - 1: dump user goroutines
  - 2: dump run time and user goroutines
  - crash: dump user and run time goroutines and crash in OS specific manner (core dump)
    - You can also get this manually by issuing a `SIGABRT` to a running process

- `gdb -d $(go env GOROOT) -c core <your app>`

- Gotchas:
  - Don't forget to set `ulimit -c` to non-zero value
  - TODO: Getting ulimit -c working on OS X


# Presenter Notes
