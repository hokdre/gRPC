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

# Summary Compile
| Type | Proto | Client | Server |
|------|------- | ----- | ------ |
| Unary | rpc NamaMethod(Input) (Output) {}; | NamaMethod(ctx,*Input, ...) (Output, error) | NamaMethod(ctx,*Input, ...) (Output, error) |
| Server Stream | rpc NamaMethod(Input) (stream Output) {}; | NamaMethod(ctx, *Input, ...) (StreamClient, error) | NamaMethod(*Input, StreamServer) error |
<br />

## What's Stream? <br />
> | Properties| StreamClient | StreamServer |
> | --------- | ------------ | ------------ |
> | Interface Name | NamaService_NamaMethodClient | NamaService_NamaMethodServer |
> | Signature | Recv() (*Output, error) | Send(*Ouput) error |

# Reference 
https://github.com/grpc
https://github.com/grpc/grpc/blob/master/doc/statuscodes.md