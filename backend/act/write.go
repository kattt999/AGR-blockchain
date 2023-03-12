package act

import (
	/*"database/sql"
	"fmt"
	"log"
	"os"

	"../github.com/gorilla/mux"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"github.com/awslabs/amazon-qldb-driver-go/qldbdriver"*/

	"context"
	"io/ioutil"
	"strings"

	"github.com/awslabs/amazon-qldb-driver-go/qldbdriver"
)

func Write(driver *qldbdriver.QLDBDriver) {
	str, _ := ioutil.ReadFile("db.sql")
	list := strings.Split(string(str), "\n")
	for _, s := range list {
		driver.Execute(context.Background(), func(txn qldbdriver.Transaction) (interface{}, error) {
			txn.Execute(s)
			return nil, nil
		})

	}
	return
}
