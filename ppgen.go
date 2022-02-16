package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Client struct {
	XMLName     xml.Name `xml:"client"`
	ID          int      `xml:"id"`
	CompanyName string   `xml:"companyname"`
	Adress      string   `xml:"adress"`
	Tarif       int      `xml:"tarif"`
	Balance     int      `xml:"balance"`
}

type Clients struct {
	XMLName xml.Name `xml:"clients"`
	Clients []Client `xml:"client"`
}

func main() {
	filename := "db.xml"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data, _ := ioutil.ReadAll(file)

	var c Clients
	xml.Unmarshal(data, &c)
	//fmt.Println(c.Clients)
	for _, cl := range c.Clients {
		fileNamePP := cl.CompanyName + ".txt"
		filePP, err := os.Create(fileNamePP)
		if err != nil {
			fmt.Println("Ошибка создания файла")
			panic(err)
		}
		defer filePP.Close()
		var dataToFilePP string
		toPay := cl.Tarif - cl.Balance
		if toPay <= 0 {
			dataToFilePP = "Денежных средств достаточно"

		} else {
			strToPay := strconv.Itoa(toPay)
			dataToFilePP = "Необходимо заплатить " + strToPay

		}
		strId := strconv.Itoa(cl.ID)
		strBalance := strconv.Itoa(cl.Balance)
		filePP.WriteString("ID: " + strId + " Клиент: " + cl.CompanyName + " Баланс: " + strBalance + "\n")
		filePP.WriteString(dataToFilePP + "\n")
		//fmt.Printf("Клиент %s требуется заплатить %d\n", cl.CompanyName, cl.Tarif-cl.Balance)
	}
}
