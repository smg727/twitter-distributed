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
		return &pb.LoginReply{Status:false}, errors.New("Wrong password")
	}
}

func (s *server) AddTweet(ctx context.Context, in *pb.AddTweetRequest) (*pb.AddTweetReply, error){
	user, ok := userdata[in.Username]
	if(!ok){
		debugPrint("No such user")
		return &pb.AddTweetReply{Status:false}, errors.New("No such User")
	}
	newTweet := tweet{text: in.TweetText}
	user.tweets = append(user.tweets, newTweet)
	userdata[in.Username] = user
	return &pb.AddTweetReply{Status:true},nil
}

func (s *server) OwnTweets(ctx context.Context, in *pb.OwnTweetsRequest) (*pb.OwnTweetsReply, error){
	user, ok := userdata[in.Username]
	if(!ok){
		debugPrint("No such user")
		return nil, errors.New("No such User")
	}
	response := pb.OwnTweetsReply{}
	for _,i := range user.tweets{
		tweetToAdd := pb.Tweet{Text:i.text}
		response.TweetList = append(response.TweetList, &tweetToAdd)
	}
	debugPrint("your tweets")
	fmt.Println(response)
	return &response,nil
}

func (s *server) UserExists(ctx context.Context, in *pb.UserExistsRequest) (*pb.UserExistsReply, error) {
	username:=in.Username
	_, ok := userdata[username]
	if(!ok) {
		debugPrint("No such user")
		return &pb.UserExistsReply{Status: false}, errors.New("No such user exists")
	}else{
		return &pb.UserExistsReply{Status:true}, nil
	}
}

func (s *server) DeleteUser(ctx context.Context, in *pb.Credentials) (*pb.DeleteReply, error){

	//TODO: for later stages, we might have to add Locks here
	debugPrint("Deleting User: " + in.Uname +"Account")
	delete(userdata,in.Uname)
	return &pb.DeleteReply{DeleteStatus:false}, errors.New("User deleted")

	//user, ok := userdata[in.Uname]
	//if(!ok){
	//	debugPrint("No such user")
	//	return &pb.LoginReply{Status:false}, errors.New("No such User")
	//}
	//if(in.Pwd==user.password){
	//	return &pb.LoginReply{Status:true}, nil
	//}else {
	//	return &pb.LoginReply{Status:false}, errors.New("Wrong Password")
	//}
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
