package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page int
	Users []string
}

type Response2 struct {
	Page int  `json:"page"`
	User []string `json:"user"`
}

func main() {

	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(9)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(1.141)
	fmt.Println(string(floatB))

	strB, _ := json.Marshal("php is best 朗柜机")
	fmt.Println(string(strB))

	s1D := []string{"phper", "javaer", "桌椅维修工程师"}
	s1B ,_:= json.Marshal(s1D)
	fmt.Println(string( s1B) )

	map1D := map[int]string{ 1:"使用php的小红", 2:"使用java的小绿", 3:"使用golang的秃头" }
	map1B, _ := json.Marshal(map1D)
	fmt.Println(string(map1B))

	res1D := &Response1{
		Page:  1,
		Users: []string{"小王", "老刘", "小八"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page: 0,
		User: []string{"赵慧", "王会", "丁会"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//现在来看看解码
	byt := []byte(`{"name":"halweg", "age":19.2, "language":["lisp", "东北话"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)


	// 为了使用解码 map 中的值，我们需要将他们进行适当的类型转换。
	// 例如这里我们将 `num` 的值转换成 `float64` 类型。
	num := dat["age"].(float64)
	fmt.Println(num)

	// 访问嵌套的值需要一系列的转化。
	strs := dat["language"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们也可以解码 JSON 值到自定义类型。这个功能的好处就
	// 是可以为我们的程序带来额外的类型安全加强，并且消除在
	// 访问数据时的类型断言。
	str := `{"page": 1, "user": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.User[0])

	// 在上面的例子中，我们经常使用 byte 和 string 作为使用
	// 标准输出时数据和 JSON 表示之间的中间值。我们也可以和
	// `os.Stdout` 一样，将 JSON 编码直接输出至 `os.Writer`
	// 流中，或者作为 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)



}