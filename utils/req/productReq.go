package req

type ProductRequest struct {
	Name  string `json:"name" from:"name"`
	Harga string `json:"harga" from:"harga"`
}
