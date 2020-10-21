module github.com/tendermint/cosmos-rosetta-gateway

go 1.14

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/coinbase/rosetta-sdk-go v0.3.4
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/google/go-cmp v0.4.0
	github.com/irisnet/irishub v0.16.3
	github.com/magiconair/properties v1.8.3 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	github.com/rs/zerolog v1.20.0 // indirect
	github.com/spf13/afero v1.4.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.31.0
	github.com/vektra/mockery/v2 v2.2.1
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/sys v0.0.0-20200922070232-aee5d888a860 // indirect
	golang.org/x/tools v0.0.0-20200921210052-fa0125251cc4 // indirect
	gopkg.in/ini.v1 v1.61.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace (
	github.com/tendermint/iavl => github.com/irisnet/iavl v0.12.3
	github.com/tendermint/tendermint => github.com/irisnet/tendermint v0.32.2
	golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
)
