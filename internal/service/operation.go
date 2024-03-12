package service

import (
	"context"

	pb "review-o/api/operation/v1"
	"review-o/internal/biz"
)

type OperationService struct {
	pb.UnimplementedOperationServer
	uc *biz.OperationUsecase
}

// 这里 这里要传入一个参数才能调用 不然wire无法识别 因为我们自定义的service中绑定了这个结构体
func NewOperationService(uc *biz.OperationUsecase) *OperationService {
	return &OperationService{uc: uc} //这里一定要 返回值
}

func (s *OperationService) AuditReview(ctx context.Context, req *pb.AuditReviewRequest) (*pb.AuditReviewReply, error) {
	return &pb.AuditReviewReply{}, nil
}

// AuditAppeal O端审核评价申述

func (s *OperationService) AuditAppeal(ctx context.Context, req *pb.AuditAppealRequest) (*pb.AuditAppealReply, error) {
	err := s.uc.AuditAppeal(ctx, &biz.AuditAppealParam{
		AppealID:  req.GetAppealID(),
		ReviewID:  req.GetReviewID(),
		Status:    int(req.GetStatus()),
		OpReason:  req.GetOpReason(),
		OpRemarks: req.GetOpRemarks(),
		OpUser:    req.GetOpUser(),
	})
	return &pb.AuditAppealReply{}, err
}
