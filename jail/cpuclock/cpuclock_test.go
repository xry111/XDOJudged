// Unitest of package cpuclock
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

package cpuclock_test

import (
	"os"
	"testing"

	"github.com/xry111/XDOJudged/jail/cpuclock"
)

func TestTrivial(t *testing.T) {
	proc, err := os.FindProcess(1)
	if err != nil {
		t.Fatalf("Can not find process 1: %v", err)
	}

	clock, err := cpuclock.New(proc)
	if err != nil {
		t.Fatalf("Can not create CPU-time clock of process 1: %v", err)
	}

	time, err := clock.GetTime()
	if err != nil {
		t.Fatalf("Can not get the time of the CPU-time clock", err)
	}

	t.Logf("result = %v", time)
}
