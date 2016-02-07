package main

import (
	"fmt"
	"runtime"
	"strconv"
	"syscall"
)


// TODO: Import gopsutil package and use it instead of
// reinventing the wheel.
func getSystemInformation() map[string]string {
	m := make(map[string]string)

	m["OS"] = runtime.GOOS
	m["Arch"] = runtime.GOARCH
	m["Number of CPUs"] = strconv.Itoa(runtime.NumCPU())
	m["Mem Page Size (bytes)"] = strconv.Itoa(syscall.Getpagesize())

	return m
}

func main() {
	for k, v := range getSystemInformation() {
		fmt.Printf("%-26s: %s\n", k, v)
	}
}
