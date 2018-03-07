package sysutil

/*
#include <unistd.h>

int getJiffies() {
    return sysconf(_SC_CLK_TCK);
}
*/
import "C"

// GetJiffies get jiffies of current version kernel
func GetJiffies() int {
	return int(C.getJiffies())
}
