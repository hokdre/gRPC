# pre require <br />
Sebelumnya kita telah mengkompile proto file yang berisi message dengan bantuan plugin "google.golang.org/protobuf/cmd/protoc-gen-go" dan mengkompile file tersebut dengan perintah <br />
```
protoc -I=. --go_out=<destination dir> --go_opt=paths=source_relative *.proto;
```
, untuk mengkompile file berisi sebuah service maka dibutuhkan plugin lain yaitu "google.golang.org/grpc/cmd/protoc-gen-go-grpc" dan untuk mengkompile file service dilakukan dengan perintah
```
protoc -I=. --go-grpc_out=<destination dir> --go-grpc_opt=paths=source_relative *.proto
```
sehingga apabila file proto kita mendefinisikan message dan service bersamaan maka kita harus menggabungkan ke 2 perintah kompile, sehingga menjadi 
```
protoc -I=. --go_out=<destination dir message> --go_opt=paths=source_relative  --go-grpc_out=<destination dir services> --go-grpc_opt=paths=source_relative *.proto 
```
> summary : <br />
> google.golang.org/protobuf/cmd/protoc-gen-go untuk mengkompile message <br />
> google.golang.org/grpc/cmd/protoc-gen-go-grpc untuk mengkompile service <br />
> --go_out=<destination dir> --go_opt=paths=source_relative konfigurasi message <br />
> --go-grpc_out=<destination dir services> --go-grpc_opt=paths=source_relative konfigurasi service <br />
# Define a service
1. Buat Proto file service yang mendefinisikan service. contoh : laptop_service.proto
   ```
   message DataInputan {
       fieldsDataInputan
   }
   message DataOuput {
       fieldsDataOuput
   }

   service NamaService {
       rpc NamaFunction(DataInputan) returns (DataOuput) {};
   }
   ```
2. compile file proto 
3. setelah dicompile kita akan mendapatakan fungsionalitas :
   1. object DataInputan <br />
      object data inputan memiliki beberapa method, tetapi yang paling umum adalah method untuk mengakses field properti. seperti GetName(), GetAge() 
   2. object DataOuput <br />
      object data ouputan memiliki beberapa method, tetapi yang paling umum adalah method untuk mengakses field properti. seperti GetName(), GetAge()  
   3. object Service <br />
      akan dihasilkan 2 interface dengan method yang sama sesuai yang diskemakan pada file proto, yaitu :
      1.  serviceClient
      2.  serviceServer

# Reference Plugins
https://github.com/grpc