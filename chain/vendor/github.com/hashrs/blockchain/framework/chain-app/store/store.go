package store

import (
	dbm "github.com/hashrs/blockchain/libs/state-db/tm-db"

	"github.com/hashrs/blockchain/framework/chain-app/store/cache"
	"github.com/hashrs/blockchain/framework/chain-app/store/rootmulti"
	"github.com/hashrs/blockchain/framework/chain-app/store/types"
)

// Pruning strategies that may be provided to a KVStore to enable pruning.
const (
	PruningStrategyNothing    = "nothing"
	PruningStrategyEverything = "everything"
	PruningStrategySyncable   = "syncable"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}

func NewPruningOptionsFromString(strategy string) (opt PruningOptions) {
	switch strategy {
	case PruningStrategyNothing:
		opt = PruneNothing
	case PruningStrategyEverything:
		opt = PruneEverything
	case PruningStrategySyncable:
		opt = PruneSyncable
	default:
		opt = PruneSyncable
	}
	return
}
