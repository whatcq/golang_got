package main

import "fmt"

func main(){
	var floors [34]int
	l:=&lift{
		floor: 0,
		direction: 0,
		weight:0,
		floors:floors,//数组语法？
	}
	go l.listen()
}

type lift struct {
	floor int
	direction int
	weight int
	floors [34]int
}
//direction: 0 电梯内部请求到哪一层；1/-1 哪一层要上/下
func (l *lift)listen(){
	var floor int
	var direction bool
	for{
		fmt.Println("floor & direction: ")
	}
}
func (l *lift)move(floor int, direction bool){
	for{//这个函数的职责、逻辑不清楚
		for i:=l.floor;i< ;  {
			
		}
		for floor, direction = range l.floors{
			if 
		}
	}
}
func (l *lift)stop(floor int){

}