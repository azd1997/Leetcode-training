package unionfind

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestUnionFindV1(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := 1000
	uf := NewUnionFindV1(n)
	t1 := time.Now()
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.Union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.IsConnected(a, b)
	}
	t2 := time.Now()

	fmt.Printf("time used: %s\n", t2.Sub(t1).String())
}

func TestUnionFindV2(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := 1000
	uf := NewUnionFindV2(n)
	t1 := time.Now()
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.Union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.IsConnected(a, b)
	}
	t2 := time.Now()

	fmt.Printf("time used: %s\n", t2.Sub(t1).String())
}

func TestUnionFindV3(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := 1000
	uf := NewUnionFindV3(n)
	t1 := time.Now()
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.Union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.IsConnected(a, b)
	}
	t2 := time.Now()

	fmt.Printf("time used: %s\n", t2.Sub(t1).String())
}

func TestUnionFindV4(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := 1000
	uf := NewUnionFindV4(n)
	t1 := time.Now()
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.Union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.IsConnected(a, b)
	}
	t2 := time.Now()

	fmt.Printf("time used: %s\n", t2.Sub(t1).String())
}

func TestUnionFindV5(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := 1000
	uf := NewUnionFindV5(n)
	t1 := time.Now()
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.Union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.IsConnected(a, b)
	}
	t2 := time.Now()

	fmt.Printf("time used: %s\n", t2.Sub(t1).String())
}

func TestUnionFindV6(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := 1000
	uf := NewUnionFindV6(n)
	t1 := time.Now()
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.Union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf.IsConnected(a, b)
	}
	t2 := time.Now()

	fmt.Printf("time used: %s\n", t2.Sub(t1).String())
}

//time used: 1.402944ms
// time used: 635.838µs
// time used: 201.489µs
// time used: 183.859µs
// time used: 166.221µs
// time used: 243.278µs
// 可以看出，尽管V6版本的路径压缩比V5版本逻辑上要更好，
// 但是由于使用了递归操作，实际效果是不如V5版本的
