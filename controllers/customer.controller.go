//implementing the functions, and also in that mapping the mongodb values to server(proto msg)

package controllers

import (
	"context"
	"time"

	"github.com/umamaheswari76/netxd_customer_dal/interfaces"
	"github.com/umamaheswari76/netxd_customer_dal/models"	
	cst "github.com/umamaheswari76/netxd_customer_proto/customer"
)

type RPCServer struct{
	cst.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

// CustomerId: 101,
// FirstName:  "umamaheswari",
// SecondName: "m",
// BankId:     "1",
// Balance:    5000,

func (s *RPCServer) CreateCustomer(ctx context.Context, req *cst.Customer)(*cst.CustomerResponse, error){
	dbCustomer := &models.Customer{
		CustomerId: req.CustomerId,
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		BankId:     req.BankId,
		Balance:    req.Balance,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		IsActive:   "yes",
	}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err !=nil{
		return nil, err
	}else {
		responseCustomer := &cst.CustomerResponse{
			CustomerId: result.CustomerId,
			CreatedAt:  result.CreatedAt.String(),
		}
		return responseCustomer, nil
	}
}

