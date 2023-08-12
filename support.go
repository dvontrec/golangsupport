package support

import (
	"fmt"
	"time"
	"math/rand"
)

func GetEvens(arr []int){
	fmt.Printf("we got an array %d", arr[0])
}

func PrintArray(arr []int){
 for _, d := range arr {
	fmt.Println(d)
 }
}

func CreateRandomArray(length int) []int{
	fmt.Println("updating with rqandoms")
	arr := []int{}
	index := 0
	for index <= length {
		arr = append(arr, GenerateRandomNumber())
		index += 1
	}
	return arr
}

func GenerateRandomNumber() int {
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	return r1.Int()
}
	
