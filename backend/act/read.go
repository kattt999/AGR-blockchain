package act

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "../github.com/go-sql-driver/mysql"
)

func Read(db *sql.DB, ID int) {
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
		deviceId            int     //16
		basicData           float64 //12.5
		measurementUnitType int     //Cel.-1
		//measureUnitType string//Temperature
		gatherTime  string //2019-01-11 00:04:21
		location    string //First Monitor
		installTime string //2019-01-10 15:15:12
		farmId      int    //35
		userId      int    //7
		//ID: 59
		product_name  string //Rice
		product_batch string //super rice
		basename      string //hunan jiasui
		content       string //pesticide
		activity_date string //2019-04-04 00:00:00
		finish_date   string //2019-04-08 00:00:00
		province      string //Hunan
		city          string //Changsha
		district      string //Yuelu
		base_type     string //farm
		nickname      string //DENG Mi
		telephone     string //13810265160
		desc          string
	)
	rows, err := db.Query("select product_name, product_batch, sys_user_id, farm_id from product_batches where id = ?", ID)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&product_name, &product_batch, &userId, &farmId)
		if err != nil {
			log.Fatal(err)
		}
		fg.WriteString(product_name + "\n")
		fg.WriteString(product_batch + "\n")
	}
	rows, err = db.Query("select province, city, district, type, name, linkman, telephone from t_farm where id = ?", farmId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&province, &city, &district, &base_type, &basename, &nickname, &telephone)
		if err != nil {
			log.Fatal(err)
		}
		fg.WriteString(province + "\n")
		fg.WriteString(city + "\n")
		fg.WriteString(district + "\n")
		fg.WriteString(base_type + "\n")
		fg.WriteString(basename + "\n")
		fg.WriteString(nickname + "\n")
		fg.WriteString(telephone + "\n")
	}
	rows, err = db.Query("select id, location, installTime from t_device where farmId = ?", farmId)
	log.Println(farmId)
	log.Println("--------------------------------")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&deviceId, &location, &installTime)
		if err != nil {
			log.Fatal(err)
		}
		fg.WriteString(location + "\t" + installTime + "\n")
		rows2, err := db.Query("select basicData, measurementUnitId, gatherTime from t_device_gather where deviceId = ?", deviceId)
		for rows2.Next() {
			err := rows2.Scan(&basicData, &measurementUnitType, &gatherTime)
			if err != nil {
				log.Fatal(err)
			}
			switch measurementUnitType {
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
	rows, err = db.Query("select content, activity_date, finish_date, desc_content from t_task where product_batches_id = ?", ID)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&content, &activity_date, &finish_date, &desc)
		if err != nil {
			log.Fatal(err)
		}
		fg.WriteString(content + "\t")
		fg.WriteString(activity_date + "\t")
		fg.WriteString(finish_date + "\n" + desc + "\n")
	}
}
