# Objective <br />
1. Create Server 
   1. Buat Instance Object Server Sesuai dengan contract Interface Service Server
        ```
        type Server struct {
            Store TypeStore
        }

        func Contract(ctx, DataInput) (DataOuput, error){
            //success
            return DataOuput, nil

            //error
            return nil, status.Errorf(code, "description error: %s", err)
        }
        ```
   2. Buat Instance GRPC Server
      ```
      grpcServer := grpc.NewServer()
      ``` 
   3. Hubungkan Instance Server dengan GRPC Server
      ```
      // hasil compile service
      package_hasil_compile_service.Register<NamaService>Server(grpcServer, server)
      ``` 
   4. Buat listener/network yang akan digunakan
      ```
      // contoh melalui tcp connection
      listener, err := net.Listen("tcp", "0.0.0.0:8080")
      ``` 
   5. Serve GRPC Server
      ```
      err := grpcServer.Serve(listener)
      ``` 
    
2. Create Client : **Client Harus Tahu Address Server**.
   1. Buat Connection Dengan Server
      ```
      serverAddress := "0.0.0.0:8080"
      conn, err := grpc.Dial(serverAddress, grpc.WithInSecure())
      ``` 
   2. Buat Instance Client
      ```
      package_hasil_compile_service.New<NamaService>Client(conn)
      ``` 

3. Komunikasi Client - Server
   | Client | Server |
   | -------| ------ |
   | client.Contract(ctx, input) |  //return grpc status as error return output , status.Errorf(code, .... ) |
   | //casting error to grpc status for check err <br /> status , ok := status.FromError(err) <br />
   switch status.Code() { <br /> case .. : <br /> } | .... |
