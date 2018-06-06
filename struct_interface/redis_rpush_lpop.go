package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
)

type RedisInfo struct {
	Conn redis.Conn
}

//--------------------------------------rpush接口
type Rpush interface {
	Rpush() error
}

// 实现redis的rpush接口
func (r *RedisInfo) Rpush() error{
	sadd_match_user := time.Now().Format("20060102") + "_match_user_1"

	for i := 1; i <= 10;  i++{
		_, err := r.Conn.Do("rpush", sadd_match_user, i)


		if err != nil {
			return err
		}
	}
	return nil
}

// 接口函数调用
func PushFunc(r Rpush) error {
	return r.Rpush()
}

//--------------------------------lpop接口
type Lpop interface {
	Lpop() error
}

func (r *RedisInfo) Lpop() error {

	match_user := time.Now().Format("20060102") + "_match_user_1"
	for i := 1; i<= 100;  i++{
		value, err := redis.String(r.Conn.Do("lpop", match_user))

		if err == nil {
			fmt.Printf("value： %v - time: %s \n", value, time.Now().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf("no: %s \n", time.Now().Format("2006-01-02 15:04:05"))
		}
	}

	return nil
}

func LpopFunc(r Lpop) error {
	return r.Lpop()
}

//-------------------------------------
const tcp = "192.168.0.222:6379"
const password = "123456"

func main() {

	// 连接redis服务
	conn, err := redis.Dial("tcp", tcp, redis.DialPassword(password))
	if err != nil {
		fmt.Println("Connect to redis error", conn)
		return
	}

	// 这一种只符合单个引用
	//var rpushInterface Rpush
	//rpushInterface = &RedisInfo{
	//	Conn: conn,
	//}

	// 这种go会自动调用rpush接口和lpop接口
	rpushInterfaceAndLpopInterface := &RedisInfo{
		Conn: conn,
	}


	// 调用rpush接口
	fmt.Println(PushFunc(rpushInterfaceAndLpopInterface))

	// 调用lpop接口
	LpopFunc(rpushInterfaceAndLpopInterface)

}
