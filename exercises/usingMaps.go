package main

import "fmt"

func main() {
	iMap := make(map[string]int)
	iMap["k1"] = 12
	iMap["k2"] = 13
	fmt.Println("iMap:", iMap)

	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	}
	fmt.Println("anotherMap:", anotherMap)

	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	fmt.Println("anotherMap:", anotherMap)

	val, ok := iMap["doesItExist"]
	if ok {
		fmt.Println("Exists!", val)
	} else {
		fmt.Println("Does NOT exist", val)
	}

	for k, v := range iMap {
		fmt.Println(k, v)
	}
}
