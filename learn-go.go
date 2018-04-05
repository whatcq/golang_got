package learn_go


// comment
/*
multi lines comment (hard to type without IDE)
*/

import(
    "fmt"
    "io/ioutil"
    //mx "math"
    "net/http"
    "os"
    "strconv"

)

// 语句结尾不用“;”，一行一条语句 

// main是程序执行入口，同C lang
func main(){
    fmt.Println("")

    // 调用其他函数，函数可以写在后面
    beyondHello()
}

func beyondHello(){
    var x int // var变量声明；变量名在前类型在后
    x = 3
    y := 4 //合并上两步，程序会在编译时识别出变量类型

    sum, prod := learnMultiple(x, y)
    fmt.Println(sum, prod)
}

// 多个变量 可以一起定义一个类型
func learnMultiple(x, y int) (sum, prod int){
    return x+y, x+y //返回多个值，类型上面定义了
}

func learnTypes(){
    str := "string：字符串只能用双引号！"
    s2 := `
    多行文本
    `

    g := 'Σ' // 字符 rune类型 int32的别名，UTR-8

    f := 3.1415926
    c := 3+4i //complex126类型，内部使用两个float64表示；2d游戏用得着

    n := byte('\n') //字符转换成byte(uint8)
    
    var arr[4] int //=>[0,0,0,0]
    arr3 := [...]int{3, 1, 5}

    // slice切片(翻译不如英文)？
    s3 := []int{3, 4, 6} //相比array,slice可以动态增删
    s4 := make([]int, 4)

    var d2 [][]float64 //二维slice?
    bs := []byte("a string")

    s6 := append(s3, 6, 4, 6)
    s5 := append(s3, []int{1,2}...)

    //map基本类似与hash,dict; php中array呢，类型无限制，那是动态语言！
    m := map[string]int{"learn": 10, "exercise": 20}
    m["self-control"] = 50

    file, _ := os.Create("output.txt")
    fmt.Fprintf(file, "写入文件。")
    file.Close()

    _,_ = g,arr3

    fmt.Println(str, c, arr, s3, d2, m, s2, f, n, s4, bs, s6, s5, arr3, g)

    learnFlowControl()
}

// type是个问题，太慢，有的字符特别难敲
// 一行行照着写太慢，也容易忘记：不如
// 计划：先看一两遍，然后把代码全去掉，默写！
// https://learnxinyminutes.com/docs/zh-cn/go-cn/


// 和其他编程语言不同的是，go支持有名称的变量返回值。
// 声明返回值时带上一个名字允许我们在函数内的不同位置
// 只用写return一个词就能将函数内指定名称的变量返回
func learnNamedReturns(x, y int) (z int) {
    z = x + y
    return // z is implicit here, because we named it earlier.
}

// Go全面支持垃圾回收。Go有指针，但是不支持指针运算。
// 你会因为空指针而犯错，但是不会因为增加指针而犯错。
func learnMemory() (p, q *int) {
    // 返回int型变量指针p和q
    p = new(int)// 内置函数new分配内存
    // 自动将分配的int赋值0，p不再是空的了。
    s := make([]int, 20) // 给20个int变量分配一块内存
    s[3] = 5// 赋值
    r := -2 // 声明另一个局部变量
    return &s[3], &r// & 取地址
}

func expensiveComputation() int {
	return 1e6
}

func learnFlowControl() {
    // If需要花括号，括号就免了
    if true {
		fmt.Println("这句话肯定被执行")
	}

    // 用go fmt 命令可以帮你格式化代码，所以不用怕被人吐槽代码风格了，
    // 也不用容忍别人的代码风格。
	if false {
		// pout
	}else {
		// gloat
	}

    // 如果太多嵌套的if语句，推荐使用switch
    x := 1
	switch x {
	case 0:
	case 1:
		x = 999// 隐式调用break语句，匹配上一个即停止
	case 2:
		x = 2// 不会运行
	}

    // 和if一样，for也不用括号

    for i:=1;i<10; i++ { // ++ 自增
		fmt.Println("遍历", i)
	}

    // x在这里还是1。为什么？

    // for 是go里唯一的循环关键字，不过它有很多变种
    for { // 死循环
		break// 骗你的
		continue// 不会运行的
	}

    // 用range可以枚举 array、slice、string、map、channel等不同类型
    // 对于channel，range返回一个值，
    // array、slice、string、map等其他类型返回一对儿
	for key,value := range map[string]int{"a":1,"b":2} {
		// 打印map中的每一个键值对
		fmt.Printf("索引：%s, 值为：%d\n", key, value)
	}
    // 如果你只想要值，那就用前面讲的下划线扔掉没用的
	for _, name := range []string{"cqiu", "xx"} {
		fmt.Printf("你是。。 %s\n", name)
	}


    // 和for一样，if中的:=先给y赋值，然后再和x作比较。
	if y:=3;y>2{
		fmt.Println("ok")
	}


    // 闭包函数
    xBig := func() bool{
    	return x > 100
	}

    // x是上面声明的变量引用
    fmt.Println("xBig:", xBig()) // true （上面把y赋给x了）
    x =10 // x变成10
    fmt.Println("xBig:", xBig()) // 现在是false

    // 除此之外，函数体可以在其他函数中定义并调用，
    // 满足下列条件时，也可以作为参数传递给其他函数：
    //   a) 定义的函数被立即调用
    //   b) 函数返回值符合调用者对类型的要求
    fmt.Println("两数相加乘二: ", func(x,y int) int {
    	return (x+y)*2
	}(5,6))


    // Called with args 10 and 2
    // => Add + double two numbers: 24

    // 当你需要goto的时候，你会爱死它的！
    goto love
    love:

    learnFunctionFactory()// 返回函数的函数多棒啊
    learnDefer()// 对defer关键字的简单介绍
    learnInterfaces()// 好东西来了！
}

func learnFunctionFactory() {
    // 空行分割的两个写法是相同的，不过第二个写法比较实用
    fmt.Println(sentenceFactory("原谅")("当然选择", "她！"))

	d := sentenceFactory("原谅")
    fmt.Println(d("当然选择", "她！"))
    fmt.Println(d("你怎么可以", "她？"))
}

// Decorator在一些语言中很常见，在go语言中，
// 接受参数作为其定义的一部分的函数是修饰符的替代品
func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return before+mystring+after// new string
	}
}

func learnDefer() (ok bool) {
    // defer表达式在函数返回的前一刻执行
    defer fmt.Println("11111")
    defer fmt.Println("22222")

    // 关于defer的用法，例如用defer关闭一个文件，
    // 就可以让关闭操作与打开操作的代码更近一些
	return true
}

// 定义Stringer为一个接口类型，有一个方法String
type Stringer interface{
	String() string
}

// 定义pair为一个结构体，有x和y两个int型变量。
type pair struct{
	x,y int
}

// 定义pair类型的方法，实现Stringer接口。
func (p pair) String() string { // p被叫做“接收器”
    return fmt.Sprintf("(%d, %d)", p.x, p.y)// Sprintf是fmt包中的另一个公有函数。
    // 用 . 调用p中的元素。
}

func learnInterfaces() {
    p := pair{2, 4}// 花括号用来定义结构体变量，:=在这里将一个结构体变量赋值给p。

    fmt.Println(p.String()) // 调用pair类型p的String方法
    var i Stringer// 声明i为Stringer接口类型
    i = p// 有效！因为p实现了Stringer接口（类似java中的塑型）
    // 调用i的String方法，输出和上面一样
    fmt.Println(i.String())

    // fmt包中的Println函数向对象要它们的string输出，实现了String方法就可以这样使用了。
    // （类似java中的序列化）
    fmt.Println(p) // 输出和上面一样，自动调用String函数。
    fmt.Println(i) // 输出和上面一样。

	learnVariadicParams("a", "bx")
}

// 有变长参数列表的函数
func learnVariadicParams(myStrings ...interface{}) {
    // 枚举变长参数列表的每个参数值
    // 下划线在这里用来抛弃枚举时返回的数组索引值
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

    // 将可变参数列表作为其他函数的参数列表
    fmt.Println("params:", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	m := map[int]string{2: "BB"}
    // ", ok"用来判断有没有正常工作
    if x, ok := m[1]; !ok { // ok 为false，因为m中没有1
		fmt.Println("别找了真没有")
	}else {
		fmt.Print(x) // 如果x在map中的话，x就是那个值喽。
	}

    // 错误可不只是ok，它还可以给出关于问题的更多细节。
    if _, err:= strconv.Atoi("NaN");err != nil{// _ discards value
		// 输出"strconv.ParseInt: parsing "non-int": invalid syntax"
		fmt.Println(err)
	}

    // 待会再说接口吧。同时，
    learnConcurrency()
}

// c是channel类型，一个并发安全的通信对象。
func inc(i int, c chan int) {
    c <- i// <-把右边的发送到左边的channel。
}

// 我们将用inc函数来并发地增加一些数字。
func learnConcurrency() {
    // 用make来声明一个slice，make会分配和初始化slice，map和channel。
    c := make(chan int)

    // 用go关键字开始三个并发的goroutine，如果机器支持的话，还可能是并行执行。
    // 三个都被发送到同一个channel。
    go inc(2, c)// go is a statement that starts a new goroutine.
    go inc(5, c)
    go inc(9, c)

    // 从channel中读取结果并打印。
    // 打印出什么东西是不可预知的。
    fmt.Println(<-c, <-c, <-c) // channel在右边的时候，<-是读操作。

    cs := make(chan string)// 操作string的channel
    cc := make(chan chan string)// 操作channel的channel
    go func() {c <- 33}()// 开始一个goroutine来发送一个新的数字
    go func() {cs <- "ttt"}()// 发送给cs
	// Select类似于switch，但是每个case包括一个channel操作。
	select {
		// 它随机选择一个准备好通讯的case。
	case i := <- c: // 从channel接收的值可以赋给其他变量
		fmt.Println("这是……", i)
	case <-cs: // 或者直接丢弃
		fmt.Println("这是个字符串！")
	case <- cc:
		// 空的，还没作好通讯的准备
		fmt.Println("别瞎想")
	}
    // 上面c或者cs的值被取到，其中一个goroutine结束，另外一个一直阻塞。

    // Go很适合web编程，我知道你也想学！
    learnWebProgramming()
}

// http包中的一个简单的函数就可以开启web服务器。
func learnWebProgramming() {
    // ListenAndServe第一个参数指定了监听端口，第二个参数是一个接口，特定是http.Handler。
	go func() {
		err := http.ListenAndServe("127.0.0.1:8080", nil)
		fmt.Println(err) // 不要无视错误。
	}()

	requestServer()
}

// 使pair实现http.Handler接口的ServeHTTP方法。
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // 使用http.ResponseWriter返回数据
	w.Write([]byte("xxxx"))
}

func requestServer() {
    resp, err := http.Get("http://localhost:8080")
    if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf("\n服务器消息： `%s`", string(body))
}

// 变量，函数参数，返回值，任何地方都要注意类型的声明
// 函数库的使用是下一步，需要看手册