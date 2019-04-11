package echo2

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		fmt.Println(i)
		s += sep + arg
		sep = ""
	}
	fmt.Println(s)
}
