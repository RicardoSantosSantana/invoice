package invoice

type HourControl struct {
	WorkedHoursStr      string  // string no formato "140:00:00"
	QuantityHoursWorked float64 // quantidade de horas trabalhadas com base no WorkedHoursStr
	AmountCalculed      float64 // valor total de horas trabalhadas e calculadas HoursWorked * HourlyRate
	HourlyRate          float64 // taxa cobrada por hora X euros
}

type HourControlStr struct {
	WorkedHoursStr      string
	QuantityHoursWorked string
	AmountCalculed      string
	HourlyRate          string
}

type InvoiceData struct {
	Number  string
	Date    string
	DueDate string
}

type Image struct {
	Logo string
}

type InvoiceParams struct {
	HourControlStr
	InvoiceData
	Image
	Params
}

type Params struct {
	WorkedHoursStr string
	Tax            float64
	Prefix         string
	Url            string
}

func getInvoiceParams(h HourControlStr, i InvoiceData, p Params) InvoiceParams {

	image := Image{
		Logo: logoImage(),
	}
	return InvoiceParams{
		HourControlStr: h,
		InvoiceData:    i,
		Image:          image,
		Params:         p,
	}
}
