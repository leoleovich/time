//go:build 386 && !darwin

/*
Copyright (c) Facebook, Inc. and its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package phc

import (
	"time"

	"golang.org/x/sys/unix"
)

func timeToTimespec(t time.Time) unix.Timespec {
	return unix.Timespec{Sec: int32(t.Unix()), Nsec: int32(t.Nanosecond())}
}

func ptpClockTimeToTime(t PTPClockTime) time.Time {
	return time.Unix(int64(t.Sec), int64(t.NSec))
}