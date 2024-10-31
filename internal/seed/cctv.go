package seed

import (
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/api/device/repo"
	"sipamit-be/internal/pkg/doc"
	"time"
)

func CCTV(db *mongo.Database) {
	cctvRepo := repo.NewCCTVRepository(db)

	count, _ := cctvRepo.Count()
	if count > 0 {
		log.Info("CCTV already seeded")
		return
	}

	err := cctvRepo.InsertMany(cctvs)
	if err != nil {
		log.Errorf("Failed to seed cctv: %v", err)
	}
	log.Info("CCTV seeded")
}

var cctvs = []repo.CCTV{
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV GERBANG 1",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-DAHUA-001",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV GERBANG 2",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-HIKVISION-002",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV LASTING AB",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-003",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV G. JADI LOADING 1",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-004",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV G. JADI LOADING 2",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-005",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV G. JADI LOADING 3",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-006",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV G. JADI LOADING 4",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-007",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV G. JADI DALAM",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-008",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV GERBANG PLONG CDE",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-009",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV GERBANG G. BAHAN",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-010",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV GERBANG G. SOLE",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-011",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV DEPAN BENGKEL",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-012	",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV BENGKEL PISAU",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-013",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV JAHIT A",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-014",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV SAMPAH & GERBANG MESS",
		Lokasi: "PH1",
		Kode:   "PH1-CCTV-LVISION-015",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "DVR  GUDANG JADI PH1",
		Lokasi: "PH1",
		Kode:   "PH1-DVRLVISION-030",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "DVR  GUDANG BAHAN PH1",
		Lokasi: "PH1",
		Kode:   "PH1-HIKVISION-031",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "CCTV PRODUKSI 1",
		Lokasi: "PH2",
		Kode:   "PH2-CCTV-DAHUA-001",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "GERBANG UTAMA",
		Lokasi:    "PH2",
		Kode:      "PH2-CCTV-DAHUA-002",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV PRODUKSI 2",
		Lokasi:    "PH2",
		Kode:      "PH2-CCTV-DAHUA-003",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV GERBANG TIMUR",
		Lokasi:    "PH2",
		Kode:      "PH2-CCTV-DAHUA-004",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV GERBANG SELATAN",
		Lokasi:    "PH2",
		Kode:      "PH2-CCTV-DAHUA-005",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV GUDANG JADI",
		Lokasi:    "PH2",
		Kode:      "PH2-CCTV-DAHUA-006",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV AUTO CLAVE",
		Lokasi:    "PH2",
		Kode:      "PH2-CCTV-DAHUA-007",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV LORONG PH2",
		Lokasi:    "PH2",
		Kode:      "PH2-CCTV-DAHUA-008",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV PH-PLAN-B-1",
		Lokasi:    "PH3",
		Kode:      "PH3-CCTV-TECHMA-009",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV PH-PLAN-B-2",
		Lokasi:    "PH3",
		Kode:      "PH3-CCTV-TECHMA-010",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV PH-PLAN-B-3",
		Lokasi:    "PH3",
		Kode:      "PH3-CCTV-TECHMA-011",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV PH-PLAN-B-4",
		Lokasi:    "PH3",
		Kode:      "PH3-CCTV-TECHMA-012",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV PH-PLAN-B-5",
		Lokasi:    "PH3",
		Kode:      "PH3-CCTV-TECHMA-013",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV PH-PLAN-B-6",
		Lokasi:    "PH3",
		Kode:      "PH3-CCTV-TECHMA-014",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV PH-PLAN-B-7",
		Lokasi:    "PH3",
		Kode:      "PH3-CCTV-TECHMA-015",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "CCTV PH-PLAN-B-8",
		Lokasi:    "PH3",
		Kode:      "PH3-CCTV-TECHMA-016",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "NVR HIKVISION",
		Lokasi:    "PH3",
		Kode:      "PH3-DS-7732NI-K4-002",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
	{
		ID:        bson.NewObjectID(),
		Nama:      "DVR DAHUA",
		Lokasi:    "PH3",
		Kode:      "PH2-DH-XVR4108HS-001",
		Inserted:  doc.ByAt{At: time.Now()},
		IsDeleted: false,
	},
}
