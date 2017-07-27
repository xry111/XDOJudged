

# cpuclock
`import "github.com/xry111/XDOJudged/jail/cpuclock"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package cpuclock manipulates POSIX process CPU-time clocks, which is
more effective than reading /proc/[PID]/stat.




## <a name="pkg-index">Index</a>
* [type CPUClock](#CPUClock)
  * [func New(proc *os.Process) (clock CPUClock, err error)](#New)
  * [func (clock CPUClock) GetTime() (t time.Time, err error)](#CPUClock.GetTime)


#### <a name="pkg-files">Package files</a>
[cpuclock.go](/src/github.com/xry111/XDOJudged/jail/cpuclock/cpuclock.go) 






## <a name="CPUClock">type</a> [CPUClock](/src/target/cpuclock.go?s=1093:1130#L23)
``` go
type CPUClock struct {
    // contains filtered or unexported fields
}
```
A CPUClock represents an instant of POSIX process CPU-time clock.







### <a name="New">func</a> [New](/src/target/cpuclock.go?s=1223:1277#L29)
``` go
func New(proc *os.Process) (clock CPUClock, err error)
```
New returns a CPUClock corresponding to a POSIX CPU-time clock of the
given process.





### <a name="CPUClock.GetTime">func</a> (CPUClock) [GetTime](/src/target/cpuclock.go?s=1920:1976#L54)
``` go
func (clock CPUClock) GetTime() (t time.Time, err error)
```
GetTime returns the time of the CPUClock.

NOTE: to get the total CPU time of Process p, we need to use GetTime
before calling p.Wait().  Or the process identifier of p would be
released and we would get EINVAL.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)