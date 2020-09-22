package product

//ProductDTO ...
type ProductDTO struct {
	ID          uint    `json:"id,string,omitempty"`
	Code        string  `json:"code"`
	Price       float64 `json:"price,string"`
	Description string  `json:"description"`
}
