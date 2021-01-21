module github.com/blocktree/hypercash-adapter

go 1.12

require (
	github.com/HcashOrg/bliss v0.0.0-20180719035130-f5d53c2a9b7d // indirect
	github.com/HcashOrg/hcd v0.0.0-20190611115530-593f7eea8b25
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // indirect
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet/v2 v2.0.9
	github.com/bndr/gotabulate v1.1.2
	github.com/codeskyblue/go-sh v0.0.0-20190412065543-76bd3d59ff27
	github.com/dchest/blake256 v1.0.0 // indirect
	github.com/graarh/golang-socketio v0.0.0-20170510162725-2c44953b9b5f
	github.com/imroc/req v0.2.4
	github.com/marwan-at-work/mod v0.4.1 // indirect
	github.com/pborman/uuid v1.2.0
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
	github.com/tidwall/gjson v1.3.5
	go.etcd.io/bbolt v1.3.3 // indirect
)

//replace github.com/blocktree/go-owcdrivers => ../../go-owcdrivers
