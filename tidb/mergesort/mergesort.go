package main

import "sync"
func Divide(src []int64,start,end,depth int,wt *sync.WaitGroup){
	if start>=end {
		if depth==1 { //在16m的数组下是无用的,但保证程序可用性还是加了
			wt.Done()
		}
		return
	}
	var mid int=(start+end)/2
	Divide(src,start,mid,depth+1,wt)
	Divide(src,mid+1,end,depth+1,wt)
	Merge(src,start,mid,end)
	if depth==1 {
		wt.Done()
	}
}
func Merge(src []int64,start,mid,end int){
	tmp:=make([]int64,end-start+1)
	i,j,k:=start,mid+1,0
	for i<=mid && j<=end {
		if src[i]<src[j] {
			tmp[k]=src[i]
			k++;i++
		}else{
			tmp[k]=src[j]
			k++;j++
		}
	}
	for i<=mid {
		tmp[k]=src[i]
		k++;i++
	}
	for j<=end {
		tmp[k]=src[j]
		k++;j++
	}
	for t:=0;t<k;t++ {
		src[start+t]=tmp[t]
	}
}
func MergeSort(src []int64) {
	length:=len(src)
	var mid int=length/2
	var wt sync.WaitGroup
	wt.Add(2)
	go Divide(src,0,mid,1,&wt)
	go Divide(src,mid+1,length-1,1,&wt)
	wt.Wait()
	Merge(src,0,mid,length-1)
}