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
	"fmt"
	"os"

	"github.com/amzn/ion-go/ion"
	"github.com/awslabs/amazon-qldb-driver-go/qldbdriver"
)

func ReadQLDB(driver *qldbdriver.QLDBDriver, ID int) {
	newFile, err := os.Create("general.txt")
	newFile.Close()
	newFile, err = os.Create("air_temp.txt")
	newFile.Close()
	newFile, err = os.Create("air_hum.txt")
	newFile.Close()
	newFile, err = os.Create("light.txt")
	newFile.Close()
	newFile, err = os.Create("soil_temp.txt")
	newFile.Close()
	newFile, err = os.Create("soil_hum.txt")
	newFile.Close()
	newFile, err = os.Create("res.txt")
	newFile.Close()
	newFile, err = os.Create("ph.txt")
	newFile.Close()
	//file general, file air temp, file air hum, file light, file soil temp, file soil hum, file res, file ph
	fg, err := os.OpenFile("general.txt", os.O_APPEND, 0666)
	fat, err := os.OpenFile("air_temp.txt", os.O_APPEND, 0666)
	fah, err := os.OpenFile("air_hum.txt", os.O_APPEND, 0666)
	fl, err := os.OpenFile("light.txt", os.O_APPEND, 0666)
	fst, err := os.OpenFile("soil_temp.txt", os.O_APPEND, 0666)
	fsh, err := os.OpenFile("soil_hum.txt", os.O_APPEND, 0666)
	fr, err := os.OpenFile("res.txt", os.O_APPEND, 0666)
	fp, err := os.OpenFile("ph.txt", os.O_APPEND, 0666)
	var (
		//measureUnitType string//Temperature
		gatherTime string //2019-01-11 00:04:21
		farmId     int    //35
		//ID: 59
	)
	type First struct {
		product_name  string
		product_batch string
		userId        int
		farmId        int
	}
	p, err := driver.Execute(context.Background(), func(txn qldbdriver.Transaction) (interface{}, error) {
		result, err := txn.Execute("select product_name, product_batch, sys_user_id, farm_id from product_batches where id = ?", ID)
		if err != nil {
			return nil, err
		}

		// Assume the result is not empty
		hasNext := result.Next(txn)
		if !hasNext && result.Err() != nil {
			return nil, result.Err()
		}
		fmt.Println("?")
		ionBinary := result.GetCurrentData()

		temp := new(First)
		err = ion.Unmarshal(ionBinary, temp)
		if err != nil {
			return nil, err
		}

		return *temp, nil
	})
	if err != nil {
		panic(err)
	}
	var returnedFirst First
	returnedFirst = p.(First)

	type Second struct {
		province  string
		city      string
		district  string
		base_type string //farm
		basename  string //hunan jiasui
		nickname  string //DENG Mi
		telephone string
	}
	p, err = driver.Execute(context.Background(), func(txn qldbdriver.Transaction) (interface{}, error) {
		result, err := txn.Execute("select province, city, district, type, name, linkman, telephone from t_farm where id = ?", returnedFirst.farmId)
		if err != nil {
			return nil, err
		}

		// Assume the result is not empty
		hasNext := result.Next(txn)
		if !hasNext && result.Err() != nil {
			return nil, result.Err()
		}
		fmt.Println("?")
		ionBinary := result.GetCurrentData()

		temp := new(First)
		err = ion.Unmarshal(ionBinary, temp)
		if err != nil {
			return nil, err
		}

		return *temp, nil
	})
	if err != nil {
		panic(err)
	}
	var returnedSecond Second
	returnedSecond = p.(Second)
	fg.WriteString(returnedSecond.province + "\n")
	fg.WriteString(returnedSecond.city + "\n")
	fg.WriteString(returnedSecond.district + "\n")
	fg.WriteString(returnedSecond.base_type + "\n")
	fg.WriteString(returnedSecond.basename + "\n")
	fg.WriteString(returnedSecond.nickname + "\n")
	fg.WriteString(returnedSecond.telephone + "\n")

	type Third struct {
		deviceId    int //16
		basicData   string
		installTime string //2019-01-10 15:15:12
	}
	p, err = driver.Execute(context.Background(), func(txn qldbdriver.Transaction) (interface{}, error) {
		result, err := txn.Execute("select id, location, installTime from t_device where farmId = ?", farmId)
		if err != nil {
			return nil, err
		}

		var three []Third
		for result.Next(txn) {
			ionBinary := result.GetCurrentData()
			temp := new(Third)
			err = ion.Unmarshal(ionBinary, temp)
			if err != nil {
				return nil, err
			}

			three = append(three, *temp)
		}
		if result.Err() != nil {
			return nil, result.Err()
		}

		return three, nil
	})
	if err != nil {
		panic(err)
	}

	var three []Third
	three = p.([]Third)

	type Set struct {
		basicData           float64
		measurementUnitType int
		gatherTime          string
	}

	for i := 0; i < len(three); i++ {
		device := three[i].deviceId

		p, err = driver.Execute(context.Background(), func(txn qldbdriver.Transaction) (interface{}, error) {
			result, err := txn.Execute("select basicData, measurementUnitId, gatherTime from t_device_gather where deviceId = ?", device)
			if err != nil {
				return nil, err
			}

			var info []Set
			for result.Next(txn) {
				ionBinary := result.GetCurrentData()
				temp := new(Set)
				err = ion.Unmarshal(ionBinary, temp)
				if err != nil {
					return nil, err
				}

				info = append(info, *temp)
			}
			if result.Err() != nil {
				return nil, result.Err()
			}

			return info, nil
		})
	}
	var info []Set
	info = p.([]Set)
	for i := 0; i < len(info); i++ {
		UnitType := info[i].measurementUnitType
		basicData := info[i].basicData
		switch UnitType {
		case 1:
			fat.WriteString("\n" + gatherTime + "\t")
			fat.WriteString(fmt.Sprintf("%f", basicData))
		case 2:
			fah.WriteString("\n" + gatherTime + "\t")
			fah.WriteString(fmt.Sprintf("%f", basicData))
		case 3:
			fst.WriteString("\n" + gatherTime + "\t")
			fst.WriteString(fmt.Sprintf("%f", basicData))
		case 4:
			fsh.WriteString("\n" + gatherTime + "\t")
			fsh.WriteString(fmt.Sprintf("%f", basicData))
		case 5:
			fr.WriteString("\n" + gatherTime + "\t")
			fr.WriteString(fmt.Sprintf("%f", basicData))
		case 6:
			fp.WriteString("\n" + gatherTime + "\t")
			fp.WriteString(fmt.Sprintf("%f", basicData))
		case 7:
			fl.WriteString("\n" + gatherTime + "\t")
			fl.WriteString(fmt.Sprintf("%f", basicData))
		}
	}
}
