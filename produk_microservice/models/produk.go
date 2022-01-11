package models

import (
	"mime/multipart"
)

type CreateProduk struct {
	NamaProduk      	string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    		float64 `json:"harga_produk" form:"harga_produk"`
	KodeProduk      	string `json:"kode_produk" form:"kode_produk"`
	TipeProduk      	string `json:"tipe_produk" form:"tipe_produk"`
	BeratProduk     	string `json:"berat_produk" form:"berat_produk"`
	IdMerkProduk   	 	string `json:"id_merk_produk" form:"id_merk_produk"`
	DeskripsiProduk 	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	StokProduk      	int `json:"stok_produk" form:"stok_produk"`
	IdKategoriProduk    string `json:"id_kategori_produk" form:"id_kategori_produk"`
	LinkTokopedia  		string `json:"link_tokopedia" form:"link_tokopedia"`
	LinkBukalapak	    string `json:"link_bukalapak" form:"link_bukalapak"`
	GambarProduk      	*multipart.FileHeader `json:"gambar_produk" form:"gambar_produk"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}
type ProdukCreate struct {
	IdProduk			string `json:"id_produk" form:"id_produk"`
	NamaProduk      	string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    		float64 `json:"harga_produk" form:"harga_produk"`
	KodeProduk      	string `json:"kode_produk" form:"kode_produk"`
	TipeProduk      	string `json:"tipe_produk" form:"tipe_produk"`
	BeratProduk     	string `json:"berat_produk" form:"berat_produk"`
	IdMerkProduk   	 	string `json:"id_merk_produk" form:"id_merk_produk"`
	DeskripsiProduk 	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	StokProduk      	int `json:"stok_produk" form:"stok_produk"`
	Tayangan      		int `json:"tayangan" form:"tayangan"`
	IdKategoriProduk    string `json:"id_kategori_produk" form:"id_kategori_produk"`
	LinkTokopedia  		string `json:"link_tokopedia" form:"link_tokopedia"`
	LinkBukalapak	    string `json:"link_bukalapak" form:"link_bukalapak"`
	GambarProduk      	string `json:"gambar_produk" form:"gambar_produk"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}
type UpdateProduk struct {
	IdProduk       	 	string `json:"id_produk" form:"id_produk"`
	NamaProduk      	string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    		float64 `json:"harga_produk" form:"harga_produk"`
	KodeProduk      	string `json:"kode_produk" form:"kode_produk"`
	TipeProduk      	string `json:"tipe_produk" form:"tipe_produk"`
	BeratProduk     	string `json:"berat_produk" form:"berat_produk"`
	IdMerkProduk   	 	string `json:"id_merk_produk" form:"id_merk_produk"`
	DeskripsiProduk 	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	StokProduk      	int `json:"stok_produk" form:"stok_produk"`
	Tayangan      		int `json:"tayangan" form:"tayangan"`
	IdKategoriProduk    string `json:"id_kategori_produk" form:"id_kategori_produk"`
	GambarProduk      	*multipart.FileHeader `json:"gambar_produk" form:"gambar_produk"`
	LinkTokopedia  		string `json:"link_tokopedia" form:"link_tokopedia"`
	LinkBukalapak	    string `json:"link_bukalapak" form:"link_bukalapak"`
}
type ProdukUpdate struct {
	IdProduk			string `json:"id_produk" form:"id_produk"`
	NamaProduk      	string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    		float64 `json:"harga_produk" form:"harga_produk"`
	KodeProduk      	string `json:"kode_produk" form:"kode_produk"`
	TipeProduk      	string `json:"tipe_produk" form:"tipe_produk"`
	BeratProduk     	string `json:"berat_produk" form:"berat_produk"`
	IdMerkProduk   	 	string `json:"id_merk_produk" form:"id_merk_produk"`
	DeskripsiProduk 	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	StokProduk      	int `json:"stok_produk" form:"stok_produk"`
	Tayangan      		int `json:"tayangan" form:"tayangan"`
	IdKategoriProduk    string `json:"id_kategori_produk" form:"id_kategori_produk"`
	GambarProduk      	*string `json:"gambar_produk" form:"gambar_produk"`
	LinkTokopedia  		string `json:"link_tokopedia" form:"link_tokopedia"`
	LinkBukalapak	    string `json:"link_bukalapak" form:"link_bukalapak"`
	UpdatedAt        	string `json:"updated_at" form:"updated_at"`
}

type GetProduk struct {
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	Search       		string `json:"search" form:"search"`
	SubKategori	    	string `json:"sub_kategori" form:"sub_kategori"`
	Limit     			*int64 `json:"limit" form:"limit"`
	Offset    			*int64 `json:"offset" form:"offset"`
}


type ProdukGet struct {
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	NamaProduk      	string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    		string `json:"harga_produk" form:"harga_produk"`
	KodeProduk      	string `json:"kode_produk" form:"kode_produk"`
	TipeProduk      	string `json:"tipe_produk" form:"tipe_produk"`
	BeratProduk     	string `json:"berat_produk" form:"berat_produk"`
	IdMerkProduk   	 	string `json:"id_merk_produk" form:"id_merk_produk"`
	NamaMerk   	 		string `json:"nama_merk" form:"nama_merk"`
	GambarMerk   	 	string `json:"gambar_merk" form:"gambar_merk"`
	DeskripsiProduk 	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	StokProduk      	string `json:"stok_produk" form:"stok_produk"`
	Tayangan      		string `json:"tayangan" form:"tayangan"`
	IdKategoriProduk    string `json:"id_kategori_produk" form:"id_kategori_produk"`
	Kategori    		string `json:"kategori" form:"kategori"`
	SubKategori	    	string `json:"sub_kategori" form:"sub_kategori"`
	GambarProduk      	string `json:"gambar_produk" form:"gambar_produk"`
	LinkTokopedia  		string `json:"link_tokopedia" form:"link_tokopedia"`
	LinkBukalapak	    string `json:"link_bukalapak" form:"link_bukalapak"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
	UpdatedAt        	string `json:"updated_at" form:"updated_at"`
	DeletedAt        	string `json:"deleted_at" form:"deleted_at"`
}

type DeleteProduk struct {
	IdProduk   string `json:"id_produk" form:"id_produk"`
	DeletedAt  string `json:"deleted_at" form:"deleted_at"`
}

type CreateArtikel struct {
	JudulArtikel      	string `json:"judul_artikel" form:"judul_artikel"`
	GambarArtikel    	*multipart.FileHeader `json:"gambar_artikel" form:"gambar_artikel"`
	DeskripsiArtikel    string `json:"deskripsi_artikel" form:"deskripsi_artikel"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type ArtikelCreate struct {
	IdArtikel			string `json:"id_artikel" form:"id_artikel"`
	JudulArtikel      	string `json:"judul_artikel" form:"judul_artikel"`
	GambarArtikel    	string `json:"gambar_artikel" form:"gambar_artikel"`
	DeskripsiArtikel    string `json:"deskripsi_artikel" form:"deskripsi_artikel"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type GetArtikel struct {
	IdArtikel			string `json:"id_artikel" form:"id_artikel"`
	Search       		string `json:"search" form:"search"`
	Limit     			*int64 `json:"limit" form:"limit"`
	Offset    			*int64 `json:"offset" form:"offset"`
}

type ArtikelGet struct {
	IdArtikel			string `json:"id_artikel" form:"id_artikel"`
	JudulArtikel      	string `json:"judul_artikel" form:"judul_artikel"`
	GambarArtikel    	string `json:"gambar_artikel" form:"gambar_artikel"`
	DeskripsiArtikel    string `json:"deskripsi_artikel" form:"deskripsi_artikel"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
	UpdatedAt        	string `json:"updated_at" form:"updated_at"`
	DeletedAt        	string `json:"deleted_at" form:"deleted_at"`
}

type DeleteArtikel struct {
	IdArtikel			string `json:"id_artikel" form:"id_artikel"`
	DeletedAt        	string `json:"deleted_at" form:"deleted_at"`
}

type UpdateArtikel struct {
	IdArtikel			string `json:"id_artikel" form:"id_artikel"`
	JudulArtikel      	string `json:"judul_artikel" form:"judul_artikel"`
	GambarArtikel    	*multipart.FileHeader `json:"gambar_artikel" form:"gambar_artikel"`
	DeskripsiArtikel    string `json:"deskripsi_artikel" form:"deskripsi_artikel"`
}
type ArtikelUpdate struct {
	IdArtikel			string `json:"id_artikel" form:"id_artikel"`
	JudulArtikel      	string `json:"judul_artikel" form:"judul_artikel"`
	GambarArtikel    	*string `json:"gambar_artikel" form:"gambar_artikel"`
	DeskripsiArtikel    string `json:"deskripsi_artikel" form:"deskripsi_artikel"`
	UpdatedAt        	string `json:"updated_at" form:"updated_at"`
}