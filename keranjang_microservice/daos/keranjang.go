package daos

import (
	"fmt"
	"io/ioutil"
	"keranjang_microservice/databases"
	"keranjang_microservice/helper"
	"keranjang_microservice/models"
	"net/http"
	"os"
	_ "os"
	"path/filepath"
	_ "path/filepath"
	"strings"
	_ "strings"
)

type Keranjang struct {
	helper helper.Helper
}

func (m *Keranjang) KeranjangGet(params models.KeranjangGet) ([]models.KeranjangGet, error) {

	keranjang := []models.KeranjangGet{}

	err := databases.DatabaseSellPump.DB.Table("keranjang").Select("keranjang.*,k.nama_produk,k.harga_produk,k.deskripsi_produk,k.gambar_produk,k.berat_produk, k.harga_diskon, k.diskon, k.bool_diskon").
		Joins("join produk k on k.id_produk = keranjang.id_produk").Order("keranjang.created_at desc")

	if params.IdUser != "" {
		err = err.Where("keranjang.id_user = ?", params.IdUser)
	}
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&keranjang)

	errx := err.Error

	if errx != nil {
		return []models.KeranjangGet{}, errx
	}

	return keranjang, nil
}

func (m *Keranjang) KeranjangCreate(params models.CreateKeranjang) (models.KeranjangCreate, error) {

	keranjang := models.KeranjangCreate{}

	keranjang.IdKeranjang = m.helper.StringWithCharset()
	keranjang.IdProduk = params.IdProduk
	keranjang.IdUser = params.IdUser
	keranjang.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("keranjang").Create(&keranjang).Error

	if err != nil {
		return models.KeranjangCreate{}, err
	}

	return keranjang, nil
}

func (m *Keranjang) KeranjangDelete(params models.DeleteKeranjang) (models.DeleteKeranjang, error) {

	keranjang := models.DeleteKeranjang{}

	keranjang.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("keranjang").Where("id_keranjang = ?", params.IdKeranjang).Update(&keranjang).Error

	if err != nil {
		return models.DeleteKeranjang{}, err
	}

	return keranjang, nil

}

func (m *Keranjang) ProvinceGet(params models.RajaOngkir) (string, error) {

	url := "https://pro.rajaongkir.com/api/province"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "0e33e27e42c55799ebe174e1307f2adf")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	isi := string(body)

	return isi, nil
}

func (m *Keranjang) CityGet(params models.RajaOngkir) (string, error) {

	url := ""
	if params.IdProvinsi == "" && params.IdCity == "" {
		url = "https://pro.rajaongkir.com/api/city"
	}
	if params.IdProvinsi != "" && params.IdCity == "" {
		url = "https://pro.rajaongkir.com/api/city?province=" + params.IdProvinsi
	}
	if params.IdProvinsi == "" && params.IdCity != "" {
		url = "https://pro.rajaongkir.com/api/city?id=" + params.IdCity
	}
	if params.IdProvinsi != "" && params.IdCity != "" {
		url = "https://pro.rajaongkir.com/api/city?id=" + params.IdCity + "&province=" + params.IdProvinsi
	}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "0e33e27e42c55799ebe174e1307f2adf")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	isi := string(body)

	fmt.Println("city" + params.IdCity)
	fmt.Println("province" + params.IdProvinsi)

	return isi, nil
}

func (m *Keranjang) SubdistrictGet(params models.RajaOngkir) (string, error) {

	url := ""
	if params.IdCity != "" && params.IdKecamatan == "" {
		url = "https://pro.rajaongkir.com/api/subdistrict?city=" + params.IdCity
	}
	if params.IdCity != "" && params.IdKecamatan != "" {
		url = "https://pro.rajaongkir.com/api/subdistrict?city=" + params.IdCity + "&id=" + params.IdKecamatan
	}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "0e33e27e42c55799ebe174e1307f2adf")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	isi := string(body)

	fmt.Println("city" + params.IdCity)
	fmt.Println("province" + params.IdProvinsi)

	return isi, nil
}

func (m *Keranjang) CostGet(params models.RajaOngkir) (string, error) {

	url := "https://pro.rajaongkir.com/api/cost"

	payload := strings.NewReader("origin=" + params.Origin + "&originType=subdistrict" +
		"&destination=" + params.Destination + "&destinationType=subdistrict&weight=" + params.Weight +
		"&courier=jne:tiki:pos:sicepat:jnt:wahana:anteraja")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("key", "0e33e27e42c55799ebe174e1307f2adf")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	isi := string(body)

	fmt.Println("city" + params.IdCity)
	fmt.Println("province" + params.IdProvinsi)

	return isi, nil
}

func (m *Keranjang) PesananCreate(params models.CreatePesanan) (models.CreatePesanan, error) {

	pesanan := models.CreatePesanan{}

	pesanan.IdPesanan = m.helper.StringWithCharset()
	pesanan.NoInv = params.NoInv
	pesanan.IdProduk = params.IdProduk
	pesanan.IdUser = params.IdUser
	pesanan.Kuantitas = params.Kuantitas
	pesanan.Harga = params.Harga
	pesanan.Berat = params.Berat
	pesanan.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("pesanan").Create(&pesanan).Error

	if err != nil {
		return models.CreatePesanan{}, err
	}

	return pesanan, nil
}

func (m *Keranjang) PesananGet(params models.GetPesanan) ([]models.PesananGet, error) {

	pesanan := []models.PesananGet{}

	err := databases.DatabaseSellPump.DB.Table("pesanan").Select("pesanan.*,k.nama_produk,k.harga_produk,k.deskripsi_produk,k.gambar_produk,k.berat_produk,k.harga_diskon,k.diskon,k.bool_diskon").
		Joins("join produk k on k.id_produk = pesanan.id_produk")

	if params.IdUser != "" {
		err = err.Where("pesanan.id_user = ?", params.IdUser)
	}
	if params.NoInv != "" {
		err = err.Where("pesanan.no_inv = ?", params.NoInv)
	}
	if params.CreatedAt != "" {
		err = err.Where("pesanan.created_at::text like  ?", "%"+params.CreatedAt+"%")
	}

	err = err.Find(&pesanan)

	errx := err.Error

	if errx != nil {
		return []models.PesananGet{}, errx
	}

	return pesanan, nil
}

func (m *Keranjang) InvoiceCreate(params models.CreateInvoice) (models.CreateInvoice, error) {

	invoice := models.CreateInvoice{}

	invoice.IdInvoice = m.helper.StringWithCharset()
	invoice.NoInv = params.NoInv
	invoice.IdUser = params.IdUser
	invoice.Total = params.Total
	invoice.Ppn = params.Ppn
	invoice.NilaiPpn = params.NilaiPpn
	invoice.OngkosKirim = params.OngkosKirim
	invoice.JumlahPembayaran = params.JumlahPembayaran
	invoice.IdStatusPembayaran = params.IdStatusPembayaran
	invoice.NamaEkspedisi = params.NamaEkspedisi
	invoice.LayananEkspedisi = params.LayananEkspedisi
	invoice.Etd = params.Etd
	invoice.IdStatusPengiriman = params.IdStatusPengiriman
	invoice.DetailAlamat = params.DetailAlamat
	invoice.PesanPembeli = params.PesanPembeli
	invoice.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("invoice").Create(&invoice).Error

	if err != nil {
		return models.CreateInvoice{}, err
	}

	return invoice, nil
}

func (m *Keranjang) InvoiceGet(params models.GetInvoice) ([]models.InvoiceGet,int64, error) {

	var count int64
	invoice := []models.InvoiceGet{}
	getPesanan := models.GetPesanan{}

	err := databases.DatabaseSellPump.DB.Table("invoice").Select("invoice.*").Order("invoice.created_at desc")

	if params.IdStatusPembayaran != "" {
		err = err.Where("invoice.id_status_pembayaran = ?", params.IdStatusPembayaran)
	}
	if params.IdStatusPengiriman != "" {
		err = err.Where("invoice.id_status_pengiriman = ?", params.IdStatusPengiriman)
	}
	if params.IdUser != "" {
		err = err.Where("invoice.id_user = ?", params.IdUser)
	}
	if params.NoInv != "" {
		err = err.Where("invoice.no_inv = ?", params.NoInv)
	}
	if params.CreatedAt != "" {
		err = err.Where("invoice.created_at::text like  ?", "%"+params.CreatedAt+"%")
	}
	err.Count(&count)
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&invoice)
	errx := err.Error

	if errx != nil {
		return []models.InvoiceGet{}, count, errx
	}

	for i, _ := range invoice {

		getPesanan.NoInv = invoice[i].NoInv
		invoice[i].Pesanan, errx = m.PesananGet(getPesanan)

		if errx != nil {
			return []models.InvoiceGet{}, count, errx
		}

	}

	return invoice,count, nil
}

func (m *Keranjang) InvoiceUpdate(params models.UpdateInvoice) ([]models.InvoiceGet, int64, error) {

	var count int64
	invoice := models.InvoiceUpdate{}
	getinvoice := []models.InvoiceGet{}

	if params.BuktiBayar != nil {
		invoice.TglPembayaran = m.helper.GetTimeNow()
		path := "/keranjang/"
		pathImage := "./files/"+path
		ext := filepath.Ext(params.BuktiBayar.Filename)
		filename := strings.Replace(params.IdInvoice," ","_", -1)+"-Pembayaran"+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.BuktiBayar, pathImage+filename)
		if errx != nil{
			return []models.InvoiceGet{},count,errx
		}

		url := string(filepath.FromSlash(path+filename))

		invoice.BuktiBayar = new(string)
		*invoice.BuktiBayar = url
	}

	invoice.UpdatedAt = m.helper.GetTimeNow()
	invoice.AlasanDitolak = params.AlasanDitolak
	invoice.ExpiredPengiriman = params.ExpiredPengiriman
	invoice.IdStatusPengiriman = params.IdStatusPengiriman
	invoice.IdStatusPembayaran = params.IdStatusPembayaran
	invoice.NoPo = params.NoResi
	invoice.NoResi = params.NoResi

	err := databases.DatabaseSellPump.DB.Table("invoice").Where("id_invoice = ?", params.IdInvoice).Update(&invoice).Error

	if err != nil {
		return []models.InvoiceGet{},count, err
	}

	paraminvoice := models.GetInvoice{}
	paraminvoice.NoInv = params.NoInv
	getinvoice,count,errx := m.InvoiceGet(paraminvoice)
	if errx != nil {
		return []models.InvoiceGet{}, count, errx
	}
	return getinvoice, count, nil

}
