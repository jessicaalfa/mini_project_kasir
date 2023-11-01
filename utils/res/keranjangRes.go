package res

import (
	"kasir/model"
)

type KeranjangRes struct {
	ID           uint                `json:"id"`
	Kasir        KeranjangKasirRes   `json:"kasir"`
	Product      KeranjangProductRes `json:"product"`
	JumlahBarang uint                `json:"jumlahBarang"`
	TotalHarga   uint                `json:"totalHarga"`
	Status       string              `json:"status"`
}

type KeranjangKasirRes struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type KeranjangProductRes struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Harga string `json:"harga"`
}

func (k *KeranjangRes) GetKeranjangRes(keranjang model.Keranjang) {
	k.Kasir = KeranjangKasirRes{
		ID:    keranjang.Kasir.ID,
		Name:  keranjang.Kasir.Name,
		Email: keranjang.Kasir.Email,
	}

	k.Product = KeranjangProductRes{
		ID:    keranjang.Product.ID,
		Name:  keranjang.Product.Name,
		Harga: keranjang.Product.Harga,
	}

	k.JumlahBarang = keranjang.JumlahBarang
	k.TotalHarga = keranjang.TotalHarga
	k.Status = keranjang.Status
}

func GetKeranjangAll(keranjangs []model.Keranjang) []KeranjangRes {
	var keranjangRes []KeranjangRes
	for _, keranjang := range keranjangs {
		var k KeranjangRes
		k.ID = keranjang.ID
		k.Kasir = KeranjangKasirRes{
			ID:    keranjang.Kasir.ID,
			Name:  keranjang.Kasir.Name,
			Email: keranjang.Kasir.Email,
		}

		k.Product = KeranjangProductRes{
			ID:    keranjang.Product.ID,
			Name:  keranjang.Product.Name,
			Harga: keranjang.Product.Harga,
		}

		k.JumlahBarang = keranjang.JumlahBarang
		k.TotalHarga = keranjang.TotalHarga
		k.Status = keranjang.Status
		keranjangRes = append(keranjangRes, k)
	}
	return keranjangRes
}
