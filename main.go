package main

import (
	"fmt"
	"net/url"
)

//func main() {
//	fmt.Println("Hello World")
//
//	fz, err := services.Get44FZ("0309600004925000003")
//	if err != nil {
//		fmt.Println(err)
//
//		return
//	}
//	fmt.Println(fz.URL)
//	tender, err := services.ParseTender(fz.HTML)
//	if err != nil {
//		fmt.Println(err)
//	}
//	MongoDB.SaveTender(*tender)
//}

func buildURL(baseURL string, params map[string]string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	query := u.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	u.RawQuery = query.Encode()

	return u.String(), nil
}

func main() {
	base := "https://zakupki.gov.ru/epz/order/extendedsearch/results.html"
	params := map[string]string{
		"morphology":                            "on",
		"sortBy":                                "UPDATE_DATE",
		"pageNumber":                            "1",
		"sortDirection":                         "false",
		"recordsPerPage":                        "_10",
		"fz44":                                  "on",
		"af":                                    "on",
		"ktruCodeNameList":                      "26.20.11.110-00000001&&&Ноутбук",
		"currencyIdGeneral":                     "-1",
		"ktruSelectedPageNum":                   "1",
		"showLotsInfoHidden":                    "false",
		"priceContractAdvantages44IdNameHidden": "{}",
		"selectedSubjectsIdNameHidden":          "{}",
		"koksIdsIdNameHidden":                   "{}",
		"gws":                                   "Выберите тип закупки",
	}

	fullURL, err := buildURL(base, params)
	if err != nil {
		panic(err)
	}

	fmt.Println("Сформированный URL:")
	fmt.Println(fullURL)
}
