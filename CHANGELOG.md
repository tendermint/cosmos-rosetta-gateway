<a name="unreleased"></a>
## [Unreleased]


<a name="v1.0.0"></a>
## [v1.0.0] - 2021-05-11
### Feat 
- remove cli support
- upgrade to rosetta sdk v0.6.10

:warning: last release! see https://github.com/cosmos/cosmos-sdk/issues/9300

<a name="v0.3.0-rc2"></a>
## [v0.3.0-rc2] - 2021-01-14
### Fix
- offline network options returning nil


<a name="v0.3.0-rc1"></a>
## [v0.3.0-rc1] - 2021-01-04
### Fix
- errors code


<a name="v0.3.0-rc0"></a>
## [v0.3.0-rc0] - 2020-12-22
### Add
- ros errors Is implementation
- construction derive and hash
- finalize refactor
- refactor demo

### Change
- finalize api
- interfaces API
- use handler to enable retries

### Chore
- lint
- update docs

### Ci
- run unit tests

### Fix
- sleep before retrying node readiness

### Launchpad
- generate Tendermint client
- fix Tendermint OpenAPI def
- add openapi for tendermint

### Makefile
- add dev script

### Remove
- old implementation files

### Scripts
- config mock gen for Tendermint client

### Pull Requests
- Merge pull request [#97](https://github.com/tendermint/cosmos-rosetta-gateway/issues/97) from tendermint/alessio/simplify
- Merge pull request [#93](https://github.com/tendermint/cosmos-rosetta-gateway/issues/93) from tendermint/sahith/update-sdk-swagger
- Merge pull request [#90](https://github.com/tendermint/cosmos-rosetta-gateway/issues/90) from tendermint/sahith/update-payloads
- Merge pull request [#85](https://github.com/tendermint/cosmos-rosetta-gateway/issues/85) from tendermint/sahith/address-broadcast
- Merge pull request [#88](https://github.com/tendermint/cosmos-rosetta-gateway/issues/88) from tendermint/sahith/fix-metadata
- Merge pull request [#84](https://github.com/tendermint/cosmos-rosetta-gateway/issues/84) from tendermint/feature/cp-signed
- Merge pull request [#81](https://github.com/tendermint/cosmos-rosetta-gateway/issues/81) from tendermint/sahith/implement-combine-api
- Merge pull request [#80](https://github.com/tendermint/cosmos-rosetta-gateway/issues/80) from tendermint/jonathan/issue-43-implement-construction-hash
- Merge pull request [#79](https://github.com/tendermint/cosmos-rosetta-gateway/issues/79) from tendermint/sahith/add-metadata
- Merge pull request [#57](https://github.com/tendermint/cosmos-rosetta-gateway/issues/57) from tendermint/sahith/add-cons-parse
- Merge pull request [#76](https://github.com/tendermint/cosmos-rosetta-gateway/issues/76) from tendermint/jonathan/issue-41-implement-construction-payloads
- Merge pull request [#78](https://github.com/tendermint/cosmos-rosetta-gateway/issues/78) from tendermint/sahith/add-preprocess
- Merge pull request [#72](https://github.com/tendermint/cosmos-rosetta-gateway/issues/72) from tendermint/fix/block-reverted
- Merge pull request [#73](https://github.com/tendermint/cosmos-rosetta-gateway/issues/73) from tendermint/jonathan/issue-66-block-id-balance
- Merge pull request [#70](https://github.com/tendermint/cosmos-rosetta-gateway/issues/70) from tendermint/fix/block-op
- Merge pull request [#71](https://github.com/tendermint/cosmos-rosetta-gateway/issues/71) from tendermint/feature/codeowners
- Merge pull request [#64](https://github.com/tendermint/cosmos-rosetta-gateway/issues/64) from tendermint/sahith/validate-curve
- Merge pull request [#62](https://github.com/tendermint/cosmos-rosetta-gateway/issues/62) from tendermint/fix/data-blocks
- Merge pull request [#63](https://github.com/tendermint/cosmos-rosetta-gateway/issues/63) from tendermint/remove-launchpad-gomod
- Merge pull request [#58](https://github.com/tendermint/cosmos-rosetta-gateway/issues/58) from tendermint/rethink-supported-operations
- Merge pull request [#56](https://github.com/tendermint/cosmos-rosetta-gateway/issues/56) from tendermint/replace-version
- Merge pull request [#54](https://github.com/tendermint/cosmos-rosetta-gateway/issues/54) from tendermint/mempool-transaction
- Merge pull request [#55](https://github.com/tendermint/cosmos-rosetta-gateway/issues/55) from tendermint/feature/construction-derive
- Merge pull request [#53](https://github.com/tendermint/cosmos-rosetta-gateway/issues/53) from tendermint/sahith/add-ci-lint
- Merge pull request [#51](https://github.com/tendermint/cosmos-rosetta-gateway/issues/51) from tendermint/formatting
- Merge pull request [#48](https://github.com/tendermint/cosmos-rosetta-gateway/issues/48) from tendermint/alessio/codecoverage
- Merge pull request [#50](https://github.com/tendermint/cosmos-rosetta-gateway/issues/50) from tendermint/fix-alessio
- Merge pull request [#49](https://github.com/tendermint/cosmos-rosetta-gateway/issues/49) from tendermint/fix/block-test
- Merge pull request [#31](https://github.com/tendermint/cosmos-rosetta-gateway/issues/31) from tendermint/jonathan/mempool
- Merge pull request [#37](https://github.com/tendermint/cosmos-rosetta-gateway/issues/37) from tendermint/feature/api-block-tx
- Merge pull request [#34](https://github.com/tendermint/cosmos-rosetta-gateway/issues/34) from tendermint/feature/api-block
- Merge pull request [#35](https://github.com/tendermint/cosmos-rosetta-gateway/issues/35) from tendermint/jonathan/little-tweak-api-block
- Merge pull request [#33](https://github.com/tendermint/cosmos-rosetta-gateway/issues/33) from tendermint/jonathan/makefile-and-cosmosfolder
- Merge pull request [#32](https://github.com/tendermint/cosmos-rosetta-gateway/issues/32) from tendermint/refactor/move-cosmos-root
- Merge pull request [#30](https://github.com/tendermint/cosmos-rosetta-gateway/issues/30) from tendermint/refactor/use-generated-tendermint
- Merge pull request [#29](https://github.com/tendermint/cosmos-rosetta-gateway/issues/29) from tendermint/refactor/use-generated-client
- Merge pull request [#28](https://github.com/tendermint/cosmos-rosetta-gateway/issues/28) from tendermint/feature/codegen
- Merge pull request [#26](https://github.com/tendermint/cosmos-rosetta-gateway/issues/26) from tendermint/add-test-peers
- Merge pull request [#24](https://github.com/tendermint/cosmos-rosetta-gateway/issues/24) from tendermint/feature/network-status
- Merge pull request [#22](https://github.com/tendermint/cosmos-rosetta-gateway/issues/22) from tendermint/feature/network-options
- Merge pull request [#23](https://github.com/tendermint/cosmos-rosetta-gateway/issues/23) from tendermint/rethink-options
- Merge pull request [#14](https://github.com/tendermint/cosmos-rosetta-gateway/issues/14) from tendermint/build-networks-endpoint
- Merge pull request [#15](https://github.com/tendermint/cosmos-rosetta-gateway/issues/15) from tendermint/feature/ci-test
- Merge pull request [#11](https://github.com/tendermint/cosmos-rosetta-gateway/issues/11) from tendermint/jonathan/interface-rosetta
- Merge pull request [#12](https://github.com/tendermint/cosmos-rosetta-gateway/issues/12) from tendermint/fix/onterface-rosetta-tweaks
- Merge pull request [#10](https://github.com/tendermint/cosmos-rosetta-gateway/issues/10) from tendermint/jonathan/interface-rosetta
- Merge pull request [#8](https://github.com/tendermint/cosmos-rosetta-gateway/issues/8) from tendermint/feature/project-structure
- Merge pull request [#9](https://github.com/tendermint/cosmos-rosetta-gateway/issues/9) from tendermint/alt-structure
- Merge pull request [#5](https://github.com/tendermint/cosmos-rosetta-gateway/issues/5) from tendermint/feature/go-mod


<a name="v0.2.0"></a>
## [v0.2.0] - 2020-11-23

<a name="v0.1.1"></a>
## [v0.1.1] - 2020-10-15

<a name="v0.1.0"></a>
## v0.1.0 - 2020-09-30
### Ci
- run unit tests

### Launchpad
- generate Tendermint client
- fix Tendermint OpenAPI def
- add openapi for tendermint

### Makefile
- add dev script

### Scripts
- config mock gen for Tendermint client

### Pull Requests
- Merge pull request [#97](https://github.com/tendermint/cosmos-rosetta-gateway/issues/97) from tendermint/alessio/simplify
- Merge pull request [#93](https://github.com/tendermint/cosmos-rosetta-gateway/issues/93) from tendermint/sahith/update-sdk-swagger
- Merge pull request [#90](https://github.com/tendermint/cosmos-rosetta-gateway/issues/90) from tendermint/sahith/update-payloads
- Merge pull request [#85](https://github.com/tendermint/cosmos-rosetta-gateway/issues/85) from tendermint/sahith/address-broadcast
- Merge pull request [#88](https://github.com/tendermint/cosmos-rosetta-gateway/issues/88) from tendermint/sahith/fix-metadata
- Merge pull request [#84](https://github.com/tendermint/cosmos-rosetta-gateway/issues/84) from tendermint/feature/cp-signed
- Merge pull request [#81](https://github.com/tendermint/cosmos-rosetta-gateway/issues/81) from tendermint/sahith/implement-combine-api
- Merge pull request [#80](https://github.com/tendermint/cosmos-rosetta-gateway/issues/80) from tendermint/jonathan/issue-43-implement-construction-hash
- Merge pull request [#79](https://github.com/tendermint/cosmos-rosetta-gateway/issues/79) from tendermint/sahith/add-metadata
- Merge pull request [#57](https://github.com/tendermint/cosmos-rosetta-gateway/issues/57) from tendermint/sahith/add-cons-parse
- Merge pull request [#76](https://github.com/tendermint/cosmos-rosetta-gateway/issues/76) from tendermint/jonathan/issue-41-implement-construction-payloads
- Merge pull request [#78](https://github.com/tendermint/cosmos-rosetta-gateway/issues/78) from tendermint/sahith/add-preprocess
- Merge pull request [#72](https://github.com/tendermint/cosmos-rosetta-gateway/issues/72) from tendermint/fix/block-reverted
- Merge pull request [#73](https://github.com/tendermint/cosmos-rosetta-gateway/issues/73) from tendermint/jonathan/issue-66-block-id-balance
- Merge pull request [#70](https://github.com/tendermint/cosmos-rosetta-gateway/issues/70) from tendermint/fix/block-op
- Merge pull request [#71](https://github.com/tendermint/cosmos-rosetta-gateway/issues/71) from tendermint/feature/codeowners
- Merge pull request [#64](https://github.com/tendermint/cosmos-rosetta-gateway/issues/64) from tendermint/sahith/validate-curve
- Merge pull request [#62](https://github.com/tendermint/cosmos-rosetta-gateway/issues/62) from tendermint/fix/data-blocks
- Merge pull request [#63](https://github.com/tendermint/cosmos-rosetta-gateway/issues/63) from tendermint/remove-launchpad-gomod
- Merge pull request [#58](https://github.com/tendermint/cosmos-rosetta-gateway/issues/58) from tendermint/rethink-supported-operations
- Merge pull request [#56](https://github.com/tendermint/cosmos-rosetta-gateway/issues/56) from tendermint/replace-version
- Merge pull request [#54](https://github.com/tendermint/cosmos-rosetta-gateway/issues/54) from tendermint/mempool-transaction
- Merge pull request [#55](https://github.com/tendermint/cosmos-rosetta-gateway/issues/55) from tendermint/feature/construction-derive
- Merge pull request [#53](https://github.com/tendermint/cosmos-rosetta-gateway/issues/53) from tendermint/sahith/add-ci-lint
- Merge pull request [#51](https://github.com/tendermint/cosmos-rosetta-gateway/issues/51) from tendermint/formatting
- Merge pull request [#48](https://github.com/tendermint/cosmos-rosetta-gateway/issues/48) from tendermint/alessio/codecoverage
- Merge pull request [#50](https://github.com/tendermint/cosmos-rosetta-gateway/issues/50) from tendermint/fix-alessio
- Merge pull request [#49](https://github.com/tendermint/cosmos-rosetta-gateway/issues/49) from tendermint/fix/block-test
- Merge pull request [#31](https://github.com/tendermint/cosmos-rosetta-gateway/issues/31) from tendermint/jonathan/mempool
- Merge pull request [#37](https://github.com/tendermint/cosmos-rosetta-gateway/issues/37) from tendermint/feature/api-block-tx
- Merge pull request [#34](https://github.com/tendermint/cosmos-rosetta-gateway/issues/34) from tendermint/feature/api-block
- Merge pull request [#35](https://github.com/tendermint/cosmos-rosetta-gateway/issues/35) from tendermint/jonathan/little-tweak-api-block
- Merge pull request [#33](https://github.com/tendermint/cosmos-rosetta-gateway/issues/33) from tendermint/jonathan/makefile-and-cosmosfolder
- Merge pull request [#32](https://github.com/tendermint/cosmos-rosetta-gateway/issues/32) from tendermint/refactor/move-cosmos-root
- Merge pull request [#30](https://github.com/tendermint/cosmos-rosetta-gateway/issues/30) from tendermint/refactor/use-generated-tendermint
- Merge pull request [#29](https://github.com/tendermint/cosmos-rosetta-gateway/issues/29) from tendermint/refactor/use-generated-client
- Merge pull request [#28](https://github.com/tendermint/cosmos-rosetta-gateway/issues/28) from tendermint/feature/codegen
- Merge pull request [#26](https://github.com/tendermint/cosmos-rosetta-gateway/issues/26) from tendermint/add-test-peers
- Merge pull request [#24](https://github.com/tendermint/cosmos-rosetta-gateway/issues/24) from tendermint/feature/network-status
- Merge pull request [#22](https://github.com/tendermint/cosmos-rosetta-gateway/issues/22) from tendermint/feature/network-options
- Merge pull request [#23](https://github.com/tendermint/cosmos-rosetta-gateway/issues/23) from tendermint/rethink-options
- Merge pull request [#14](https://github.com/tendermint/cosmos-rosetta-gateway/issues/14) from tendermint/build-networks-endpoint
- Merge pull request [#15](https://github.com/tendermint/cosmos-rosetta-gateway/issues/15) from tendermint/feature/ci-test
- Merge pull request [#11](https://github.com/tendermint/cosmos-rosetta-gateway/issues/11) from tendermint/jonathan/interface-rosetta
- Merge pull request [#12](https://github.com/tendermint/cosmos-rosetta-gateway/issues/12) from tendermint/fix/onterface-rosetta-tweaks
- Merge pull request [#10](https://github.com/tendermint/cosmos-rosetta-gateway/issues/10) from tendermint/jonathan/interface-rosetta
- Merge pull request [#8](https://github.com/tendermint/cosmos-rosetta-gateway/issues/8) from tendermint/feature/project-structure
- Merge pull request [#9](https://github.com/tendermint/cosmos-rosetta-gateway/issues/9) from tendermint/alt-structure
- Merge pull request [#5](https://github.com/tendermint/cosmos-rosetta-gateway/issues/5) from tendermint/feature/go-mod


[Unreleased]: https://github.com/tendermint/cosmos-rosetta-gateway/compare/v1.0.0...HEAD
[v1.0.0]: https://github.com/tendermint/cosmos-rosetta-gateway/compare/v0.3.0-rc2...v1.0.0
[v0.3.0-rc2]: https://github.com/tendermint/cosmos-rosetta-gateway/compare/v0.3.0-rc1...v0.3.0-rc2
[v0.3.0-rc1]: https://github.com/tendermint/cosmos-rosetta-gateway/compare/v0.3.0-rc0...v0.3.0-rc1
[v0.3.0-rc0]: https://github.com/tendermint/cosmos-rosetta-gateway/compare/v0.2.0...v0.3.0-rc0
[v0.2.0]: https://github.com/tendermint/cosmos-rosetta-gateway/compare/v0.1.1...v0.2.0
[v0.1.1]: https://github.com/tendermint/cosmos-rosetta-gateway/compare/v0.1.0...v0.1.1
