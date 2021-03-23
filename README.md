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