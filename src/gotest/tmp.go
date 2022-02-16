package main

import (
	"fmt"
)

//
// import (
// 	"fmt"
// 	"strconv"
// 	"utils"
// )
//
// func main() {
// 	node, err := utils.NewNode(0)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	fmt.Println(node.Generate())
// 	fmt.Println(strconv.FormatInt(int64(node.Generate()), 10))
// }

//
// import (
// 	"fmt"
// 	"path"
// 	"runtime"
// )
//
// func main() {
// 	fileName, funcName, line := f2(3)
// 	fmt.Printf("file:%v;function:%v;line:%d\n", fileName, funcName, line)
//
// 	fmt.Printf("[%s:%d]\n", fileName, line)
// }
//
// func getLocation(skip int) (fileName, funcName string, line int) {
// 	pc, file, line, ok := runtime.Caller(skip)
// 	if !ok {
// 		fmt.Println("get info failed")
// 		return
// 	}
// 	fileName = path.Base(file)
// 	funcName = runtime.FuncForPC(pc).Name()
// 	return
// }
//
// func f1(skip int) (fileName, funcName string, line int) {
// 	fileName, funcName, line = getLocation(skip)
// 	return
// }
//
// func f2(skip int) (fileName, funcName string, line int) {
// 	return f1(skip)
// }

//
// import (
// 	"fmt"
// 	"time"
// 	"utils"
// )
//
// func main() {
// 	y, w := utils.YeadAndWeek()
// 	fmt.Println("y==", y)
// 	fmt.Println("w==", w)
//
// 	fmt.Println(utils.DayTimeZero2())
// 	fmt.Println(utils.DayTimeZero())
// 	now := time.Now().Unix()
// 	now -= now % 86400
// 	fmt.Println(now - int64(time.Now().Weekday()*86400))
//
// 	start, end := utils.MonthStartEnd()
// 	fmt.Println(start, end)
// }

//
// import (
// 	"event"
// 	"fmt"
// 	. "event"
// )
//
// func TestHandler(et EventType, args ...interface{}) {
// 	fmt.Println("test2 event handler...")
// }
//
// func main() {
// 	ev := event.NewEvent()
// 	ev.RegEvent(ET_LOGIN, TestHandler)
// 	ev.RegEvent(ET_LOGIN, TestHandler)
//
// 	ev.DispatchEvent(ET_LOGIN, 1, 3, 4)
// }

//
// import (
// 	"fmt"
// 	"utils"
// )
//
// func main() {
// 	fmt.Println(utils.CRC32("123456"))
// 	fmt.Println(utils.MD5("123456"))
// 	fmt.Println(utils.SHA1("123456"))
// }

//
// import (
// 	"fmt"
// 	. "utils"
// )
//
// func main() {
// 	list := NewLinkList()
// 	for i := 0; i < 5; i++ {
// 		node := NewLinkNode(i)
// 		list.PushFront(node)
// 	}
// 	fmt.Println(list)
// 	c := list.Count()
// 	for c > 0 {
// 		list.PopFront()
// 		fmt.Println(list)
// 		c--
// 	}
// }

//
// import (
// 	"fmt"
// 	"os"
// 	"utils"
// )
//
// func main() {
// 	fmt.Println(os.Getpid())
// 	fmt.Println(utils.NewObjID())
// 	fmt.Println(len(utils.NewObjID()))
// 	fmt.Println(utils.UniqID())
// 	fmt.Println(len(utils.UniqID()))
// }

//
// import (
// 	"fmt"
// 	"utils"
// )
//
// func main() {
// 	// node := -1 ^ (-1 << 10)
// 	// fmt.Println(node)
// 	var arr []interface{}
// 	arr = append(arr, 1)
// 	arr = append(arr, 2)
// 	arr = append(arr, 3)
// 	arr = append(arr, 4)
// 	arr = append(arr, 5)
// 	utils.GRand.Shuffle(arr)
// 	fmt.Println(arr)
// }

//
// import (
// 	"fmt"
// 	"utils"
// )
//
// func main() {
// 	bitmap := utils.NewBitMap(2)
// 	bitmap.Set(0)
// 	bitmap.Set(5)
// 	bitmap.Set(10)
// 	bitmap.Set(39)
// 	bitmap.Set(40)
// 	bitmap.Set(41)
// 	bitmap.Set(63)
// 	bitmap.Set(64)
// 	bitmap.Set(256)
// 	fmt.Println(bitmap.IsExist(40))
// 	fmt.Println(bitmap.IsExist(0))
// 	fmt.Println(bitmap.Len())
// 	fmt.Println(bitmap)
//
// 	bitmap.Clear()
// 	fmt.Println(bitmap)
// }

//
// import (
// 	"fmt"
// 	"utils/linklist"
// )
//
// func main() {
// 	circles := linklist.NewCircleQueue()
// 	circles.Enqueue(1)
// 	circles.Enqueue(2)
// 	circles.Enqueue(3)
// 	circles.Enqueue(4)
// 	circles.Enqueue(5)
// 	circles.Dequeue()
// 	circles.Dequeue()
// 	circles.Dequeue()
// 	circles.Enqueue(5)
// 	circles.Enqueue(6)
// 	circles.Enqueue(7)
//
// 	circles.ForEach(func(data interface{}) bool {
// 		fmt.Println(data)
// 		return true
// 	})
// 	fmt.Println("count=", circles.Len())
// }

//
// import (
// 	"fmt"
// 	. "global"
// )
//
// func main() {
// 	err := GConfig.LoadConfig("src/global/config/server.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(GConfig.ChatServer)
//
// 	fmt.Println(1e-4)
// }

//
// import (
// 	"context"
// 	"fmt"
// 	"time"
// )
//
// type faveContextKey string //定义别名
//
// func main() {
// 	fmt.Println("context包学习")
// 	UseContext()
// 	UseContextWithDeadline()
// 	UseContextWithTimeout()
// 	UseContextWithValue()
// }
//
// func UseContext() {
// 	// 使用withCancel生成上下文
// 	gen := func(ctx context.Context) <-chan int {
// 		dst := make(chan int) //无缓冲，无长度
// 		n := 1
// 		fmt.Println("这个函数被调用多少次：", n) //被调用一次
// 		go func() {
// 			for {
// 				select {
// 				// 多个case同时满足，就会随机执行case
// 				case <-ctx.Done():
// 					// 关闭上下文
// 					fmt.Println("ctx.Done", ctx.Err(), n)
// 					return
// 				case dst <- n: //要么运行完这个作用域，要么就不用运行
// 					n++
// 					fmt.Printf("n = %v \n", n)
// 				}
// 			}
// 		}()
// 		//发送器
// 		return dst
// 	}
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	data := gen(ctx)
// 	for n := range data { //接收
// 		fmt.Println(n)
// 		if n == 16 {
// 			break
// 		}
// 	}
// }
//
// func UseContextWithDeadline() {
// 	// withDeadline作用：设置context的存活期
// 	d := time.Now().Add(50 * time.Millisecond)
// 	ctx, cancel := context.WithDeadline(context.Background(), d)
// 	if ctx == nil {
// 		fmt.Println("ctx is nil")
// 		return
// 	}
// 	defer cancel()
// 	select {
// 	case <-time.After(time.Second * 1):
// 		fmt.Println("overslept")
// 	case <-ctx.Done():
// 		fmt.Println(ctx.Err())
// 	}
// }
//
// func UseContextWithTimeout() {
// 	// withTimeout作用：设置context的存活期
// 	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
// 	if ctx == nil {
// 		fmt.Println("ctx is nil")
// 		return
// 	}
// 	defer cancel()
// 	select {
// 	case <-time.After(time.Second * 1):
// 		fmt.Println("overslept")
// 	case <-ctx.Done():
// 		fmt.Println(ctx.Err())
// 	}
// }
//
// func UseContextWithValue() {
// 	// withValue作用：带值
// 	f := func(ctx context.Context, k faveContextKey) {
// 		if v := ctx.Value(k); v != nil {
// 			fmt.Printf("key:%v,value:%v \n", k, v)
// 		} else {
// 			fmt.Println("key not found!", k)
// 		}
// 	}
// 	k := faveContextKey("language")
// 	ctx := context.WithValue(context.Background(), k, "go")
// 	if ctx == nil {
// 		fmt.Println("ctx is nil")
// 	}
// 	f(ctx, k)
// 	f(ctx, faveContextKey("color"))
// }

//
// import (
// 	"fmt"
// 	"utils"
// )
//
// func main() {
// 	consistentHash := utils.NewConsistentHash()
// 	consistentHash.Insert(1)
// 	consistentHash.Insert(2)
// 	consistentHash.Insert(3)
//
// 	fmt.Println(consistentHash.Get("AAA"))
// 	fmt.Println(consistentHash.Get("aaa"))
// 	fmt.Println(consistentHash.Get("ccc"))
// 	fmt.Println(consistentHash.Get("FFF"))
//
// 	consistentHash.Delete(1)
// 	fmt.Println("========================")
// 	fmt.Println(consistentHash.Get("AAA"))
// 	fmt.Println(consistentHash.Get("aaa"))
// 	fmt.Println(consistentHash.Get("ccc"))
// 	fmt.Println(consistentHash.Get("FFF"))
//
// 	consistentHash.Insert(4)
// 	consistentHash.Insert(5)
//
// 	fmt.Println("========================")
// 	fmt.Println(consistentHash.Get("AAA"))
// 	fmt.Println(consistentHash.Get("aaa"))
// 	fmt.Println(consistentHash.Get("ccc"))
// 	fmt.Println(consistentHash.Get("FFF"))
// }

// package main
//
// import (
// 	"fmt"
// 	"time"
// )
//
// func main() {
// 	//now := time.Now()
// 	fmt.Println(time.Now().Format("2006/1/2 15:04:05"))
// 	str := time.Now().Format("2006/1/2 15:04:05")
// 	fmt.Println(str)
// }

//
// import (
// 	"fmt"
// 	"reflect"
// 	"runtime"
// )
//
// type Test struct {
// 	name string
// }
//
// func (t *Test) TempFunc(str1, str2 string) {
// 	fmt.Println(t.name)
// 	fmt.Println(str1, str2)
// }
//
// type Service struct {
// 	servers map[string]reflect.Method
// 	rcvr    reflect.Value
// 	typ     reflect.Type
// }
//
// func NewService(rep interface{}) *Service {
// 	service := new(Service)
// 	service.typ = reflect.TypeOf(rep)
// 	service.rcvr = reflect.ValueOf(rep)
//
// 	name := reflect.Indirect(service.rcvr).Type().Name()
// 	fmt.Println(name)
// 	service.servers = map[string]reflect.Method{}
// 	fmt.Println(service.typ.NumMethod(), service.typ.Name())
// 	for i := 0; i < service.typ.NumMethod(); i++ {
// 		method := service.typ.Method(i)
// 		// mtype := method.Type
// 		fmt.Println("name===", method.Name)
// 		service.servers[method.Name] = method
// 	}
// 	return service
// }
//
// func main() {
// 	t := new(Test)
// 	t.name = "hahaha"
// 	service := NewService(t)
//
// 	for funcName, method := range service.servers {
// 		values := []reflect.Value{service.rcvr, reflect.ValueOf("aaaaa"), reflect.ValueOf("bbbbbbbbb")}
// 		// replyv := reflect.ValueOf("xxxxxxxx")
// 		function := method.Func
// 		fmt.Println("call func: ", funcName)
// 		function.Call(values)
// 	}
// 	// funcName := "TempFunc"
// 	// if method, ok := service.servers[funcName]; ok {
// 	// 	//replyType := method.Type.In(1).Elem()
// 	// 	replyv := reflect.ValueOf("xxxxxxxx")
// 	// 	function := method.Func
// 	// 	function.Call([]reflect.Value{service.rcvr, replyv})
// 	// }
//
// 	method := reflect.ValueOf(t)
// 	for i := 0; i < method.NumMethod(); i++ {
// 		me := method.Method(i)
// 		fmt.Println(me.Pointer(), me.Type(), me.String(), reflect.TypeOf(me.Interface()).Name())
// 		name := runtime.FuncForPC(me.Pointer()).Name()
// 		fmt.Println(name)
// 	}
// }

//
// import (
// 	"fmt"
// 	"reflect"
// 	"runtime"
// 	"runtime/debug"
// 	"strings"
// )
//
// func foo() {}
//
// func GetFunctionName(i interface{}, seps ...rune) string {
// 	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
// 	fields := strings.FieldsFunc(fn, func(sep rune) bool {
// 		for _, s := range seps {
// 			if sep == s {
// 				return true
// 			}
// 		}
// 		return false
// 	})
//
// 	if size := len(fields); size > 0 {
// 		return fields[size-1]
// 	}
// 	return ""
// }
//
// func main() {
// 	fmt.Println("name:", GetFunctionName(foo))
// 	fmt.Println("=================================================")
// 	fmt.Println(GetFunctionName(debug.FreeOSMemory))
// 	fmt.Println(GetFunctionName(debug.FreeOSMemory, '.'))
// 	fmt.Println(GetFunctionName(debug.FreeOSMemory, '/', '.'))
// }

//
// import (
// 	"fmt"
// 	"time"
// 	"utils/com"
// )
//
// func main() {
// 	fmt.Println(com.GoID())
// 	go func() {
// 		fmt.Println("new go routine:", com.GoID())
// 	}()
//
// 	time.Sleep(time.Second * 5)
// }

//
// import (
// 	"fmt"
// 	"utils"
// )
//
// type test struct {
// }
//
// func NewTest() interface{} {
// 	return &test{}
// }
//
// func (t *test) Init(a, b string) {
// 	fmt.Println(a, b)
// }
//
// func (t *test) print() {
// 	fmt.Println("test: print")
// }
//
// func main() {
// 	objPool := utils.NewObjPool(NewTest, 0, 15)
// 	for i := 0; i < 20; i++ {
// 		o := objPool.GetObj().(*test)
// 		fmt.Printf("%v\n", o)
// 	}
// 	fmt.Println(objPool.UseCount())
// 	fmt.Println(objPool.Count())
// 	// o := objPool.GetObj("aa", "bb").(*test)
// 	// o.print()
// }

//
// import (
// 	"encoding/json"
// 	"fmt"
// )
//
// type Settings struct {
// 	setList []interface{}
// }
//
// func main() {
// 	// json_str2 := "[false,false,false,false,false,false,true,true,1,false,false,false,80,20,true,{\"nLevel\":1,\"bIsOpen\":true},{\"nLevel\":1,\"bIsOpen\":true},{\"nLevel\":1,\"bIsOpen\":true},{\"tPickMaterialSetList\":[false,false,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true],\"bIsOpen\":true},false,true,true,true,true,true,true,false,false,false,false]"
// 	json_str2 := "[false,false,false,false,false,false,true,true,1,false,false,false,80,20,true,{\"nLevel\":30009001,\"bIsOpen\":true},{\"nLevel\":1,\"bIsOpen\":true},{\"nLevel\":1,\"bIsOpen\":true},{\"tPickMaterialSetList\":[false,false,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true],\"bIsOpen\":true},false,true,true,true,true,true,true,false,false,false,false]"
// 	settings := &Settings{
// 		setList: make([]interface{}, 100),
// 	}
// 	err := json.Unmarshal([]byte(json_str2), &settings.setList)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for k, v := range settings.setList {
// 		fmt.Printf("index=%d value=%v \n", k+1, v)
// 	}
// }

//
// import (
// 	"db"
// 	"fmt"
// 	"time"
// )
//
// func main() {
// 	glocker := db.NewGLocker("test")
// 	count := 0
// 	go func() {
// 		glocker.Lock()
// 		for i := 0; i < 10; i++ {
// 			count++
// 			fmt.Println("1:count==", count)
// 		}
// 		//glocker.UnLock()
// 	}()
// 	go func() {
// 		glocker.Lock()
// 		for i := 0; i < 10; i++ {
// 			count++
// 			fmt.Println("2:count==", count)
// 		}
// 		//glocker.UnLock()
// 	}()
// 	go func() {
// 		glocker.Lock()
// 		for i := 0; i < 10; i++ {
// 			count++
// 			fmt.Println("3:count==", count)
// 		}
// 		//glocker.UnLock()
// 	}()
//
// 	time.Sleep(time.Second * 10)
// 	glocker.Release()
// 	time.Sleep(time.Second * 5)
// }

//
// import (
// 	"fmt"
// )
//
// type test struct {
// 	name string
// }
//
// func NewTest(name string) *test {
// 	return &test{
// 		name: name,
// 	}
// }
// func main() {
// 	t := NewTest("test")
// 	t1 := t
// 	//t = nil
// 	fmt.Println("t1===", t1.name)
// 	t1.name = "bbbbbbb"
// 	fmt.Println("t===", t.name)
// 	// t1 = nil
// 	// fmt.Println("t======", t)
// }

//
// import (
// 	"fmt"
// 	. "utils/com"
// )
//
// func main() {
// 	// DefaultTrimChars := string([]byte{
// 	// 	'\t', // Tab.
// 	// 	'\v', // Vertical tab.
// 	// 	'\n', // New line (line feed).
// 	// 	'\r', // Carriage return.
// 	// 	'\f', // New page.
// 	// 	' ',  // Ordinary space.
// 	// 	0x00, // NUL-byte.
// 	// 	0x85, // Delete.
// 	// 	0xA0, // Non-breaking space.
// 	// })
// 	//
// 	// fmt.Println("trimChars==", DefaultTrimChars)
// 	// tmpString := "\nnnnsdrgere\r\n"
// 	// tmpString = strings.Trim(tmpString, DefaultTrimChars)
// 	// fmt.Println(tmpString)
// 	var t int8 = 1
// 	fmt.Println("t==", ToString(t))
// 	var t1 float64 = 18.04
// 	fmt.Println("t1==", ToString(t1))
//
//
// }

//
// type IService interface {
// 	myprint()
// }
//
// type Service struct {
// }
//
// func (s *Service) myprint() {
//
// }
//
// type kService struct {
// 	Service
// }
//
// func NewService() IService {
// 	k := &kService{}
// 	return k
// }
//
// func main() {
// 	service := NewService()
// 	service.myprint()
// }

//
// import (
// 	"fmt"
// 	"os"
// 	"path"
// 	"utils/file"
// )
//
// func main() {
// 	logPath := "./gamelog/gamelog"
// 	dir := path.Dir(logPath)
// 	if !file.Exists(dir) {
// 		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// }

//
// import (
// 	"fmt"
// 	"strings"
// )
//
// func main() {
// 	url := "aa/bb/cc"
// 	if strings.HasSuffix(url, "/") {
// 		fmt.Println("==============")
// 	} else {
// 		fmt.Println("+++++++++++++++++")
// 	}
// }

// func GetUnixTimeDiffDays(time1 int64, time2 int64) int {
// 	_, offset := time.Now().Zone()
// 	zeroTime1 := time1 - (time1+int64(offset))%86400
// 	zeroTime2 := time2 - (time2+int64(offset))%86400
// 	return int(math.Abs(float64(zeroTime1-zeroTime2)) / 86400)
// }
//
// func main() {
// 	days := GetUnixTimeDiffDays(1640847456, 1640793456)
// 	fmt.Println("days = ", days)
// }

func main() {
	array := make([]int, 0)
	fmt.Println(len(array))
	array = nil
	for _, v := range array {
		fmt.Println(v)
	}
	fmt.Println("aaaaaaaaaaa")
}
