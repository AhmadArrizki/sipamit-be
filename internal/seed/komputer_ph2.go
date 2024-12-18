package seed

import (
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/api/device/repo"
	"sipamit-be/internal/pkg/doc"
	"time"
)

func KomputerPH2(db *mongo.Database) {
	kph2Repo := repo.NewKomputerPH2Repository(db)

	count, _ := kph2Repo.Count()
	if count > 0 {
		log.Info("Komputer ph2 already seeded")
		return
	}

	err := kph2Repo.InsertMany(komputerPh2)
	if err != nil {
		log.Errorf("Failed to seed komputer ph2: %v", err)
	}
	log.Info("Komputer ph2 seeded")
}

var komputerPh2 = []repo.KomputerPH2{
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER SERVER 1 PACKING",
		Merk:      "BUILD UP DELL  INSPIRON 3881",
		PC:        "PH2-CPUCOREI3-021",
		Monitor:   "PH2-MONITORDELLE2219HN-021",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 500 GB",
		Lokasi:    "SERVER 1 PACKING",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER SERVER 2 PACKING",
		Merk:      "BUILD UP DELL POWER EDGE T30",
		PC:        "PH2-CPUCOREI3-022",
		Monitor:   "PH2-MONITORDELLE2219HN-022",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 500 GB",
		Lokasi:    "SERVER 2 PACKING",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER ZAKIYAH",
		Merk:      "Rakitan",
		PC:        "PH2-CPUCOREI3-001",
		Monitor:   "PH2-BENQG16HDPL-003",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 1 TB",
		Lokasi:    "EXIM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "LAPTOP CANTIKA",
		Merk:      "TOSHIBA SATELITE L505",
		PC:        "PH2-LAPTOPTOSHIBASATELLITEL505-002",
		Monitor:   "-",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "EXIM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "LAPTOP EVA",
		Merk:      "HP 250-G7",
		PC:        "PH2-LAPTOPHP25067-006",
		Monitor:   "-",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 250 GB",
		Lokasi:    "EXIM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "LAPTOP ZAKIYA",
		Merk:      "DELL INSPIRON L5300 SERIES",
		PC:        "PH2-LAPTOPDELLINSPIRON15-003",
		Monitor:   "-",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 1 TB",
		Lokasi:    "EXIM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER AMELIA",
		Merk:      "Rakitan TSX Pro II",
		PC:        "PH2-CPUCOREI3-027",
		Monitor:   "PH2-MONITORLG22MK400H-022",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "EXIM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER RAHAYU",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-005",
		Monitor:   "PH2-LG19M38A-005",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "EXIM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER ALFI",
		Merk:      "RAKITAN DAZUMBA",
		PC:        "PH2-CPUCOREI3-008",
		Monitor:   "PH2-MONITORDELLE2216HV-008",
		CPU:       "Intel Core I5",
		RAM:       "4 GB",
		Internal:  "HDD:250GB",
		Lokasi:    "EXIM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER GIRINDRA",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-025",
		Monitor:   "PH2-MONITORLG19M38A-025",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "EXIM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER SELVI",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-041",
		Monitor:   "PH2-MONITORLG19M38A-032",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "EXIM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER ELY",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-003",
		Monitor:   "PH2-DELL2219HN-021",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD 500 GB + SSD 500GB",
		Lokasi:    "EXIM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER NENI",
		Merk:      "BUILD UP ASUSPRO S641MD",
		PC:        "PH2-CPUCOREI3-017",
		Monitor:   "PH2-MONITORBENQT52WA-008",
		CPU:       "Intel CoreI3",
		RAM:       "4 GB",
		Internal:  "SSD : 480 GB",
		Lokasi:    "PPIC",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER DELLA",
		Merk:      "Rakitan",
		PC:        "PH2-CPUCOREI3-009",
		Monitor:   "PH2-MONITORLG19M38A-009",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 500 GB",
		Lokasi:    "PPIC",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER HENI",
		Merk:      "Rakitan",
		PC:        "PH2-CPUDUALCORE-016",
		Monitor:   "PH2-MONITORDELL1918H-016",
		CPU:       "Intel Core2Duo",
		RAM:       "2 GB",
		Internal:  "SSD 480GB",
		Lokasi:    "PPIC",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER NOVI",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-042",
		Monitor:   "PH2-MONITORLG19M38A-034",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "PPIC",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER FEBRI",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-045",
		Monitor:   "PH2-BENQET0005B-035",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "PPIC",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER AHUN",
		Merk:      "Rakitan",
		PC:        "PH2-CPUCOREI3-028",
		Monitor:   "PH2-MONITORASUSVX207-017",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD 240 GB + SSD 500 GB",
		Lokasi:    "PPIC",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER DELA",
		Merk:      "HP PRO 3340 MT",
		PC:        "PH2-CPUCOREI5-073",
		Monitor:   "PH2-MONITORDELLE2016HV-066",
		CPU:       "Intel Core I5",
		RAM:       "6 GB",
		Internal:  "SSD 250 GB + HDD 500GB",
		Lokasi:    "ACCOUNTING",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER NOVITA",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-068",
		Monitor:   "PH2-HPP191-055",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "ACCOUNTING",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER DWI RATNA",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-029",
		Monitor:   "PH2-LG24MK430H-029",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "ADMIN LABEL",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER NIKE",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-040",
		Monitor:   "PH2-MONITORLG19M38A-036",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "ADMIN - K3",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER FARISAH",
		Merk:      "Rakitan AERO",
		PC:        "PH2-CPUCOREI3-026",
		Monitor:   "PH2-MONITORLG19M38A-026",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "PEMBELIAN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER AYU",
		Merk:      "Rakitan AERO",
		PC:        "PH2-CPUCOREI3-033",
		Monitor:   "PH2-MONITORLG19M38A-027",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "PEMBELIAN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER ALFINA",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-047",
		Monitor:   "PH2-MONITORLG19M38A-037",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "PEMBELIAN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER ROIHATUL",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-046",
		Monitor:   "PH2-MONITORLG19M38A-038",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "PEMBELIAN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER WIWIK",
		Merk:      "Rakitan TSX PRO II",
		PC:        "PH2-CPUCOREI3-023",
		Monitor:   "PH2-MONITORLG19M38A-023",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "HRD",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER FIFIN",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-024",
		Monitor:   "PH2-MONITORLG19M38A-024",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "HRD",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER MELATI",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-030",
		Monitor:   "PH2-MONITORLG19M38A-030",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "HRD",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "LAPTOP OFFICE",
		Merk:      "DELL INSPIRON L5300 SERIES",
		PC:        "PH2-LAPTOPDELLINSPIRONL5-001",
		Monitor:   "-",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "OFFICE",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER SCAN 2",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-018",
		Monitor:   "PH2-MONITORIFOUNDFD199B-018",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 500 GB",
		Lokasi:    "PACKING",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER SCAN 1",
		Merk:      "Rakitan Power Up",
		PC:        "PH2-CPUCORE2DUO-019",
		Monitor:   "PH2-MONITORLG19M38A-022",
		CPU:       "Intel Core2Duo",
		RAM:       "3 GB",
		Internal:  "HDD: 300 GB",
		Lokasi:    "PACKING",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER SCAN 3 (RESTU UMUM)",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-022",
		Monitor:   "PH2-DELL2219HN-019",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "RESTU",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER YUNITA",
		Merk:      "RAKITAN SIMBADDA (MSIPROH410M-B)",
		PC:        "PH2-CPUCOREI3-048",
		Monitor:   "PH2-MONITORLG19M38A-040",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "QA/QC INSPEK ATAS",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER BC1",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-013",
		Monitor:   "PH2-MONITORDELLE2216-013",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "HANGGAR BEA CUKAI",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER BC2",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-014",
		Monitor:   "PH2-MONITORDELLE2216-014",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 1 TB + SSD 240GB",
		Lokasi:    "HANGGAR BEA CUKAI",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER CCTV",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-015",
		Monitor:   "PH2-TVSHARPAQUAOS-095",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 1 TB",
		Lokasi:    "HANGGAR BEA CUKAI",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER TATIK",
		Merk:      "Rakitan",
		PC:        "PH2-CPUCOREI3-007",
		Monitor:   "PH2-DELLE2219HN-007",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 250 GB",
		Lokasi:    "GUDANG BAHAN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER PUTRI",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-043",
		Monitor:   "PH2-LG19M38A-039",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "GUDANG BAHAN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER INTAN",
		Merk:      "Rakitan JETWAY",
		PC:        "PH2-CPUCORE2DUO-005",
		Monitor:   "PH2-DELLE1916HV-005",
		CPU:       "Intel Core2Duo",
		RAM:       "4 GB",
		Internal:  "HDD: 320 GB + SSD: 500 GB",
		Lokasi:    "GUDANG BAHAN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER HERMIN (GRADING 1)",
		Merk:      "Rakitan",
		PC:        "PH2-CPUCOREI5-020",
		Monitor:   "PH2-LEDHPM24FWA-020",
		CPU:       "Intel Core I5",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER JUKI JANNATIN",
		Merk:      "Rakitan",
		PC:        "PH2-CPUCOREI3-017",
		Monitor:   "PH2-LG22MK400H-022",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN JAHIT KOMP",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER FIFI",
		Merk:      "RAKITAN SIMBADDA (MSIPROH410M-B)",
		PC:        "PH2-CPUCOREI3-052",
		Monitor:   "PH2-MONITORLG19M38A-042",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER SHERLY",
		Merk:      "Rakitan",
		PC:        "PH2-CPUCORI3-004",
		Monitor:   "PH2-MONITORLG19M38A-004",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER DYAH",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-044",
		Monitor:   "PH2-MONITORLG19M38A-041",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER SAYUTI",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-036",
		Monitor:   "PH2-MONITORLG19M38A-030",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER AWAN",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-038",
		Monitor:   "PH2-MONITORLG19M38A-031",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER IDA",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-039",
		Monitor:   "PH2-MONITORLG19M38A-033",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN SOP",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER MIKA",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-049",
		Monitor:   "PH2-MONITORLG19M38A-043",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN SOP",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER STEFANI",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-051",
		Monitor:   "PH2-MONITORLG19M38A-050",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER DWI",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-031",
		Monitor:   "PH2-MONITORDELL2219H-012",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN BORDIR",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER AMELIA GRADING 2",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-032",
		Monitor:   "PH2-MONITORLG19M38A-032",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 240 GB",
		Lokasi:    "DESIGN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER LULUK",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-037",
		Monitor:   "PH2-LG22MK400H-030",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "SABLON / FILM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER FILM TAIWAN",
		Merk:      "RAKITAN ARMAGEDDON",
		PC:        "PH2-CPUCOREI3-036",
		Monitor:   "PH2-BENQT52WA-030",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "SABLON / FILM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER SISKIA",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-054",
		Monitor:   "PH2-MONITORLG19M38A-044",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "HDD: 500GB",
		Lokasi:    "OUTSOLE",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER HENGJIANG GRADING 2",
		Merk:      "Rakitan TAIWAN",
		PC:        "PH2-CPUCOREGOLD-050",
		Monitor:   "PH2-MONITORDELLE2219HN-002",
		CPU:       "Intel Core GOLD",
		RAM:       "16 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "DESIGN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER ADMIN GUDANG BAHAN",
		Merk:      "RAKITAN SIMBADDA",
		PC:        "PH2-CPUCOREI3-052",
		Monitor:   "PH2-LG19M38A-045",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500 GB",
		Lokasi:    "GUDANG BAHAN",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER FIDHO",
		Merk:      "RAKITAN GAMEMAX",
		PC:        "PH2-CPUCOREI3-053",
		Monitor:   "PH2-MONITORLG19M38A-046",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD: 500GB",
		Lokasi:    "OUTSOLE",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER MR. CHOU",
		Merk:      "RAKITAN ENLIGHT",
		PC:        "PH2-CPUCOREI3-054",
		Monitor:   "PH2-LG22MR410-047",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD-M2: 500GB",
		Lokasi:    "MANAGER PPIC PH3",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "LAPTOP SINTA",
		Merk:      "HP 14S 1315U",
		PC:        "PH2-LAPTOPHP14S1315U-007",
		Monitor:   "-",
		CPU:       "Intel Core I3",
		RAM:       "4 GB",
		Internal:  "SSD: 500GB",
		Lokasi:    "INSPEK",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "KOMPUTER IRA LAB BOUNDING",
		Merk:      "RAKITAN INFINITY",
		PC:        "PH2-CPUCOREI3-055",
		Monitor:   "PH2-LG19M38A-056",
		CPU:       "Intel Core I3",
		RAM:       "8 GB",
		Internal:  "SSD-M2: 500GB",
		Lokasi:    "LABORATORIUM",
		Inserted:  doc.ByAt{At: time.Now(), ID: &SuperAdminID},
		IsDeleted: false,
	},
}
