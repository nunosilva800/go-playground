package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	{
		input := ""
		b, err := NewBook(input)
		if err != nil {
			fmt.Println(err)
		} else {
			panic("wanted an error, got none")
		}

		report(b)
	}

	println("\n******************************************\n")

	{
		input := "Perky Blenders Coffee, 56101, 100000, B, 8%\nLaird Hatters LTD, 14190, 25000, C, 15%, 75000, 20%"
		b, err := NewBook(input)
		if err != nil {
			fmt.Println(err)
		} else {
			panic("wanted an error, got none")
		}

		report(b)
	}

	println("\n******************************************\n")

	{
		input := "Perky Blenders Coffee, 56101, 100000, B, 8.12345654765879%"
		b, err := NewBook(input)
		if err != nil {
			panic("got error: " + err.Error())
		}

		report(b)
	}

	println("\n******************************************\n")

	{
		input := "Perky Blenders Coffee, 56101, 100000, B, 8%\nLaird Hatters LTD, 14190, 25000, C, 15%, 75000, D, 20%"
		b, err := NewBook(input)
		if err != nil {
			panic("got error: " + err.Error())
		}
		report(b)
	}

	println("\n******************************************\n")

	{

		input := "Perky Blenders Coffee, 56101, 100000.10, B, 8%\nLaird Hatters, 14190, 25000, C, 15%, 75000, D, 20%\nBobbin Bicycles, 47990, 100000, C, 15%, 50000, B, 7.5%, 80000, A, 3.2%\nArapina Bakery, 56101, 25000, A+, 1.8%\nCanine Creche, 75000, 60000, B, 8%"
		b, err := NewBook(input)
		if err != nil {
			panic("got error: " + err.Error())
		}
		report(b)
	}
}

func report(b *book) {
	tot := b.TotalAmount()
	println("\n\nTotal amount is: ", tot)

	println("\nCompany average amounts: ")
	compAvg := b.BusinessAverageLoans()
	for name, amount := range compAvg {
		fmt.Printf("%v\t%.2f\n", name, amount)
	}

	println("\nLending by nature: ")
	byNature := b.PercLentPerNature()
	for nature, perc := range byNature {
		fmt.Printf("%v\t%.2f\n", nature, perc)
	}
}

type book struct {
	entries []entry
}

type entry struct {
	name string
	code string

	loans []triplet
}

type triplet struct {
	loanAmount uint
	riskBand   string
	interest   uint // 20% is '20'
}

func NewBook(input string) (*book, error) {
	b := &book{
		entries: []entry{},
	}

	err := b.parse(input)

	return b, err
}

// TotalAmount returns the the total amount we have loaned.
func (b *book) TotalAmount() int {
	total := 0

	for _, e := range b.entries {
		for _, l := range e.loans {
			total += int(l.loanAmount)
		}
	}

	return total / 100
}

// The business name with the average loan amount we have issued them.
func (b *book) BusinessAverageLoans() map[string]float32 {
	result := make(map[string]float32)

	for _, e := range b.entries {
		total := 0
		for _, l := range e.loans {
			total += int(l.loanAmount)
		}

		result[e.name] = float32(total/100) / float32(len(e.loans))
	}

	return result
}

// PercLentPerNature returns the percentage of money lent per each business nature.
func (b *book) PercLentPerNature() map[string]float32 {
	naturesAndTotals := make(map[string]float32)
	for _, e := range b.entries {
		total := 0
		for _, l := range e.loans {
			total += int(l.loanAmount)
		}

		naturesAndTotals[e.code] += float32(total)
	}

	grandTotal := b.TotalAmount()

	result := make(map[string]float32)
	for nature, total := range naturesAndTotals {
		result[nature] = (total / 100) / float32(grandTotal)
	}

	return result
}

// The Total interest earned
// func (b *book) TotalInterestEarned(input string) error {
// }

func (b *book) parse(input string) error {
	lines := strings.Split(input, "\n")

	for idx, l := range lines {
		entry, err := ParseEntry(l)
		if err != nil {
			return fmt.Errorf("parsing entry line #%v: %w", idx, err)
		}
		b.entries = append(b.entries, entry)
	}

	return nil
}

// Laird Hatters LTD, 14190,     25000,        C,        15%,      75000,       D,        20%
// |---------------||-------||-----------||---------||--------||-----------||---------||--------|
//       NAME         CODE    LOAN AMOUNT  RISK BAND  INTEREST  LOAN AMOUNT  RISK BAND  INTEREST
//                           |__________Triplet 1_____________||__________Triplet 2_____________|
func ParseEntry(line string) (entry, error) {
	if line == "" {
		return entry{}, errors.New("line is empty")
	}

	parts := strings.Split(line, ",")
	entry := entry{
		name:  parts[0],
		code:  parts[1],
		loans: nil,
	}

	if (len(parts)-2)%3 != 0 {
		return entry, errors.New("format of tripets is incorrect (missing data)")
	}

	for i := 2; i < len(parts); i += 3 {
		loanAmount := strings.TrimSpace(parts[i])
		riskBand := strings.TrimSpace(parts[i+1])
		interest := strings.TrimSpace(parts[i+2])

		loanAmountVal, err := strconv.ParseFloat(loanAmount, 64)
		if err != nil {
			return entry, fmt.Errorf("extracting loan amount: %w", err)
		}

		interestParts := strings.Split(interest, "%")
		interestVal, err := strconv.ParseFloat(interestParts[0], 64)
		if err != nil {
			return entry, fmt.Errorf("extracting interest value: %w", err)
		}

		newLoan := triplet{
			loanAmount: uint(loanAmountVal * 100),
			riskBand:   riskBand,
			interest:   uint(interestVal),
		}

		entry.loans = append(entry.loans, newLoan)
	}

	return entry, nil
}
