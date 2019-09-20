package repos


type BaseTableQuery interface {
	TableName() string
}

type BasePartitionKeyQuery interface {
	BaseTableQuery
	PartitionKeyName() string
	PartitionKey() string
}

type BaseUniqueKeyQuery interface {
	BasePartitionKeyQuery
	SortKeyName() string
	SortKey() string
}
