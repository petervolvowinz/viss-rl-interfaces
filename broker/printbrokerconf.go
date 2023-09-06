/******** Peter Winzell (c), 8/31/23 *********************************************/

package broker

import (
	"context"
	"fmt"
	"github.com/petervolvowinz/viss-rl-interfaces/base"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// print current configuration to the console
func PrintSignalTree(clientconnection *grpc.ClientConn) error {
	systemServiceClient := base.NewSystemServiceClient(clientconnection)
	md := metadata.Pairs(
		"x-api-key", "abcdefg",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	configuration, err := systemServiceClient.GetConfiguration(ctx, &base.Empty{})

	infos := configuration.GetNetworkInfo()
	for _, element := range infos {
		printSignals(element.Namespace.Name, clientconnection)
	}

	if err != nil {
		log.Debug("could not retrieve configuration ", err)
		return err
	}

	return nil
}

// print signal tree(s) to console , using fmt for this.
func printSpaces(number int) {
	for k := 1; k < number; k++ {
		fmt.Print(" ")
	}
}

func printTreeBranch() {
	fmt.Print("|")
}

func getFirstNameSpace(frames []*base.FrameInfo) string {
	element := frames[0]
	return element.SignalInfo.Id.Name
}

func printSignals(zenamespace string, clientconnection *grpc.ClientConn) {
	systemServiceClient := base.NewSystemServiceClient(clientconnection)
	md := metadata.Pairs(
		"x-api-key", "abcdefg",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	signallist, err := systemServiceClient.ListSignals(ctx, &base.NameSpace{Name: zenamespace})

	frames := signallist.GetFrame()

	rootstring := "|[" + zenamespace + "]---|"
	rootstringlength := len(rootstring)
	fmt.Println(rootstring)

	for _, element := range frames {

		printTreeBranch()
		printSpaces(rootstringlength - 1)

		framestring := "|---[" + element.SignalInfo.Id.Name + "]---|"
		framestringlength := len(framestring)

		fmt.Println(framestring)
		childs := element.ChildInfo

		for _, childelement := range childs {
			outstr := childelement.Id.Name
			printTreeBranch()
			printSpaces(rootstringlength - 1)
			printTreeBranch()
			printSpaces(framestringlength - 1)
			fmt.Println("|---{", outstr, "}")
		}
	}

	if err != nil {
		log.Debug(" could not list signals ", err)
	}
}
