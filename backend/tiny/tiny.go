package tiny

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"net/http"

	"../act"
	"../entity"

	"context"
	//"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"github.com/awslabs/amazon-qldb-driver-go/qldbdriver"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func calc(data [30]float64, rang [4]float64) float64 {
	var ret float64
	ret = 0
	for i := 0; i < 30; i++ {
		if data[i] <= rang[2] && data[i] >= rang[1] {
			ret = ret + 10
		} else if data[i] <= rang[1] && data[i] >= rang[0] {
			ret = ret + 10*(data[i]-rang[0])/(rang[1]-rang[0])
		} else if data[i] <= rang[3] && data[i] >= rang[2] {
			ret = ret + 10*(rang[3]-data[i])/(rang[3]-rang[2])
		}
	}
	return ret / 30
}
func Rate(cond entity.Article, rang [7][4]float64) float64 {
	/*var all_rate [7]float64
	all_rate[0] = calc(cond.PH, rang[0])
	all_rate[1] = calc(cond.L, rang[1])
	all_rate[2] = calc(cond.SH, rang[2])
	all_rate[3] = calc(cond.ST, rang[3])
	all_rate[4] = calc(cond.AH, rang[4])
	all_rate[5] = calc(cond.AT, rang[5])
	all_rate[6] = calc(cond.R, rang[6])
	var ret float64
	ret = 0
	for i := 0; i < 7; i++ {
		ret += all_rate[i]
	}
	return ret / 7*/
	return 10
}

func Fake (str string) []entity.KlineData{
	temp := strings.Split(str, "\n")
	var kd []entity.KlineTemp
	timerepeat := 0
	for tempsplit := range temp {
		//fmt.Println(tempsplit,temp[tempsplit])
		if temp[tempsplit] == "" {
			continue
		}
		dat := strings.Split(temp[tempsplit], "	")
		f, _ := strconv.ParseFloat(dat[1], 64)
		if f != 0 {
			DnT := strings.Split(dat[0], " ")
			index := 0
			needappend := true
			for timerepeat = range kd{
				if kd[timerepeat].Kline.Date == DnT[0]{
					needappend = false
					index = timerepeat
					break
				}
			}
			if needappend == true {
				kd=append(kd, entity.KlineTemp{Kline: entity.KlineData{Date: DnT[0], Data: [4]float64{0,0,100000000,0}}, Btime: "23:59:59", Etime: "00	:00:00"})
				index = timerepeat
			}
			//fmt.Println(DnT[0],DnT[1])
			if DnT[1] < kd[index].Btime{
				kd[index].Btime = DnT[1]
				kd[index].Kline.Data[0] = f
			}
			if DnT[1] > kd[index].Etime{
				kd[index].Etime = DnT[1]
				kd[index].Kline.Data[1] = f
			}
			if f < kd[index].Kline.Data[2]{
				kd[index].Kline.Data[2] = f
			}
			if f > kd[index].Kline.Data[3]{
				kd[index].Kline.Data[3] = f
			}
		}
	}
	var ret []entity.KlineData
	for returnTemp := range kd{
		ret = append (ret, kd[returnTemp].Kline)
	}
	return ret
}
func Fake_old(str string) [30]float64 {
	temp := strings.Split(str, "\n")
	var ret [30]float64
	var tempsplit []string
	a := 30
	for i := 0; i < a; i++ {
		for ; len(temp[100+i]) == 0; i++ {
			a++
		}
		tempsplit = strings.Split(temp[100+i], "	")
		ret[i+30-a], _ = strconv.ParseFloat(tempsplit[1], 64)
	}
	return ret

}

func Find(id int) entity.Article {
	db, err := sql.Open("mysql", "root:3150@tcp(127.0.0.1:3306)/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println(id)
	act.Read(db, id)
	/*general, err := ioutil.ReadFile("general.txt")
	if err != nil {
		log.Fatal(err)
	}*/
	ph, err := ioutil.ReadFile("ph.txt")
	PH := Fake(string(ph))
	light, err := ioutil.ReadFile("light.txt")
	LIGHT := Fake(string(light))
	soil_hum, err := ioutil.ReadFile("soil_hum.txt")
	SOIL_HUM := Fake(string(soil_hum))
	soil_temp, err := ioutil.ReadFile("soil_temp.txt")
	SOIL_TEMP := Fake(string(soil_temp))
	air_hum, err := ioutil.ReadFile("air_hum.txt")
	AIR_HUM := Fake(string(air_hum))
	air_temp, err := ioutil.ReadFile("air_temp.txt")
	AIR_TEMP := Fake(string(air_temp))
	res, err := ioutil.ReadFile("res.txt")

	RES := Fake(string(res))
	art := entity.Article{ /*Funda: string(general), */ PH: PH, L: LIGHT, SH: SOIL_HUM, ST: SOIL_TEMP, AH: AIR_HUM, AT: AIR_TEMP, R: RES}
	return art
}

func FindQLDB(id int) entity.Article {
	awsSession := session.Must(session.NewSession(aws.NewConfig().WithRegion("us-east-1")))
	qldbSession := qldbsession.New(awsSession)

	driver, err := qldbdriver.New(
		"AIoTDB",
		qldbSession,
		func(options *qldbdriver.DriverOptions) {
			options.LoggerVerbosity = qldbdriver.LogInfo
		})
	if err != nil {
		panic(err)
	}

	defer driver.Shutdown(context.Background())
	log.Println(id)
	act.ReadQLDB(driver, id)
	/*general, err := ioutil.ReadFile("general.txt")
	if err != nil {
		log.Fatal(err)
	}*/
	ph, err := ioutil.ReadFile("ph.txt")
	PH := Fake(string(ph))
	light, err := ioutil.ReadFile("light.txt")
	LIGHT := Fake(string(light))
	soil_hum, err := ioutil.ReadFile("soil_hum.txt")
	SOIL_HUM := Fake(string(soil_hum))
	soil_temp, err := ioutil.ReadFile("soil_temp.txt")
	SOIL_TEMP := Fake(string(soil_temp))
	air_hum, err := ioutil.ReadFile("air_hum.txt")
	AIR_HUM := Fake(string(air_hum))
	air_temp, err := ioutil.ReadFile("air_temp.txt")
	AIR_TEMP := Fake(string(air_temp))
	res, err := ioutil.ReadFile("res.txt")

	RES := Fake(string(res))
	art := entity.Article{ /*Funda: string(general), */ PH: PH, L: LIGHT, SH: SOIL_HUM, ST: SOIL_TEMP, AH: AIR_HUM, AT: AIR_TEMP, R: RES}
	return art
}
func Trig() {
	awsSession := session.Must(session.NewSession(aws.NewConfig().WithRegion("us-east-1")))
	qldbSession := qldbsession.New(awsSession)

	driver, err := qldbdriver.New(
		"quick-start",
		qldbSession,
		func(options *qldbdriver.DriverOptions) {
			options.LoggerVerbosity = qldbdriver.LogInfo
		})
	if err != nil {
		panic(err)
	}

	defer driver.Shutdown(context.Background())
	act.Write(driver)
}
func Uploadchart(w http.ResponseWriter, _ *http.Request, kd []entity.KlineData, str string) {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([]opts.KlineData, 0)
	for i := 0; i < len(kd); i++ {
		x = append(x, kd[i].Date)
		y = append(y, opts.KlineData{Value: kd[i].Data})
	}

	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: str,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)

	kline.SetXAxis(x).AddSeries("kline", y)
	kline.Render(w)
}
