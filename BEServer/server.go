package main



import (
	"sync"
	pb "twitter-distributed/utils/ProtoDef"
	"golang.org/x/net/context"
	"errors"
	"log"
	"fmt"
)





// As this is a tester function, not converting it to RPC for now. Let's transorm it if needed

// IsCommitted is called by tester to check whether an index position
// has been considered committed by this server

func (srv *server) IsCommitted(index int) (committed bool) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	if srv.commitIndex >= index {
		return true
	}
	return false
}

// As this is a tester function, not converting it to RPC for now. Let's transorm it if needed

// ViewStatus is called by tester to find out the current view of this server
// and whether this view has a status of NORMAL.

func (srv *server) ViewStatus() (currentView int, statusIsNormal bool) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	return srv.currentView, srv.status == NORMAL
}

// As this is a tester function, not converting it to RPC for now. Let's transorm it if needed

// GetEntryAtIndex is called by tester to return the command replicated at
// a specific log index. If the server's log is shorter than "index", then
// ok = false, otherwise, ok = true

func (srv *server) GetEntryAtIndex(index int) (ok bool, command interface{}) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	if len(srv.log) > index {
		return true, srv.log[index]
	}
	return false, command
}

// As this is a tester function, not converting it to RPC for now. Let's transorm it if needed

// Kill is called by tester to clean up (e.g. stop the current server)
// before moving on to the next test
func (srv *server) Kill() {
	// Your code here, if necessary
}

// Make is called by tester to create and initalize a PBServer
// peers is the list of RPC endpoints to every server (including self)
// me is this server's index into peers.
// startingView is the initial view (set to be zero) that all servers start in
//func Make(peers []*ClientEnd, me int, startingView int) *server {
//	srv := &server{
//		//peers:          peers,
//		me:             me,
//		currentView:    startingView,
//		lastNormalView: startingView,
//		status:         NORMAL,
//		opNo:		0,
//	}
//	// all servers' log are initialized with a dummy command at index 0
//	var v interface{}
//	srv.log = append(srv.log, v)
//
//	// Your other initialization code here, if there's any
//	return srv
//}

// Start() is invoked by tester on some replica server to replicate a
// command.  Only the primary should process this request by appending
// the command to its log and then return *immediately* (while the log is being replicated to backup servers).
// if this server isn't the primary, returns false.
// Note that since the function returns immediately, there is no guarantee that this command
// will ever be committed upon return, since the primary
// may subsequently fail before replicating the command to all servers
//
// The first return value is the index that the command will appear at
// *if it's eventually committed*. The second return value is the current
// view. The third return value is true if this server believes it is
// the primary.

//func (srv *server) Start(command interface{}) (
//	index int, view int, ok bool) {
//	srv.mu.Lock()
//	defer srv.mu.Unlock()
//	// do not process command if status is not NORMAL
//	// and if i am not the primary in the current view
//	if srv.status != NORMAL {
//		return -1, srv.currentView, false
//	} else if GetPrimary(srv.currentView, len(srv.peers)) != srv.me {
//		return -1, srv.currentView, false
//	}
//	srv.log = append(srv.log,command)
//	srv.opNo=srv.opNo+1
//	i:=0
//	count:=0
//	length:=len(srv.peers)
//	done:=make(chan int)
//	for i<length {
//		if(i!=srv.me){
//			val:=i
//			go func(){
//
//				inArgs:=&pb.PrepareArgs{
//					View:int32(srv.currentView),
//					PrimaryCommit:int32(srv.commitIndex),
//					Index:int32(srv.opNo),
//					Entry:command,
//				}
//				outArgs:=&pb.PrepareReply{
//					View:0,
//					Success:false,
//				}
//				okPrep:=srv.sendPrepare(val,&inArgs,&outArgs)
//				if(okPrep==true){
//					if(outArgs.Success==true){
//						count=count+1
//					}else{
//
//					}
//				}else{
//
//				}
//				done<-val
//			}()
//		}
//		i++
//	}
//
//
//	for i:=0;i<(length-1);i++ {
//		<-done
//	}
//	if(count>=(length)/2){
//	//	srv.log = append(srv.log,command)
//		srv.commitIndex=srv.opNo
//	//	index=srv.commitIndex
//		ok=true
//	}
//	ok=true
//	index=srv.opNo
//	return index, view, ok
//}

// exmple code to send an AppendEntries RPC to a server.
// server is the index of the target server in srv.peers[].
// expects RPC arguments in args.
// The RPC library fills in *reply with RPC reply, so caller should pass &reply.
// the types of the args and reply passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// The labrpc package simulates a lossy network, in which servers
// may be unreachable, and in which requests and replies may be lost.
// Call() sends a request and waits for a reply. If a reply arrives
// within a timeout interval, Call() returns true; otherwise
// Call() returns false. Thus Call() may not return for a while.
// A false return can be caused by a dead server, a live server that
// can't be reached, a lost request, or a lost reply.
func (srv *server) sendPrepare(server int, args *pb.PrepareArgs, reply *pb.PrepareReply) bool {
	ok := srv.peers[server].Call("server.Prepare", args, reply)
	return ok
}

// Prepare is the RPC handler for the Prepare RPC
//func (srv *server) Prepare(args *pb.PrepareArgs) (reply *pb.PrepareReply, err error) {
//	srv.mu.Lock()
//	defer srv.mu.Unlock()
//	reply.View=int32(srv.currentView)
//	reply.Success=false;
//	if(int(args.View) < srv.currentView){
//		return
//	}
//        if(int(args.Index)<=srv.commitIndex){
//		return
//	}
//	if(int(args.PrimaryCommit)>srv.commitIndex){
//		srv.commitIndex=int(args.PrimaryCommit)
//	}
//	if(int(args.Index)!=srv.opNo+1||int(args.View)>srv.currentView){
//		fmt.Println("Debug: Server needs to recover")
//		log.Fatal("Debug: Server needs to recover")
//		//srv.status = RECOVERING
//		//PrimaryIndex:=GetPrimary(int(args.View), len(srv.peers))
//		//RecoveryInArgs:=pb.RecoveryArgs{
//		//	View:args.View,
//		//	Server:int32(srv.me),
//		//}
//		////RecoveryoutArgs:=RecoveryReply{
//		////	View:0,
//		////	Entries:srv.log,
//		////	PrimaryCommit:srv.commitIndex,
//		////	Success:false,
//		////}
//		//ok := srv.peers[PrimaryIndex].Call("server.Recovery",&RecoveryInArgs,&RecoveryoutArgs)
//		//if(ok==true) {
//		//	if(RecoveryoutArgs.Success){
//		//		srv.log=RecoveryoutArgs.Entries
//		//		srv.commitIndex=RecoveryoutArgs.PrimaryCommit
//		//		srv.currentView=RecoveryoutArgs.View
//		//		srv.status = NORMAL
//		//		srv.opNo=len(srv.log)-1
//		//		srv.commitIndex=args.PrimaryCommit
//		//		reply.Success=true
//		//		return
//		//	}else{
//		//
//		//	}
//		//}
//		//return
//	}
//	//srv.commitIndex=args.PrimaryCommit
//	if(int(args.Index)==len(srv.log)){
//		srv.log = append(srv.log,args.Entry)
//		srv.opNo=srv.opNo+1
//		srv.commitIndex=int(args.PrimaryCommit)
//		reply.Success=true
//		return
//	}
//	return
//
//}

// Recovery is the RPC handler for the Recovery RPC
//func (srv *server) Recovery(args *pb.RecoveryArgs) (reply *pb.RecoveryReply, err error) {
//	// Your code here
//	reply.View=int32(srv.currentView)
//	reply.Entries=srv.log
//	reply.PrimaryCommit=int32(srv.commitIndex)
//	reply.Success=true
//	return
//}

// Some external oracle prompts the primary of the newView to
// switch to the newView.
// PromptViewChange just kicks start the view change protocol to move to the newView
// It does not block waiting for the view change process to complete.
func (srv *server) PromptViewChange(newView int) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	newPrimary := GetPrimary(newView, len(srv.peers))

	if newPrimary != srv.me { //only primary of newView should do view change
		return
	} else if newView <= srv.currentView {
		return
	}
	vcArgs := &ViewChangeArgs{
		View: newView,
	}
	vcReplyChan := make(chan *ViewChangeReply, len(srv.peers))
	// send ViewChange to all servers including myself
	for i := 0; i < len(srv.peers); i++ {
		go func(server int) {
			var reply ViewChangeReply
			ok := srv.peers[server].Call("server.ViewChange", vcArgs, &reply)
			// fmt.Printf("node-%d (nReplies %d) received reply ok=%v reply=%v\n", srv.me, nReplies, ok, r.reply)
			if ok {
				vcReplyChan <- &reply
			} else {
				vcReplyChan <- nil
			}
		}(i)
	}

	// wait to receive ViewChange replies
	// if view change succeeds, send StartView RPC
	go func() {
		var successReplies []*ViewChangeReply
		var nReplies int
		majority := len(srv.peers)/2 + 1
		for r := range vcReplyChan {
			nReplies++
			if r != nil && r.Success {
				successReplies = append(successReplies, r)
			}
			if nReplies == len(srv.peers) || len(successReplies) == majority {
				break
			}
		}
		ok, log := srv.determineNewViewLog(successReplies)
		if !ok {
			return
		}
		svArgs := &StartViewArgs{
			View: vcArgs.View,
			Log:  log,
		}
		// send StartView to all servers including myself
		for i := 0; i < len(srv.peers); i++ {
			var reply StartViewReply
			go func(server int) {
				//fmt.Printf("Debug: node-%d sending StartView v=%d to node-%d\n", srv.me, svArgs.View, server)
				srv.peers[server].Call("server.StartView", svArgs, &reply)
			}(i)
		}
	}()
}

// determineNewViewLog is invoked to determine the log for the newView based on
// the collection of replies for successful ViewChange requests.
// if a quorum of successful replies exist, then ok is set to true.
// otherwise, ok = false.
func (srv *server) determineNewViewLog(successReplies []*ViewChangeReply) (
	ok bool, newViewLog []interface{}) {
	// Your code here
	lenSucess:=len(successReplies)
	Majority:=(len(srv.peers)-1)/2+1
	if(lenSucess<Majority){
		ok=false
		return
	}
	Index:=0
	MaxView:=0
	MaxLength:=0
	for i,reply :=  range successReplies{
		if(reply.LastNormalView>MaxView){
			Index=i
			MaxView=reply.LastNormalView
			MaxLength=len(reply.Log)
		}
		if(reply.LastNormalView==MaxView && len(reply.Log)>MaxLength){
			Index=i
			MaxView=reply.LastNormalView
			MaxLength=len(reply.Log)
		}
	}
	newViewLog=successReplies[Index].Log
	ok=true
	return ok, newViewLog
}

// ViewChange is the RPC handler to process ViewChange RPC.
func (srv *server) ViewChange(args *ViewChangeArgs, reply *ViewChangeReply) {
	// Your code here
	if(args.View<=srv.currentView){
		reply.Success=false
		return
	}
	reply.LastNormalView=srv.currentView
	reply.Log=srv.log
	reply.Success=true
	srv.lastNormalView=srv.currentView
	srv.currentView=args.View
	srv.status=VIEWCHANGE
	return
}

// StartView is the RPC handler to process StartView RPC.
func (srv *server) StartView(args *StartViewArgs, reply *StartViewReply) {
	if(srv.currentView>args.View){
		return
	}
	srv.currentView=args.View
	srv.log=args.Log
	srv.status=NORMAL
	srv.opNo=len(srv.log)-1
	return

}
