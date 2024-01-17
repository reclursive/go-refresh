package main
import (
	"fmt"
)
//syntax GO review 

type node struct{
 data int
 next *node
}

type linkedList struct{
	head *node
	length int
}
 
func (l *linkedList) prepend (n *node){
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData(){
	toPrint := l.head
	for l.length != 0{
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length --
	}
	fmt.Println("\n")
}
func (l *linkedList) deletewValue(value int){
	if l.length == 0 {
		return 
	}
	if l.head.data == value{
		l.head = l.head.next
		l.length --
		return
	}

	prevToDelete := l.head
	for prevToDelete.next.data != value{
		if prevToDelete.next.next == nil{
			return
		}
		prevToDelete = prevToDelete.next
	}
	prevToDelete = prevToDelete.next.next
	l.length--
}
func main(){
	myList := linkedList{}
	node1 := &node{data:48}
	node2 := &node{data:18}
	node3 := &node{data:16}
	node4 := &node{data:4}
	node5 := &node{data:1}
	node6 := &node{data:6}
	node7 := &node{data:33}
	node8 := &node{data:22}
	node9 := &node{data:99}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.prepend(node7)
	myList.prepend(node8)
	myList.prepend(node9)
	myList.printListData()

	myList.deletewValue(99)
	myList.printListData()

}
	// name, age := "doireann", 32
	// saying := "good morn"
	// time := []string{"morn", "eve", "afternoon"}
	
	// fav := 4
	// letters := []string{"m", "a", "r", "t", "y"}
	// fmt.Println(fav >= 8)
	// fmt.Println(fav >= 10)
	// if strings.Contains(saying, time[0]) {
	// 	fmt.Println("Good morning!")
	// } else if strings.Contains(saying, time[1]) {
	// 	fmt.Println("good evening")
	// } else {
	// 	fmt.Println("Hey")
	// }

	// for i := 0; i < len(letters); i++ {
	// 	fmt.Println(letters[i])
	// } 

	// for index, value := range letters{
	// 	fmt.Printf("%v is at index: %v \n", value, index)
	// }


