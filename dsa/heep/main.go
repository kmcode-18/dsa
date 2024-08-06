package main

import(
	"fmt"
)
//MaxHeap struct has a slice 
type MaxHeap struct{
	array []int
}
//Insert insert key in MaxHeap
func (h *MaxHeap)Insert(key int){
	h.array=append(h.array,key)
	h.maxHeapifyUp(len(h.array)-1)
}
//Extract extract the largest element from heap and remove from heap
func(h *MaxHeap)Extract()int{
	l:=len(h.array)-1
	if len(h.array) == 0 {
		fmt.Println("cannot extract because length of array is zero")
	}
	extract:=h.array[0]
	h.array[0]=h.array[l]
	h.array=h.array[:l]
	h.maxHeapifyDown(0)
	return extract
}

//maxHeapifyDown will heapify from top to bottom
func(h *MaxHeap)maxHeapifyDown(index int){
	lastIndex:=len(h.array)-1
	l,r:=left(index),right(index)
	childToCompare:=0
	//will loop through till index has atleast one child
	for l<=lastIndex{
		//when left child is only child
		if l == lastIndex{
			childToCompare=l
		}else if h.array[l]>h.array[r]{//when left child is larger
			childToCompare=l
		}else{//when right child is larger
			childToCompare=r
		}
		//compare the value at current index with larger child i.e childToCompare
		if h.array[childToCompare]>h.array[index]{
			//swap
			h.swap(childToCompare,index)
			index=childToCompare
			l,r=left(index),right(index)
		}else{
			return
		}
	}

}

//maxHeapifyUp reagange elements in heap in upward direction comparing with parent
func (h *MaxHeap)maxHeapifyUp(key int){
	for h.array[parent(key)]<h.array[key]{
		h.swap(parent(key),key)
		key=parent(key)
	}

}
//get parent index
func parent(i int)int{
	return (i-1)/2
}
//get left child index
func left(i int)int{
	return i*2+1
}
//get right child index
func right(i int)int{
	return i*2+2
}
//swap keys in array
func (h *MaxHeap)swap(i1,i2 int){
	h.array[i1],h.array[i2]=h.array[i2],h.array[i1]
}
func main(){
	m:=&MaxHeap{}
	buildHeap:=[]int{10,20,30,5,7,9,11,13,15,17}
	fmt.Println(buildHeap)
	for _,v:= range buildHeap{
		m.Insert(v)
		fmt.Println(m)
	}
	fmt.Println(m)
	for i:=0;i<5;i++{
		m.Extract()
		fmt.Println(m)
	}
}