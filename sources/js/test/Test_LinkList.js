import LinkList from "../list/LinkList";

function TestLinkList(){

    console.log("********** 顺序表测试开始 **********");

    let l = new LinkList();
    l.display()

    console.log("增加元素-------------");
    l.append(17)
    l.append(20)
    l.display()

    console.log("插入元素-------------");
    l.insert(2, 18)
    l.insert(3, 19)
    l.display()

    console.log("修改元素-------------");
    l.update(2, 22)
    l.display()

    console.log("查找第二个元素前驱：", l.priorElem(22));
    console.log("查找第二个元素后继：", l.nextElem(22));

    console.log("删除元素-------------");
    l.delete(2)
    l.display()

    console.log("********** 顺序表测试结束 **********");
}

export default TestLinkList