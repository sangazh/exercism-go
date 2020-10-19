package ledger

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type LocaleType string
type Currency string

const (
	localeUS LocaleType = "en-US"
	localeNL LocaleType = "nl-NL"

	currencyUS  Currency = "USD"
	currencyEUR Currency = "EUR"

	USD = "$"
	EUR = "€"
)

func (c Currency) Symbol() string {
	switch c {
	case currencyUS:
		return USD
	case currencyEUR:
		return EUR
	}
	return ""
}

var emptyErr = errors.New("")

type EntryFormatter interface {
	Locale() LocaleType
	FormatDate(date time.Time) string
	FormatCurrency(change int) string
	Title() string
}

func NewFormatter(locale LocaleType, currency Currency) (EntryFormatter, error) {
	if len(currency.Symbol()) == 0 {
		return nil, emptyErr
	}

	switch locale {
	case localeUS:
		return &USEntry{
			locale:   locale,
			currency: currency,
		}, nil
	case localeNL:
		return &NLEntry{
			locale:   locale,
			currency: currency,
		}, nil
	default:
		return nil, emptyErr
	}
}

type USEntry struct {
	locale   LocaleType
	currency Currency
}

func (USEntry) Locale() LocaleType {
	return localeUS
}
func (e USEntry) Currency() Currency {
	return e.currency
}

func (e USEntry) Title() string {
	var (
		date   = "Date"
		desc   = "Description"
		change = "Change"
	)
	return fmt.Sprintf("%-10s | %-25s | %s\n", date, desc, change)
}

func (USEntry) FormatDate(date time.Time) string {
	return date.Format("01/02/2006")
}

func (e USEntry) FormatCurrency(cents int) string {
	negative := false
	if cents < 0 {
		cents = cents * -1
		negative = true
	}

	a := formatMoney(cents, ".", ",")

	aa := fmt.Sprintf("%s%s", e.currency.Symbol(), a)
	if negative {
		aa = fmt.Sprintf("(%s)", aa)
	} else {
		aa += " "
	}

	return aa
}

type NLEntry struct {
	locale   LocaleType
	currency Currency
}

func (NLEntry) Locale() LocaleType {
	return localeNL
}

func (e NLEntry) Currency() Currency {
	return e.currency
}

func (e NLEntry) FormatDate(date time.Time) string {
	return date.Format("02-01-2006")
}

func (e NLEntry) Title() string {
	var (
		date   = "Datum"
		desc   = "Omschrijving"
		change = "Verandering"
	)
	return fmt.Sprintf("%-10s | %-25s | %s\n", date, desc, change)
}

func (e NLEntry) FormatCurrency(cents int) string {
	negative := false
	if cents < 0 {
		cents = cents * -1
		negative = true
	}

	a := formatMoney(cents, ",", ".")
	aa := fmt.Sprintf("%s %s", e.currency.Symbol(), a)
	if negative {
		aa += "-"
	} else {
		aa += " "
	}

	return aa
}

func formatMoney(money int, separatorCents, separatorThousand string) string {
	cents := money % 100
	dollars := money / 100
	parts := make([]string, 0)
	for dollars > 1000 {
		i := dollars % 1000
		parts = append([]string{strconv.Itoa(i)}, parts...)

		dollars = dollars / 1000
	}
	parts = append([]string{strconv.Itoa(dollars)}, parts...)

	a := fmt.Sprintf("%s%s%02d", strings.Join(parts, separatorThousand), separatorCents, cents)
	return a
}

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
	date        time.Time
}

func (e *Entry) CheckValid() bool {
	if err := e.ParseDate(); err != nil {
		return false
	}
	return true
}

func (e *Entry) ParseDate() (err error) {
	e.date, err = time.Parse("2006-01-02", e.Date)
	if err != nil {
		return err
	}
	return nil
}

type chanStruct struct {
	i    int
	text string
	err  error
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	formatter, err := NewFormatter(LocaleType(locale), Currency(currency))
	if err != nil {
		return "", emptyErr
	}

	var entriesCopy []Entry
	for _, e := range entries {
		entriesCopy = append(entriesCopy, e)
	}
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}
	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	es := entriesCopy
	for len(es) > 1 {
		first, rest := es[0], es[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
					m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
					m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
					es[0], es[i+1] = es[i+1], es[0]
					success = false
				}
			}
		}
		es = es[1:]
	}
	s := formatter.Title()

	// Parallelism, always a great idea
	co := make(chan chanStruct)
	for i, et := range entriesCopy {
		go handle(formatter, i, et, co)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.err != nil {
			return "", v.err
		}
		ss[v.i] = v.text
	}
	for i := 0; i < len(entriesCopy); i++ {
		s += ss[i]
	}
	return s, nil
}

func handle(formatter EntryFormatter, i int, entry Entry, co chan chanStruct) {
	if !entry.CheckValid() {
		co <- chanStruct{err: emptyErr}
	}

	dateS := formatter.FormatDate(entry.date)
	a := formatter.FormatCurrency(entry.Change)

	line := FormatLine(dateS, entry.Description, a)
	co <- chanStruct{i: i, text: line}
}

func FormatLine(date, desc, change string) string {
	if len(desc) > 25 {
		desc = desc[:22] + "..."
	}
	return fmt.Sprintf("%-10s | %-25s | %13s\n", date, desc, change)
}
