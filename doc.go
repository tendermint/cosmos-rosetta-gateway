// Package crg holds the source code of Rosetta API bindings for Cosmos SDK.
// it also provides utils and a service under service/ to expose Rosetta implementation
// to different interfaces (network layers, command line interfaces, etc.)
//
// project structure:
//
// .
// ├── doc.go
// ├── interface                      > expose service/ functionalities to different targets.
// │   ├── cli
// │   │   └── crg
// │   │       ├── cmd                > Gateway exposed as a cli command `crg`.
// │   │       │   └── gateway.go
// │   │       └── main.go            > Gateway exposed as standalone program.
// │   └── http                       > Gateway exposed as an HTTP service.
// │       └── http.go
// ├── pkg                            > reusable, generic utils.
// │   └── doc.go
// └── service                        > the actual Rosetta API implementation as high level Go funcs.
//	└── online.go

package crg
