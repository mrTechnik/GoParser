package main

import (
	"fmt"
	//"net/http"
	"src"
)

func main() {
	//cli := &http.Client{}

	// req, err := src.MyRequest(*cli, "https://benchmark.best/en/cpu_table.html?numofcores=8&scrl=300")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(req)

	var myQueue src.FreshLinksContainer = src.FreshLinksContainer{MaxRamBuffer: 1000, MinimalRamBuffer: 100}
	var i int = 100
	for i >= 1 {
		err := myQueue.AddNewElem(i)
		if err != nil {
			fmt.Println(err)
		}
		i--
	}

	var flag bool = true
	for flag {
		var elem, err any = myQueue.GetFirstElemAndDelete()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(elem)

		if elem == nil {
			flag = false
		}
	}
}
