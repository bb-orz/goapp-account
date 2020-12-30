module goapp

go 1.15

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.47.0
	// github.com/bb-orz/goinfras => /Users/fun/Code/MyProject/goinfras
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	go.uber.org/atomic => github.com/uber-go/atomic v1.5.0
	go.uber.org/multierr => github.com/uber-go/multierr v1.4.0
	go.uber.org/tools => github.com/uber-go/tools v0.0.0-20190618225709-2cfd321de3ee
	go.uber.org/zap => github.com/uber-go/zap v1.12.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/exp => github.com/golang/exp v0.0.0-20191030013958-a1ab85dbe136
	golang.org/x/image => github.com/golang/image v0.0.0-20191009234506-e7c1f5e7dbb8
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190930215403-16217165b5de
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20191031020345-0945064e013a
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20190327025741-74e053c68e29
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190322080309-f49334f85ddc
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20191024005414-555d28b269f0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190330180304-aef51cc3777c
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191011141410-1b5146add898
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.13.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.5
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20191028173616-919d9bdd9fe6
	google.golang.org/grpc => github.com/grpc/grpc-go v1.24.0

)

require (
	github.com/bb-orz/goinfras v1.2.0
	github.com/garyburd/redigo v1.6.2
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	github.com/segmentio/ksuid v1.0.3
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.16.0
	gopkg.in/go-playground/validator.v9 v9.31.0
)
