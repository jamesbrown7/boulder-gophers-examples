package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

//
// See http://msdn.microsoft.com/en-us/library/windows/desktop/aa370651%28v=vs.85%29.aspx
//
func DeleteUser(userName string) (retValue uintptr, err error) {
	dll := syscall.NewLazyDLL("netapi32.dll")

	var netUserDel = dll.NewProc("NetUserDel")

	namePtr, err := syscall.UTF16PtrFromString(userName)
	if err != nil {
		err = fmt.Errorf("Could not convert userName='%s' to UTF16ptr, err='%s'", userName, err)
		return
	}

	retValue, _, err = netUserDel.Call(0, uintptr(unsafe.Pointer(namePtr)))

	return
}

func main() {
	userToDel := os.Args[1]

	retValue, err := DeleteUser(userToDel)
	if err != nil {
		log.Fatalf("Could not delete user, retValue=%d, err='%s'", retValue, err.Error())
	}
}
