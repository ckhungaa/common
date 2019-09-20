package entities

import (
	"context"
	"github.com/ckhungaa/proto/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"time"
)

//Audit audit object for every entity
type Audit struct {
	Id           string
	CreatedDate  time.Time
	ModifiedDate time.Time
}

func NewAudit(ctx context.Context, id string) *Audit {
	return &Audit{
		Id:           id,
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
	}
}

func (m *Audit) Modified() {
	m.ModifiedDate = time.Now()
}

//ToProto to gRPC proto object
func (m *Audit) ToProto() (*proto.Audit, error) {
	createDate, err := ptypes.TimestampProto(m.CreatedDate)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid timestamp format:%s", m.CreatedDate)
	}
	modifiedDate, err := ptypes.TimestampProto(m.ModifiedDate)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid timestamp format:%s", m.ModifiedDate)
	}

	return &proto.Audit{
		Id:           m.Id,
		CreatedDate:  createDate,
		ModifiedDate: modifiedDate,
	}, nil
}
