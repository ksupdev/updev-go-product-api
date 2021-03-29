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


> sig := <-sigChain ref https://tour.golang.org/concurrency/2
```GO
	sigChain := make(chan os.Signal)
	signal.Notify(sigChain, os.Interrupt)
	signal.Notify(sigChain, os.Kill)

	sig := <-sigChain
	l.Println("Recieved terminal , graceful shutdown", sig)
```

> channel : คือวิธีในการนำค่าออกมาจาก go Routines โดยเราสามารถสร้าง channel โดยใช้ ``sigChain := make(chan os.Signal)`` และเราสามารถทำการ set ค่าใส่ channel โดยใช้ ``sigChain <- value ....`` และเราสามารถเอาค่าออกจาก channel โดยใช้ `` sig := <-sigChain``

## ep3
> REST : Representational State Transfer (REST)

เราสามารถทำการ Convert struct array to json โดยใช้ ``json.Marshal(.. data ..)`` 
``` GO
    [filename:./data/products.go]
    func GetProducts() []*Product {
	return productList
    }

    ----- end file ----
    [filename:./handler/products.go]

    func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
        lp := data.GetProducts()

        d, err := json.Marshal(lp)
        if err != nil {
            http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
        }

        rw.Write(d)

    }

```

การเลือกใช้ convert JSON
> The only difference is if you want to play with string or bytes use marshal, and if any data you want to read or write to some writer interface, use encodes and decode. [ref](https://stackoverflow.com/questions/33061117/in-golang-what-is-the-difference-between-json-encoding-and-marshalling#:~:text=Marshal%20and%20Unmarshal%20convert%20a,into%20JSON%20and%20vice%20versa.&text=The%20only%20difference%20is%20if,interface%2C%20use%20encodes%20and%20decode.)

## ep4
test with post method 

``` powershell
curl http://localhost:9090 -d '{"id":1, "name":"tea", "description":"a nice cup of tea"}' 
curl -v http://localhost:9090/1 -XPUT -d '{ "name":"update tea", "description":"a nice cup of tea"}'
```


