mockgen:
	cd ./service/athUser && mockgen -package=mock -destination=./mock/ath_user_mock.go backend/service/athUser AthUserService
	cd ./service/athOtp && mockgen -package=mock -destination=./mock/ath_otp_mock.go backend/service/athOtp AthOtpService
	cd ./service/athCompany && mockgen -package=mock -destination=./mock/ath_company_mock.go backend/service/athCompany AthCompanyService
	cd ./service/athCustomer && mockgen -package=mock -destination=./mock/ath_customer_mock.go backend/service/athCustomer AthCustomerService
	cd ./service/athFacility && mockgen -package=mock -destination=./mock/ath_facility_mock.go backend/service/athFacility AthFacilityService
	cd ./service/athVenue && mockgen -package=mock -destination=./mock/ath_venue_mock.go backend/service/athVenue AthVenueService
	cd ./service/athAdmin && mockgen -package=mock -destination=./mock/ath_admin_mock.go backend/service/athAdmin AthAdminUserService

run_tests:
	go test -v ./pkg/handler/...

generate_protoc:
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. api/athUser/v1/athUser.proto
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. api/athUser/v1/athUser.proto
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. api/athUser/v1/athUser.proto

	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. api/athCustomer/v1/athCustomer.proto
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. api/athCustomer/v1/athCustomer.proto
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. api/athCustomer/v1/athCustomer.proto

	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. api/athFacility/v1/athFacility.proto
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. api/athFacility/v1/athFacility.proto
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. api/athFacility/v1/athFacility.proto

	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. api/athVenue/v1/athVenue.proto
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. api/athVenue/v1/athVenue.proto
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. api/athVenue/v1/athVenue.proto

	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. api/athAdmin/v1/athAdmin.proto
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. api/athAdmin/v1/athAdmin.proto
	protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. api/athAdmin/v1/athAdmin.proto