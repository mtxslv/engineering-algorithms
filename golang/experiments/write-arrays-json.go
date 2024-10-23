package main

import (
	"encoding/json"
	"fmt"
	// "math/rand/v2"
	"os"
)

type Dados struct {
	n []int
	y []int
}

func generate_array(n int, reversed bool) []int {
	arr := make([]int, n)
	for i := 0 ; i < n; i++ {
		if reversed {
			arr[i] = n - i - 1
		} else {
			arr[i] = i
		}
	} 
	return arr
}

func check(e error){
	if e != nil {
		panic(e)
	}
}

func write_array_json(arr_n, arr_y []int) {
	tempos := Dados{arr_n, arr_y}
	fmt.Println(tempos)
	b, err := json.Marshal(tempos)
	check(err)
	err = os.WriteFile("./a.json", b, 0644)
	check(err)
}

func main(){
	array_n := []int{1,2,3}
	array_y := []int{1,2,3}
	write_array_json(array_n, array_y)
}