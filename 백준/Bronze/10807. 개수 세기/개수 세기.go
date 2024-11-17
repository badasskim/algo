package main
import "fmt"

func main(){
	var length int
	var targetNum int

	fmt.Scanf("%d", &length)
	temp := make([]int, length)
	
	for i := 0; i < length; i++{
		fmt.Scanf("%d", &temp[i])
	}
	fmt.Scanf("%d", &targetNum)
	
	var count int = 0

	for _, value := range temp {
		if value == targetNum {
			count++
		}
	}
	
	fmt.Print(count)
	

}

