package sysutil

import (
	"bytes"
	"io/ioutil"
)

const thread_siblings_list_file = `/sys/devices/system/cpu/cpu0/topology/thread_siblings_list`

// IsHyperThreadingEnable
// -1 error occur
// 1 hyper threading is enable
// 0 hyper threading is disable
func IsHyperThreadingEnable() int {
	bs, err := ioutil.ReadFile(thread_siblings_list_file)
	if err != nil {
		return -1
	}

	if bytes.ContainsAny(bs, ",") {
		return 1
	}

	return 0
}
