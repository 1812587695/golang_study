package main

import "./util" // 导入时使用相对路径

func main() {
	util.Test("0Test message")

	var u = &util.Util{} // 实例化类

	u.Test("1Test message")
}
