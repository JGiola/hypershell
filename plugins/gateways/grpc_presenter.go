package gateways

import (
	pb "github.com/openshift-online/hypershell/pkg/api/grpc/hypershell/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func gatewayToProto(d *Gateway) *pb.Gateway {
	return &pb.Gateway{
		Metadata: &pb.ObjectReference{
			Id:        d.ID,
			CreatedAt: timestamppb.New(d.CreatedAt),
			UpdatedAt: timestamppb.New(d.UpdatedAt),
			Kind:      "Gateway",
			Href:      "/api/hypershell/v1/gateways/" + d.ID,
		},
		Name:        d.Name,
		FleetId:     d.FleetId,
		ClusterId:   d.ClusterId,
		ReleaseId:   d.ReleaseId,
		DatabaseId:  d.DatabaseId,
		Namespace:   d.Namespace,
		ExternalDns: d.ExternalDns,
		TlsMode:     d.TlsMode,
		ServiceType: d.ServiceType,
		Status:      d.Status,
		Phase:       d.Phase,
	}
}
