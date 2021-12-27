package daos

import (
	"fmt"
	"io/ioutil"
	"keranjang_microservice/databases"
	"keranjang_microservice/helper"
	"keranjang_microservice/models"
	"net/http"
	_ "os"
	_ "path/filepath"
	"strings"
	_ "strings"
)

type Keranjang struct {
	helper helper.Helper
}

func (m *Keranjang) KeranjangGet(params models.KeranjangGet) ([]models.KeranjangGet, error) {

	keranjang := []models.KeranjangGet{}

	err := databases.DatabaseSellPump.DB.Table("keranjang").Select("keranjang.*,k.nama_produk,k.harga_produk,k.deskripsi_produk,k.gambar_produk,k.berat_produk").
		Joins("join produk k on k.id_produk = keranjang.id_produk")

	if params.IdUser != "" {
		err = err.Where("keranjang.id_user = ?", params.IdUser)
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
	if params.IdProvinsi == "" && params.IdCity == ""{
		url = "https://pro.rajaongkir.com/api/city"
	}
	if params.IdProvinsi != "" && params.IdCity == ""{
		url = "https://pro.rajaongkir.com/api/city?province="+params.IdProvinsi
	}
	if params.IdProvinsi == "" && params.IdCity != ""{
		url = "https://pro.rajaongkir.com/api/city?id="+params.IdCity
	}
	if params.IdProvinsi != "" && params.IdCity != ""{
		url = "https://pro.rajaongkir.com/api/city?id="+params.IdCity+"&province="+params.IdProvinsi
	}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "0e33e27e42c55799ebe174e1307f2adf")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	isi := string(body)

	fmt.Println("city"+params.IdCity)
	fmt.Println("province"+params.IdProvinsi)

	return isi, nil
}

func (m *Keranjang) SubdistrictGet(params models.RajaOngkir) (string, error) {

	url := ""
	if params.IdCity != "" && params.IdKecamatan == ""{
		url = "https://pro.rajaongkir.com/api/subdistrict?city="+params.IdCity
	}
	if params.IdCity != "" && params.IdKecamatan != ""{
		url = "https://pro.rajaongkir.com/api/subdistrict?city="+params.IdCity+"&id="+params.IdKecamatan
	}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "0e33e27e42c55799ebe174e1307f2adf")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	isi := string(body)

	fmt.Println("city"+params.IdCity)
	fmt.Println("province"+params.IdProvinsi)

	return isi, nil
}

func (m *Keranjang) CostGet(params models.RajaOngkir) (string, error) {

	url := "https://pro.rajaongkir.com/api/cost"

	payload := strings.NewReader("origin="+params.Origin+"&originType=subdistrict"+
		"&destination="+params.Destination+"&destinationType=subdistrict&weight="+params.Weight+
		"&courier=jne:tiki:pos")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("key", "0e33e27e42c55799ebe174e1307f2adf")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	isi := string(body)

	fmt.Println("city"+params.IdCity)
	fmt.Println("province"+params.IdProvinsi)

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

	err := databases.DatabaseSellPump.DB.Table("pesanan").Select("pesanan.*,k.nama_produk,k.harga_produk,k.deskripsi_produk,k.gambar_produk,k.berat_produk").
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