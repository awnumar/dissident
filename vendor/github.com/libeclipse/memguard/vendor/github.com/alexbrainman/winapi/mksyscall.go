// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zsyscall_windows.go winapi.go serial.go tls.go memory.go
