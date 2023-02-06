package main

import "fmt"

func main() {
	//map[one:php three:java two:golang]
	var test1 map[string]string
	test1 = make(map[string]string, 10)
	test1["one"] = "php"
	test1["two"] = "golang"
	test1["three"] = "java"
	fmt.Println(test1)
	fmt.Println("------------------")
	//map[one:php2 three:java2 two:golang2]
	test2 := make(map[string]string)
	test2["one"] = "php2"
	test2["two"] = "golang2"
	test2["three"] = "java2"
	fmt.Println(test2)
	fmt.Println("------------------")
	//map[one:php3 three:java3 two:golang3]
	test3 := map[string]string{
		"one":   "php3",
		"two":   "golang3",
		"three": "java3",
	}
	fmt.Println(test3)
	fmt.Println("------------------")
	//map[golang:map[desc:C++ is very good id:2] php:map[desc:PHP是世界上最难写的语言 id:1]]
	lang := make(map[string]map[string]string)
	lang["php"] = make(map[string]string, 2)
	lang["php"]["id"] = "1"
	lang["php"]["desc"] = "PHP是世界上最难写的语言"
	lang["golang"] = make(map[string]string, 2)
	lang["golang"]["id"] = "2"
	lang["golang"]["desc"] = "C++ is very good"
	fmt.Println(lang)
	fmt.Println("------------------")
	//C++会被删掉
	//map[golang:map[desc:C++ is very good id:2] php:map[desc:PHP是世界上最难写的语言 id:1]]
	delete(lang, "desc")
	fmt.Println(lang)
	fmt.Println("------------------")
	//
	lang["golang"]["id"] = "3"
	for list, item := range lang {
		fmt.Println("list = ", list)
		for itemA, ValueA := range item {
			fmt.Println("Set =", itemA)
			fmt.Println("Value =", ValueA)
		}
		fmt.Println("---------------")
	}

}
