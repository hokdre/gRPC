package unary_service_test

import (
	"context"
	"testing"

	"github.com/gRPC/pb"
	"github.com/gRPC/sample"
	"github.com/gRPC/store"
	unary_service "github.com/gRPC/unary"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestLaptopServer(t *testing.T) {
	t.Parallel()

	laptop := sample.NewLaptop()

	laptopWithNoId := sample.NewLaptop()
	laptopWithNoId.Id = ""

	laptopWithInvalidID := sample.NewLaptop()
	laptopWithNoId.Id = "abcd"

	duplicateLaptop := sample.NewLaptop()
	inMemoryStorage := store.NewInMemoryLaptopStore()
	if err := inMemoryStorage.Save(duplicateLaptop); err != nil {
		t.Fatalf("cannot seed duplicate laptop test : %s", err)
		return
	}

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  store.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: laptop,
			store:  store.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopWithNoId,
			store:  store.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "error_invalid_uuid_id",
			laptop: laptopWithInvalidID,
			store:  store.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "error_not_unique_uuid_id",
			laptop: duplicateLaptop,
			store:  store.NewInMemoryLaptopStore(),
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := unary_service.NewServer(tc.store)
			res, err := server.CreateLaptop(context.Background(), &req)
			if tc.code == codes.OK {
				require.Nil(err)
			} else {

			}
		})
	}

}
