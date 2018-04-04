package learn_go


// comment
/*
multi lines comment (hard to type without IDE)
*/

import(
    "fmt"
    "io/ioutil"
    m "math"
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
    mx := map[string]int{"learn": 10, "exercise": 20}
    mx["self-control"] = 50

    file, _ := os.Create("output.txt")
    fmt.Fprintf(file, "写入文件。")
    file.Close()

    fmt.Println(str, c, arr, s3, d2, mx)

    learnFlowControl()
}

func learnFlowControl(){
    if true {
    	fmt.Println("打印出来")
	}
}

// type是个问题，太慢，有的字符特别难敲
// 一行行照着写太慢，也容易忘记：不如
// 计划：先看一两遍，然后把代码全去掉，默写！
// https://learnxinyminutes.com/docs/zh-cn/go-cn/

















































































