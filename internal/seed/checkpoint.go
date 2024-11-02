package seed

import (
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/api/device_cp/repo"
	"sipamit-be/internal/pkg/doc"
)

func Checkpoint(db *mongo.Database) {
	cpRepo := repo.NewCheckpointRepository(db)

	count, _ := cpRepo.Count()
	if count > 0 {
		log.Info("Checkpoint already seeded")
		return
	}

	err := cpRepo.InsertMany(cps)
	if err != nil {
		log.Errorf("Failed to seed checkpoints: %v", err)
	}
	log.Info("Checkpoint seeded")
}

var cps = []repo.Checkpoint{
	{
		ID:     bson.NewObjectID(),
		Device: doc.CCTV,
		Checkpoint: []string{
			"Kebersihan Perangkat",
			"Penyimpanan histori",
			"Tampilan Gambar",
		},
	},
	{
		ID:     bson.NewObjectID(),
		Device: doc.Fingerprint,
		Checkpoint: []string{
			"Jam Finger",
			"Kondisi Sensor",
			"Kondisi Tombol",
			"Kondisi Baterai",
			"Kapasitas Foto",
			"Kebersihan perangkat",
		},
	},
	{
		ID:     bson.NewObjectID(),
		Device: doc.KomputerPH1,
		Checkpoint: []string{
			"Kebersihan perangkat",
			"Kipas pendingin",
			"Fungsi keyboard, mouse",
			"Jam CMOS",
			"Kondisi Harddisk",
			"Software/aplikasi",
			"Antivirus",
			"Backup",
		},
	},
	{
		ID:     bson.NewObjectID(),
		Device: doc.KomputerPH2,
		Checkpoint: []string{
			"Kebersihan perangkat",
			"Kipas pendingin",
			"Fungsi keyboard, mouse",
			"Jam CMOS",
			"Kondisi Harddisk",
			"Software/aplikasi",
			"Antivirus",
			"Backup",
		},
	},
	{
		ID:     bson.NewObjectID(),
		Device: doc.Printer,
		Checkpoint: []string{
			"Kebersihan perangkat",
			"Hasil cetak / scan",
			"Kondisi head print",
			"Kondisi mata pisau",
			"Pembuangan tinta",
			"Software / aplikasi",
		},
	},
	{
		ID:     bson.NewObjectID(),
		Device: doc.Telepon,
		Checkpoint: []string{
			"Kebersihan perangkat",
			"Suara Dering",
			"Suara Panggilan",
			"Fungsi Tombol",
		},
	},
	{
		ID:     bson.NewObjectID(),
		Device: doc.Toa,
		Checkpoint: []string{
			"Kebersihan Perangkat",
			"Suara",
		},
	},
	{
		ID:     bson.NewObjectID(),
		Device: doc.Ups,
		Checkpoint: []string{
			"Kebersihan Perangkat",
			"Pemeriksaan daya",
			"Pengujian daya",
		},
	},
}
