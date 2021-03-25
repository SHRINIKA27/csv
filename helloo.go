package main

import (
    "fmt"
    "math/big"
    "math/rand"
    "time"
    "strconv"
    "os"
    "encoding/csv"
)

var (

    ID string   
    HGB string  
    SAO2 string   
    SVO2 string  
    PAO2 string  
    PVO2 string  
    CO string 
)
var remainder = big.NewInt(0);
var pk = big.NewInt(0);
var ten1 = big.NewInt(0);
var cao2_range1 = big.NewInt(0);
var cvo2_range1 = big.NewInt(0);
var co_range1 = big.NewInt(0);
var co_range2 = big.NewInt(0);
var ards string ;
var b1 int =0;


var IDe = big.NewInt(0) 
var HGBe = big.NewInt(0) 
var SAO2e = big.NewInt(0) 
var SVO2e = big.NewInt(0) 
var PAO2e = big.NewInt(0) 
var PVO2e = big.NewInt(0) 
var COe = big.NewInt(0) 
var DO2calce = big.NewInt(0) 
var VO2calce = big.NewInt(0) 

const MAX int = 10000000;
var get_key int =0;
var arr_key int =0;
var arr_key1 int =3;
var get_keyvalues int =0;
var a1 int =0;    
//var b1 int =0;
var j int =0;

var do2_range  = make([]*big.Int, 0)
var vo2_range  = make([]*big.Int, 0)
var hgb_range  = make([]*big.Int, 0)
var pkey  = make([]*big.Int, 0)

var key_hard = make([]int,0)
var keylist = make([]int,0)
var list = make([]int,0)

func Encryption_Initialization() {
           
            //prime get  
            get_key=get_key+1;
            get_keyvalues=get_key+1;
            key_generation := new(big.Int)  
            
            key_generation.Mul(big.NewInt(int64(list[get_key])),big.NewInt(int64(list[get_keyvalues])));
            fmt.Println(get_key);
            pkey=append(pkey,big.NewInt(int64(list[get_keyvalues])));
            //random key
            var result int =key_hard[arr_key];
            arr_key+=1;
            
            //power 
            power_values := new(big.Int)  

            power_values.Exp((big.NewInt(int64(list[get_keyvalues]))),big.NewInt(int64(list[get_key])),nil);
            //final power
            final_power:= new(big.Int)  

            final_power.Mul(power_values,big.NewInt(int64(result)));

            //mod
            remainder.Mod(final_power,key_generation);
    }

    func Encryption_process(val string) (ncrypted_unit *big.Int) {

            id1:= new(big.Int)  
            ncrypted_unit = new(big.Int)  
            id1,ok:= id1.SetString(val, 10)
            if !ok {
             fmt.Println("SetString: error")
             return
            }
            ncrypted_unit.Add(id1,remainder);
            return            
    }


func Encryption() {
    
    //random number generation

    get_key =1000+rand.Intn(400);
        var i int 
        for  i = 0;i<MAX;i++ {
                      
                    var key int = 1200+rand.Intn(200000);
                    if key>0 {
                        
                        key_hard = append(key_hard,key);
                        }       
        }
    //  prime number gneration

    var beg=1500
    var end=20000 
    var n int 
    var j int 
            for  n = beg; n <= end; n++ {
            var prime = true;
            for  j = 2; j <= n / 2; j++ {
                if n % j == 0 && n != j {
                    prime = false;
                }
            }
            if prime {
                list = append(list,n);
            }
            }    

    //converting into float for cao2 calculation
    f1, _ := strconv.ParseFloat(HGB, 32)
    f2, _ := strconv.ParseFloat(SAO2, 32)
    f3, _ := strconv.ParseFloat(PAO2, 32)
    f4, _ := strconv.ParseFloat(SVO2, 32)
    f5, _ := strconv.ParseFloat(PVO2, 32)


    //cao2 calculation
    cao2calc :=(1.34*f1*(f2/100))+(f3*0.0031)
    //fmt.Print("CAO2 :")
    //fmt.Println(int64(cao2calc))
    cao2s := strconv.Itoa(int(cao2calc))

    //cvo2 calculation
    cvo2calc :=(1.34*f1*(f4/100))+(f5*0.0031)
    //fmt.Print("CVO2 :")
    //fmt.Println(int64(cvo2calc))


    //cao2-cvo2
    cao2calc1 :=int64(cao2calc)
    cvo2calc1 :=int64(cvo2calc)
    dif := cao2calc1 - cvo2calc1
    //fmt.Print("DIFF :")
    //fmt.Println(int64(dif))
    difs := strconv.Itoa(int(dif))

    //encryption starts

    Encryption_Initialization();

    fmt.Println("----------------Encrypted Values----------------")
    //IDe:= new(big.Int) 
    IDe=Encryption_process(ID) 
    fmt.Print("ID :")
    fmt.Println(IDe)
    //HGBe:= new(big.Int) 
    HGBe=Encryption_process(HGB) 
    fmt.Print("HBG :")
    fmt.Println(HGBe)
    //SAO2e:= new(big.Int) 
    SAO2e=Encryption_process(SAO2) 
    fmt.Print("SAO2 :")
    fmt.Println(SAO2e)
    //SVO2e:= new(big.Int) 
    SVO2e=Encryption_process(SVO2) 
    fmt.Print("SVO2 :")
    fmt.Println(SVO2e)
    //PAO2e:= new(big.Int) 
    PAO2e=Encryption_process(PAO2)
    fmt.Print("PAO2 :")
    fmt.Println(PAO2e) 
    //PVO2e:= new(big.Int) 
    PVO2e=Encryption_process(PVO2) 
    fmt.Print("PVO2 :")
    fmt.Println(PVO2e)
    //COe:= new(big.Int) 
    COe=Encryption_process(CO)
    fmt.Print("CO :")
    fmt.Println(COe) 
    CAO2e:= new(big.Int) 
    CAO2e=Encryption_process(cao2s) 
    //fmt.Println(CAO2e)
    DIFe:= new(big.Int) 
    DIFe=Encryption_process(difs) 
    //fmt.Println(DIFe)
    tene:= new(big.Int) 
    tene=Encryption_process("10") 
    

    //DO2 Calculation
    
    DO2calce.Mul(DO2calce.Mul(CAO2e,COe),tene)
    fmt.Print("DO2 :")
    fmt.Println(DO2calce)
    
    //VO2 Calc
    //VO2calce := new(big.Int) 
    VO2calce.Mul(VO2calce.Mul(COe,DIFe),tene)
    fmt.Print("VO2 :")
    fmt.Println(VO2calce)
    //fmt.Print("VO2 :")
    fmt.Println("------------------------------------------------")
    
    cao2_r1 := "11";
    cvo2_r1 := "8";
    co_r1 := "3";
    
    hgb_r2 := "11";

    cao2_re:= new(big.Int) 
    cao2_re=Encryption_process(cao2_r1)

    cvo2_re:= new(big.Int) 
    cvo2_re=Encryption_process(cvo2_r1)

    co_r1e:= new(big.Int) 
    co_r1e=Encryption_process(co_r1)

    
    hgb_re:= new(big.Int) 
    hgb_re=Encryption_process(hgb_r2)

    hgb_range = append(hgb_range,hgb_re);

    DO2ra := new(big.Int) 
    DO2ra.Mul(DO2ra.Mul(cao2_re,co_r1e),tene)

    do2_range = append(do2_range,DO2ra);

    VO2ra := new(big.Int) 
    VO2ra.Mul(VO2ra.Mul(co_r1e,cvo2_re),tene)

    vo2_range = append(vo2_range,VO2ra);
    
}

func Analysis()  {


        var temp =0;
        var temp1 =0;
        var temp2 =0;
        //var ards="No"

       var c1 =0;
       var c2 =0;
       var c3 =0;
       
       do2_ranged := new(big.Int) 
       do2_ranged.Mod(do2_range[b1],pkey[b1])
       //fmt.Println("do2 range:")
       //fmt.Print(do2_ranged)

       do2_calcd := new(big.Int) 
       do2_calcd.Mod(DO2calce,pkey[b1])
       //fmt.Println("do2 :")
       //fmt.Print(do2_calcd)

       vo2_ranged := new(big.Int) 
       vo2_ranged.Mod(vo2_range[b1],pkey[b1])
       //fmt.Println("vo2 range:")
       //fmt.Print(vo2_ranged)

       vo2_calcd := new(big.Int) 
       vo2_calcd.Mod(VO2calce,pkey[b1])
       //fmt.Println("vo2:")
       //fmt.Print(vo2_calcd)

       hgb_ranged := new(big.Int) 
       hgb_ranged.Mod(hgb_range[b1],pkey[b1])
       //fmt.Println("hgb range:")
       //fmt.Print(hgb_ranged)

       hgb_calcd := new(big.Int) 
       hgb_calcd.Mod(HGBe,pkey[b1])
       //fmt.Println("hgb:")
       //fmt.Print(hgb_calcd)


       c1=do2_range[b1].Cmp(DO2calce)
       //fmt.Println(c1)

       c2=vo2_range[b1].Cmp(VO2calce)
       //fmt.Println(c2)

       c3=hgb_range[b1].Cmp(HGBe)
       //fmt.Println(c3)

       if (c1==0 || c1==1) {
        temp=1
       }
       if (c2==0 || c2==1) {
        temp1=1
       }
       if (c3==0 || c3==1) {
        temp2=1
       }
     
       if ( (temp==1 && temp2==1) || (temp==1 && temp1==1) || (temp1==1 || temp==1 )) {
                ards="Yes";
        }else {
            ards="No"
        }
       //fmt.Println(ards)

}

func Decryption_process(id1 *big.Int) (n1 *big.Int) {
        
    n1 = new(big.Int) 
    n1.Mod(id1,pkey[b1]);
    return 
}

func Decryption() {
    
        
        IDd:= new(big.Int) 
        IDd =Decryption_process(IDe)
        
        HGBd:= new(big.Int)
        HGBd =Decryption_process(HGBe)
        
        SAO2d:= new(big.Int) 
        SAO2d =Decryption_process(SAO2e)
        
        SVO2d:= new(big.Int) 
        SVO2d =Decryption_process(SVO2e)
        
        PAO2d:= new(big.Int) 
        PAO2d =Decryption_process(PAO2e)
        
        PVO2d:= new(big.Int) 
        PVO2d =Decryption_process(PVO2e)
        
        COd:= new(big.Int) 
        COd =Decryption_process(COe)
        
        DO2d:= new(big.Int) 
        DO2d =Decryption_process(DO2calce)
        
        VO2d:= new(big.Int) 
        VO2d =Decryption_process(VO2calce)


    fmt.Println()
    fmt.Println("---------------Decrypted values---------------")
    fmt.Print("ID :")
    fmt.Println(IDd)
    fmt.Print("HGB :")
    fmt.Println(HGBd) 
    fmt.Print("SAO2 :")
    fmt.Println(SAO2d)
    fmt.Print("SVO2 :" )
    fmt.Println(SVO2d)
    fmt.Print("PAO2 :")
    fmt.Println(PAO2d)
    fmt.Print("PVO2 :")
    fmt.Println(PVO2d)
    fmt.Print("CO :")
    fmt.Println(COd)
    fmt.Print("DO2 :")
    fmt.Println(DO2d)
    fmt.Print("VO2 :")
    fmt.Println(VO2d)
    fmt.Print("ARDS :")
    fmt.Println(ards)
    fmt.Println("----------------------------------------------")
    fmt.Println()
    b1++;
}

func main() {

    rand.Seed(time.Now().UnixNano())
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
        
            ID = line[0]
            HGB = line[1]
            SAO2 = line[2]
            SVO2 = line[3]
            PAO2 = line[4]
            PVO2 = line[5]
            CO = line[6]
        

        if i==0 {
            i = i + 1
        } else {
        Encryption()
        Analysis()
        Decryption()
    }
    }
    
    
}

