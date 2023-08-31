package controllers

import (
	"context"

	"github.com/umamaheswari76/netxd_customer_dal/interfaces"
	"github.com/umamaheswari76/netxd_customer_dal/models"
	tsn "github.com/umamaheswari76/netxd_customer_proto/transaction"
)

type TransactionServer struct{
	tsn.UnimplementedTransactionServiceServer
}

var (
	TransactionService interfaces.ITransaction
)

func (t * TransactionServer) Transfer(ctx context.Context,  req *tsn.Transaction)(*tsn.StringMessage, error){
	dbTransaction := &models.Transaction{
		Fromaccount: int(req.Fromaccount),
		Toaccount:   int(req.Toaccount),
		Amount:      int(req.Amount),
	}
	
	_, err := TransactionService.Transfer(dbTransaction.Fromaccount, dbTransaction.Toaccount, dbTransaction.Amount)
	if err != nil{
		return nil, err
	}else{
		stringmessage := &tsn.StringMessage{
			Message: "stringMssage: Transaction successfull",
		}
		return stringmessage, nil
	}


}