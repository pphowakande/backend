package main

import (
	db2 "backend/pkg/db"
	"backend/pkg/handler"
	"context"
	"log"
	"net"
	"net/http"

	userpb "backend/api/athUser/v1"

	adminpb "backend/api/athAdmin/v1"
	customerpb "backend/api/athCustomer/v1"
	facilitypb "backend/api/athFacility/v1"
	venuepb "backend/api/athVenue/v1"

	adminService "backend/service/athAdmin"
	companyService "backend/service/athCompany"
	customerService "backend/service/athCustomer"
	facilityService "backend/service/athFacility"
	otpService "backend/service/athOtp"
	userService "backend/service/athUser"
	venueService "backend/service/athVenue"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const grpcPort = ":8082"
const httpPort = ":8081"

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", grpcPort)
		//logger.Log("fatal", "datastore", "ab-101", "error listening to grpcport	", "connectionError", "", map[string]interface{}{}, true)
	}

	s := grpc.NewServer()
	db := db2.New(db2.Connect())

	athUserService := userService.NewBasicAthUserService(db)
	athCompanyService := companyService.NewBasicAthCompanyService(db)
	athAdminUserService := adminService.NewBasicAthAdminService(db)
	athOtpService := otpService.NewBasicAthOtpService(db)
	athVenueService := venueService.NewBasicAthVenueService(db)
	athFacilityService := facilityService.NewBasicAthFacilityService(db)
	athCustomerService := customerService.NewBasicAthCustomerService(db)

	server := handler.NewGrpcService(athCompanyService, athUserService, athAdminUserService, athOtpService, athVenueService, athFacilityService, athCustomerService)
	//customerServer := handler.NewGrpcService(athCustomerService)

	userpb.RegisterUserServer(s, server)
	customerpb.RegisterCustomerServer(s, server)
	facilitypb.RegisterFacilityServer(s, server)
	venuepb.RegisterVenueServer(s, server)
	adminpb.RegisterAdminServer(s, server)

	log.Println("grrpc server has started on", grpcPort)

	// Start the gRPC server in goroutine
	go s.Serve(lis)

	// Start the HTTP server for Rest
	log.Println("Starting HTTP server on port " + httpPort)
	run()

}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	//swagger := http.FileServer(http.Dir("./3rdparty/swagger-ui"))
	// fmt.Println("request", r.URL.Path)
	// p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	// p = path.Join("3rdparty/swagger-ui/", p)
	// fmt.Println("request map ", p)
	// http.ServeFile(w, r, p)

	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "default.swagger.json")
}

func run() error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithInsecure()}

	var conn *grpc.ClientConn
	err := userpb.RegisterUserHandler(ctx, gw, conn)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", serveSwagger)
	//curdir, _ := os.Getwd()
	mux.Handle("/api/", gw)

	return http.ListenAndServe(httpPort, mux)
}
