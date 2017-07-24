// Manipulate POSIX process CPU-time clocks.
// Copyright (C) 2017  Laboratory of ACM/ICPC, Xidian University

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warramty of
// MERCHANTABILITY or FITNESS FOR A PARICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Author: Xi Ruoyao <ryxi@stu.xidian.edu.cn>

// +build linux

// Package cpuclock manipulates POSIX process CPU-time clocks, which is
// more effective than reading /proc/[PID]/stat.
package cpuclock

import (
	"os"
	"syscall"
	"time"
	"unsafe"
)

// A CPUClock represents an instant of POSIX process CPU-time clock.
type CPUClock struct {
	clockID int
}

// New returns a CPUClock corresponding to a POSIX CPU-time clock of the
// given process.
func New(proc *os.Process) (clock CPUClock, err error) {
	/* This magic expression is from Linux kernel ABI for CPU clocks.  */
	var clockID int = (^proc.Pid)<<3 | 2

	/* Do a clock_getres call to validate it.  */
	_, _, errno := syscall.Syscall(syscall.SYS_CLOCK_GETRES,
		uintptr(clockID), uintptr(0), uintptr(0))
	if errno == syscall.EINVAL {
		errno = syscall.ESRCH
	}

	if errno != 0 {
		return CPUClock{}, errno
	}

	return CPUClock{
		clockID: clockID,
	}, nil
}

// GetTime returns the time of the CPUClock.
func (clock CPUClock) GetTime() (t time.Time, err error) {
	var ts syscall.Timespec

	_, _, errno := syscall.Syscall(syscall.SYS_CLOCK_GETTIME,
		uintptr(clock.clockID), uintptr(unsafe.Pointer(&ts)), uintptr(0))
	if errno != 0 {
		return time.Time{}, errno
	}

	return time.Unix(ts.Unix()), nil
}
