package examples

import (
	"fmt"
	"sorting-algorithms/structures"
)

func HashMap() {
	hashmap := structures.NewHashMap[string](3)
	hashmap.Set("george", "russel")
	hashmap.Set("lewis", "hamilton")
	hashmap.Set("max", "verstappen")
	fmt.Println(hashmap.Get("george"))
	fmt.Println(hashmap.Get("max"))
	fmt.Println(hashmap.Get("lewis"))
}
