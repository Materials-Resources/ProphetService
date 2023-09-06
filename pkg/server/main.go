package server

import (
	"fmt"
	"github.com/materials-resources/s_prophet/pkg/database"
	"github.com/materials-resources/s_prophet/pkg/database/db_receiving"
	"github.com/materials-resources/s_prophet/pkg/grpc_service"
	rpc_order "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
	product "github.com/materials-resources/s_prophet/proto/product/v1alpha0"
	rpc_receiving "github.com/materials-resources/s_prophet/proto/receiving/v1alpha0"
	rpc_warehouse "github.com/materials-resources/s_prophet/proto/warehouse/v1alpha0"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
)

func Serve() error {

	listenOn := "0.0.0.0:50058"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to start grpc_service on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()

	db := connectBun()

	gdb := connectSQL()

	registerGRPCServices(server, db, gdb)

	reflection.Register(server)

	log.Println("Listening on", listenOn)

	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}
	return nil
}

func registerGRPCServices(server *grpc.Server, db *bun.DB, gdb *gorm.DB) {
	product.RegisterProductServiceServer(server, &grpc_service.ProductServer{ProductHandler: database.NewProductHandler(db)})
	rpc_warehouse.RegisterWarehouseServiceServer(server, &grpc_service.WarehouseServer{WarehouseHandler: database.NewWarehouseHandler(db)})
	rpc_order.RegisterOrderServiceServer(server, &grpc_service.OrderServer{OrderHandler: database.NewOrderHandler(db)})
	rpc_receiving.RegisterReceivingServiceServer(server, &grpc_service.ReceivingServer{DBHandler: db_receiving.NewReceivingHandler(gdb)})

	//accountV1.RegisterAccountServiceServer(server, &serviceAccount.Server{DB: db})
	//inventoryV1.RegisterInventoryServiceServer(server, &serviceInventory.Server{DB: db})
	//orderV1.RegisterOrderServiceServer(server, &serviceOrder.Server{DB: db})
	//warehousev1grpc.RegisterWarehouseServiceServer(server, &serviceWarehouse.Server{DB: db})
	//catalogv1grpc.RegisterCatalogServiceServer(server, &serviceCatalog.Server{DB: db})

}
