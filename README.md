# updev-go-product-api
Example project with Nic Jackson (https://www.youtube.com/watch?v=VzBGi_n65iU&amp;t=485s&amp;ab_channel=NicJackson)

## get start ep1

Found error ``import net/http `` เกี่ยวกับการ Import นั้นแสดงว่าคุณยังไม่ได้ทำการกำหนด ``GO module``

และนีคือขั้นตอนในการกำหนด go module
    
```
go mod init github.com/ksupdev/updev-go-product-api
```

และ Error ที่เจอก่อนหน้าก็จะหายไป

> กรณีที่เราต้องทำการพัฒนา GO ด้านนนอก GO workspace จะต้องมีการนำ go module มาช่วยในการพัฒนา

ลองทำการ run server ของเราโดยการใช้

```
    go run main.go
```

และมา ``CURL`` กันหน่อย

``` powersshell
    curl -v localhost:9090
    *   Trying ::1...
    * TCP_NODELAY set
    * Connected to localhost (::1) port 9090 (#0)
    > GET / HTTP/1.1
    > Host: localhost:9090
    > User-Agent: curl/7.64.1
    > Accept: */*
    > 
    < HTTP/1.1 404 Not Found
    < Content-Type: text/plain; charset=utf-8
    < X-Content-Type-Options: nosniff
    < Date: Tue, 23 Mar 2021 06:16:02 GMT
    < Content-Length: 19
    < 
    404 page not found
    * Connection #0 to host localhost left intact
    * Closing connection 0
```

> log.Println("Hello world") จะเป็น Command สำหรับการแสดงผลผ่าน Terminal ตัวอย่างข้อมูล ``2021/03/23 13:24:39 Hello world``

> ถ้าเราทำการกำหนด ``http.HandleFunc("/"..... `` นั้นจะหมายความว่าไม่ว่ามันจะเป็น default request หรือก็คือถ้าไม่ตรงกับใครก็จะมาที่เรานี้เอง

เราสามารถใช้ ``d, _ := ioutil.ReadAll(r.Body)`` สำหรับการอ่านค่า Request body เพื่อนำมาแสดงใน log ได้
```go
http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		d, _ := ioutil.ReadAll(r.Body)

		log.Printf("Data %s\n", d)
})
```

ทดสอบการ request โดยให้มีการโยน body data มาด้วย
```powershell
    curl -v -d 'pu' localhost:9090
    ..
    ..
    ..
    //output
    2021/03/24 23:52:34 Data pu
```

> ในการที่เราจะ write log เพื่อแสดงผลใน terminal เราสามารถใช้ ``log.Printf("",.....)`` แต่ถ้าเราต้องการให้มีการ write log แทนที่จะผ่าน terminal แต่ให้ Response กลับไปที่ user เราก็สามารถใช้ ``rw http.ResponseWriter`` และ ``fmt.Fprintf()`` สำหรับกำหนด format ของข้อมูล เพื่อ Response กลับไป

## ep2
    เราทำการแยกการทำงานของ Hello และ Good bye ออกจาก ``main.go`` แลำทำการสร้างของ package ``handlers`` เพื่อทำการเก็บ Hello,GoodBye 

    > การในส่วนของ http.Handler มีการกำหนด interface ในภาษา GO มันไม่จำเป็นที่ต้องมีการระบุว่า func ของเรานั้น Implement interface ไหน Go จะสนใจแค่ว่าขอให้มี ชื่อและ parameter ให้ตรงตาม interface ก็จะสามารถใช้งานได้ทันที

```GO
[filename : server.go]
package http

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
// ----- end file
[filename : hello.go]
type Hello struct {
    ......
}

func NewHello(l *log.Logger) *Hello {
    ......
}
// create func has structure same Interface http.Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    .......
}

// ----- end file

[filename : main.go]

l := log.New(os.Stdout, "product-api", log.LstdFlags)
hh := handlers.NewHello(l)

sm := http.NewServeMux()
//Method Handle requir string parameter and Interface http.Handler
sm.Handle("/", hh)

```

จากตัวอย่างเราจะเห็นว่าถ้ามีทำการ implement func ServeHTTP สำหรับ Hello เราก็จะไม่สามารถทำการ mapping request กับ Hello action ได้


