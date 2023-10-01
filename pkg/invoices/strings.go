package invoice

import (
	"fmt"
	"strconv"
)

func (h *HourControl) String() string {
	return fmt.Sprintf(
		"WorkedHoursStr: %s\nQuantityHoursWorked: %s\nAmountCalculed: %s\nHourlyRate: %s\n",
		h.WorkedHoursStr,
		strconv.FormatFloat(h.QuantityHoursWorked, 'f', 2, 64),
		strconv.FormatFloat(h.AmountCalculed, 'f', 2, 64),
		strconv.FormatFloat(h.HourlyRate, 'f', 2, 64),
	)
}

func (h HourControlStr) String() string {
	return fmt.Sprintf(
		"WorkedHoursStr: %s\nQuantityHoursWorked: %s\nAmountCalculed: %s\nHourlyRate: %s\n",
		h.WorkedHoursStr,
		h.QuantityHoursWorked,
		h.AmountCalculed,
		h.HourlyRate,
	)
}

func (i InvoiceData) String() string {
	return fmt.Sprintf(
		"Number: %s\nDate: %s\nDueDate: %s\n",
		i.Number,
		i.Date,
		i.DueDate,
	)
}

func (p Params) String() string {
	return fmt.Sprintf(
		"WorkedHoursStr: %s\nTax: %.2f\nPrefix: %s\nUrl: %s\n",
		p.WorkedHoursStr,
		p.Tax,
		p.Prefix,
		p.Url,
	)
}
