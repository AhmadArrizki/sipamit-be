package seed

import (
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/api/device/repo"
	"sipamit-be/internal/pkg/doc"
	"time"
)

func TOA(db *mongo.Database) {
	toaRepo := repo.NewTOARepository(db)

	count, _ := toaRepo.Count()
	if count > 0 {
		log.Info("TOA already seeded")
		return
	}

	err := toaRepo.InsertMany(toas)
	if err != nil {
		log.Errorf("Failed to seed toa: %v", err)
	}
	log.Info("TOA seeded")
}

var toas = []repo.TOA{
	{
		ID:     bson.NewObjectID(),
		Nama:   "AMPLIFIER",
		Lokasi: "PH1",
		Kode:   "PH1-AMPTOA-001",
		Posisi: "MEJA OPERATOR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "AMPLIFIER",
		Lokasi: "PH1",
		Kode:   "PH2-AMPTOA-002",
		Posisi: "MEJA OPERATOR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "HALAMAN LUAR",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-003",
		Posisi: "POJOK BARAT GEDUNG LINE A",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LASTING B",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-004",
		Posisi: "DIATAS KONVEYOR LASTING B",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "RUANG BORDIR",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-005",
		Posisi: "DIATAS RUANG BORDIR SISI TIMUR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LASTING AB",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-006",
		Posisi: "DIBELAKANG LUAR RUANG EXIM",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GERBANG ATAS JAHIT A",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-007",
		Posisi: "DIATAS GERBANG SELATAN JAHIT A",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "DEPAN KANTOR PRODUKSI",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-008",
		Posisi: "DIATAS DEPAN KANTOR PRODUKSI",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "PACKING C. D. E",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-009",
		Posisi: "DIATAS PINTU LUAR LORONG CRIME",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LASTING C. D. E",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-010",
		Posisi: "DIBELAKANG LUAR RUANG HRD",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GERBANG LINE E",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-011",
		Posisi: "DIATAS PANEL LISTRIK LINE E",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "DEPAN SABLON",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-012",
		Posisi: "DIATAS LEMARI TAS ( DEPAN SABLON)",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GUDANG BAHAN",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-013",
		Posisi: "DIATAS POJOK BARAT LUAR OFFICE GUDANG BAHAN",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GUDANG SOLE",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-014",
		Posisi: "DIATAS POJOK TIMUR OFFICE GUDANG OUTSOLE",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "SAMPAH",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-015",
		Posisi: "DIATAS POJOK UTARA GEDUNG OUTSOLE",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GUDANG JADI",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-016",
		Posisi: "DIATAS LUAR GERBANG GUDANG JADI",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "SABLON INSOLE MIDSOLE",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-017",
		Posisi: "DIATAS TANGGA TIMUR INSOLE MIDSOLE",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "OPERATOR",
		Lokasi: "PH1",
		Kode:   "PH1-MICROPHONETOA-001",
		Posisi: "MEJA OPERATOR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LORONG CRIME",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-018",
		Posisi: "DIATAS UJUNG LORONG DEKAT PINTU CRIME",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LORONG OFFICE ( DEPAN DAPUR )",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-019",
		Posisi: "DEPAN PINTU DAPUR LORONG OFFICE",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "JAHIT DESIGN BAWAH",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-020",
		Posisi: "DIATAS PANEL JAHIT DESIGN BAWAH",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LORONG DEPAN K3",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-021",
		Posisi: "DEPAN PINTU K3 LORONG LAB DAN ACCOUNTING",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LINE H",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-022",
		Posisi: "DIATAS KABEL TRAY SISI TIMUR SELATAN LINE H",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "FOXING",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-023",
		Posisi: "DIATAS TEMBOK SISI UTARA FOXING MESIN PRESS",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LINE F",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-024",
		Posisi: "DIATAS POJOK TEMBOK SISI UTARA",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "BENGKEL PISAU",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-025",
		Posisi: "DIATAS DEPAN PINTU BENGKEL PISAU",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "AMPLIFIER TOA",
		Lokasi: "PH2",
		Kode:   "PH2-AMPTOA-001",
		Posisi: "RAK/MEJA OPERATOR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "AMPLIFIER",
		Lokasi: "PH1",
		Kode:   "PH1-AMPTOA-001",
		Posisi: "MEJA OPERATOR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "AMPLIFIER",
		Lokasi: "PH1",
		Kode:   "PH2-AMPTOA-002",
		Posisi: "MEJA OPERATOR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "HALAMAN LUAR",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-003",
		Posisi: "POJOK BARAT GEDUNG LINE A",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LASTING B",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-004",
		Posisi: "DIATAS KONVEYOR LASTING B",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "RUANG BORDIR",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-005",
		Posisi: "DIATAS RUANG BORDIR SISI TIMUR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LASTING AB",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-006",
		Posisi: "DIBELAKANG LUAR RUANG EXIM",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GERBANG ATAS JAHIT A",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-007",
		Posisi: "DIATAS GERBANG SELATAN JAHIT A",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "DEPAN KANTOR PRODUKSI",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-008",
		Posisi: "DIATAS DEPAN KANTOR PRODUKSI",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "PACKING C. D. E",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-009",
		Posisi: "DIATAS PINTU LUAR LORONG CRIME",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LASTING C. D. E",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-010",
		Posisi: "DIBELAKANG LUAR RUANG HRD",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GERBANG LINE E",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-011",
		Posisi: "DIATAS PANEL LISTRIK LINE E",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "DEPAN SABLON",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-012",
		Posisi: "DIATAS LEMARI TAS ( DEPAN SABLON)",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GUDANG BAHAN",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-013",
		Posisi: "DIATAS POJOK BARAT LUAR OFFICE GUDANG BAHAN",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GUDANG SOLE",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-014",
		Posisi: "DIATAS POJOK TIMUR OFFICE GUDANG OUTSOLE",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "SAMPAH",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-015",
		Posisi: "DIATAS POJOK UTARA GEDUNG OUTSOLE",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GUDANG JADI",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-016",
		Posisi: "DIATAS LUAR GERBANG GUDANG JADI",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "SABLON INSOLE MIDSOLE",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-017",
		Posisi: "DIATAS TANGGA TIMUR INSOLE MIDSOLE",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "OPERATOR",
		Lokasi: "PH1",
		Kode:   "PH1-MICROPHONETOA-001",
		Posisi: "MEJA OPERATOR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LORONG CRIME",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-018",
		Posisi: "DIATAS UJUNG LORONG DEKAT PINTU CRIME",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LORONG OFFICE ( DEPAN DAPUR )",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-019",
		Posisi: "DEPAN PINTU DAPUR LORONG OFFICE",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "JAHIT DESIGN BAWAH",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-020",
		Posisi: "DIATAS PANEL JAHIT DESIGN BAWAH",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LORONG DEPAN K3",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-021",
		Posisi: "DEPAN PINTU K3 LORONG LAB DAN ACCOUNTING",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LINE H",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-022",
		Posisi: "DIATAS KABEL TRAY SISI TIMUR SELATAN LINE H",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "FOXING",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-023",
		Posisi: "DIATAS TEMBOK SISI UTARA FOXING MESIN PRESS",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LINE F",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-024",
		Posisi: "DIATAS POJOK TEMBOK SISI UTARA",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "BENGKEL PISAU",
		Lokasi: "PH1",
		Kode:   "PH1-TOA-025",
		Posisi: "DIATAS DEPAN PINTU BENGKEL PISAU",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "AMPLIFIER TOA",
		Lokasi: "PH2",
		Kode:   "PH2-AMPTOA-001",
		Posisi: "RAK/MEJA OPERATOR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "EQUALIZER TOA",
		Lokasi: "PH2",
		Kode:   "PH2-EQUALIZER-002",
		Posisi: "RAK/MEJA OPERATOR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LORONG OFFICE PH2",
		Lokasi: "PH2",
		Kode:   "PH2-TOA-001",
		Posisi: "DIBAWAH TANGGA LORONG OFFICE",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GUDANG JADI ATAS PH2",
		Lokasi: "PH2",
		Kode:   "PH2-TOA-002",
		Posisi: "DIATAS LORONG SISI SELATAN",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "PLONG A PH2",
		Lokasi: "PH2",
		Kode:   "PH2-TOA-003",
		Posisi: "DIATAS POJOK TIMUR PLONG A",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "DESIGN PH2",
		Lokasi: "PH2",
		Kode:   "PH2-TOA-004",
		Posisi: "DIATAS LORONG SISI GUDANG BAHAN (ZAHRO)",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "PRODUKSI LINE A PH2",
		Lokasi: "PH2",
		Kode:   "PH2-TOA-005",
		Posisi: "DIATAS PLONG A HADAP UTARA LINE A",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "PRODUKSI LINE B PH2",
		Lokasi: "PH2",
		Kode:   "PH2-TOA-006",
		Posisi: "DIATAS PLONG A HADAP SELATAN LINE B",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LASTING C PH2",
		Lokasi: "PH2",
		Kode:   "PH2-TOA-007",
		Posisi: "DIATAS SISI TIMUR TEMBOK LASTING C",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "JAHIT KOMPUTER D PH2",
		Lokasi: "PH2",
		Kode:   "PH2-TOA-008",
		Posisi: "DIATAS JAHIT KOMPUTER D HADAP UTARA",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "JAHIT D / OUTSOLE PH2",
		Lokasi: "PH2",
		Kode:   "PH2-TOA-009",
		Posisi: "DIATAS DEPAN LIFT BARANG LINE D HADAP TIMUR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "OPERATOR",
		Lokasi: "PH2",
		Kode:   "PH2-MICROPHONETOA-001",
		Posisi: "DIMEJA KOMPUTER RESTU AMELIA",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LUAR ATAS HANGGAR BC",
		Lokasi: "PH3",
		Kode:   "PH3-TOA-001",
		Posisi: "DISISI LUAR GEDUNG ATAS HANGGAR BC HADAP TIMUR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LAMINATING",
		Lokasi: "PH3",
		Kode:   "PH3-TOA-002",
		Posisi: "DI TIANG TENGAH LAMINATING HADAP UTARA",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "INSOLE",
		Lokasi: "PH3",
		Kode:   "PH3-TOA-003",
		Posisi: "DI TIANG TENGAH INSOLE HADAP UTARA",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "OUTSOLE",
		Lokasi: "PH3",
		Kode:   "PH3-TOA-004",
		Posisi: "DI TIANG TENGAH OUTSOLE SISI DEPAN/UTARA",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "OUTSOLE TENGAH",
		Lokasi: "PH3",
		Kode:   "PH3-TOA-005",
		Posisi: "DI TIANG TENGAH OUTSOLE SISI TENGAH/SELATAN",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "SESEP OUTSOLE",
		Lokasi: "PH3",
		Kode:   "PH3-TOA-006",
		Posisi: "DI POJOK TIMUR SELATAN ATAS HADAP UTARA",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LINE C.D SELATAN",
		Lokasi: "PH3",
		Kode:   "PH3-TOA-007",
		Posisi: "DIPOJOK BARAT SELATAN ATAS HADAP UTARA",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LINE C.D TENGAH",
		Lokasi: "PH3",
		Kode:   "PH3-TOA-008",
		Posisi: "DI ATAS GERBANG BARAT HADAP TIMUR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "LINE A",
		Lokasi: "PH3",
		Kode:   "PH3-TOA-009",
		Posisi: "DI ATAS GERBANG BARAT HADAP TIMUR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
	{
		ID:     bson.NewObjectID(),
		Nama:   "GUDANG BAHAN BARAT",
		Lokasi: "PH3",
		Kode:   "PH3-TOA-010",
		Posisi: "DI ATAS GERBANG BARAT HADAP TIMUR",
		Inserted: doc.ByAt{
			At: time.Now(),
		},
		IsDeleted: false,
	},
}
