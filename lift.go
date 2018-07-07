package main

import (
	"fmt"
	"time"
)

func main() {
	var floors [34]int
	l := &lift{
		floor:     0,
		direction: 0,
		weight:    0,
		floors:    floors, //数组语法？
	}
	go l.listen()
}

type lift struct {
	floor     int
	direction int
	weight    int
	floors    [34]int
}

//direction: 0 电梯内部请求到哪一层；1/-1 哪一层要上/下
func (l *lift) listen() {
	var floor int
	var direction bool
	for {
		fmt.Println("floor & direction: ")
		// 等待输入？
	}
}
// 扫描请求，并move
func (l *lift) run() {
	// 死循环该怎样写?
	// l.floors会改变吗？需要channel吗？
	for { //这个函数的职责、逻辑不清楚
		for floor, direction := range l.floors {
			if direction != 0 && floor > l.floor*l.direction {
				//move 1 floor with 1 second
				l.floor += l.direction
				time.Sleep(1 * time.Second)
				fmt.Println("-> ", l.floor)
				break
			}
		}
	}
}
