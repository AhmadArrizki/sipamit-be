package doc

const (
	AdminRole      = "admin"
	SuperAdminRole = "superadmin"
)

const (
	CCTV        = "cctv"
	Fingerprint = "fingerprint"
	KomputerPH1 = "komputer_ph1"
	KomputerPH2 = "komputer_ph2"
	Printer     = "printer"
	Telepon     = "telepon"
	Toa         = "toa"
	Ups         = "ups"
)

func ValidDevice(device string) bool {
	switch device {
	case CCTV, Fingerprint, KomputerPH1, KomputerPH2, Printer, Telepon, Toa, Ups:
		return true
	default:
		return false
	}
}
