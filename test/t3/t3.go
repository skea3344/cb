package main

import (
	"errors"
	"fmt"
	"time"
	"yf/platform/yflogger"

	"github.com/garyburd/redigo/redis"
	"github.com/satori/go.uuid"
)

// ErrTimeout 获取redis锁超时
var ErrTimeout = errors.New("get redis key lock timeout")

func main() {
	defer yflogger.Join()
	log := yflogger.Get("Test")
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Debug("%v", err)

		return
	}
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Debug("%v", err)
		return
	}
	defer c1.Close()
	c.Send("SET", "k1", 100)
	c.Send("GET", "k1")
	c.Flush()
	fmt.Println(c.Receive())
	fmt.Println(redis.Int(c.Receive()))
	fmt.Println("Hello world")

	c.Send("MULTI")
	c.Send("SET", "k1", 200)
	c.Send("GET", "k1")
	fmt.Println(c.Do("EXEC"))
	c.Send("SET", "k1", 200)
	c.Send("GET", "k1")
	fmt.Println(c.Do(""))

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	c.Do("PING")
	// }()

	// c.Send("SUBSCRIBE", "example")
	// c.Flush()
	// for i := 0; i < 100; i++ {
	// 	reply, err := redis.Values(c.Receive())
	// 	if err != nil {
	// 		log.Debug("%v", err)
	// 	} else {
	// 		for j := 0; j < len(reply); j++ {
	// 			switch reply[j].(type) {
	// 			case string:
	// 				fmt.Println("字符串")
	// 				fmt.Println(reply[j].(string))
	// 			case int64:
	// 				fmt.Println("int64")
	// 				fmt.Println(reply[j])
	// 			case []interface{}:
	// 				fmt.Println("[]int64")
	// 			default:
	// 				fmt.Println(redis.String(reply[j], nil))
	// 			}
	// 		}
	//
	// 	}
	// }
	// var d chan int
	// <-d
	// time.Sleep(4 * time.Second)

	// psc := redis.PubSubConn{
	// 	Conn: c,
	// }
	// psc.Subscribe("example")
	// for {
	// 	switch v := psc.Receive().(type) {
	// 	case redis.Message:
	// 		fmt.Printf("%s:message:%s\n", v.Channel, v.Data)
	// 	case redis.Subscription:
	// 		fmt.Printf("%s:%s %d\n", v.Channel, v.Kind, v.Count)
	// 	case error:
	// 		fmt.Println(v)
	// 		return
	// 	}
	// }
	c.Send("MULTI")
	c.Send("SET", "k1", 1000)
	c.Send("SET", "k2", 1000)
	fmt.Println(c.Do("EXEC"))
	c.Do("SET", "k3", 0)
	d := make(chan int, 1)
	go func() {
		addThousand(c1, "k3")
		d <- 1
	}()
	addThousand(c, "k3")
	fmt.Println(redis.Int(c.Do("GET", "k3")))

	c.Do("DEL", "k4")
	fmt.Println(c.Do("SET", "k4", 100))
	fmt.Println(c.Do("SETNX", "k4", 200))
	// lock, err := getLock(c, "k3", 3)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(lock)
	<-d

	c.Do("SET", "k3", 0)
	go func() {
		addThousandWatch(c1, "k3")
		d <- 1
	}()
	addThousandWatch(c, "k3")
	<-d
	fmt.Println(redis.Int(c.Do("GET", "k3")))

}

func addThousand(c redis.Conn, key string) {
	for i := 0; i < 1000; i++ {
		value, err := redis.Int(c.Do("GET", key))
		if err != nil {
			if err.Error() == redis.ErrNil.Error() {
				c.Do("SET", key, 1)
				// fmt.Println(&c, "设置值", 1)
				continue
			}
			panic(err)
		}
		c.Do("SET", key, value+1)
		//
	}
}

func addThousandWatch(c redis.Conn, key string) (err error) {
	defer func() {
		if err != nil {
			c.Do("DISCARD")
		}
	}()
	for i := 0; i < 100; i++ {
		for {
			if _, err = c.Do("WATCH", key); err != nil {
				return
			}
			value, err := redis.Int(c.Do("GET", key))
			if err != nil {
				panic(err)
			}
			// fmt.Println(&c, "取得值", value)
			c.Send("MULTI")
			c.Send("SET", key, value+1)
			ret, err := c.Do("EXEC")
			if err != nil {
				panic(err)
			}
			if ret != nil {
				break
			}
			// fmt.Println("事务失败")
		}
	}
	return
}

// getLock 获取某个key的锁
func getLock(c redis.Conn, key string, timeout int) (string, error) {
	lock := uuid.Must(uuid.NewV4()).String()
	endtime := time.Now().Add(time.Duration(timeout) * time.Second)
	for time.Now().Before(endtime) {
		ret, err := redis.Int(c.Do("SETNX", "lock:"+key, lock))
		if err != nil {
			return "", err
		}
		if ret == 1 {
			return lock, nil
		}
		time.Sleep(5 * time.Millisecond)
	}
	return "", ErrTimeout
}

// zpop 从一个有序集合里面取出排名最前的对象
// 使用乐观锁和流水型事务线实现
func zpop(c redis.Conn, key string) (result string, err error) {
	defer func() {
		if err != nil {
			c.Do("DISCARD")
			// 注意DISCARD和UNWATCH的区别
		}
	}()
	for {
		if _, err = c.Do("WATCH", key); err != nil {
			return "", err
		}
		members, err := redis.Strings(c.Do("ZRANGE", key, 0, 0))
		if err != nil {
			return "", err
		}
		if len(members) != 1 {
			return "", redis.ErrNil
		}
		c.Send("MULTI")
		c.Send("ZREM", key, members[0])
		queued, err := c.Do("EXEC")
		if err != nil {
			return "", err
		}
		if queued != nil {
			result = members[0]
			break
		}
		// 这里还有失败可能然后循环重试的  事务失败
	}
	return result, nil
}
