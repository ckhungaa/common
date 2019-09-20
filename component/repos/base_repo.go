package repos

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ckhungaa/common/component/configs"
	"github.com/google/wire"
	"github.com/guregu/dynamo"
)

var (
	WireSet = wire.NewSet(
		ProvideConfig,
		ProvideBaseRepository,
	)
)

// TODO: add non-local config
// Config BaseRepository configuration
type Config struct {
	EndPoint string `config:"DYNAMO_DB_END_POINT" config_def:"http://localhost:4569"`
	Region   string `config:"DYNAMO_DB_REGION" config_def:"ap-southeast-1"`
}

// ProvideConfig service config provider
func ProvideConfig(ctx context.Context, configStore configs.Store) (*Config, error) {
	var result Config
	if err := configStore.GetConfig(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type BaseRepository struct {
	db *dynamo.DB
}

func ProvideBaseRepository(ctx context.Context, cnf *Config) (*BaseRepository, error) {
	db := dynamo.New(session.New(), &aws.Config{Endpoint: &cnf.EndPoint, Region: aws.String(cnf.Region)}) //TODO: fix when go prod
	return &BaseRepository{db: db}, nil
}

func (r *BaseRepository) DB() *dynamo.DB {
	return r.db
}

//TODO wrap aws error to our error
//FindById find unique object by BaseUniqueKeyQuery
func (r *BaseRepository) FindById(ctx context.Context, idQuery BaseUniqueKeyQuery, result interface{}) error {
	return r.DB().Table(idQuery.TableName()).Get(idQuery.PartitionKeyName(), idQuery.PartitionKey()).Range(idQuery.SortKeyName(), dynamo.Equal, idQuery.SortKey()).OneWithContext(ctx, result)
}

//FindById objects by BasePartitionKeyQuery
func (r *BaseRepository) FindAll(ctx context.Context, idQuery BasePartitionKeyQuery, result interface{}) error {
	return r.DB().Table(idQuery.TableName()).Get(idQuery.PartitionKeyName(), idQuery.PartitionKey()).AllWithContext(ctx, result)
}