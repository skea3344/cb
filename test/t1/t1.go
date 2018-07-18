package main

/*
int Sum(int a, int b){
return a+b;
}
*/
import "C"

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode/utf8"
)

// import "fmt"

const (
	A = iota
	B
	C
	D
	E
)

type T1 struct {
	m map[string]uint32
}

func NewT1() {
	return
}

func (t *T1) f(code string) uint32 {
	if t.m["aaa"] == 0 {
		fmt.Println("哈哈")
	}
	return t.m[code]

}

func main() {
	defer func() {
		fmt.Println("中断信号")
	}()
	fmt.Println(C.Sum(100, 200))

	t := &T1{m: make(map[string]uint32)}

	fmt.Println(t.f("abs"))
	var a uint32
	a = math.MaxUint32
	fmt.Println(a + 2)
	a++
	fmt.Println(a)

	testUint32 := uint32(0)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if atomic.LoadUint32(&testUint32) == 1 {
				fmt.Println("已经是1了")
				return
			}
			// time.Sleep(2 * time.Millisecond)
			atomic.AddUint32(&testUint32, 1)
			fmt.Println("加1")
		}()
	}

	wg.Wait()
	fmt.Println(testUint32)

	c := make(chan int, 10)

	go func() {
		c <- 10
		c <- 20
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
	path := os.Getenv("GOPATH")
	fmt.Println(strings.Split(path, string(os.PathListSeparator)))
	var found []string
	fmt.Println(len(found))

	// fi, err := os.Stat("/home/cb/gocode/src/caibo/test/t1/t1.go")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(filepath.Ext("/home/cb/gocode/src/caibo/test/t1/t1.go"))
	fmt.Println(1 << '\t')

	fmt.Println('\t')
	fmt.Println('\n')
	fmt.Println('\r')
	fmt.Println(' ')

	var buff bytes.Buffer
	buff.WriteString("aaa")
	reader := strings.NewReader("aaa")
	for i := 0; i < 5; i++ {
		c, err := reader.ReadByte()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(c)
	}
	a = 63
	fmt.Println(uint(1 << a))
	p := pos{
		a: 10,
		b: 10,
	}

	p1 := p

	p1.a = 20
	p1.b = 20

	fmt.Println(p1)
	fmt.Println(p)

	fmt.Println(filepath.Base("/a/b/c"))
	fmt.Println(strings.Title("good job are here"))
	test1("ccc")

	buff1 := make([]byte, 10)
	utf8.EncodeRune(buff1, '我')
	fmt.Println(buff1)
	var str string
	str = "abc"
	if str == "abc" {
		fmt.Println("真的")
	}
	if str == string("abc") {
		fmt.Println("真的到底")
	}

	fmt.Println(strings.TrimLeft(".test.Error", "."))
	switch {
	case true:
		break
	}

	// var lock sync.Mutex
	// var wg1 sync.WaitGroup
	// wg1.Add(1)
	// fmt.Println("加锁前")
	// lock.Lock()
	// fmt.Println("加锁后")
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Println("异步进程")
	// 	wg1.Done()
	// }()
	// fmt.Println("解锁前")
	// fmt.Println("解锁后")
	//
	// wg1.Wait()

	fmt.Println(runtime.NumCPU())

	byteArray := []byte("abcde")
	readBuff := bytes.NewBuffer(byteArray)
	fmt.Println(readBuff.String())
	readBuff.ReadByte()
	fmt.Println(readBuff.String())

	var c1 chan int
	c1 = make(chan int, 10)
	go func() {
		c1 <- 10
	}()
	select {
	case num, ok := <-c1:
		fmt.Println(num, ok)
		// default:
		// 	fmt.Println("默认")
	}

	fmt.Println("退出")

	// var f1 TestFunc
	// f1 = nil
	// f1(1)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%d\n", uint(math.Pow(55, 5))%13)
	fmt.Printf("%d\n", uint(math.Pow(55, 2))%13)
	fmt.Printf("%d\n", uint(math.Pow(9, 2))%13)
	fmt.Printf("%d\n", uint(math.Pow(9, 5))%13)

	last := make(chan bool)
	go func() {
		time.Sleep(20 * time.Second)
		last <- true
	}()

	mmp := make(map[string]int, 20)
	mmp["key"] = 1
	fmt.Println(mmp)
	delete(mmp, "key")
	fmt.Println(mmp)
	<-last
}

type TestFunc func(int) bool

type pos struct {
	a int
	b int
}

func test1(s string) {
	if s == "aaa" {
		fmt.Println(s)
	} else {
		s := "bbb"
		fmt.Println(s)
	}
	fmt.Println(s)
}
