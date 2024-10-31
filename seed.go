package main

import (
	_db "sipamit-be/internal/db"
	"sipamit-be/internal/seed"
)

func main() {
	db := _db.Connect()

	seed.SuperAdmin(db)
	seed.Checkpoint(db)
	seed.CCTV(db)
	seed.FingerPrint(db)
	seed.KomputerPH1(db)
	seed.KomputerPH2(db)
	seed.Printer(db)
	seed.Telepon(db)
	seed.TOA(db)
	seed.UPS(db)

	return
}
