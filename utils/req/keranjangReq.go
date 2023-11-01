package req

type KeranjangRequest struct {
	KasirID      uint `json:"kasirID" from:"KasirID"`
	ProductID    uint `json:"productID" from:"ProductID"`
	JumlahBarang uint `json:"jumlahBarang" from:"JumlahBarang"`
}

type UpdateKeranjangReq struct {
	Status string `json:"status" from:"Status"`
}

