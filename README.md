# twitter-distributed

Team Members:
Sangram Ghuge - smg727
Geetesh Nikhde - gpn218

Setup Steps:
1. Clone repoistory to a folder with GoPath set

Back-End Server:
1. go to BEServer folder and run the back-end server using: go run BEsrv.go
2. TO run the back-end server, we need GRPC set up on the machine
3. Ensure the following grpc libraries are present at the path GoPATH/src/ :
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
4. The above libraries can be obtained as shown here: https://grpc.io/docs/quickstart/go.html


Front-End Server:
1. The following files need to be built to run: Data.go & srv.go
2. Have the back-end server running before the front-end server starts
3. Go to - http://localhost:9090/home If you are not logged in, you will be redirected to the login page.

