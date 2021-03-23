# updev-go-product-api
Example project with Nic Jackson (https://www.youtube.com/watch?v=VzBGi_n65iU&amp;t=485s&amp;ab_channel=NicJackson)

## get start ep1
    Found error ``import "net/http" `` เกี่ยวกับการ Import นั้นแสดงว่าคุณยังไม่ได้ทำการกำหนด ``GO module``

    และนีคือขั้นตอนในการกำหนด go module
    ```powershell
        go mod init github.com/ksupdev/updev-go-product-api
    ```

    และ Error ที่เจอก่อนหน้าก็จะหายไป

    > กรณีที่เราต้องทำการพัฒนา GO ด้านนนอก GO workspace จะต้องมีการนำ go module มาช่วยในการพัฒนา