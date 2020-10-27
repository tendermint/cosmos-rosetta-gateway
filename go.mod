module github.com/tendermint/cosmos-rosetta-gateway

go 1.14

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d // indirect
	github.com/coinbase/rosetta-sdk-go v0.3.4
	github.com/cosmos/cosmos-sdk v0.34.4-0.20201026150813-bd3a29bdc1b5
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/google/go-cmp v0.5.0
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0 // indirect
	github.com/spf13/afero v1.4.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/iavl v0.14.0 // indirect
	github.com/tendermint/starport v0.0.11-0.20200924143929-4cbe2b57c65c
	github.com/tendermint/tendermint v0.34.0-rc5
	github.com/vektra/mockery/v2 v2.2.1
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/tools v0.0.0-20200921210052-fa0125251cc4 // indirect
	gopkg.in/ini.v1 v1.61.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
