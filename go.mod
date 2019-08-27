module github.com/CortexFoundation/CortexTheseus

go 1.12

require (
	github.com/Azure/azure-pipeline-go v0.0.0-20180607212504-7571e8eb0876 // indirect
	github.com/Azure/azure-storage-blob-go v0.0.0-20180712005634-eaae161d9d5e
	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
	github.com/allegro/bigcache v0.0.0-20190218064605-e24eb225f156
	github.com/anacrolix/dht v0.0.0-20180808005204-cae37fd18420 // indirect
	github.com/anacrolix/missinggo v1.2.1
	github.com/anacrolix/tagflag v0.0.0-20180803105420-3a8ff5428f76
	github.com/anacrolix/torrent v0.0.0-20190823021928-95a521bad6f5
	github.com/aristanetworks/goarista v0.0.0-20170210015632-ea17b1a17847
	github.com/bitly/go-simplejson v0.0.0-20171023175154-0c965951289c
	github.com/boltdb/bolt v1.3.1
	github.com/bradfitz/iter v0.0.0-20190303215204-33e6a9893b0c
	github.com/btcsuite/btcd v0.0.0-20171128150713-2e60448ffcc6
	github.com/cespare/cp v0.1.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v0.0.0-20180603214616-504e848d77ea
	github.com/dgrijalva/jwt-go v0.0.0-20170201225849-2268707a8f08 // indirect
	github.com/docker/docker v0.0.0-20180625184442-8e610b2b55bf
	github.com/edsrzf/mmap-go v1.0.0
	github.com/elastic/gosigar v0.0.0-20180330100440-37f05ff46ffa
	github.com/elgatito/upnp v0.0.0-20180711183757-2f244d205f9a // indirect
	github.com/fatih/color v1.3.0
	github.com/fjl/memsize v0.0.0-20180418122429-ca190fb6ffbc
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gizak/termui v0.0.0-20170117222342-991cd3d38091
	github.com/go-ole/go-ole v1.2.1 // indirect
	github.com/go-stack/stack v1.8.0
	github.com/golang/protobuf v1.2.0
	github.com/golang/snappy v0.0.1
	github.com/hashicorp/golang-lru v0.0.0-20160813221303-0a025b7e63ad
	github.com/huin/goupnp v0.0.0-20161224104101-679507af18f3
	github.com/influxdata/influxdb v0.0.0-20180221223340-01288bdb0883
	github.com/jackpal/go-nat-pmp v0.0.0-20160603034137-1fa385a6f458
	github.com/julienschmidt/httprouter v0.0.0-20170430222011-975b5c4c7c21
	github.com/karalabe/hid v0.0.0-20170821103837-f00545f9f374
	github.com/maruel/panicparse v0.0.0-20160720141634-ad661195ed0e // indirect
	github.com/mattn/go-colorable v0.0.0-20170210172801-5411d3eea597
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/naoina/toml v0.0.0-20170918210437-9fafd6967416
	github.com/nsf/termbox-go v0.0.0-20170211012700-3540b76b9c77 // indirect
	github.com/olekukonko/tablewriter v0.0.0-20190409134802-7e037d187b0c
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/opentracing/opentracing-go v0.0.0-20180606204148-bd9c31933947
	github.com/pborman/uuid v0.0.0-20170112150404-1b00554d8222
	github.com/peterh/liner v0.0.0-20190123174540-a2c9a5303de7
	github.com/prometheus/tsdb v0.0.0-20190402121629-4f204dcbc150
	github.com/rjeczalik/notify v0.9.1
	github.com/robertkrimen/otto v0.0.0-20170205013659-6a77b7cbc37d
	github.com/rs/cors v0.0.0-20160617231935-a62a804a8a00
	github.com/rs/xhandler v0.0.0-20160618193221-ed27b6fd6521 // indirect
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/goleveldb v0.0.0-20190318030020-c3a204f8e965
	github.com/uber/jaeger-client-go v0.0.0-20180607151842-f7e0d4744fa6 // indirect
	github.com/uber/jaeger-lib v0.0.0-20180615202729-a51202d6f4a7 // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/sys v0.0.0-20190712062909-fae7ac547cb7
	golang.org/x/tools v0.0.0-20170215214335-be0fcc31ae23
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20180302121509-abf0ba0be5d5
	gopkg.in/resty.v1 v1.0.0-20180529033006-fc3ad735b556
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/urfave/cli.v1 v1.20.0
)
