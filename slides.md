# Debugging Go

## Jesse Szwedko
## j.szwedko@modcloth.com
## @jesse_szwedko
## This talk: [github.com/jszwedko/gsp-go-debugging](https://github.com/jszwedko/gsp-go-debugging)

### Location: ModCloth, Inc. (Pittsburgh Office)
### Date: 7/31/2014

# ![modcloth](modcloth.png)

---

# Debugging

## Tracing
## Interactive debugger
## Post-mortem debugging
## Remote debugger

# Presenter Notes
- Any other debugging styles?

---

# Go Debugging

## Tracing
- `log`
- `logrus`

## Interactive debugger
- `gdb`
- LiteIDE

## Post-mortem debugging
- `GOTRACEBACK`
- `SIGABRT`

# Presenter Notes
- No remote debugger

---

# Factorial web service

---

# Tracing

---

# Tracing

- [Logs are streams](http://adam.herokuapp.com/past/2011/4/1/logs_are_streams_not_files/)
- stdlib [`log`](http://golang.org/pkg/log)
- [`logrus`](https://github.com/sirupsen/logrus)
- [`GODEBUG`](http://golang.org/pkg/runtime)
    - Can be used to output information about garbage collection and the scheduler

# Presenter Notes

---

# Tracing: `log`

- Offers much of what you need from a logger
- Recommend (conditonally?) setting `Flags` to include source line
- Recommend allowing conditional enabling of logs (integration tests)

---

# `log` example

# Presenter Notes
- Introduce sleep and use concurrent requests

---

# Tracing: `logrus`

- Multiple output formats (text ([l2met](http://r.32k.io/l2met-introduction)), json)
- Multiple levels (can configure output level)
- Hooks (e.g. sending to airbrake when an error is logged)
- Fields (can be used for metrics or context)

---

# `logrus` example

# Presenter Notes
- Show request UUID
- Any other logging packages people like using?

---

# Interactive debugging

---

# Interactive debugging

## [GDB](https://sourceware.org/gdb/current/onlinedocs/gdb/)
- The GNU Project Debugger
## [LiteIDE](https://github.com/visualfc/liteide)
- Cross-platform Go IDE

---

# GDB

---

# Interactive debugging: GDB

- Build application with `go build -gcflags "-N -l"` to disable optimizations
- For tests, run `go test -gcflags "-N -l" -c` to have Go compile the test, but not run it
- Run `gdb -d $(go env GOROOT) <your binary>`
- Go outputs DWARF debugging information with the binary
    - See [Introduction to the DWARF Debugging Format](http://www.dwarfstd.org/doc/Debugging%20using%20DWARF-2012.pdf) for a good intro to DWARF

# Presenter Notes
- `objdump -h`
- `readelf -w`

---

# Interactive debugging: GDB

## Gotchas:
- OSX ships with GDB 6.0 (you'll want > 7.0)
- Wasn't able to get Go GDB extensions working with versions built somewhere other than install location (ignores `GOROOT`)
- Hit or miss about which locals are available to be printed
- Unusable in Go 1.3 (for me at least)
    - [gdb: Wrong values for local variables](https://code.google.com/p/go/issues/detail?id=8256)
    - [gdb: breakpoints break things](https://code.google.com/p/go/issues/detail?id=7803)
    - [gdb: nothing works (windows amd64)](https://code.google.com/p/go/issues/detail?id=5552)
- [Debugging Go Code with GDB](http://golang.org/doc/gdb)
- [GDB user manual](https://sourceware.org/gdb/current/onlinedocs/gdb/)

---

# Interactive debugging: GDB

- `list` Print next 10 source lines
- `break`
  - `break <file.go line>` Insert breakpoint in `file.go` at line `line`
  - `break <p.Func>` Insert breakpoint at function `Func` of package `p`
- `bt` Print backtrace
- `continue` Continue to next break point
- `step` Run until next source line
- `next` Run until next source line in the same stack frame

---

# Interactive debugging: GDB

- `info args` Show function arguments
- `info locals` Show local variables
- `p <variable>` Print variable value
    - `<slc>->array[<x>]` print element `x` of `slc`
- `whatis <variable>` Print type of variable
- Leave empty to execute last command again


---

# Interactive debugging: GDB

## [Go extensions](http://golang.org/src/pkg/runtime/runtime-gdb.py)
- Adds pretty printing of some non-standard types
- `info goroutines` Print goroutines
- `goroutine <n> <cmd>` Execute `cmd` in the context of goroutine `n`
- `$len()` Gets len
- `$cap()` Gets capacity
- `$dtype()` Get underlying type of interface value
- [Walkthrough of installing gdb/LiteIDE on OSX](http://www.goinggo.net/2013/06/installing-go-gocode-gdb-and-liteide.html)

---

# GDB example

# Presenter Notes
- Introduce bug into Factorial function
- Compile test package and step through

---

# Interactive debugging: GDB

## From http://golang.org/doc/gdb

## "GDB does not understand Go programs well. The stack management, threading, and runtime contain aspects that differ enough from the execution model GDB expects that they can confuse the debugger, even when the program is compiled with gccgo. [...] Moreover, it is not a priority for the Go project to address these issues, which are difficult. [...]"

---

# LiteIDE

---

# Interactive debugging: LiteIDE

## Graphical editor for Golang
## Integrates with GDB
## https://github.com/visualfc/liteide

---

# Interactive debugging: LiteIDE

## Setup
- View -> Edit Environment
    - set GOPATH and GOROOT
- Click gear icon when file with main() is open
    - Set `BUILDARGS` to `-gcflags "-N -l"`
- Set breakpoints with the red orb
- Debug -> Start Debugging
- [Walkthrough of installing gdb/LiteIDE on OSX](http://www.goinggo.net/2013/06/installing-go-gocode-gdb-and-liteide.html)

# Presenter Notes
- Relatively simple to get up and running

---

# LiteIDE Example

---

# Post-mortem debugging

---

# Post-mortem debugging

## `GOTRACEBACK`: Dump stack traces when Go program encounters unrecovered panic or unexpected runtime condition
- `0`: no stack traces
- `1`: dump user goroutines backtrace (default)
- `2`: dump run time and user goroutines backtrace
- `crash`: dump user and run time goroutines and crash in OS specific manner (core dump)
    - You can also get this manually by issuing a `SIGABRT` to a running process
- `gdb -d $(go env GOROOT) -c core <your app>`

---

# Post-mortem debugging

## Gotchas:
- Still want to compile with `-gcflags "-N -l"` to get debugging information
- Don't forget to set `ulimit -c` to non-zero value

---

# Post-mortem debugging example

# Presenter Notes
- Remove checking of error from log initialization
- Also demonstrate SIGABRT

---

# Go Debugging

## Tracing
- `log`
- `logrus`

## Interactive debugger
- `gdb`
- LiteIDE

## Post-mortem debugging
- `GOTRACEBACK`
- `SIGABRT`


---

# Questions

## Jesse Szwedko
## j.szwedko@modcloth.com
## @jesse_szwedko
## This talk: [github.com/jszwedko/gsp-go-debugging](https://github.com/jszwedko/gsp-go-debugging)
