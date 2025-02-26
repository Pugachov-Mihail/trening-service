package trening

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	trening_v1 "trening/protos/gen/trening.v1"
)

type TreningServiceApi struct {
	trening_v1.UnimplementedTreningServiceServer
	trening Trening
}

type Trening interface {
	TreningsListService(ctx context.Context, page, offset int32) ([]*trening_v1.GetTreningList, error)
	AddTreningsUserService(ctx context.Context, idTrening, userId int64) ([]*trening_v1.GetTreningList, error)
	GetUserTreningService(ctx context.Context, userId int64) ([]*trening_v1.GetTreningList, error)
	DeletedTreningUserService(ctx context.Context, idTrening, userId int64) ([]*trening_v1.GetTreningList, error)
	GetCurrentTreningService(ctx context.Context, idTrening, id int64) (trening_v1.GetTreningList, error)
}

func RegisterTreningServiceApi(grpc *grpc.Server, trening Trening) {
	trening_v1.RegisterTreningServiceServer(grpc, &TreningServiceApi{trening: trening})
}

func (t *TreningServiceApi) GetTreningList(ctx context.Context, req *trening_v1.GetTreningListRequest) (*trening_v1.GetTreningListResponse, error) {
	ctx = context.Background()

	var limit int32

	if req.GetPage() == 0 {
		limit = 10
	} else {
		limit = req.GetPage()
	}

	treningList, err := t.trening.TreningsListService(ctx, limit, req.GetOffset())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("GetTreningList: %v", err))
	}

	return &trening_v1.GetTreningListResponse{TreningList: treningList}, nil
}

func (t *TreningServiceApi) AddTreningUser(ctx context.Context, req *trening_v1.AddTreningRequest) (*trening_v1.GetTreningListResponse, error) {
	if req.GetUserId() == 0 || req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "UserId is required")
	}

	res, err := t.trening.AddTreningsUserService(ctx, req.GetUserId(), req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("AddTreningUser: %v", err))
	}

	return &trening_v1.GetTreningListResponse{TreningList: res}, nil
}

func (t *TreningServiceApi) GetUserTrenings(ctx context.Context, req *trening_v1.AddTreningRequest) (*trening_v1.GetTreningListResponse, error) {
	if req.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "UserId is required")
	}
	res, err := t.trening.GetUserTreningService(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("GetUserTreningService: %v", err))
	}
	return &trening_v1.GetTreningListResponse{TreningList: res}, nil
}

func (t *TreningServiceApi) DeletedTreningUser(ctx context.Context, req *trening_v1.AddTreningRequest) (*trening_v1.GetTreningListResponse, error) {
	if req.GetUserId() == 0 || req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "UserId is required")
	}

	res, err := t.trening.DeletedTreningUserService(ctx, req.GetId(), req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("DeletedTreningUser: %v", err))
	}
	return &trening_v1.GetTreningListResponse{TreningList: res}, nil
}

func (t *TreningServiceApi) GetCurrentTrening(ctx context.Context, req *trening_v1.TreningIdRequest) (*trening_v1.GetCurrentTreningResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "ProgrammId is required")
	}

	res, err := t.trening.GetCurrentTreningService(ctx, req.GetId(), req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("GetCurrentTreningService: %v", err))
	}

	return &trening_v1.GetCurrentTreningResponse{
		Id:           res.Id,
		Title:        res.Title,
		Image:        res.Image,
		Description:  res.Description,
		Raiting:      res.Raiting,
		TrenningInfo: res.TrenningInfo,
		TrenerInfo:   res.TrenerInfo,
	}, nil
}
