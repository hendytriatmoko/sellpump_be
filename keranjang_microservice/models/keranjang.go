package models

type GetKeranjang struct {
	IdUser      string `json:"id_user" form:"id_user"`
	Email      	string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
}

type KeranjangGet struct {
	IdKeranjang    	string `json:"id_keranjang" form:"id_keranjang"`
	IdUser         	string `json:"id_user" form:"id_user"`
	IdProduk   		string `json:"id_produk" form:"id_produk"`
	NamaProduk      string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    	float64 `json:"harga_produk" form:"harga_produk"`
	DeskripsiProduk	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	GambarProduk	string `json:"gambar_produk" form:"gambar_produk"`
	CreatedAt       string `json:"created_at" form:"created_at"`
	UpdatedAt       string `json:"updated_at" form:"updated_at"`
	DeletedAt       string `json:"deleted_at" form:"deleted_at"`
}

type CreateKeranjang struct {
	IdKeranjang string `json:"id_keranjang" form:"id_keranjang"`
	IdUser      string `json:"id_user" form:"id_user"`
	IdProduk   	string `json:"id_produk" form:"id_produk"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

type KeranjangCreate struct {
	IdKeranjang string `json:"id_keranjang" form:"id_keranjang"`
	IdUser      string `json:"id_user" form:"id_user"`
	IdProduk   	string `json:"id_produk" form:"id_produk"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

type DeleteKeranjang struct {
	IdKeranjang   	string `json:"id_keranjang" form:"id_keranjang"`
	DeletedAt  		string `json:"deleted_at" form:"deleted_at"`
}

type RajaOngkir struct {
	IdProvinsi			string `json:"id_provinsi" form:"id_provinsi"`
	IdCity     			string `json:"id_city" form:"id_city"`
	IdKecamatan			string `json:"id_kecamatan" form:"id_kecamatan"`
	Origin				string `json:"origin" form:"origin"`
	OriginType 			string `json:"origin_type" form:"origin_type"`
	Destination			string `json:"destination" form:"destination"`
	DestinationType		string `json:"destination_type" form:"destination_type"`
	Weight				string `json:"weight" form:"weight"`
}
