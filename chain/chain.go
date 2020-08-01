package chain

import (
	"strings"

	abci "github.com/hashrs/blockchain/core/consensus/dpos-pbft/abci/types"
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/libs/log"
	sdk "github.com/hashrs/blockchain/framework/chain-app/types"
	starter "github.com/hashrs/blockchain/framework/chain-starter"
	dbm "github.com/hashrs/blockchain/libs/state-db/tm-db"

	//import greeter types
	main_net "github.com/hashrs/blockchain/chain/x/greeter"
	cli "github.com/hashrs/blockchain/framework/chain-cli"
)

const ChainName = "HashRs"

var (
	// ModuleBasics holds the AppModuleBasic struct of all modules included in the app
	ModuleBasics = starter.ModuleBasics
)

// Add the keeper and its key to our app struct
type ChainApp struct {
	*starter.AppStarter                 // helloChainApp extends starter.AppStarter
	greeterKey          *sdk.KVStoreKey // the store key for the greeter module
	greeterKeeper       main_net.Keeper // the keeper for the greeter module
}

// NewHelloChainApp returns a fully constructed SDK application
func NewChain(logger log.Logger, db dbm.DB) abci.Application {

	// pass greeter's AppModuleBasic to be included in the ModuleBasicsManager
	appStarter := starter.NewAppStarter(ChainName, logger, db, main_net.AppModuleBasic{})

	// create the key for greeter's store
	greeterKey := sdk.NewKVStoreKey(main_net.StoreKey)

	// construct the keeper
	greeterKeeper := main_net.NewKeeper(greeterKey, appStarter.Cdc)

	// compose our app with greeter
	var app = &ChainApp{
		appStarter,
		greeterKey,
		greeterKeeper,
	}

	// Add greeters' complete AppModule to the ModuleManager
	greeterMod := main_net.NewAppModule(greeterKeeper)
	app.Mm.Modules[greeterMod.Name()] = greeterMod

	// create a subspace for greeter's data in the main store.
	app.MountStore(greeterKey, sdk.StoreTypeDB)

	// do some final configuration...
	app.InitializeStarter()

	return app
}

func NewNode() {
	params := starter.NewServerCommandParams(
		strings.ToLower(ChainName)+"-node",
		ChainName+"Chain Node",
		starter.NewAppCreator(NewChain),
		starter.NewAppExporter(NewChain),
	)

	serverCmd := starter.NewServerCommand(params)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(serverCmd, "HR", starter.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func NewCli() {
	starter.BuildModuleBasics(ChainName, main_net.AppModuleBasic{})

	rootCmd := starter.NewCLICommand(ChainName)

	txCmd := starter.TxCmd(starter.Cdc)
	queryCmd := starter.QueryCmd(starter.Cdc)

	// add more Tx and Query commands
	ModuleBasics.AddTxCommands(txCmd, starter.Cdc)
	ModuleBasics.AddQueryCommands(queryCmd, starter.Cdc)
	rootCmd.AddCommand(txCmd, queryCmd)

	executor := cli.PrepareMainCmd(rootCmd, "HR", starter.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
