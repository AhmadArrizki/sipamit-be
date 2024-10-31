package seed

import (
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/api/device/repo"
	"sipamit-be/internal/pkg/doc"
	"time"
)

func UPS(db *mongo.Database) {
	upsRepo := repo.NewUPSRepository(db)

	count, _ := upsRepo.Count()
	if count > 0 {
		log.Info("UPS already seeded")
		return
	}

	err := upsRepo.InsertMany(ups)
	if err != nil {
		log.Errorf("Failed to seed ups: %v", err)
	}
	log.Info("UPS seeded")
}

var ups = []repo.UPS{
	{
		ID:         bson.NewObjectID(),
		Nama:       "IT",
		Departemen: "IT",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH1-UPSICASE3100-001",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "QA LAB",
		Departemen: "QA LAB",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH1-UPSICASE3100-002",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "WILDA",
		Departemen: "DESIGN",
		Tipe:       "ICA CT682B",
		NoSeri:     "PH1-ICACT682B-003",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "IZZATIR",
		Departemen: "DESIGN",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH1-PROLINK700SFC-004",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "MESIN GRADING",
		Departemen: "DESIGN",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH1-UPSICASE3100-005",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "PPIC",
		Departemen: "PPIC",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH1-UPSICASE3100-006",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "FEBRI",
		Departemen: "FOXING WARNA",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH1-PROLINK700SFC-012",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "DJUHEDI",
		Departemen: "G. BAHAN",
		Tipe:       "KENIKA KS600",
		NoSeri:     "PH1-KENIKAKS600-013",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "LENI",
		Departemen: "G. BAHAN",
		Tipe:       "PROLINK PRO1200",
		NoSeri:     "PH1-PROLINK1200-036",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "SUSANTI",
		Departemen: "SABLON ATAS",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH1-UPSICASE3100-014",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "DAVID",
		Departemen: "CNC MATRAS",
		Tipe:       "PROLINK PRO1201",
		NoSeri:     "PH1-PROLINK1201-015",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "EXIM",
		Departemen: "EXIM",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH1-UPSICASE3100-016",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "LINA",
		Departemen: "UMUM",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH1-PROLINK700SFC-017",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "PURCHASING",
		Departemen: "PURCHASING",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH1-UPSICASE3100-018",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "HANGGAR",
		Departemen: "BEA CUKAI",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH1-PROLINK700SFC-019",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "WAHYU",
		Departemen: "BORDIR",
		Tipe:       "ICA CT1328B",
		NoSeri:     "PH1-ICACT1328B-020",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "SANTI",
		Departemen: "BORDIR",
		Tipe:       "PROLINK PRO1200",
		NoSeri:     "PH1-PROLINKPRO1200-021",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "AGUS PC 1",
		Departemen: "JAHIT KOMPUTER",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH1-PROLINK700SFC-022",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "AGUS PC 1",
		Departemen: "JAHIT KOMPUTER",
		Tipe:       "KENIKA KS1200",
		NoSeri:     "PH1-KENIKAKS1200-023",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "NINIK",
		Departemen: "JAHIT KOMPUTER",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH1-PROLINK700SFC-024",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "DAVID",
		Departemen: "CNC PC",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH1-PROLINK700SFC-025",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "KIRI (IKA)",
		Departemen: "ACCOUNTING",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH1-ICASE3100-026",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "KANAN (NORMA)",
		Departemen: "ACCOUNTING",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH1-ICASE3100-027",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "SERVER AJC",
		Departemen: "ACCOUNTING",
		Tipe:       "ICA CT1082B",
		NoSeri:     "PH1-ICACT1082B-028",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "DEVI",
		Departemen: "ACCOUNTING",
		Tipe:       "PROLINK 230SFC",
		NoSeri:     "PH1-PROLINK230SFC-029",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "DILLA",
		Departemen: "ACCOUNTING",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH1-PROLINK700SFC-030",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "RINI",
		Departemen: "ACCOUNTING",
		Tipe:       "ICA CT1082B",
		NoSeri:     "PH1-ICACT1082B-031",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "RACHEL",
		Departemen: "HRD",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH1-PROLINK700SFC-032",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "RASMINI",
		Departemen: "HRD",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH1-PROLINK700SFC-033",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "SUCI",
		Departemen: "HRD",
		Tipe:       "KENIKA KS1200",
		NoSeri:     "PH1-KENIKAKS1200-034",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "ENDAH",
		Departemen: "ACCOUNTING",
		Tipe:       "APC BVX700",
		NoSeri:     "PH1-APCBVX700-035",
		Lokasi:     "PH1",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "OFICE",
		Departemen: "OFFICE",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH2-UPSICASE3100-001",
		Lokasi:     "PH2",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "RUANG IT",
		Departemen: "IT",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH2-UPSICASE3100-002",
		Lokasi:     "PH2",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "HERMIN",
		Departemen: "DESIGN",
		Tipe:       "ICA SE3100",
		NoSeri:     "PH2-UPSICASE3100-003",
		Lokasi:     "PH2",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "DWI",
		Departemen: "DESIGN",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH2-PROLINK700SFC-004",
		Lokasi:     "PH2",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "HANGGAR PC1",
		Departemen: "BEA CUKAI",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH2-PROLINK700SFC-005",
		Lokasi:     "PH2",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "HANGGAR PC2",
		Departemen: "BEA CUKAI",
		Tipe:       "PROLINK 700SFC",
		NoSeri:     "PH2-PROLINK700SFC-006",
		Lokasi:     "PH2",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:         bson.NewObjectID(),
		Nama:       "OFFICE (NEW)",
		Departemen: "OFFICE",
		Tipe:       "APC SRT5K",
		NoSeri:     "PH2-APCSRT5K-007",
		Lokasi:     "PH2",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
}