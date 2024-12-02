module py_go_bridge_example2

go 1.21

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/sirupsen/logrus v1.9.3
	github.com/xuri/excelize/v2 v2.9.0
	py_go_bridge v0.1.0
)

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.4 // indirect
	github.com/xuri/efp v0.0.0-20240408161823-9ad904a10d6d // indirect
	github.com/xuri/nfp v0.0.0-20240318013403-ab9948c2c4a7 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)

replace py_go_bridge => ../../py_go_bridge/go
