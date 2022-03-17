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

	insertproduk.IdProduk = m.helper.StringWithCharset()
	path := "/produk/"
	pathImage := "./files/"+path
	ext := filepath.Ext(params.GambarProduk.Filename)
	//filename := insertproduk.IdProduk+ext
	filename := strings.Replace(insertproduk.IdProduk," ","_", -1)+ext

	os.MkdirAll(pathImage, 0777)
	errx := m.helper.SaveUploadedFile(params.GambarProduk, pathImage+filename)
	if errx != nil{
		return models.ProdukCreate{},errx
	}

	url := string(filepath.FromSlash(path+filename))


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
	insertproduk.HargaDiskon = params.HargaDiskon
	insertproduk.Diskon = params.Diskon
	insertproduk.BoolDiskon = params.BoolDiskon
	insertproduk.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("produk").Create(&insertproduk).Error

	if err != nil {
		return models.ProdukCreate{}, err
	}

	return insertproduk, nil
}

func (m *Produk) ProdukUpdate(params models.UpdateProduk) ([]models.ProdukGet, int64, error) {

	var count int64
	produk := models.ProdukUpdate{}
	getproduk := []models.ProdukGet{}

	if params.GambarProduk != nil {
		path := "/produk/"
		pathImage := "./files/"+path
		ext := filepath.Ext(params.GambarProduk.Filename)
		//filename := params.IdProduk+ext
		filename := strings.Replace(params.IdProduk," ","_", -1)+ext


		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.GambarProduk, pathImage+filename)
		if errx != nil{
			return []models.ProdukGet{}, count,errx
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
	produk.HargaDiskon = params.HargaDiskon
	produk.Diskon = params.Diskon
	produk.BoolDiskon = params.BoolDiskon
	produk.IdKategoriProduk = params.IdKategoriProduk

	err := databases.DatabaseSellPump.DB.Table("produk").Where("id_produk = ?", params.IdProduk).Update(&produk).Error

	if err != nil {
		return []models.ProdukGet{}, count, err
	}

	paramproduk := models.GetProduk{}
	paramproduk.IdProduk = params.IdProduk
	getproduk, count,errx := m.ProdukGet(paramproduk)
	if errx != nil {
		return []models.ProdukGet{}, count, errx
	}
	return getproduk, count, nil

}

func (m *Produk) ProdukGet(params models.GetProduk) ([]models.ProdukGet, int64, error) {

	var count int64
	produk := []models.ProdukGet{}

	err := databases.DatabaseSellPump.DB.Table("produk").Select("produk.*,k.kategori,k.sub_kategori, m.nama_merk, m.gambar_merk").
		Joins("join kategori k on k.id_kategori = produk.id_kategori_produk").
		Joins("join merk m on m.id_merk = produk.id_merk_produk").
		Order("produk.created_at desc")

	if params.IdProduk != "" {
		err = err.Where("produk.id_produk = ?", params.IdProduk)
	}
	if params.Search != "" {
		err = err.Where("produk.nama_produk ilike '%"+params.Search+"%' OR produk.kode_produk ilike '%"+params.Search+"%' OR produk.tipe_produk ilike '%"+params.Search+"%' OR m.nama_merk ilike '%"+params.Search+"%'")
	}
	if params.SubKategori != "" {
		err = err.Where("k.sub_kategori = ?", params.SubKategori)
	}
	err.Count(&count)
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&produk)

	errx := err.Error


	if errx != nil {
		return []models.ProdukGet{}, count, errx
	}

	return produk, count, nil
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

func (m *Produk) ArtikelGet(params models.GetArtikel) ([]models.ArtikelGet,int64, error) {

	var count int64
	artikel := []models.ArtikelGet{}

	err := databases.DatabaseSellPump.DB.Table("artikel").Order("created_at desc")

	if params.IdArtikel != "" {
		err = err.Where("id_artikel = ?", params.IdArtikel)
	}
	if params.Search != "" {
		err = err.Where("judul_artikel ilike '%"+params.Search+"%'")
	}
	err.Count(&count)
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&artikel)

	errx := err.Error


	if errx != nil {
		return []models.ArtikelGet{},count, errx
	}

	return artikel,count, nil
}

func (m *Produk) ArtikelUpdate(params models.UpdateArtikel) ([]models.ArtikelGet,int64, error) {

	var count int64
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
			return []models.ArtikelGet{},count,errx
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
		return []models.ArtikelGet{},count, err
	}

	paramartikel := models.GetArtikel{}
	paramartikel.IdArtikel = params.IdArtikel
	getartikel,count,errx := m.ArtikelGet(paramartikel)
	if errx != nil {
		return []models.ArtikelGet{},count, errx
	}
	return getartikel,count, nil

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

func (m *Produk) KhususCreate(params models.CreateKhusus) (models.KhususCreate, error) {

	insertkhusus := models.KhususCreate{}

	insertkhusus.IdKhusus = m.helper.StringWithCharset()
	insertkhusus.IdUser = params.IdUser
	insertkhusus.NamaKhusus = params.NamaKhusus
	insertkhusus.HargaKhusus = params.HargaKhusus
	insertkhusus.HargaAsli = params.HargaAsli
	insertkhusus.Diskon = params.Diskon
	insertkhusus.BeratKhusus = params.BeratKhusus
	insertkhusus.DeskripsiKhusus = params.DeskripsiKhusus
	insertkhusus.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("mst_produk_khusus").Create(&insertkhusus).Error

	if err != nil {
		return models.KhususCreate{}, err
	}

	return insertkhusus, nil
}

func (m *Produk) KhususGet(params models.GetKhusus) ([]models.KhususGet, int64, error) {

	var count int64
	khusus := []models.KhususGet{}
	getProdukKhusus := models.GetProdukKhusus{}

	err := databases.DatabaseSellPump.DB.Table("mst_produk_khusus").Select("mst_produk_khusus.*").Order("mst_produk_khusus.created_at desc")

	if params.IdUser != "" {
		err = err.Where("mst_produk_khusus.id_user = ?", params.IdUser)
	}
	if params.IdKhusus != "" {
		err = err.Where("mst_produk_khusus.id_khusus = ?", params.IdKhusus)
	}
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&khusus)
	errx := err.Error

	if errx != nil {
		return []models.KhususGet{}, count, errx
	}

	for i, _ := range khusus {

		getProdukKhusus.IdKhusus = khusus[i].IdKhusus
		khusus[i].ProdukKhusus, count, errx = m.ProdukKhususGet(getProdukKhusus)

		if errx != nil {
			return []models.KhususGet{}, count, errx
		}

	}

	return khusus, count, nil
}

func (m *Produk) KhususDelete(params models.DeleteKhusus) (models.DeleteKhusus, error) {

	khusus := models.DeleteKhusus{}

	khusus.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("mst_produk_khusus").Where("id_khusus = ?", params.IdKhusus).Update(&khusus).Error

	if err != nil {
		return models.DeleteKhusus{}, err
	}

	return khusus, nil

}

func (m *Produk) ProdukKhususCreate(params models.CreateProdukKhusus) (models.ProdukKhususCreate, error) {

	insertprodukkhusus := models.ProdukKhususCreate{}

	insertprodukkhusus.IdProdukKhusus = m.helper.StringWithCharset()
	insertprodukkhusus.IdKhusus = params.IdKhusus
	insertprodukkhusus.IdProduk = params.IdProduk
	insertprodukkhusus.Kuantitas = params.Kuantitas
	insertprodukkhusus.Harga = params.Harga
	insertprodukkhusus.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("produk_khusus").Create(&insertprodukkhusus).Error

	if err != nil {
		return models.ProdukKhususCreate{}, err
	}

	return insertprodukkhusus, nil
}

func (m *Produk) ProdukKhususGet(params models.GetProdukKhusus) ([]models.ProdukKhususGet, int64, error) {

	var count int64
	produkkhusus := []models.ProdukKhususGet{}
	getProduk := models.GetProduk{}

	err := databases.DatabaseSellPump.DB.Table("produk_khusus").Select("produk_khusus.*").Order("produk_khusus.created_at desc")

	if params.IdProdukKhusus != "" {
		err = err.Where("produk_khusus.id_produk_khusus = ?", params.IdProdukKhusus)
	}
	if params.IdKhusus != "" {
		err = err.Where("produk_khusus.id_khusus = ?", params.IdKhusus)
	}
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&produkkhusus)
	errx := err.Error

	if errx != nil {
		return []models.ProdukKhususGet{}, count, errx
	}

	for i, _ := range produkkhusus {

		getProduk.IdProduk = produkkhusus[i].IdProduk
		produkkhusus[i].Produk,count, errx = m.ProdukGet(getProduk)

		if errx != nil {
			return []models.ProdukKhususGet{}, count, errx
		}

	}

	return produkkhusus, count, nil
}

func (m *Produk) RatingCreate(params models.CreateRating) (models.RatingCreate, error) {

	rating := models.RatingCreate{}

	rating.IdRating = m.helper.StringWithCharset()
	rating.IdProduk = params.IdProduk
	rating.IdUser = params.IdUser
	rating.Nama = params.Nama
	rating.Komentar = params.Komentar
	rating.Rating = params.Rating
	rating.CreatedAt = m.helper.GetTimeNow()


	err := databases.DatabaseSellPump.DB.Table("rating").Create(&rating).Error

	if err != nil {
		return models.RatingCreate{}, err
	}

	return rating, nil
}

func (m *Produk) RatingGet(params models.GetRating) ([]models.RatingGet, error) {

	rating := []models.RatingGet{}

	err := databases.DatabaseSellPump.DB.Table("rating").Order("created_at desc")

	if params.IdRating != "" {
		err = err.Where("id_rating = ?", params.IdRating)
	}
	if params.IdProduk != "" {
		err = err.Where("id_produk = ?", params.IdProduk)
	}
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&rating)

	errx := err.Error


	if errx != nil {
		return []models.RatingGet{}, errx
	}

	return rating, nil
}

func (m *Produk) FileCreate(params models.CreateFile) (models.FileCreate, error) {

	file := models.FileCreate{}

	path := "/file/"
	pathFile := "./files/"+path
	ext := filepath.Ext(params.File.Filename)
	filename := strings.Replace(params.NamaFile," ","_", -1)+ext

	os.MkdirAll(pathFile, 0777)
	errx := m.helper.SaveUploadedFile(params.File, pathFile+filename)
	if errx != nil{
		return models.FileCreate{},errx
	}

	url := string(filepath.FromSlash(path+filename))

	file.IdFile = m.helper.StringWithCharset()
	file.NamaFile = params.NamaFile
	file.File = url
	file.CreatedAt = m.helper.GetTimeNow()


	err := databases.DatabaseSellPump.DB.Table("file").Create(&file).Error

	if err != nil {
		return models.FileCreate{}, err
	}

	return file, nil
}

func (m *Produk) FileGet(params models.GetFile) ([]models.FileGet, error) {

	file := []models.FileGet{}

	err := databases.DatabaseSellPump.DB.Table("file").Order("created_at desc")

	if params.IdFile != "" {
		err = err.Where("id_file = ?", params.IdFile)
	}
	if params.Search != "" {
		err = err.Where("nama_file ilike '%"+params.Search+"%'")
	}
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&file)

	errx := err.Error


	if errx != nil {
		return []models.FileGet{}, errx
	}

	return file, nil
}

func (m *Produk) FileDelete(params models.DeleteFile) (models.DeleteFile, error) {

	file := models.DeleteFile{}

	file.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("file").Where("id_file = ?", params.IdFile).Update(&file).Error

	if err != nil {
		return models.DeleteFile{}, err
	}

	return file, nil

}