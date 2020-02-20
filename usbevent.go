package main

import (
	"fmt"
	"os"
	"syscall"
)

// for netlink details see https://www.unix.com/man-page/linux/7/PF_NETLINK/
func init_hotplug_event() {
	sock, _ := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_DGRAM, syscall.NETLINK_KOBJECT_UEVENT)
	addr := syscall.SockaddrNetlink{
		Family: syscall.AF_NETLINK,
		Pad:    0,
		Pid:    uint32(os.Getpid()),
		Groups: 1,
	}
	syscall.Bind(sock, &addr)
	for {
		var p = make([]byte, 2048)
		_, _, err := syscall.Recvfrom(sock, p, 0)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
		fmt.Println(string(p))
	}
}
