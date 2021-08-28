package handler

import (
	"context"
	"taskRestAPI/internal/app/service"
	pb "taskRestAPI/proto"
)

type CRUDHandler struct{
	pb.UnimplementedCRUDServiceServer
	service service.CRUDService
}

func (C CRUDHandler) mustEmbedUnimplementedCRUDServiceServer() {
	panic("implement me")
}

func NewCRUDHandler(service service.CRUDService) *CRUDHandler {
	return &CRUDHandler{service: service}
}

func (C CRUDHandler) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	return C.service.CreateUser(ctx, req)
}

func (C CRUDHandler) GetUserByUUID(ctx context.Context, req *pb.GetUserByUUIDReq) (*pb.GetUserByUUIDRes, error) {
	return C.service.GetUserByUUID(ctx, req)
}

func (C CRUDHandler) UpdateUserByUUID(ctx context.Context, req *pb.UpdateUserByUUIDReq) (*pb.UpdateUserByUUIDRes, error) {
	return C.service.UpdateUserByUUID(ctx, req)
}




