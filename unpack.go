package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"unsafe"
)

func Unpack() {
	var bzImageLen int32
	var ramdiskLen int32
	bufs := make([]byte, 8192)
	//var bufs [8192]byte
	var missing int32

	originFile := origin /*"recovery.img" */

	fin, err := os.Open(originFile)
	if err != nil {
		fmt.Println(originFile, err)

	}
	defer fin.Close()

	// zimage unpack
	zImagefile := bzImage //"zImage"

	fzimage, fz := os.Create(zImagefile)
	if fz != nil {
		fmt.Println(fz)
	}
	defer fzimage.Close()

	//ramdisk.cpio.gz
	ramdiskfile := ramdisk //"ramdisk.cpio.gz"

	framdisk, fr := os.Create(ramdiskfile)
	if fr != nil {
		fmt.Println(fr)
	}
	defer framdisk.Close()

	//Read the zImage len to unpack
	_, e := fin.Seek(2016 /*CMDLINE_END*/, os.SEEK_SET)
	if e != nil {
		fmt.Println(e)
	}
	err = binary.Read(fin, binary.LittleEndian, &bzImageLen)
	if err != nil {
		fmt.Println(err)
	}

	var tmp uint32
	//Read ramdisk length from the image to unpack
	_, e = fin.Seek(2016+int64(unsafe.Sizeof(tmp)), os.SEEK_SET)
	if e != nil {
		fmt.Println(e)
	}

	err = binary.Read(fin, binary.LittleEndian, &ramdiskLen)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("size of zImage  = %d\n", bzImageLen)
	fmt.Printf("size of ramdisk = %d\n", ramdiskLen)

	// Copy zImage
	//missing := bzImageLen

	_, err = fin.Seek(9184 /* 9184 */, os.SEEK_SET) //int64(unsafe.Sizeof(bootheader struct) = 9184
	if err != nil {
		fmt.Println(err)
	}

	missing = bzImageLen
	//err = binary.write(bufs,binary.LittleEndian, fin)

	//fmt.Printf("zImage size  = %d\n", missing)
	var size int

	for missing > 0 {

		size, err = fin.Read(bufs)
		if size != 0 && err == nil {
			_, e = fzimage.Write(bufs)
		}
		if e != nil {
			fmt.Println(e)
		}

		missing = missing - int32(size)

	}

	//copy ramdisk

	_, err = fin.Seek(9184+int64(bzImageLen), os.SEEK_SET) //int64(unsafe.Sizeof(bootheader struct) = 9184
	if err != nil {
		fmt.Println(err)
	}

	missing = ramdiskLen
	//fmt.Printf("ramdisk.cpio.gz size = %d\n", missing)
	for missing > 0 {
		size, err = fin.Read(bufs)
		if size != 0 && err == nil {
			_, e = framdisk.Write(bufs)
			if e != nil {
				fmt.Println(e)
			}
		}

		missing = missing - int32(size)

	}

}
