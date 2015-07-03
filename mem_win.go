// +build windows

package mem

/*
//#cgo release CFLAGS: -O2 -DNDEBUG
//#cgo windows CFLAGS: -std=gnu99
//#cgo windows LDFLAGS: -lrt

#include "windows.h"

static DWORDLONG get_mem_available() {
	MEMORYSTATUSEX status;
	memset(&status, 0, sizeof(status));
	status.dwLength = sizeof(status);
	
	GlobalMemoryStatusEx(&status);
	
	return status.ullAvailPhys;
}


static DWORDLONG get_mem_total() {
	MEMORYSTATUSEX status;
	memset(&status, 0, sizeof(status));
	status.dwLength = sizeof(status);
	
	GlobalMemoryStatusEx(&status);
	
	return status.ullTotalPhys;
}

*/
import "C"

func GetAvailable() uint64 {
	avail := uint64(C.get_mem_available())
	return avail
}

func GetTotal() uint64 {
	total := uint64(C.get_mem_total())
	return total
}
