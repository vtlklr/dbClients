package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {

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

	v := &Clients{}
	var answer string
	for answer != "n" {
		fmt.Printf("Добавить клиента? (y/n)")

		fmt.Scan(&answer)
		if answer == "n" {
			fmt.Println("Выход")
			continue
		}
		var c1 Client
		fmt.Println("Введите ID:")
		fmt.Scan(&c1.ID)
		fmt.Println("Введите название компании:")
		fmt.Scan(&c1.CompanyName)
		fmt.Println("Введите адрес:")
		fmt.Scan(&c1.Adress)
		fmt.Println("Введите тариф:")
		fmt.Scan(&c1.Tarif)
		fmt.Println("Введите баланс:")
		fmt.Scan(&c1.Balance)
		v.Clients = append(v.Clients, c1)
	}
	filename := "db.xml"
	file, _ := os.Create(filename)

	xmlWriter := io.Writer(file)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

}
