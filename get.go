package goroutineid

import (
	"sync/atomic"
	"unsafe"
)

func Get() uint64 {
	ptr := uintptr(goroutinePtr())
	runtimeG := (*[32]uint64)(unsafe.Pointer(ptr))

	offset := atomic.LoadInt64(&goRoutineIdOffset)
	if offset > -1 {
		return runtimeG[int(offset)]
	}

	id := goroutineID()

	matchedCount := 0
	matchedOffset := 0
	for offset, value := range runtimeG[:] {
		if value == id {
			matchedOffset = offset
			matchedCount++
			if matchedCount == 2 {
				break
			}
		}
	}

	if matchedCount == 1 {
		atomic.StoreInt64(&goRoutineIdOffset, int64(matchedOffset))
	}

	return id
}

func goroutinePtr() uint64

var goRoutineIdOffset int64 = -1
