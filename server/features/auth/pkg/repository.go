package pkg

import pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"

// Repository defines the methods that any concrete implementation of this
// repository must satisfy.
type Repository interface {
	Login(req *pb.LoginRequest) (*pb.LoginResponse, error)
	// @todo -> add other methods here
}
