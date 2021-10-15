package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	sr := strings.NewReader("嗨 0x0844, your ID is 1234567890!")
	br := bufio.NewReader(sr)
	// 4096
	fmt.Printf("size: %d\n", br.Size())

	// 嗨  不会移动read指针
	bs, _ := br.Peek(3)
	fmt.Printf("%s\n",bs)


	// 嗨 3 读取一个rune,并返回rune的长度
	r, size, _ := br.ReadRune()
	fmt.Printf("%c %d\n", r, size)

	br.Discard(1)

	// 0x0844,   ","也会被读取
	s, _ := br.ReadString(byte(','))
	fmt.Println(s)

	br.Discard(1)

	bs,_ = br.ReadBytes(byte('!'))
	fmt.Printf("%s\n",bs)
}
