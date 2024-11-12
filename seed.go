package main

import (
	_db "sipamit-be/internal/db"
	"sipamit-be/internal/seed"
)

func main() {
	seed.SuperAdmin(_db.Client)
	seed.Checkpoint(_db.Client)
	seed.CCTV(_db.Client)
	seed.FingerPrint(_db.Client)
	seed.KomputerPH1(_db.Client)
	seed.KomputerPH2(_db.Client)
	seed.Printer(_db.Client)
	seed.Telepon(_db.Client)
	seed.TOA(_db.Client)
	seed.UPS(_db.Client)

	return
}
