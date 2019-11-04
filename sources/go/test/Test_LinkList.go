package test

import (
	"fmt"
	"structure/list/linklist"
)

func TestLinkList() {

	fmt.Printf("********** 单链表测试开始 **********\n")

	l := linklist.NewLinkList()

	fmt.Println("追加元素-------------")
	l.Append(17)
	l.Append(20)
	l.Display()

	fmt.Println("插入元素-------------")
	l.Insert(2, 18)
	l.Insert(3, 19)
	l.Display()

	fmt.Println("修改元素-------------")
	l.Update(2, 22)
	l.Display()

	fmt.Println("查找元素-------------")
	te, _ := l.LocateElem(19)
	pe, _ := l.PriorElem(22)
	ne, _ := l.NextElem(22)
	fmt.Printf("查找第三个元素位置：%d\n", te)
	fmt.Printf("查找第二个元素前驱：%d\n", pe)
	fmt.Printf("查找第二个元素后继：%d\n", ne)

	fmt.Println("删除元素-------------")
	l.Delete(2)
	l.Display()

	fmt.Printf("********** 测试结束 **********\n")

}
