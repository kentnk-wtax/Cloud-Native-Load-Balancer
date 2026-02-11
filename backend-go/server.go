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
// Hash 9586
// Hash 4799
// Hash 3530
// Hash 4753
// Hash 2864
// Hash 4973
// Hash 8250
// Hash 7697
// Hash 6114
// Hash 1526
// Hash 5810
// Hash 1222
// Hash 2486
// Hash 8596
// Hash 5975
// Hash 2103
// Hash 5654
// Hash 9790
// Hash 5415
// Hash 4828
// Hash 5609
// Hash 3592
// Hash 8249
// Hash 7377
// Hash 7533
// Hash 7413
// Hash 5351
// Hash 1883
// Hash 8564
// Hash 2819
// Hash 6382
// Hash 1734
// Hash 6002
// Hash 6817
// Hash 7329
// Hash 4962
// Hash 2902
// Hash 4723
// Hash 1819
// Hash 9438
// Hash 3785
// Hash 6598
// Hash 6205
// Hash 7062
// Hash 3143
// Hash 4840
// Hash 8987
// Hash 9394
// Hash 8071
// Hash 7877
// Hash 4717
// Hash 5021
// Hash 9556
// Hash 9602
// Hash 6728
// Hash 3412
// Hash 2072
// Hash 8580
// Hash 1217
// Hash 4392
// Hash 5900
// Hash 8111
// Hash 1850
// Hash 6995
// Hash 3077
// Hash 6048
// Hash 3217
// Hash 2987
// Hash 9536
// Hash 9957
// Hash 5694
// Hash 3684
// Hash 7255
// Hash 2250
// Hash 8857
// Hash 5972
// Hash 2820
// Hash 7069
// Hash 1307
// Hash 9909
// Hash 1473
// Hash 8672
// Hash 1218