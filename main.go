package main

import (
	"flag"

	"github.com/k0kubun/pp"
	"gorm.io/gorm/clause"
)

func main() {
	importBool := flag.Bool("import", false, "Import data from CSV")
	filePath := flag.String("fp", "", "file path")
	lotString := flag.String("lot", "", "lot to find")

	flag.Parse()

	if (*importBool && *filePath == "") || (!*importBool && *lotString == "") {
		flag.Usage()
		return
	}

	if *importBool && *filePath != "" {
		cleaned := clean_lot_csv(filePath)
		lots := unmarshalToStruct(&cleaned)

		db.Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(&lots)

		return
	}

	if !*importBool && *lotString != "" {
		l := Lot{Lot: *lotString}
		l.getFromDatabase()
		pp.Print(l)

		return
	}
}
