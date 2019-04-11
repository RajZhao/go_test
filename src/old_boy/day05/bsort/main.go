package main

import "fmt"

func bsort(a [10]int)[10]int{
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] > a[j-1]{
				a[j-1],a[j] = a[j],a[j-1]
			}
		}
	}
	return a
}
func main(){
	var b = [10]int{1,2,33,41,12,31,4,8,11,22}
	b = bsort(b)
	fmt.Println(b)


}
