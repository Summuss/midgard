// Copyright 2020-2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a GPL-3.0
// license that can be found in the LICENSE file.

package daemon

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"changkun.de/x/midgard/internal/config"
	"changkun.de/x/midgard/internal/office"
	"changkun.de/x/midgard/internal/types"
	"changkun.de/x/midgard/internal/types/proto"
	"changkun.de/x/midgard/internal/utils"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

// Daemon is the midgard daemon that interact with midgard server.
type Daemon struct {
	sync.Mutex

	ID          string
	status      office.Status
	forceUpdate chan struct{}
	s           *grpc.Server
	ws          *websocket.Conn
	readChs     sync.Map                     // {string: chan *types.WebsocketMessage}
	writeCh     chan *types.WebsocketMessage // writeCh is used for sending message along ws.

	proto.UnimplementedMidgardServer
}

// NewDaemon creates a new midgard daemon
func NewDaemon() *Daemon {
	id, err := os.Hostname()
	if err != nil {
		id, err = utils.NewUUIDShort()
		if err != nil {
			panic(fmt.Errorf("failed to initialize deamon: %v", err))
		}
	}
	return &Daemon{
		ID:          id,
		forceUpdate: make(chan struct{}, 1),
		writeCh:     make(chan *types.WebsocketMessage, 10),
	}
}

// Serve serves Daemon daemon:
// 1. maintaining midgard daemon rpc;
// 2. maintaining midgard daemon to server websocket.
func (m *Daemon) Serve() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			log.Println("graceful shutdown assistant is terminated.")
		}()
		q := make(chan os.Signal, 1)
		signal.Notify(q, os.Interrupt)
		sig := <-q
		log.Printf("%v", sig)
		log.Printf("shutting down midgard daemon ...")
		m.s.GracefulStop()
		cancel()
	}()
	go func() {
		defer log.Println("websocket is terminated.")
		m.wsConnect()
		m.handleIO(ctx)
		m.wsClose()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Println("clipboard watcher is terminated.")
		m.watchLocalClipboard(ctx)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Println("office watcher is terminated.")
		m.watchOfficeStatus(ctx)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Println("rpc server is terminated.")
		m.serveRPC()
	}()
	wg.Wait()

	log.Printf("daemon is down, good bye!")
}

const maxMessageSize = 10 << 20 // 10 MB

func (m *Daemon) serveRPC() {
	l, err := net.Listen("tcp", config.D().Addr)
	if err != nil {
		log.Fatalf("fail to initalize midgard daemon, err: %v", err)
	}

	m.s = grpc.NewServer(
		grpc.MaxRecvMsgSize(maxMessageSize),
		grpc.MaxSendMsgSize(maxMessageSize),
		grpc.ConnectionTimeout(time.Minute*5),
	)
	proto.RegisterMidgardServer(m.s, m)
	log.Printf("daemon running at rpc://%s", config.D().Addr)
	if err := m.s.Serve(l); err != nil {
		log.Fatalf("fail to serve midgard daemon, err: %v", err)
	}
}
