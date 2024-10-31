package seed

import (
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/api/device/repo"
	"sipamit-be/internal/pkg/doc"
	"time"
)

func Printer(db *mongo.Database) {
	printerRepo := repo.NewPrinterRepository(db)

	count, _ := printerRepo.Count()
	if count > 0 {
		log.Info("Printer already seeded")
		return
	}

	err := printerRepo.InsertMany(printers)
	if err != nil {
		log.Errorf("Failed to seed printer: %v", err)
	}
	log.Info("Printer seeded")
}

var printers = []repo.Printer{
	{
		ID:          bson.NewObjectID(),
		Nama:        "LINA",
		Departemen:  "UMUM",
		TipePrinter: "BROTHERDCP-T700-001",
		NoSeri:      "PH1-BROTHERDCP-T700-001",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "NANCY",
		Departemen:  "EXIM",
		TipePrinter: "BROTHERDCP-T420w-006",
		NoSeri:      "PH1-BROTHERDCP-T420W-006",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "CHENNEDY",
		Departemen:  "EXIM",
		TipePrinter: "BROTHERDCP-T220-003",
		NoSeri:      "PH1-BROTHERDCP-T220-003",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "VIOCHA",
		Departemen:  "EXIM",
		TipePrinter: "BROTHERMFC-J3720-002",
		NoSeri:      "PH1-BROTHERMFC-J3720-002",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "AYU",
		Departemen:  "PPIC",
		TipePrinter: "BROTHERDCP-T520W-030",
		NoSeri:      "PH1-BROTHERDCP-T520W-030",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "RASMINI",
		Departemen:  "HRD",
		TipePrinter: "FARGODTC1500-022",
		NoSeri:      "PH1-FARGODTC1500-022",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "RACHEL",
		Departemen:  "HRD",
		TipePrinter: "EPSONLX-310-023",
		NoSeri:      "PH1-EPSONLX-310-023",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "RACHEL",
		Departemen:  "HRD",
		TipePrinter: "EPSONLX-300+II-027",
		NoSeri:      "PH1-EPSONLX-300+II-027",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "RENA",
		Departemen:  "HRD",
		TipePrinter: "BROTHERDCP-T710W-021",
		NoSeri:      "PH1-BROTHERDCP-T710W-021",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "SUCI",
		Departemen:  "HRD",
		TipePrinter: "BROTHERDCP-T220-070",
		NoSeri:      "PH1-BROTHERDCP-T220-070",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "ANISA",
		Departemen:  "QALAB",
		TipePrinter: "BROTHERDCP-T710W-024",
		NoSeri:      "PH1-BROTHERDCP-T710W-024",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "APRIL",
		Departemen:  "PEMBELIAN",
		TipePrinter: "BROTHERDCP-T420W-050",
		NoSeri:      "PH1-BROTHERDCP-T420W-040",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "NIA",
		Departemen:  "PEMBELIAN",
		TipePrinter: "BROTHERDCP-T710W-052",
		NoSeri:      "PH1-BROTHERDCP-T710W-052",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "DESY",
		Departemen:  "PEMBELIAN",
		TipePrinter: "CANON PIXMA IE G3770",
		NoSeri:      "PH1-CANONPIXMAG3770-066",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "ARUM",
		Departemen:  "PEMBELIAN",
		TipePrinter: "CANONLBP6030-009",
		NoSeri:      "PH1-CANONLBP6030-009",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "DHIMA",
		Departemen:  "PEMBELIAN",
		TipePrinter: "EPSONLQ-2180-015",
		NoSeri:      "PH1-EPSONLQ-2180-015",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "META",
		Departemen:  "DC",
		TipePrinter: "EPSON-L3210-051",
		NoSeri:      "PH1-EPSON-L3210-041",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "ETI",
		Departemen:  "PEMBELIAN",
		TipePrinter: "SATOCL4NXPLUS-056",
		NoSeri:      "PH1-SATOCL4NXPLUS-056",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "RICKY",
		Departemen:  "PEMBELIAN",
		TipePrinter: "EPSONC6550A-012",
		NoSeri:      "PH1-EPSONC6550A-012",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "RICKY",
		Departemen:  "PEMBELIAN",
		TipePrinter: "ZEBRAZT420-011",
		NoSeri:      "PH1-ZEBRAZT420-011",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "HERA",
		Departemen:  "DESIGN",
		TipePrinter: "HPLASERJETP1102-016",
		NoSeri:      "PH1-HPLASERJETP1102-016",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "TINA",
		Departemen:  "DESIGN",
		TipePrinter: "EPSON L121",
		NoSeri:      "PH1-EPSONL121-068",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "ANGEL",
		Departemen:  "DESIGN",
		TipePrinter: "PRINTER ZEBRA ZT411NC",
		NoSeri:      "PH2-ZEBRAZT411NC-013",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "VALEN",
		Departemen:  "K3",
		TipePrinter: "EPSON L121-064",
		NoSeri:      "PH1-EPSONL121-064",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "TUTUT",
		Departemen:  "ACCOUNTING",
		TipePrinter: "EPSONLX-300+II-040",
		NoSeri:      "PH1-EPSONLX-300+II-040",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "LILIS",
		Departemen:  "ACCOUNTING",
		TipePrinter: "HPLASERJETPRO-M404N-056",
		NoSeri:      "PH1-HPLASERJETPRO-M404N-052",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "RINI",
		Departemen:  "ACCOUNTING",
		TipePrinter: "EPSONLQ-2190-025",
		NoSeri:      "PH1-EPSONLQ-2190-025",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "ATHIKA",
		Departemen:  "ACCOUNTING",
		TipePrinter: "CANONLBP6030-053",
		NoSeri:      "PH1-CANONLBP6030-053",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "NORMA",
		Departemen:  "ACCOUNTING",
		TipePrinter: "EPSONLX-310-026",
		NoSeri:      "PH1-EPSONLX-310-026",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "FITRI",
		Departemen:  "ACCOUNTING",
		TipePrinter: "EPSONLX-300+II-041",
		NoSeri:      "PH1-EPSONLX-300+II-041",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "SUSWATI",
		Departemen:  "ACCOUNTING",
		TipePrinter: "HPLASERJETPRO-MP706N-028",
		NoSeri:      "PH1-HPLASERJETPRO-MP706N-028",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "ERNA",
		Departemen:  "ACCOUNTING",
		TipePrinter: "EPSONLX-300+II-030",
		NoSeri:      "PH1-EPSONLX-300+II-030",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "ERNA",
		Departemen:  "ACCOUNTING",
		TipePrinter: "EPSONLQ-2190-033",
		NoSeri:      "PH1-EPSONLQ-2190-033",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "IKA",
		Departemen:  "ACCOUNTING",
		TipePrinter: "HPLASERJETPRO-M4003DN-034",
		NoSeri:      "PH1-HPLASERJETPRO-M4003DN-034",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "UMI",
		Departemen:  "ACCOUNTING",
		TipePrinter: "HPLASERJETPRO-M404N-054",
		NoSeri:      "PH1-HPLASERJETPRO-M404N-054",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "DEVI",
		Departemen:  "ACCOUNTING",
		TipePrinter: "BROTHERDCP-T720DW-055",
		NoSeri:      "PH1-BROTHERDCP-T720DW-055",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "KIKI/LIA",
		Departemen:  "OFFICE PRODUKSI",
		TipePrinter: "BROTHERDCP-T220",
		NoSeri:      "PH1-BROTHERDCP-T220-000",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "AGUS",
		Departemen:  "QC\\INSPECT",
		TipePrinter: "EPSON-L110-056",
		NoSeri:      "PH1-EPSON-L110-056",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "WAHYU",
		Departemen:  "BORDIR",
		TipePrinter: "EPSON-L110-049",
		NoSeri:      "PH1-EPSON-L110-049",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "RENI",
		Departemen:  "GUDANG BAHAN",
		TipePrinter: "BROTHERDCP-T220-030",
		NoSeri:      "PH1-BROTHERDCP-T220-030",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "RINA",
		Departemen:  "GUDANG BAHAN",
		TipePrinter: "CANONLBP6030-060",
		NoSeri:      "PH1-CANONLBP6030-060",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "INTAN",
		Departemen:  "SABLON ATAS",
		TipePrinter: "CANONLBP6030-057",
		NoSeri:      "PH1-CANONLBP6030-057",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "SUSI",
		Departemen:  "SABLON ATAS",
		TipePrinter: "BROTHERDCP-T220-058",
		NoSeri:      "PH1-BROTHERDCP-T220-058",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "RENA",
		Departemen:  "GUDANG SOLE",
		TipePrinter: "BROTHERDCP-T520W-059",
		NoSeri:      "PH1-BROTHERDCP-T520W-059",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "Mr.Lie",
		Departemen:  "FOXING",
		TipePrinter: "CANON PIXMA IE G3770",
		NoSeri:      "PH1-CANONPIXMAG3770-065",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "FEBRY",
		Departemen:  "FOXING WARNA",
		TipePrinter: "CANON MG2470",
		NoSeri:      "PH1-CANON-MG2470-069",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "ANGGRIANI",
		Departemen:  "PACKING LINE A",
		TipePrinter: "ZEBRA ZT-410",
		NoSeri:      "PH1-ZEBRA-ZT410-006",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "ANGGRIANI",
		Departemen:  "PACKING LINE B",
		TipePrinter: "AVERY DENNISON",
		NoSeri:      "PH1-AVERY DENNISON-042",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "GITA",
		Departemen:  "PACKING LINE F",
		TipePrinter: "ZEBRA ZT-230",
		NoSeri:      "PH1-ZEBRA-ZT230-007",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "CICI",
		Departemen:  "PACKING LINE H",
		TipePrinter: "AVERY DENNISON",
		NoSeri:      "PH1-AVERY DENNISON-043",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "DESIGN",
		Departemen:  "DESIGN",
		TipePrinter: "CANON PIXMA IE G3770",
		NoSeri:      "PH1-CANONPIXMAG3770-067",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "AGUS / LILIS",
		Departemen:  "QALAB",
		TipePrinter: "EPSON L121",
		NoSeri:      "PH1-EPSON L121-068",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "LINDA",
		Departemen:  "QALAB",
		TipePrinter: "BROTHERDCP-T700W-004",
		NoSeri:      "PH1-BROTHERDCP-T700W-004",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "HANGGAR BC",
		Departemen:  "HANGGAR BC",
		TipePrinter: "BROTHERDCP-T700W-000",
		NoSeri:      "PH1-BROTHERDCP-T700W-004",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "OFFICE",
		Departemen:  "OFFICE",
		TipePrinter: "PRINTER BROTHER DCP-T720W",
		NoSeri:      "PH2-BROTHER-T720W-010",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "OFFICE",
		Departemen:  "OFFICE",
		TipePrinter: "PRINTER CANON MF643CDW",
		NoSeri:      "PH2-CANONMF643CDW-019",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "AYU",
		Departemen:  "PEMBELIAN",
		TipePrinter: "PRINTER LX310 II",
		NoSeri:      "PH2-EPSONLX310ii-011",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "DWI RATNA",
		Departemen:  "OFFICE",
		TipePrinter: "PRINTER SATO CL4NXPLUS",
		NoSeri:      "PH2-SATOCL4NXPLUS-012",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "DWI RATNA",
		Departemen:  "OFFICE",
		TipePrinter: "PRINTER EPSON L121",
		NoSeri:      "PH2-EPSONL121-027",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "FARISAH",
		Departemen:  "PEMBELIAN",
		TipePrinter: "PRINTER BROTHER DCP-T720W",
		NoSeri:      "PH2-BROTHER-T720W-020",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "WIWIK",
		Departemen:  "HRD",
		TipePrinter: "PRINTER BROTHER DCP-T220",
		NoSeri:      "PH2-BROTHERT220-015",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "SELVY",
		Departemen:  "EXIM",
		TipePrinter: "PRINTER BROTHER MFCT4500DW",
		NoSeri:      "PH2-BROTHER MFCT4500DW-018",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "ALFY",
		Departemen:  "EXIM",
		TipePrinter: "PRINTER CANON LBP6030",
		NoSeri:      "PH2-CANONLBP6030-023",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "LABEL OFFICE",
		Departemen:  "OFFICE",
		TipePrinter: "PRINTER SATO CL4NXPLUS",
		NoSeri:      "PH2-SATOCL4NXPLUS-020",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "LABEL OFFICE",
		Departemen:  "OFFICE",
		TipePrinter: "PRINTER ZEBRA ZT420",
		NoSeri:      "PH2-ZEBRA-ZT420-021",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "LABEL OFFICE",
		Departemen:  "OFFICE",
		TipePrinter: "PRINTER ZEBRA ZT420",
		NoSeri:      "PH2-ZEBRA-ZT420-022",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "Printer Scan 1",
		Departemen:  "OFFICE / LABEL",
		TipePrinter: "PRINTER ZEBRA ZM600",
		NoSeri:      "PH2-ZEBRA-ZM600-006",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "IRA",
		Departemen:  "QA LAB",
		TipePrinter: "PRINTER BROTHER DCP-T310",
		NoSeri:      "PH2-BROTHERDCP-T310-008",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "HERMIN",
		Departemen:  "DESIGN",
		TipePrinter: "PRINTER HP LASERJET 107A",
		NoSeri:      "PH2-HPLASER107A-017",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "STEFANI",
		Departemen:  "DESIGN",
		TipePrinter: "PRINTER BROTHER DCP-T420W",
		NoSeri:      "PH2-BROTHER-T420W-014",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "SOP",
		Departemen:  "DESIGN",
		TipePrinter: "PRINTER BROTHER DCP-T720DW",
		NoSeri:      "PH2-BROTHER-T720DW-026",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "Printer Scan 2",
		Departemen:  "PACKING A",
		TipePrinter: "PRINTER ZEBRA ZT410",
		NoSeri:      "PH2-ZEBRA-ZT410-007",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "FIDHO",
		Departemen:  "OUTSOLE",
		TipePrinter: "PRINTER EPSON L121",
		NoSeri:      "PH2-EPSONL121-016",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "TATIK",
		Departemen:  "GUDANG BAHAN",
		TipePrinter: "PRINTER BROTHER DCP-T310",
		NoSeri:      "PH3-BROTHERDCP-T310-004",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "DELA",
		Departemen:  "ACCOUNTING",
		TipePrinter: "PRINTER CANON LBP6030",
		NoSeri:      "PH3-CANONLBP6030-024",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "PPIC",
		Departemen:  "OFFICE PH3",
		TipePrinter: "PRINTER BROTHER DCP-T720DW",
		NoSeri:      "PH3-BROTHER-T720DW-025",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:          bson.NewObjectID(),
		Nama:        "BC PH2",
		Departemen:  "BEA CUKAI",
		TipePrinter: "PRINTER BROTHER DCP-T710W",
		NoSeri:      "PH2-BROTHERDCP-T710W-005",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
}
