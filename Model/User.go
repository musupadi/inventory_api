package Model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type UserData struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Kontak        string `json:"kontak"`
	Id_role       string `json:"id_role"`
	Role_name     string `json:"role_name"`
	Level         string `json:"level"`
	Photo         string `json:"photo"`
	Id_shop       string `json:"id_shop"`
	Shop_name     string `json:"shop_name"`
}

type ProductData struct {
	Id            string `json:"id"`
	Img           string `json:"img"`
	Label         string `json:"label"`
	Price         string `json:"price"`
	Stock         string `json:"stock"`
	Description   string `json:"description"`
	Id_category   string `json:"id_category"`
	Category_name string `json:"category_name"`
	Unit          string `json:"unit"`
	Id_shop       string `json:"id_shop"`
	Shop_name     string `json:"Shop_name"`
}
type Logins struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Transaction
type Transaction struct {
	Id_user    string `json:"id_user"`
	Id_shop    string `json:"id_shop"`
	Id_product string `json:"id_product"`
	Quantity   string `json:"quantity"`
	Note       string `json:"Note"`
	Id_invoice string `json:"Id_invoice"`
	Created_by string `json:"created_by"`
	Updated_by string `json:"updated_by"`
}
