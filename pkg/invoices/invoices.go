package invoice

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func GenerateInvoiceData(prefix string) InvoiceData {
	// Obter a data atual
	currentTime := time.Now()

	// Encontrar o primeiro dia do próximo mês
	nextMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)

	// Encontrar o sétimo dia do próximo mês para o Due Date
	dueDate := time.Date(nextMonth.Year(), nextMonth.Month(), 7, 0, 0, 0, 0, time.UTC)

	// Gerar o número da fatura com base no mês e ano atual
	invoiceNumber := fmt.Sprintf(prefix+"-%d-%02d", currentTime.Year(), int(currentTime.Month()-1))

	// Formatar as datas para o formato desejado (DD MMM YYYY)
	invoiceDateStr := nextMonth.Format("02 Jan 2006")
	dueDateStr := dueDate.Format("02 Jan 2006")

	// Preenchendo os dados da fatura
	invoice := InvoiceData{
		Date:    strings.ToUpper(invoiceDateStr),
		DueDate: strings.ToUpper(dueDateStr),
		Number:  invoiceNumber,
	}
	return invoice
}

func (h *HourControl) getTotalValue() float64 {
	return h.HourlyRate * h.QuantityHoursWorked
}
func (h *HourControl) convertStringTimeToDecimalValue() (float64, error) {

	parts := strings.Split(h.WorkedHoursStr, ":")

	if len(parts) != 3 {
		return 0, fmt.Errorf("formato inválido")
	}

	hours, err := strconv.Atoi(parts[0])

	if err != nil {
		return 0, err
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, err
	}

	hoursWorked := float64(hours) + float64(minutes)/60 + float64(seconds)/3600
	return hoursWorked, nil
}

func roundUp(input float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))
	return math.Ceil(input*multiplier) / multiplier
}

func HoursControlStr(workedHoursStr string, tax float64) HourControlStr {

	hour := HoursControl(workedHoursStr, tax)
	hourStr := HourControlStr{}
	hourStr.AmountCalculed = fmt.Sprintf("%.2f", hour.AmountCalculed)
	hourStr.HourlyRate = fmt.Sprintf("%.2f", hour.HourlyRate)
	hourStr.QuantityHoursWorked = fmt.Sprintf("%.2f", hour.QuantityHoursWorked)
	hourStr.WorkedHoursStr = hour.WorkedHoursStr
	return hourStr
}

func HoursControl(workedHoursStr string, tax float64) *HourControl {
	hour := HourControl{}
	hour.WorkedHoursStr = workedHoursStr

	//Quantidade em decimal das horas trabalhas
	totalHour, err := hour.convertStringTimeToDecimalValue()
	if err != nil {
		panic(err)
	}

	hour.QuantityHoursWorked = totalHour
	// quanto será recebido
	hour.HourlyRate = tax
	amoundCalculated := hour.getTotalValue()
	hour.AmountCalculed = roundUp(amoundCalculated, 2)

	return &hour

}
