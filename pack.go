package main

import (
	//"encoding/binary"
	"flag"
	"fmt"
	"os"
	//"unsafe"
)

var (
	origin      string
	bzImage     string
	ramdisk     string
	output      string
	unpack_pack string
)

var BUFSIZ int64 = 8192

func init() {
	flag.StringVar(&unpack_pack, "unpack-pack", "", "unpack/pack the boot/recovery image")
	flag.StringVar(&origin, "origin", "", "Origin boot/recovery image")
	flag.StringVar(&bzImage, "zImage", "", "kernel need to pack into image")
	flag.StringVar(&output, "output", "", "ouput of repack image")
	flag.StringVar(&ramdisk, "ramdisk", "", "ramdisk.cpio.gz need to pack into image")

	flag.Usage = func() {
		fmt.Printf("unpack/pack recovery/boot image tools\n")
		flag.PrintDefaults()
	}
}

// intel-pack-unpack -unpack-pack="pack" -origin="./recovery.img" -zImage="./zImage" -ramdisk="./ramdisk.cpio.gz" -output="./new-recovery.img"

func Pack() {

	//file := make(bootheader, unsafe.Sizeof(bootheader))
	var file *bootheader = new(bootheader)
	oriImage := origin
	ramdiskImage := ramdisk
	kernel := bzImage
	op := output
	bufs := make([]byte, 8192)
	buf := make([]byte, 9184)

	var size int

	fin, err := os.Open(oriImage)
	if err != nil {
		fmt.Println(err)
	}
	defer fin.Close()

	fout, ferr := os.Create(op)
	if ferr != nil {
		fmt.Println(ferr)
	}
	defer fout.Close()

	fram, fre := os.Open(ramdiskImage)
	if fre != nil {
		fmt.Println(fre)
	}
	defer fram.Close()

	fbz, fbzr := os.Open(kernel)

	if fbzr != nil {
		fmt.Println(fbzr)
	}

	size, err = fin.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	//Figure out the bzImage size and set it
	fi, err := os.Stat(kernel)
	if err != nil {
		fmt.Println(err)
	}
	file.initrdSize = fi.Size()

	//Figure out the ramdisk size and set it
	fi, err = os.Stat(ramdiskImage)
	if err != nil {
		fmt.Println(err)
	}
	file.bzImageSize = fi.Size()

	//Write the patched bootstub to the new image
	fout.Write(buf)

	// Then copy the new bzImage
	for {
		size, err = fbz.Read(bufs)
		if size != 0 && err == nil {
			_, e := fout.Write(bufs)
			if e != nil {
				fmt.Println(e)
			}
		}

		if size == 0 {
			break
		}
	}

	//And finally copy the ramdisk
	for {
		size, err = fram.Read(bufs)
		if size != 0 && err == nil {
			_, e := fout.Write(bufs)
			if e != nil {
				fmt.Println(e)
			}
		}
		if size == 0 {
			break
		}
	}

}

func main() {
	flag.Parse()
	//flag.Args

	if unpack_pack == "unpack" {
		if len(bzImage) == 0 || len(ramdisk) == 0 {
			flag.Usage()
		}
		Unpack()

	}

	if unpack_pack == "pack" {
		if len(output) == 0 ||
			len(origin) == 0 || len(bzImage) == 0 ||
			len(ramdisk) == 0 {
			flag.Usage()
			return
		}
		Pack()
	}

}
