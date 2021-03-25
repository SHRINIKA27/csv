package main

import (
    "math/big"
    "encoding/csv"
    "os"
    "fmt"
    "math/rand"
    "time"
    _ "github.com/go-sql-driver/mysql"
)

type oxidata struct {
    
    ID string   
    HGB string  
    SAO2 string   
    SVO2 string  
    PAO2 string  
    PVO2 string  
    CO string  
}
var (
    SNO1 string
    ID1 string
    HGB1 string
    SAO21 string
    SVO21 string
    PAO21 string
    PVO21 string
    CO1 string
    DO21 string
    VO21 string
    ards string
)
var remainder = big.NewInt(0);
var pk = big.NewInt(0);
var ten1 = big.NewInt(0);
var cao2_range1 = big.NewInt(0);
var cvo2_range1 = big.NewInt(0);
var co_range1 = big.NewInt(0);
var co_range2 = big.NewInt(0);

const MAX int = 10000000;
var get_key int =0;
var arr_key int =0;
var arr_key1 int =3;
var get_keyvalues int =0;
var a1 int =0;    
var b1 int =0;
var j int =0;

var do2_range  = make([]*big.Int, 0)
var vo2_range  = make([]*big.Int, 0)
var hgb_range  = make([]*big.Int, 0)
var pkey  = make([]*big.Int, 0)

var key_hard = make([]int,0)
var keylist = make([]int,0)
var list = make([]int,0)

func GetData(){
	    csvFile, err := os.Open("oximeter_data.csv")
    if err != nil {
        fmt.Println(err)
    }

	    fmt.Println("Successfully Opened CSV file")
    defer csvFile.Close()
    
    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }   
    var i int
    i = 0
    for _, line := range csvLines {
        oximeter_data := oxidata{
            ID: line[0],
            HGB:  line[1],
            SAO2: line[2],
            SVO2:  line[3],
            PAO2: line[4],
            PVO2:  line[5],
            CO: line[6],
        }
        
        id := oximeter_data.ID
        hgb := oximeter_data.HGB
        sao2 := oximeter_data.SAO2
        svo2 := oximeter_data.SVO2
        pao2 := oximeter_data.PAO2
        pvo2 := oximeter_data.PVO2
        co := oximeter_data.CO

        if i==0 {
            i = i + 1
        } else {
        fmt.Println(id + " " + hgb + " " + sao2+ " " + svo2+ " " + pao2+ " " + pvo2+ " " + co)
    }
    }

}
func main() {

    
    rand.Seed(time.Now().UnixNano())
    GetData()
    
    
}