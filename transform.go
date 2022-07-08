package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func unmarshalToStruct(src *string) []*Lot {
	lots := []*Lot{}

	csv, err := os.OpenFile(*src, os.O_RDONLY, 0666)
	check(err)
	defer csv.Close()

	scanner := bufio.NewScanner(csv)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")

		_lot := split[0]
		_item := split[1]

		lot := &Lot{
			Lot:  _lot,
			Item: _item,
		}

		if split[2] != "" {
			_purchase_order := split[2]
			lot.PurchaseOrder = _purchase_order
		}
		if split[4] != "" {
			_sterilization_date := dateConverter(split[4])
			lot.SterilizationDate = _sterilization_date
		}
		if split[5] != "" {
			clean_string := strings.Split(strings.ReplaceAll(split[5], ",", ""), ".")[0]
			_on_hand_qty, err := strconv.Atoi(clean_string)
			check(err)
			lot.OnHandQuantity = _on_hand_qty
		}
		if split[6] != "" {
			_use_exp_date := dateConverter(split[6])
			lot.UseExpirationDate = _use_exp_date
		}

		lots = append(lots, lot)

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lots
}
