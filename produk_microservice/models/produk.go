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
	HargaDiskon		    string `json:"harga_diskon" form:"harga_diskon"`
	Diskon			    string `json:"diskon" form:"diskon"`
	BoolDiskon		    string `json:"bool_diskon" form:"bool_diskon"`
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
	HargaDiskon		    string `json:"harga_diskon" form:"harga_diskon"`
	Diskon			    string `json:"diskon" form:"diskon"`
	BoolDiskon		    string `json:"bool_diskon" form:"bool_diskon"`
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
	HargaDiskon		    string `json:"harga_diskon" form:"harga_diskon"`
	Diskon			    string `json:"diskon" form:"diskon"`
	BoolDiskon		    string `json:"bool_diskon" form:"bool_diskon"`
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
	HargaDiskon		    string `json:"harga_diskon" form:"harga_diskon"`
	Diskon			    string `json:"diskon" form:"diskon"`
	BoolDiskon		    string `json:"bool_diskon" form:"bool_diskon"`
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
	HargaDiskon		    string `json:"harga_diskon" form:"harga_diskon"`
	Diskon			    string `json:"diskon" form:"diskon"`
	BoolDiskon		    string `json:"bool_diskon" form:"bool_diskon"`
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

type CreateKhusus struct {
	IdUser       		string `json:"id_user" form:"id_user"`
	NamaKhusus      	string `json:"nama_khusus" form:"nama_khusus"`
	HargaKhusus    		float64 `json:"harga_khusus" form:"harga_khusus"`
	BeratKhusus     	string `json:"berat_khusus" form:"berat_khusus"`
	HargaAsli    		float64 `json:"harga_asli" form:"harga_asli"`
	Diskon	    		float64 `json:"diskon" form:"diskon"`
	DeskripsiKhusus 	string `json:"deskripsi_khusus" form:"deskripsi_khusus"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type KhususCreate struct {
	IdKhusus			string `json:"id_khusus" form:"id_khusus"`
	IdUser       		string `json:"id_user" form:"id_user"`
	NamaKhusus      	string `json:"nama_khusus" form:"nama_khusus"`
	HargaKhusus    		float64 `json:"harga_khusus" form:"harga_khusus"`
	HargaAsli    		float64 `json:"harga_asli" form:"harga_asli"`
	Diskon	    		float64 `json:"diskon" form:"diskon"`
	BeratKhusus     	string `json:"berat_khusus" form:"berat_khusus"`
	DeskripsiKhusus 	string `json:"deskripsi_khusus" form:"deskripsi_khusus"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type GetKhusus struct {
	IdKhusus			string `json:"id_khusus" form:"id_khusus"`
	IdUser       		string `json:"id_user" form:"id_user"`
	Limit     			*int64 `json:"limit" form:"limit"`
	Offset    			*int64 `json:"offset" form:"offset"`
}


type KhususGet struct {
	IdKhusus			string `json:"id_khusus" form:"id_khusus"`
	IdUser       		string `json:"id_user" form:"id_user"`
	NamaKhusus      	string `json:"nama_khusus" form:"nama_khusus"`
	HargaKhusus    		float64 `json:"harga_khusus" form:"harga_khusus"`
	HargaAsli    		float64 `json:"harga_asli" form:"harga_asli"`
	Diskon	    		float64 `json:"diskon" form:"diskon"`
	BeratKhusus     	string `json:"berat_khusus" form:"berat_khusus"`
	DeskripsiKhusus 	string `json:"deskripsi_khusus" form:"deskripsi_khusus"`
	ProdukKhusus		[]ProdukKhususGet `json:"produk_khusus" form:"produk_khusus"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
	UpdatedAt        	string `json:"updated_at" form:"updated_at"`
}

type DeleteKhusus struct {
	IdKhusus			string `json:"id_khusus" form:"id_khusus"`
	DeletedAt        	string `json:"deleted_at" form:"deleted_at"`
}

type CreateProdukKhusus struct {
	IdKhusus			string `json:"id_khusus" form:"id_khusus"`
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	Kuantitas       	string  `json:"kuantitas" form:"kuantitas"`
	Harga		       	string  `json:"harga" form:"harga"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type ProdukKhususCreate struct {
	IdProdukKhusus		string `json:"id_produk_khusus" form:"id_produk_khusus"`
	IdKhusus			string `json:"id_khusus" form:"id_khusus"`
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	Kuantitas       	string  `json:"kuantitas" form:"kuantitas"`
	Harga		       	string  `json:"harga" form:"harga"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type GetProdukKhusus struct {
	IdProdukKhusus		string `json:"id_produk_khusus" form:"id_produk_khusus"`
	IdKhusus			string `json:"id_khusus" form:"id_khusus"`
	Limit     			*int64 `json:"limit" form:"limit"`
	Offset    			*int64 `json:"offset" form:"offset"`
}

type ProdukKhususGet struct {
	IdProdukKhusus		string `json:"id_produk_khusus" form:"id_produk_khusus"`
	IdKhusus			string `json:"id_khusus" form:"id_khusus"`
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	Kuantitas       	string  `json:"kuantitas" form:"kuantitas"`
	Harga		       	string  `json:"harga" form:"harga"`
	Produk				[]ProdukGet `json:"produk" form:"produk"`
	CreatedAt          	string `json:"created_at" form:"created_at"`
	UpdatedAt          	string `json:"updated_at" form:"updated_at"`
}

type CreateRating struct {
	IdUser				string `json:"id_user" form:"id_user"`
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	Nama        		string `json:"nama" form:"nama"`
	Komentar        	string `json:"komentar" form:"komentar"`
	Rating        		string `json:"rating" form:"rating"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type RatingCreate struct {
	IdRating			string `json:"id_rating" form:"id_rating"`
	IdUser				string `json:"id_user" form:"id_user"`
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	Nama        		string `json:"nama" form:"nama"`
	Komentar        	string `json:"komentar" form:"komentar"`
	Rating        		string `json:"rating" form:"rating"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type GetRating struct {
	IdRating			string `json:"id_rating" form:"id_rating"`
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	Limit     			*int64 `json:"limit" form:"limit"`
	Offset    			*int64 `json:"offset" form:"offset"`
}

type RatingGet struct {
	IdRating			string `json:"id_rating" form:"id_rating"`
	IdUser				string `json:"id_user" form:"id_user"`
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	Nama        		string `json:"nama" form:"nama"`
	Komentar        	string `json:"komentar" form:"komentar"`
	Rating        		string `json:"rating" form:"rating"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
	UpdatedAt        	string `json:"updated_at" form:"updated_at"`
	DeletedAt        	string `json:"deleted_at" form:"deleted_at"`
}

type CreateFile struct {
	NamaFile	      	string `json:"nama_file" form:"nama_file"`
	File		    	*multipart.FileHeader `json:"file" form:"file"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type FileCreate struct {
	IdFile				string `json:"id_file" form:"id_file"`
	NamaFile	      	string `json:"nama_file" form:"nama_file"`
	File		    	string `json:"file" form:"file"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}

type GetFile struct {
	IdFile				string `json:"id_file" form:"id_file"`
	Search       		string `json:"search" form:"search"`
	Limit     			*int64 `json:"limit" form:"limit"`
	Offset    			*int64 `json:"offset" form:"offset"`
}

type FileGet struct {
	IdFile				string `json:"id_file" form:"id_file"`
	NamaFile	      	string `json:"nama_file" form:"nama_file"`
	File		    	string `json:"file" form:"file"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
	DeletedAt        	string `json:"deleted_at" form:"deleted_at"`
}