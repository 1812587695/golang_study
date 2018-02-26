package main

//初始化一个结构体：一是得到结构体的对象，一是得到结构的对象指针，分别有三种方式
import "fmt"

type Player struct {
	userid   int
	username string
	password string
}

func main() {
	//第1种方式，先声明对象，再初始化
	var player1 Player
	player1.userid = 1
	player1.username = "lina1"
	player1.password = "123456"
	fmt.Println(player1)

	//第2种方式，声明同时初始化
	player2 := Player{2, "lina2", "123456"}

	fmt.Println(player2)

	//第3种方式，通过 field:value 形式初始化，该方式可以灵活初始化字段的顺序
	player3 := Player{username: "lina3", password: "123456", userid: 3}

	fmt.Println(player3)

	//上面三种初始化方式都是生产对象的，相应如果想初始化得到对象指针的三种方法如下：
	//第1种方式，使用 new 关键字
	player4 := new(Player)
	player4.userid = 4
	player4.username = "lina4"
	player4.password = "123456"
	fmt.Println(player4)

	//第2种方式，声明同时初始化
	player5 := &Player{5, "lina2", "123456"}
	fmt.Println(player5)

	//第3种方式，通过 field:value 形式初始化，该方式可以灵活初始化字段的顺序
	player6 := &Player{username: "lina3", password: "123456", userid: 6}
	fmt.Println(player6)
}
