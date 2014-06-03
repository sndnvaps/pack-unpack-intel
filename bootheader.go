package main

import (
//"fmt"
//"unsafe"
//"bytes"
)

var (
	HEAD_PADDING       int64 = 0x3E0 // if device is moto razr i 
	                                 // HEAD_PADDING = 0 
	CMDLINE_SIZE       int64 = 0x400
	PADDING1_SIZE      int64 = 0xC00
	BOOTSTUBSTACK_SIZE int64 = 0x1000
	CMDLINE_END        int64 = 0x7E0 // HEAD_PADDING + CMDLINE_SIZE
)

/*
func init() {
	//HEAD_PADDING = 0 // Motorola RAZR i
	HEAD_PADDING = 0x3e0 // common Android Intel Image
	CMDLINE_SIZE = 0x400
	PADDING1_SIZE = 0xC00 // 0x1000 - 0x400
	BOOTSTUBSTACK_SIZE = 0x1000
	CMDLINE_END = HEAD_PADDING + CMDLINE_SIZE

}
*/

type bootheader struct {
	//cmdline            [CMDLINE_SIZE]byte
	//cmdline            string
	padding0 [0x3e0]byte

	cmdline            [0x400]byte
	bzImageSize        int64
	initrdSize         int64
	SPIUARTSuppression int64
	SPIType            int64
	padding1           [0xC00]byte
	//padding1 string
	bootstubStack [0x1000]byte
	//bootstubStack string
}

type testT struct {
	padding0 [0x3e0]byte
	cmdline  [0x400]byte
	padding1 [0xC00]byte
	//padding1 string
	bootstubStack [0x1000]byte
}

//type z [(len(bootheader) == 0x2000 + HEAD_PADDING)? 1 : -1]byte
/*
func main() {
	//fmt.Println("len(bootheader) = ", (bootheader))
	var bh bootheader

	fmt.Println(unsafe.Sizeof(bh))
}
*/
