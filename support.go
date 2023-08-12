package support

import "fmt"

func GetEvens(arr []int){
	fmt.Printf("we got an array %d", arr[0])
}

func PrintArray(arr []int){
 for _, d := range arr {
	fmt.Println(d)
 }
}

func CreateRandomArray(length int) []int{
	arr := []int{}
	index := 0
	for index <= length {
		arr = append(arr, index)
		index += 1
	}
	return arr
}

func main(){
	fmt.Println("hello cruel world")
	nums := CreateRandomArray(13)
	PrintArray(nums)
	GetEvens(nums)
}
	
	
