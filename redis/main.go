package main

import (
	"fmt"
	"lld/redis/services"
)

func main() {

	storage := services.CreateStorage()
	storage.Put("sde_bootcamp", [][]interface{}{{"title", "SDE-Bootcamp"}, {"price", 30000.00}, {"enrolled", false}, {"estimated_time", 30}})
	output1 := storage.Get("sde_bootcamp")
	fmt.Println(output1)
	keys := storage.Keys()
	fmt.Println(keys)
	storage.Put("sde_kickstart", [][]interface{}{{"title", "SDE-Kickstart"}, {"price", 4000}, {"enrolled", true}, {"estimated_time", 8}})
	storage.Get("sde_kickstart")
	keys2 := storage.Keys()
	fmt.Println(keys2)
	storage.Put("sde_kickstart", [][]interface{}{{"title", "SDE-Kickstart"}, {"price", 4000.00}, {"enrolled", true}, {"estimated_time", 8}})
	output2 := storage.Get("sde_kickstart")
	fmt.Println(output2)
	keys3 := storage.Keys()
	fmt.Println(keys3)
}
