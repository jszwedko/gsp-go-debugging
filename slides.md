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

- [Debugging Go Code with GDB](http://golang.org/doc/gdb)
- [GDB User mangual](https://sourceware.org/gdb/current/onlinedocs/gdb/)

# Presenter Notes
- Introduce bug into Factorial function
- Compile test package and step throug

---

# Interactive debugging

## GDB
- `list` Print next 10 source lines
- `break`
  - `break <file.go line>` Insert breakpoint in `file.go` at line `line`
  - `break <p.Func>` Insert breakpoint at function `Func` of package `p`
- `bt` Print backtrace
- `info locals` Show local variables
- `p <variable>` Print variable value
- `whatis <variable>` Print type of variable
- `continue` Continue to next break point
- `step` Run until next source line
- `next` Run until next source line in the same stack frame
- `` Execute last command again

- Go extensions (http://golang.org/src/pkg/runtime/runtime-gdb.py)
  - `info goroutines` Print goroutines
  - `goroutine <n> <cmd>` Execute `cmd` in the context of goroutine `n`
  - Adds pretty printing of some non-standard types
  - `$len()` Gets len
  - `$dtype()` Get underlying type of interface value

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

# Presenter Notes
- Relatively simple to get up and running

---

# Post-mortem debugging

# Presenter Notes
