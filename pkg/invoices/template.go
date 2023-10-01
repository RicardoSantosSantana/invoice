package invoice

import (
	"os"
	"strings"
	"text/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GenerateHTML(h HourControlStr, i InvoiceData, p Params) error {

	InvoiceParams := getInvoiceParams(h, i, p)

	// Carregar o arquivo de template
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		return err
	}

	// Criar um novo arquivo HTML
	filename := strings.Trim(InvoiceParams.InvoiceData.Number, "")
	newFile, err := os.Create("invoice-" + filename + ".html")
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Preencher o template com os dados da fatura
	err = tmpl.Execute(newFile, InvoiceParams)
	if err != nil {
		return err
	}

	// Agora, vamos criar o PDF a partir do arquivo HTML gerado.
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	// Adicionar a página
	page := wkhtmltopdf.NewPage("invoice-" + filename + ".html")
	pdfg.AddPage(page)

	// Criar PDF

	err = pdfg.Create()
	if err != nil {
		panic(err)
	}

	// Salvar PDF

	err = pdfg.WriteFile("invoice-" + filename + ".pdf")
	if err != nil {
		panic(err)
	}

	return nil
}
