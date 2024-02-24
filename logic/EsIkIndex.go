package logic

import (
	"context"
	"server/es"
	"server/proto"
)

func (s *ServerRpc) EsIkIndex(ctx context.Context, in *proto.EsIkIndexReq) (*proto.EsIkIndexRes, error) {
	var rr proto.EsIkIndexRes
	err := es.EsIkIndex(in.Index)
	if err != nil {
		return &rr, err
	}
	return &rr, nil
}
