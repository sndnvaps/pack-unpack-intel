
# pack && unpack x86 arch boot/recovery image 

origin source 

	http://forum.xda-developers.com/showthread.php?t=1972589

	https://github.com/turl/razr-i-boot-tools

origin code is language c , i port it to golang


Get it :
	go get github.com/sndnvaps/unpack-pack-intel

usage:

unpack/pack recovery/boot image tools
 	 -origin="": Origin boot/recovery image
 	 -output="": ouput of repack image
 	 -ramdisk="": ramdisk.cpio.gz need to pack into image
 	 -unpack-pack="": unpack/pack the boot/recovery image
	  -zImage="": kernel need to pack into image

example :pack 

	 ./pack-intel -origin="./recovery.img" -unpack-pack="pack" -ramdisk="./ramdisk.cpio.gz" -zImage="./zImage" -output="./new-recovery.img"

: unpack
	
	./pack-intel -origin="./recovery.img" -unpack-pack="unpack" -ramdisk="./ramdisk.cpio.gZ" -zImage="/.zImage" 


	

	


