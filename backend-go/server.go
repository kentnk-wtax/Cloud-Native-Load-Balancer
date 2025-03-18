package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 1922
// Hash 2816
// Hash 6746
// Hash 9228
// Hash 8824
// Hash 5907
// Hash 7260
// Hash 4037
// Hash 7979
// Hash 9922
// Hash 9173
// Hash 3828
// Hash 2361
// Hash 8081
// Hash 4986
// Hash 3444
// Hash 1221
// Hash 7920
// Hash 4765
// Hash 2589
// Hash 5066
// Hash 4565
// Hash 7552
// Hash 2761
// Hash 4897
// Hash 8713
// Hash 3981
// Hash 1929
// Hash 6114
// Hash 2587
// Hash 8055
// Hash 1159
// Hash 6578
// Hash 8629
// Hash 5337
// Hash 2894
// Hash 1344
// Hash 7416
// Hash 1808
// Hash 3319
// Hash 4561
// Hash 9451
// Hash 7968
// Hash 9803
// Hash 6705
// Hash 4537
// Hash 6695
// Hash 3996
// Hash 4469
// Hash 3043
// Hash 7639
// Hash 4420
// Hash 2953
// Hash 8799
// Hash 1580
// Hash 2404
// Hash 4596
// Hash 1460
// Hash 5272
// Hash 2275
// Hash 7208
// Hash 3564
// Hash 2650