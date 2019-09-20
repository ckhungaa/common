package repos

import (
	"context"
	"github.com/ckhungaa/common/utils/entities"
	"github.com/ckhungaa/common/utils/errs"
	"github.com/pkg/errors"
	"strings"
)

type PartitionDBKey struct {
	partitionKey string
	sortKey      string
}

// PartitionDBKeyFromAudit parse audit id into PartitionDBKey
func PartitionDBKeyFromAudit(ctx context.Context, audit *entities.Audit) (*PartitionDBKey, error) {
	return PartitionDBKeyFromId(ctx, audit.Id)
}

// PartitionDBKeyFromId parse id into PartitionDBKey, the id format should be '{partition_key}_{sort_key}'
func PartitionDBKeyFromId(ctx context.Context, id string) (*PartitionDBKey, error) {
	ids := strings.Split(id, "_")
	if len(ids) != 2 {
		return nil, errors.Wrapf(errs.InvalidId, "invalid id format:%s", id)
	}
	return NewPartitionDBKey(ctx, ids[0], ids[1]), nil
}

// NewPartitionDBKey constructor
func NewPartitionDBKey(ctx context.Context, partitionKey string, sortKey string) *PartitionDBKey {
	return &PartitionDBKey{
		partitionKey: partitionKey,
		sortKey:      sortKey,
	}
}

func (k *PartitionDBKey) PartitionKey() string {
	return k.partitionKey
}

func (k *PartitionDBKey) SortKey() string {
	return k.sortKey
}
