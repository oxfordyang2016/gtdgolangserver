package models

import (
	// "fmt"

	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/gomodule/redigo/redis"
)

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

// set_status, err1 := redis.Int(redisDBcon.Do("SET", "initscore", total_score_float_precision))
func SetStruct(c redis.Conn, key string, value interface{}) error {
	fmt.Printf("%#v\n", value)
	datas, _ := json.Marshal(value)
	//缓存数据
	c.Do("set", key, datas)
	// //读取数据
	// rebytes, _ := redis.Bytes(conn.Do("get", "struct3"))
	// //json反序列化
	// object := &TestStruct{}
	// json.Unmarshal(rebytes, object)

	//
	// _, err := c.Do("hmset", redis.Args{key}.AddFlat(value)...)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }

	return nil
}

func GetStructofemail(c redis.Conn, key string) (interface{}, interface{}) {

	rebytes, _ := redis.Bytes(c.Do("get", key))
	//json反序列化
	object := new(EmailVerify)
	err := json.Unmarshal(rebytes, &object)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Printf("%#v\n", object)
	return object, nil
}
