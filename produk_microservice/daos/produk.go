package daos

import (
	"os"
	_ "os"
	"path/filepath"
	_ "path/filepath"
	"produk_microservice/databases"
	"produk_microservice/helper"
	"produk_microservice/models"
	"strings"
	_ "strings"
)

type Produk struct {
	helper helper.Helper
}

func (m *Produk) ProdukCreate(params models.CreateProduk) (models.ProdukCreate, error) {

	insertproduk := models.ProdukCreate{}

	path := "/produk/"
	pathImage := "./files/"+path
	ext := filepath.Ext(params.GambarProduk.Filename)
	filename := strings.Replace(params.NamaProduk," ","_", -1)+params.KodeProduk+params.TipeProduk+ext

	os.MkdirAll(pathImage, 0777)
	errx := m.helper.SaveUploadedFile(params.GambarProduk, pathImage+filename)
	if errx != nil{
		return models.ProdukCreate{},errx
	}

	url := string(filepath.FromSlash(path+filename))

	insertproduk.IdProduk = m.helper.StringWithCharset()
	insertproduk.NamaProduk = params.NamaProduk
	insertproduk.HargaProduk = params.HargaProduk
	insertproduk.KodeProduk = params.KodeProduk
	insertproduk.TipeProduk = params.TipeProduk
	insertproduk.BeratProduk = params.BeratProduk
	insertproduk.IdMerkProduk = params.IdMerkProduk
	insertproduk.DeskripsiProduk = params.DeskripsiProduk
	insertproduk.StokProduk = params.StokProduk
	insertproduk.Tayangan = 0
	insertproduk.IdKategoriProduk = params.IdKategoriProduk
	insertproduk.GambarProduk = url
	insertproduk.LinkTokopedia = params.LinkTokopedia
	insertproduk.LinkBukalapak = params.LinkBukalapak
	insertproduk.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("produk").Create(&insertproduk).Error

	if err != nil {
		return models.ProdukCreate{}, err
	}

	return insertproduk, nil
}


func (m *Produk) ProdukUpdate(params models.UpdateProduk) ([]models.ProdukGet, error) {

	produk := models.ProdukUpdate{}
	getproduk := []models.ProdukGet{}

	if params.GambarProduk != nil {
		path := "/produk/"
		pathImage := "./files/"+path
		ext := filepath.Ext(params.GambarProduk.Filename)
		filename := strings.Replace(params.NamaProduk," ","_", -1)+params.KodeProduk+params.TipeProduk+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.GambarProduk, pathImage+filename)
		if errx != nil{
			return []models.ProdukGet{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		produk.GambarProduk = new(string)
		*produk.GambarProduk = url
	}

	produk.UpdatedAt = m.helper.GetTimeNow()
	//if params.NamaProduk != "" {
	//	produk.NamaProduk = new(string)
	//	*produk.NamaProduk = params.NamaProduk
	//}
	produk.NamaProduk = params.NamaProduk
	produk.HargaProduk = params.HargaProduk
	produk.KodeProduk = params.KodeProduk
	produk.TipeProduk = params.TipeProduk
	produk.BeratProduk = params.BeratProduk
	produk.IdMerkProduk = params.IdMerkProduk
	produk.DeskripsiProduk = params.DeskripsiProduk
	produk.StokProduk = params.StokProduk
	produk.Tayangan = params.Tayangan
	produk.LinkTokopedia = params.LinkTokopedia
	produk.LinkBukalapak = params.LinkBukalapak
	produk.IdKategoriProduk = params.IdKategoriProduk

	err := databases.DatabaseSellPump.DB.Table("produk").Where("id_produk = ?", params.IdProduk).Update(&produk).Error

	if err != nil {
		return []models.ProdukGet{}, err
	}

	paramproduk := models.GetProduk{}
	paramproduk.IdProduk = params.IdProduk
	getproduk,errx := m.ProdukGet(paramproduk)
	if errx != nil {
		return []models.ProdukGet{}, errx
	}
	return getproduk, nil

}

func (m *Produk) ProdukGet(params models.GetProduk) ([]models.ProdukGet, error) {

	produk := []models.ProdukGet{}

	err := databases.DatabaseSellPump.DB.Table("produk").Select("produk.*,k.kategori,k.sub_kategori, m.nama_merk, m.gambar_merk").
		Joins("join kategori k on k.id_kategori = produk.id_kategori_produk").
		Joins("join merk m on m.id_merk = produk.id_merk_produk")

	if params.IdProduk != "" {
		err = err.Where("produk.id_produk = ?", params.IdProduk)
	}
	if params.Search != "" {
		err = err.Where("produk.nama_produk ilike '%"+params.Search+"%' OR produk.kode_produk ilike '%"+params.Search+"%' OR produk.tipe_produk ilike '%"+params.Search+"%' OR m.nama_merk ilike '%"+params.Search+"%'")
	}
	if params.SubKategori != "" {
		err = err.Where("k.sub_kategori = ?", params.SubKategori)
	}
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&produk)

	errx := err.Error


	if errx != nil {
		return []models.ProdukGet{}, errx
	}

	return produk, nil
}

func (m *Produk) ProdukDelete(params models.DeleteProduk) (models.DeleteProduk, error) {

	produk := models.DeleteProduk{}

	produk.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("produk").Where("id_produk = ?", params.IdProduk).Update(&produk).Error

	if err != nil {
		return models.DeleteProduk{}, err
	}

	return produk, nil

}

func (m *Produk) ArtikelCreate(params models.CreateArtikel) (models.ArtikelCreate, error) {

	artikel := models.ArtikelCreate{}

	path := "/artikel/"
	pathImage := "./files/"+path
	ext := filepath.Ext(params.GambarArtikel.Filename)
	filename := strings.Replace(params.JudulArtikel," ","_", -1)+ext

	os.MkdirAll(pathImage, 0777)
	errx := m.helper.SaveUploadedFile(params.GambarArtikel, pathImage+filename)
	if errx != nil{
		return models.ArtikelCreate{},errx
	}

	url := string(filepath.FromSlash(path+filename))

	artikel.IdArtikel = m.helper.StringWithCharset()
	artikel.JudulArtikel = params.JudulArtikel
	artikel.DeskripsiArtikel = params.DeskripsiArtikel
	artikel.GambarArtikel = url
	artikel.CreatedAt = m.helper.GetTimeNow()


	err := databases.DatabaseSellPump.DB.Table("artikel").Create(&artikel).Error

	if err != nil {
		return models.ArtikelCreate{}, err
	}

	return artikel, nil
}

func (m *Produk) ArtikelGet(params models.GetArtikel) ([]models.ArtikelGet, error) {

	artikel := []models.ArtikelGet{}

	err := databases.DatabaseSellPump.DB.Table("artikel")

	if params.IdArtikel != "" {
		err = err.Where("id_artikel = ?", params.IdArtikel)
	}
	if params.Search != "" {
		err = err.Where("judul_artikel ilike '%"+params.Search+"%'")
	}
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&artikel)

	errx := err.Error


	if errx != nil {
		return []models.ArtikelGet{}, errx
	}

	return artikel, nil
}

func (m *Produk) ArtikelUpdate(params models.UpdateArtikel) ([]models.ArtikelGet, error) {

	artikel := models.ArtikelUpdate{}
	getartikel := []models.ArtikelGet{}

	if params.GambarArtikel != nil {
		path := "/artikel/"
		pathImage := "./files/"+path
		ext := filepath.Ext(params.GambarArtikel.Filename)
		filename := strings.Replace(params.JudulArtikel," ","_", -1)+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.GambarArtikel, pathImage+filename)
		if errx != nil{
			return []models.ArtikelGet{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		artikel.GambarArtikel = new(string)
		*artikel.GambarArtikel = url
	}

	artikel.UpdatedAt = m.helper.GetTimeNow()
	//if params.NamaProduk != "" {
	//	produk.NamaProduk = new(string)
	//	*produk.NamaProduk = params.NamaProduk
	//}
	artikel.JudulArtikel = params.JudulArtikel
	artikel.DeskripsiArtikel = params.DeskripsiArtikel

	err := databases.DatabaseSellPump.DB.Table("artikel").Where("id_artikel = ?", params.IdArtikel).Update(&artikel).Error

	if err != nil {
		return []models.ArtikelGet{}, err
	}

	paramartikel := models.GetArtikel{}
	paramartikel.IdArtikel = params.IdArtikel
	getartikel,errx := m.ArtikelGet(paramartikel)
	if errx != nil {
		return []models.ArtikelGet{}, errx
	}
	return getartikel, nil

}

func (m *Produk) ArtikelDelete(params models.DeleteArtikel) (models.DeleteArtikel, error) {

	artikel := models.DeleteArtikel{}

	artikel.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("artikel").Where("id_artikel = ?", params.IdArtikel).Update(&artikel).Error

	if err != nil {
		return models.DeleteArtikel{}, err
	}

	return artikel, nil

}