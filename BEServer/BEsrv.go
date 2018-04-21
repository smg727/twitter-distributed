package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "twitter-distributed/utils/ProtoDef"
	"google.golang.org/grpc/reflection"
	"fmt"
	"errors"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}
// SayHello implements helloworld.GreeterServer

//userdata
var userdata = make(map[string]User)

type User struct {
	username string
	password string
	tweets []tweet
	follows map[string]bool
}

type tweet struct {
	text string
}
//userdataend

//debugfuntion
var debugon = true //if set to true debug outputs are printed

//Function to print debug outputs if debugon=true
func debugPrint(text string){
	if(debugon){
		fmt.Println(text)
	}
}

//test function
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

//test function2
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again friend " + in.Name}, nil
}


//registeruser function
func (s *server) Register(ctx context.Context, in *pb.Credentials) (*pb.RegisterReply, error){

	usrname:=in.Uname
	pwd:=in.Pwd

	_, ok := userdata[usrname]
	if(ok){
		debugPrint("Debug: User already exists")
		return &pb.RegisterReply{Message:"User already exists"},errors.New("user already exists")
	}
	usr := User{username:usrname,password:pwd}
	usr.follows = make(map[string]bool)
	userdata[usrname] = usr
	debugPrint("Debug: User added")
	return &pb.RegisterReply{Message:"User succesfully added"},nil
	}

func (s *server) Login(ctx context.Context, in *pb.Credentials) (*pb.LoginReply, error){
	user, ok := userdata[in.Uname]
	if(!ok){
		debugPrint("No such user")
		return &pb.LoginReply{Status:false}, errors.New("No such User")
	}
	if(in.Pwd==user.password){
		return &pb.LoginReply{Status:true}, nil
	}else {
		return &pb.LoginReply{Status:false}, errors.New("Wrong Password")
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
