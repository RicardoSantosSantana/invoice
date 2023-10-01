package main

import (
	"encoding/base64"
	"fmt"
	invoice "invoices/pkg/invoices"
	"log"
	"os"
)

func main2() {
	imgData, err := os.ReadFile("logo-need.png")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	// Codificar o byte slice para uma string Base64
	encoded := base64.StdEncoding.EncodeToString(imgData)

	// Imprimir a string Base64
	fmt.Println(encoded)
}
func main() {

	params := invoice.RequestParams()

	// params := invoice.Params{}
	// params.Prefix = "RS"
	// params.Tax = 14
	// params.Url = "file:///home/mirian/Downloads/RS-2023-08%20-%20Ricardo%20Santana%20-%20.pdf"
	// params.WorkedHoursStr = "154:05:54"

	hour := invoice.HoursControlStr(params.WorkedHoursStr, params.Tax)
	invoiceData := invoice.GenerateInvoiceData(params.Prefix)

	invoice.GenerateHTML(hour, invoiceData, params)

	fmt.Println(hour)
	fmt.Println(invoiceData)
	fmt.Println(params)

}
