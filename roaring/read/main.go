package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/RoaringBitmap/roaring"
	"github.com/bits-and-blooms/bitset"
)

const (
	SERIAL_COOKIE_NO_RUNCONTAINER = 12346
	SERIAL_COOKIE                 = 12347
	NO_OFFSET_THRESHOLD           = 4
)

type KeyCard struct {
	key  uint16
	card uint16
}

func readCookieHeader(r io.Reader) (cookie uint16, containerNum uint32, runFlagBitset []byte) {
	binary.Read(r, binary.LittleEndian, &cookie)
	switch cookie {
	case SERIAL_COOKIE_NO_RUNCONTAINER:
		var dummy uint16
		binary.Read(r, binary.LittleEndian, &dummy)
		binary.Read(r, binary.LittleEndian, &containerNum)

	case SERIAL_COOKIE:
		var u16 uint16
		binary.Read(r, binary.LittleEndian, &u16)
		containerNum = uint32(u16)
		buf := make([]uint8, (containerNum+7)/8)
		r.Read(buf)
		runFlagBitset = buf[:]

	default:
		log.Fatal("unknown cookie")
	}

	fmt.Println(cookie, containerNum, runFlagBitset)
	return
}

func readOffsetHeader(r io.Reader, cookie uint16, containerNum uint32) {
	if cookie == SERIAL_COOKIE_NO_RUNCONTAINER ||
		(cookie == SERIAL_COOKIE && containerNum >= NO_OFFSET_THRESHOLD) {
		// have offset header
		var offset uint32
		for i := 0; i < int(containerNum); i++ {
			binary.Read(r, binary.LittleEndian, &offset)
			fmt.Println("offset", i, offset)
		}
	}
}

func readDescriptiveHeader(r io.Reader, containerNum uint32) []KeyCard {
	var keycards []KeyCard
	var key uint16
	var card uint16
	for i := 0; i < int(containerNum); i++ {
		binary.Read(r, binary.LittleEndian, &key)
		binary.Read(r, binary.LittleEndian, &card)
		card += 1
		fmt.Println("container", i, "key", key, "card", card)

		keycards = append(keycards, KeyCard{key, card})
	}

	return keycards
}

func readArrayContainer(r io.Reader, key, card uint16, bm *roaring.Bitmap) {
	var value uint16
	for i := 0; i < int(card); i++ {
		binary.Read(r, binary.LittleEndian, &value)
		bm.Add(uint32(key)<<16 | uint32(value))
	}
}

func readBitmapContainer(r io.Reader, key, card uint16, bm *roaring.Bitmap) {
	var u64s [1024]uint64
	for i := 0; i < 1024; i++ {
		binary.Read(r, binary.LittleEndian, &u64s[i])
	}

	bs := bitset.From(u64s[:])
	for i := uint32(0); i < 8192; i++ {
		if bs.Test(uint(i)) {
			bm.Add(uint32(key)<<16 | i)
		}
	}
}

func readRunContainer(r io.Reader, key uint16, bm *roaring.Bitmap) {
	var runNum uint16
	binary.Read(r, binary.LittleEndian, &runNum)

	var startNum uint16
	var length uint16
	for i := 0; i < int(runNum); i++ {
		binary.Read(r, binary.LittleEndian, &startNum)
		binary.Read(r, binary.LittleEndian, &length)
		length += 1
		for j := uint16(0); j < length; j++ {
			bm.Add(uint32(key)<<16 | uint32(startNum+j))
		}
	}
}

func main() {
	data, err := ioutil.ReadFile("../roaring.bin")
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(data)
	cookie, containerNum, runFlagBitset := readCookieHeader(r)

	keycards := readDescriptiveHeader(r, containerNum)
	readOffsetHeader(r, cookie, containerNum)

	bm := roaring.New()
	for i := uint32(0); i < uint32(containerNum); i++ {
		if runFlagBitset != nil && runFlagBitset[i/8]&(1<<(i%8)) != 0 {
			// run
			readRunContainer(r, keycards[i].key, bm)
		} else if keycards[i].card <= 4096 {
			// array
			readArrayContainer(r, keycards[i].key, keycards[i].card, bm)
		} else {
			// bitmap
			readBitmapContainer(r, keycards[i].key, keycards[i].card, bm)
		}
	}

	fmt.Println(bm.String())
}
