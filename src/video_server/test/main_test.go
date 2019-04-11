package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

//func testPrint(t testing.T){
//	fmt.Println("test")
//}
func TestPrint2(t *testing.T) {
	res := Print1to20()
	res++
	if res != 222 {
		t.Errorf("Test print2 failed")
	}
}

func TestAll(t *testing.T) {
	t.Run("TestPrint", TestPrint)
	t.Run("TestPrint2", TestPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("Tests begints...")
	//m.Run()
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print1to20()
	}
}
