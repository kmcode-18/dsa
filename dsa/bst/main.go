package main

import "fmt"

func bst(list []int,element int)(int,bool){
	low,high:=1,len(list)
	i:=1
	for low<=high{
		mid:=(low+high)/2
		guess:=list[mid-1]
		fmt.Println("called",i,mid)
		i=i+1
		if guess == element{
			return element,true
		}else if guess>element{
			high=mid-1
		}else{
			low=mid+1
		}
	}
	return 0,false
}

func main(){
	ele,ifExists:=bst([]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16},16)
	fmt.Println(ele,ifExists)
}

// func search(nums []int, target int) int {
//     low,high:=0,len(nums)-1
//     for i:=0;low<=high;i++{
//         mid:=(high+low)/2
//         guess:=nums[mid]
//         if guess == target{
//             return mid
//         }else if guess<target{
//             low=mid+1
//         }else{
//             high=mid-1
//         }
//     }
//     return -1
// }
