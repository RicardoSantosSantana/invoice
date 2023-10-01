package invoice

import (
	"fmt"
	"regexp"
	"strings"
)

func green(s string) string {
	return "\033[32m" + s + "\033[0m"

}
func isValidWorkedHours(workedHours string) bool {
	// Regex para o formato HHH:MM:SS
	validWorkedHours := regexp.MustCompile(`\d{2,3}:\d{2}:\d{2}`)
	return validWorkedHours.MatchString(workedHours)
}

func RequestParams() Params {
	var workedHoursStr, prefix, url string
	var tax float64

	for {
		fmt.Print(green("Por favor, informe a quantidade de horas que você trabalhou no formato 00:00:00: "))
		fmt.Scanf("%s", &workedHoursStr)

		if isValidWorkedHours(workedHoursStr) {
			break
		} else {
			fmt.Println("Formato de horas inválido. Tente novamente.")
		}
	}

	for {
		fmt.Print(green("Informe o seu valor hora no formato 00.00: "))
		_, err := fmt.Scanf("%f", &tax)

		if err == nil && tax > 0 {
			break
		} else {
			fmt.Println("Taxa inválida. Tente novamente.")
		}
	}

	fmt.Print(green("Informe o prefixo do seu invoice number, como RS: "))
	fmt.Scanf("%s", &prefix)

	fmt.Print(green("Informe o caminho fisico do seu arquivo de invoice que vc baixou do Hubstaff: "))
	fmt.Scanf("%s", &url)

	return Params{
		WorkedHoursStr: workedHoursStr,
		Tax:            tax,
		Prefix:         strings.ToUpper(prefix),
		Url:            url,
	}
}
