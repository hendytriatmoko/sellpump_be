package models

type GetKeranjang struct {
	IdUser   string `json:"id_user" form:"id_user"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserHtml struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type KeranjangGet struct {
	IdKeranjang     string  `json:"id_keranjang" form:"id_keranjang"`
	IdUser          string  `json:"id_user" form:"id_user"`
	IdProduk        string  `json:"id_produk" form:"id_produk"`
	NamaProduk      string  `json:"nama_produk" form:"nama_produk"`
	HargaProduk     float64 `json:"harga_produk" form:"harga_produk"`
	DeskripsiProduk string  `json:"deskripsi_produk" form:"deskripsi_produk"`
	GambarProduk    string  `json:"gambar_produk" form:"gambar_produk"`
	BeratProduk     string  `json:"berat_produk" form:"berat_produk"`
	CreatedAt       string  `json:"created_at" form:"created_at"`
	UpdatedAt       string  `json:"updated_at" form:"updated_at"`
	DeletedAt       string  `json:"deleted_at" form:"deleted_at"`
}

type CreateKeranjang struct {
	IdKeranjang string `json:"id_keranjang" form:"id_keranjang"`
	IdUser      string `json:"id_user" form:"id_user"`
	IdProduk    string `json:"id_produk" form:"id_produk"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

type KeranjangCreate struct {
	IdKeranjang string `json:"id_keranjang" form:"id_keranjang"`
	IdUser      string `json:"id_user" form:"id_user"`
	IdProduk    string `json:"id_produk" form:"id_produk"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

type DeleteKeranjang struct {
	IdKeranjang string `json:"id_keranjang" form:"id_keranjang"`
	DeletedAt   string `json:"deleted_at" form:"deleted_at"`
}

type RajaOngkir struct {
	IdProvinsi      string `json:"id_provinsi" form:"id_provinsi"`
	IdCity          string `json:"id_city" form:"id_city"`
	IdKecamatan     string `json:"id_kecamatan" form:"id_kecamatan"`
	Origin          string `json:"origin" form:"origin"`
	OriginType      string `json:"origin_type" form:"origin_type"`
	Destination     string `json:"destination" form:"destination"`
	DestinationType string `json:"destination_type" form:"destination_type"`
	Weight          string `json:"weight" form:"weight"`
}

type CreatePesanan struct {
	IdPesanan string  `json:"id_pesanan" form:"id_pesanan"`
	IdUser    string  `json:"id_user" form:"id_user"`
	IdProduk  string  `json:"id_produk" form:"id_produk"`
	NoInv     string  `json:"no_inv" form:"no_inv"`
	Kuantitas string  `json:"kuantitas" form:"kuantitas"`
	Harga     float64 `json:"harga" form:"harga"`
	Berat     string  `json:"berat" form:"berat"`
	CreatedAt string  `json:"created_at" form:"created_at"`
}

type GetPesanan struct {
	IdUser    string `json:"id_user" form:"id_user"`
	NoInv     string `json:"no_inv" form:"no_inv"`
	CreatedAt string `json:"created_at" form:"created_at"`
}

type PesananGet struct {
	IdPesanan       string  `json:"id_pesanan" form:"id_pesanan"`
	IdUser          string  `json:"id_user" form:"id_user"`
	IdProduk        string  `json:"id_produk" form:"id_produk"`
	NamaProduk      string  `json:"nama_produk" form:"nama_produk"`
	HargaProduk     float64 `json:"harga_produk" form:"harga_produk"`
	DeskripsiProduk string  `json:"deskripsi_produk" form:"deskripsi_produk"`
	GambarProduk    string  `json:"gambar_produk" form:"gambar_produk"`
	BeratProduk     string  `json:"berat_produk" form:"berat_produk"`
	NoInv           string  `json:"no_inv" form:"no_inv"`
	Kuantitas       string  `json:"kuantitas" form:"kuantitas"`
	Harga           float64 `json:"harga" form:"harga"`
	Berat           string  `json:"berat" form:"berat"`
	CreatedAt       string  `json:"created_at" form:"created_at"`
}

type CreateInvoice struct {
	IdInvoice          string `json:"id_invoice" form:"id_invoice"`
	NoInv              string `json:"no_inv" form:"no_inv"`
	IdUser             string `json:"id_user" form:"id_user"`
	Ppn                string `json:"ppn" form:"ppn"`
	NilaiPpn           string `json:"nilai_ppn" form:"nilai_ppn"`
	Total              string `json:"total" form:"total"`
	JumlahPembayaran   string `json:"jumlah_pembayaran" form:"jumlah_pembayaran"`
	IdStatusPembayaran string `json:"id_status_pembayaran" form:"id_status_pembayaran"`
	OngkosKirim        string `json:"ongkos_kirim" form:"ongkos_kirim"`
	NamaEkspedisi      string `json:"nama_ekspedisi" form:"nama_ekspedisi"`
	LayananEkspedisi   string `json:"layanan_ekspedisi" form:"layanan_ekspedisi"`
	Etd                string `json:"etd" form:"etd"`
	IdStatusPengiriman string `json:"id_status_pengiriman" form:"id_status_pengiriman"`
	DetailAlamat       string `json:"detail_alamat" form:"detail_alamat"`
	PesanPembeli       string `json:"pesan_pembeli" form:"pesan_pembeli"`
	CreatedAt          string `json:"created_at" form:"created_at"`
}

type GetInvoice struct {
	IdUser    string `json:"id_user" form:"id_user"`
	NoInv     string `json:"no_inv" form:"no_inv"`
	IdStatusPengiriman string `json:"id_status_pengiriman" form:"id_status_pengiriman"`
	IdStatusPembayaran string `json:"id_status_pembayaran" form:"id_status_pembayaran"`
	CreatedAt          string `json:"created_at" form:"created_at"`
}

type InvoiceGet struct {
	IdInvoice          string `json:"id_invoice" form:"id_invoice"`
	NoInv              string `json:"no_inv" form:"no_inv"`
	NoPo              string `json:"no_po" form:"no_po"`
	IdUser             string `json:"id_user" form:"id_user"`
	Ppn                string `json:"ppn" form:"ppn"`
	NilaiPpn           string `json:"nilai_ppn" form:"nilai_ppn"`
	Total              string `json:"total" form:"total"`
	JumlahPembayaran   string `json:"jumlah_pembayaran" form:"jumlah_pembayaran"`
	BuktiBayar		   string `json:"bukti_bayar" form:"bukti_bayar"`
	TglPembayaran	   string `json:"tgl_pembayaran" form:"tgl_pembayaran"`
	IdStatusPembayaran string `json:"id_status_pembayaran" form:"id_status_pembayaran"`
	AlasanDitolak		string `json:"alasan_ditolak" form:"alasan_ditolak"`
	OngkosKirim        string `json:"ongkos_kirim" form:"ongkos_kirim"`
	NamaEkspedisi      string `json:"nama_ekspedisi" form:"nama_ekspedisi"`
	LayananEkspedisi   string `json:"layanan_ekspedisi" form:"layanan_ekspedisi"`
	Etd                string `json:"etd" form:"etd"`
	NoResi             string `json:"no_resi" form:"no_resi"`
	IdStatusPengiriman string `json:"id_status_pengiriman" form:"id_status_pengiriman"`
	ExpiredPengiriman	string `json:"expired_pengiriman" form:"expired_pengiriman"`
	DetailAlamat       string `json:"detail_alamat" form:"detail_alamat"`
	PesanPembeli       string `json:"pesan_pembeli" form:"pesan_pembeli"`
	Pesanan		       []PesananGet `json:"pesanan" form:"pesanan"`
	CreatedAt          string `json:"created_at" form:"created_at"`
	UpdatedAt          string `json:"updated_at" form:"updated_at"`
}