package main

import (
	"fmt"
	"math/rand"
	"time"
	//"strconv"
	//"strings"
)

type Pos struct{
	x,y int 
}

var OPCODES = []int{
	101,102,103,104,
}

func PrintArr( arr [4][4]int ){
	fmt.Println("----------")
	for i:=0; i<len(arr); i++{
		fmt.Println(arr[i])
	}
	fmt.Println("----------")
}


func RotateArr( arrInput [4][4]int , code int ) [4][4]int {
	c := [4][4]int{}
	if code == 0 {    // 90C
		for i:=0; i<4; i++{
			for j:=0; j<4; j++{
				c[j][3-i] = arrInput[i][j]
			}
		}
	}else if code == 1{    // 180
		for i:=0; i<4; i++{
			for j:=0; j<4; j++{
				c[3-i][3-j] = arrInput[i][j]
			}
		}
	}else if code == 2{    // 270
		for i:=0; i<4; i++{
			for j:=0; j<4; j++{
				c[3-j][i] = arrInput[i][j]
			}
		}
	}
	
	return c
}

func RandomPos( arrInput [4][4]int  ) [4][4]int {
	tmpArr := make([]Pos, 0)
	for i:=0; i<4; i++{
		for j:=0; j<4; j++{
			if arrInput[i][j] == 0{
				tmpArr = append(tmpArr, Pos{i,j} )
			}
		}
	}
	//fmt.Println( tmpArr )
	
	rand.Seed( time.Now().UnixNano() % time.Now().Unix()  )
	target := tmpArr[rand.Intn( len( tmpArr ))]
	arrInput[ target.x ][ target.y ] = 2
	return arrInput
}

func FuckRow( arrRow[4]int ) [4]int{
	tmpStack := make([]int, 0)
	res := [4]int{}
	tmp := -1
	idx := 0

	for   idx < 4   {
		//fmt.Println("-------------")
		if arrRow[idx] != 0  {
			tmpStack = append( tmpStack, arrRow[idx] )
			tmp = arrRow[idx]
			fmt.Println(idx)
			idx ++
			for idx<4{
				if tmp == arrRow[idx]{
					tmpStack[ len(tmpStack)-1 ] = tmpStack[ len(tmpStack)-1 ] * 2
					tmp = -1
				}else if tmp != arrRow[idx] &&  arrRow[idx] != 0{
					break
				}
				idx ++
			}
			//break
			//fmt.Println( arrRow[idx] )
		}else{
			idx ++
		}
		
	}
	//fmt.Println( idx )
	
	for idx , val := range tmpStack{
		res[idx] = val
	}
	return res
}

func AssembleArr( arrInput [4][4]int , opcode int ) [4][4]int {
	res := [4][4]int{}
	
	switch opcode {
	case 101:
		for i , row := range arrInput{
			tmp := FuckRow( row )
			for j:=0;j<4;j++{
				res[i][j] = tmp[j]
			}
		}
	case 102:
		arrInput = RotateArr( arrInput , 2 )
		for i , row := range arrInput{
			tmp := FuckRow( row )
			for j:=0;j<4;j++{
				res[i][j] = tmp[j]
			}
		}
		res = RotateArr( res , 0 )
	case 103:
		arrInput = RotateArr( arrInput , 1 )
		for i , row := range arrInput{
			tmp := FuckRow( row )
			for j:=0;j<4;j++{
				res[i][j] = tmp[j]
			}
		}
		res = RotateArr( res , 1 )
	case 104:
		arrInput = RotateArr( arrInput , 0 )
		for i , row := range arrInput{
			tmp := FuckRow( row )
			for j:=0;j<4;j++{
				res[i][j] = tmp[j]
			}
		}
		res = RotateArr( res , 2 )
	}
	return res
}
	
func Init( resArr [4][4]int , opcode int) [4][4]int {
	arrCode := [4]int{ 101,102,103,104 }
	flag :=0
	for _,v := range arrCode{
		if opcode == v{
			flag = 1
		}
	}
	if flag ==0{
		fmt.Println( "xxxxxxxx" )
		return [4][4]int{}
	}
	
	

	resArr = AssembleArr( resArr , opcode )
	PrintArr(resArr)

	resArr = RandomPos( resArr )
	PrintArr(resArr)

	
	return resArr

}

func main()  {
	inTxt := ""

	arrMan := RandomPos( [4][4]int{} )
	PrintArr( arrMan )
	fmt.Scanf( "%s", &inTxt )
	for inTxt != "q"{
		fmt.Println( "====== "+inTxt )
		switch inTxt {
		case "1":
			arrMan = Init( arrMan,  101 )
		case "2":
			arrMan = Init( arrMan,  102 )
		case "3":
			arrMan = Init( arrMan,  103 )
		case "4":
			arrMan = Init( arrMan,  104 )
		default:
			fmt.Println("error")
		}
		inTxt = ""
		fmt.Scanf( "%s", &inTxt )
		//fmt.Println("here")
	}
	

}