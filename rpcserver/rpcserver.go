package rpcserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/gorilla/mux"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"golang.org/x/xerrors"
)

type StopFunc func(context.Context) error

// "github.com/filecoin-project/go-jsonrpc"

func ServerRPC(h http.Handler, id string, addr multiaddr.Multiaddr) (StopFunc, error) {

	// start listening to the address ;
	lst, err := manet.Listen(addr)
	if err != nil {
		return nil, xerrors.Errorf("could not listen: %w", err)
	}

	srv := &http.Server{
		Handler: h,
		// BaseContext: func(listener net.Listener) context.Context {
		// 	ctx,_ := tag.New(context.Background(),tag.Upsert(metrics.Ap))
		// }
	}

	go func() {
		err = srv.Serve(manet.NetListener(lst))
		if err != nil {
			fmt.Printf("rpc server failed :%v\n", err)
		}
	}()

	return srv.Shutdown, err
}

func FlutterHandler(fn FullNode) (http.Handler, error) {

	m := mux.NewRouter()

	//初始化一个jsonrpc server对象
	rpcServer := jsonrpc.NewServer()
	rpcServer.Register("Filecoin", fn)

	var handler http.Handler = rpcServer
	m.Handle("/rpc/v1", handler)

	return m, nil
}

type FullNode interface {
	TestHandler(ctx context.Context, input string) error
}

type ImplementFullNode struct {
	Test string
}

func (IFn *ImplementFullNode) TestHandler(ctx context.Context, input string) error {
	fmt.Println("...... input is:", input)
	return nil
}

var _ FullNode = (*ImplementFullNode)(nil)
