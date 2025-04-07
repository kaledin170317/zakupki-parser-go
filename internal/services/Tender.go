package services

import (
	"ZakupkiParser/internal/models"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseTender(html string) (*models.Tender, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	tender := &models.Tender{}
	params := parseParameters(doc)

	tender.Title = params["title"]
	tender.Subtitle = params["subtitle"]

	tender.NoticeNumber = params["Номер извещения"]
	tender.ObjectName = params["Наименование объекта закупки"]
	tender.Method = params["Способ определения поставщика (подрядчика, исполнителя)"]
	tender.PlacingOrganization = params["Размещение осуществляет"]
	tender.DeliveryPlace = params["Место поставки товара, выполнения работы или оказания услуги"]

	tender.EPlatform = models.ElectronicPlatform{
		Name: params["Наименование электронной площадки в информационно-телекоммуникационной сети «Интернет»"],
		URL:  params["Адрес электронной площадки в информационно-телекоммуникационной сети «Интернет»"],
	}

	tender.Contact = models.ContactInfo{
		Organization:      params["Организация, осуществляющая размещение"],
		PostalAddress:     params["Почтовый адрес"],
		Location:          params["Место нахождения"],
		ResponsiblePerson: params["Ответственное должностное лицо"],
		Email:             params["Адрес электронной почты"],
		Phone:             params["Номер контактного телефона"],
		Fax:               params["Факс"],
	}

	tender.Procedure = models.ProcedureInfo{
		ApplicationDeadline: params["Дата и время окончания подачи заявок"],
		ProposalDate:        params["Дата рассмотрения заявок"],
		ResultsDate:         params["Дата подведения итогов определения поставщика (подрядчика, исполнителя)"],
	}

	tender.Contract = models.ContractInfo{
		MaxPrice: parseMoney(params["Начальная (максимальная) цена контракта"]),
		Currency: "RUB",
	}

	tender.Execution = models.ExecutionInfo{
		StartDate:    params["Дата начала исполнения контракта"],
		EndDate:      params["Срок исполнения контракта"],
		BudgetFunded: params["Закупка за счет бюджетных средств"],
		OwnFunded:    params["Закупка за счет собственных средств организации"],
	}

	tender.Security = models.ApplicationSecurityInfo{
		Amount: parseMoney(params["Обеспечение исполнения контракта"]),
	}

	items, err := parseItemsAndCharacteristics(doc)
	if err != nil {
		return nil, err
	}

	tender.Items = items
	return tender, nil
}

func parseParameters(doc *goquery.Document) map[string]string {
	params := make(map[string]string)
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		// Параметры с классами parameter/parameterValue
		param := strings.TrimSpace(s.Find("p.parameter").Text())
		value := normalizeWhitespace(s.Find("p.parameterValue").Text())
		if param != "" {
			params[param] = value
		}

		// Параметры с caption + следующий параметр
		if caption := strings.TrimSpace(s.Find("p.caption").Text()); caption != "" {
			next := s.Next()
			if next != nil {
				val := normalizeWhitespace(next.Find("p.parameter").Text())
				if val != "" {
					params[caption] = val
				}
			}
		}
	})

	params["title"] = normalizeWhitespace(strings.TrimSpace(doc.Find("p.title").Text()))
	params["subtitle"] = normalizeWhitespace(strings.TrimSpace(doc.Find("p.subtitle").Text()))

	return params
}

func parseMoney(value string) float64 {

	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, "\u00a0", "")
	value = strings.ReplaceAll(value, " ", "")
	value = strings.ReplaceAll(value, ",", ".")

	for i := 0; i < len(value); i++ {
		if (value[i] < '0' || value[i] > '9') && value[i] != '.' {
			value = value[:i]
			break
		}
	}

	amount, _ := strconv.ParseFloat(value, 64)

	return amount
}

func normalizeWhitespace(input string) string {
	// Удаляет лишние пробелы, символы табуляции и переводы строк

	fields := strings.Fields(input)
	text := strings.Join(fields, " ")
	text = strings.TrimSpace(text)
	return text
}

func parseItemsAndCharacteristics(doc *goquery.Document) ([]models.ProcurementItem, error) {
	var items []models.ProcurementItem
	var currentCharacteristics [][]models.ItemCharacteristic

	// Парсинг таблиц с характеристиками
	doc.Find("table.table.font14").Each(func(i int, table *goquery.Selection) {

		header := table.Find("tr").First().Text()

		if strings.Contains(header, "Характеристики товара") {

			icols := make(map[string]int)

			table.Find("tr").Next().Find("th").FilterFunction(func(i int, th *goquery.Selection) bool {
				return th.Find("table").Length() == 0
			}).Each(func(i int, th *goquery.Selection) {
				icols[normalizeWhitespace(th.Text())] = i
				//text := strings.TrimSpace(th.Text())
				//fmt.Printf("Столбец %d: %s\n", i, text)
			})

			var characteristics []models.ItemCharacteristic
			table.Find("tr").NextAll().Each(func(j int, row *goquery.Selection) {
				tds := row.Find("td")
				if tds.Length() == 0 {
					return
				}
				if tds.Length() == 1 {
					if len(characteristics) > 0 {
						characteristics[len(characteristics)-1].Value += "; " + strings.TrimSpace(tds.Eq(0).Text())
					}
				} else {

					c := models.ItemCharacteristic{
						Name:        normalizeWhitespace(tds.Eq(icols["Наименование характеристики"]).Text()),
						Value:       normalizeWhitespace(tds.Eq(icols["Значение характеристики"]).Text()),
						Unit:        normalizeWhitespace(tds.Eq(icols["Единица измерения характеристики"]).Text()),
						Instruction: normalizeWhitespace(tds.Eq(icols["Инструкция по заполнению характеристики в заявке"]).Text()),
					}

					characteristics = append(characteristics, c)
				}
			})
			currentCharacteristics = append(currentCharacteristics, characteristics)
		}
	})

	charIndex := 0
	doc.Find("table.table.font14").Each(func(i int, table *goquery.Selection) {

		header := table.Find("tr").First()

		if strings.Contains(header.Text(), "Наименование товара") {

			icols := make(map[string]int)

			table.Find("tr").First().Find("th").FilterFunction(func(i int, th *goquery.Selection) bool {
				return th.Find("table").Length() == 0
			}).Each(func(i int, th *goquery.Selection) {
				icols[normalizeWhitespace(th.Text())] = i
			})

			table.Find("tr").NextAll().Each(func(j int, row *goquery.Selection) {

				cols := row.Find("td")

				if cols.Length() < 7 {
					return
				}

				nameHtml, _ := cols.Eq(icols["Наименование товара, работы, услуги"]).Html()
				nameParts := strings.Split(nameHtml, "<br/>")

				name := normalizeWhitespace(nameParts[0])
				identifier := normalizeWhitespace(nameParts[1])

				code := normalizeWhitespace(cols.Eq(icols["Код позиции"]).Text())
				itemType := normalizeWhitespace(cols.Eq(icols["Тип позиции"]).Text())
				unit := normalizeWhitespace(cols.Eq(icols["Единица измерения"]).Text())
				unitPrice := parseMoney(cols.Eq(icols["Цена за единицу"]).Text())
				customer := normalizeWhitespace(cols.Eq(icols["Заказчик"]).Text())
				quantity := parseMoney(cols.Eq(icols["Количество (объем работы, услуги)"]).Text())
				totalPrice := parseMoney(cols.Eq(icols["Стоимость позиции"]).Text())

				item := models.ProcurementItem{
					Name:            name,
					Identifier:      identifier,
					Code:            code,
					ItemType:        itemType,
					Unit:            unit,
					Customer:        customer,
					UnitPrice:       unitPrice,
					Quantity:        quantity,
					TotalPrice:      totalPrice,
					Characteristics: []models.ItemCharacteristic{},
				}

				if charIndex < len(currentCharacteristics) {
					item.Characteristics = currentCharacteristics[charIndex]
					charIndex++
				}

				items = append(items, item)
			})
		}

	})

	return items, nil
}
