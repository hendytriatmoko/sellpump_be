package daos

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"user_microservices/databases"
	"user_microservices/helper"
	"user_microservices/middleware"
	"user_microservices/models"
)

type User struct {
	helper helper.Helper
}

func (m *User) UserCreate(params models.CreateUser) (models.UserCreate, error) {

	user := models.UserCreate{}

	if params.Foto != nil {
		path := "/user/"

		pathImage := "./files/"+path
		ext := filepath.Ext(params.Foto.Filename)
		filename := strings.Replace(params.Nama," ","_", -1)+params.NoTelp+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.Foto, pathImage+filename)
		if errx != nil{
			return models.UserCreate{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		//user.Foto = new(string)
		user.Foto = url
	}

	user.IdUser = m.helper.StringWithCharset()
	user.Nama = params.Nama
	user.NoTelp = params.NoTelp
	user.Email = params.Email
	user.Password,_ = EncryptPassword(params.Password)
	user.Status = "pembeli"
	user.Verifikasi = "Y"
	user.CreatedAt = m.helper.GetTimeNow()
	user.TokenRegister = m.helper.StringWithCharset()

	err := databases.DatabaseSellPump.DB.Table("user").Create(&user).Error

	if err != nil {
		return models.UserCreate{}, err
	}

	dataverifikasi := models.VerifikasiUser{}
	dataverifikasi.IdUser = user.IdUser
	dataverifikasi.Email = user.Email
	dataverifikasi.TokenRegister = user.TokenRegister
	errx := m.VerifikasiUser(dataverifikasi)
	if errx != nil {
		return models.UserCreate{}, errx
	}

	dataprofil := models.ProfilUser{}
	dataprofil.IdUser = user.IdUser
	erry := m.ProfilUser(dataprofil)
	if erry != nil {
		return models.UserCreate{}, erry
	}

	return user, nil
}

func (m *User) VerifikasiUser(params models.VerifikasiUser) error {
	verifikasi := models.VerifikasiUser{}

	verifikasi.IdVerifikasi = m.helper.StringWithCharset()
	verifikasi.TokenRegister = params.TokenRegister
	verifikasi.IdUser = params.IdUser
	verifikasi.Email = params.Email
	verifikasi.Status = "Y"
	verifikasi.CreatedAt = m.helper.GetTimeNow()

	ti := time.Now()
	ti_n := ti.AddDate(0, 0, 7)
	next := string(ti_n.Format("2006-01-02 15:04:05.999999"))
	verifikasi.ExpiredAt = next

	err := databases.DatabaseSellPump.DB.Table("verifikasi").Create(&verifikasi).Error

	if err != nil {
		return err
	}

	//err = m.helper.SendEmailVerifikasi(verifikasi.Email, verifikasi.IdUser, verifikasi.TokenRegister)
	//
	//if err != nil {
	//	return err
	//}

	return nil
}

func (m *User) ProfilUser(params models.ProfilUser) error {
	profil := models.ProfilUser{}

	profil.IdProfil = m.helper.StringWithCharset()
	profil.IdUser = params.IdUser
	profil.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("profil").Create(&profil).Error

	if err != nil {
		return err
	}

	//err = m.helper.SendEmailVerifikasi(verifikasi.Email, verifikasi.IdUser, verifikasi.TokenRegister)
	//
	//if err != nil {
	//	return err
	//}

	return nil
}

func (m *User) UserGet(params models.GetUser) ([]models.UserGet, error) {

	user := []models.UserGet{}

	err := databases.DatabaseSellPump.DB.Table("user")
	if params.IdUser != "" {
		err = err.Where("id_user = ?", params.IdUser)
	}
	if params.Email != "" {
		err = err.Where("email = ?", params.Email)
	}
	if params.Search != "" {
		err = err.Where("nama ilike '%"+params.Search+"%' OR email ilike '%"+params.Search+"%'")
	}
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&user)

	errx := err.Error


	if errx != nil {
		return []models.UserGet{}, errx
	}

	return user, nil
}

func (m *User) UserUpdate(params models.UpdateUser) ([]models.UserGet, error) {

	user := models.UserUpdate{}
	getuser := []models.UserGet{}

	if params.Foto != nil {
		path := "/user/"
		pathImage := "./files/"+path
		ext := filepath.Ext(params.Foto.Filename)
		filename := strings.Replace(params.Nama," ","_", -1)+params.NoTelp+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.Foto, pathImage+filename)
		if errx != nil{
			return []models.UserGet{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		user.Foto = new(string)
		*user.Foto = url
	}
	user.UpdatedAt = m.helper.GetTimeNow()
	user.Nama = params.Nama
	user.Email = params.Email
	user.NoTelp = params.NoTelp
	if params.Password != "" {
		user.Password,_ = EncryptPassword(params.Password)
	}

	err := databases.DatabaseSellPump.DB.Table("user").Where("id_user = ?", params.IdUser).Update(&user).Error

	if err != nil {
		return []models.UserGet{}, err
	}

	paramuser := models.GetUser{}
	paramuser.IdUser = params.IdUser
	getuser,errx := m.UserGet(paramuser)
	if errx != nil {
		return []models.UserGet{}, errx
	}
	return getuser, nil

}

func (m *User) UserDelete(params models.DeleteUser) (models.DeleteUser, error) {

	user := models.DeleteUser{}

	user.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("user").Where("id_user = ?", params.IdUser).Update(&user).Error

	if err != nil {
		return models.DeleteUser{}, err
	}

	return user, nil

}

func (m *User) UserCheckAkun(params models.CheckAkunUser) error {

	checkakun := models.CheckAkunRead{}
	var check bool
	today := m.helper.GetTimeNow()

	check = databases.DatabaseSellPump.DB.Table("verifikasi v").
		Where("v.email = ?", params.Email).Find(&checkakun).RecordNotFound()

	if check == true {
		err := errors.New("Email Tidak Ditemukan")
		return err
	}else if check == false {
		//if today > checkakun.ExpiredAt {
		//	err := errors.New("Verification Expired")
		//	return err
		//}else if today < checkakun.ExpiredAt {
		//	if checkakun.Status == "N" {
		//		err := errors.New("Silahkan Verifikasi Email Anda")
		//		return err
		//	}
		//}
		if checkakun.Status == "N" {
			if today > checkakun.ExpiredAt {
				err := errors.New("Verification Expired")
				return err
			} else if today < checkakun.ExpiredAt {
				err := errors.New("Silahkan Verifikasi Email Anda")
				return err
			}
		}
	}

	return error(nil)

}

func (m *User) UserResendVerification(params models.CheckAkunUser) (models.CheckAkunRead, error) {

	updateverifikasi := models.CheckAkunRead{}
	user := models.UserUpdate{}

	ti := time.Now()
	ti_n := ti.AddDate(0, 0, 7)
	next := string(ti_n.Format("2006-01-02 15:04:05.999999"))
	updateverifikasi.ExpiredAt = next
	updateverifikasi.UpdatedAt = m.helper.GetTimeNow()
	updateverifikasi.TokenRegister = m.helper.StringWithCharset()

	user.UpdatedAt = m.helper.GetTimeNow()
	user.TokenRegister = updateverifikasi.TokenRegister

	err := databases.DatabaseSellPump.DB.Table("verifikasi").Where("email = ?", params.Email).Update(&updateverifikasi).Error

	if err != nil {
		return models.CheckAkunRead{}, err
	}

	errs := databases.DatabaseSellPump.DB.Table("user").Where("email = ?", params.Email).Update(&user).Error

	if errs != nil {
		return models.CheckAkunRead{}, errs
	}

	errx := databases.DatabaseSellPump.DB.Table("verifikasi").Where("email = ?", params.Email).Find(&updateverifikasi).Error

	if errx != nil {
		return models.CheckAkunRead{}, errx
	}



	erry := m.helper.SendEmailVerifikasi(updateverifikasi.Email, updateverifikasi.IdUser, updateverifikasi.TokenRegister)
	if erry != nil {
		return updateverifikasi, erry
	}

	return updateverifikasi, nil

}

func (m *User) UserVerificationRegister(params models.VerificationUpdate) ([]models.UserGet, error) {

	user := models.UserUpdate{}
	updateverifikasi := models.CheckAkunRead{}
	getuser := []models.UserGet{}

	user.UpdatedAt = m.helper.GetTimeNow()
	user.Verifikasi = "Y"

	updateverifikasi.UpdatedAt = m.helper.GetTimeNow()
	updateverifikasi.Status = "Y"


	errz := databases.DatabaseSellPump.DB.Table("verifikasi").Where("id_user = ?", params.IdUser).Where("token_register = ?", params.TokenRegister).Update(&updateverifikasi).Error

	if errz != nil {
		return []models.UserGet{}, errz
	}

	err := databases.DatabaseSellPump.DB.Table("user").Where("id_user = ?", params.IdUser).Where("token_register = ?", params.TokenRegister).Update(&user).Error

	if err != nil {
		return []models.UserGet{}, err
	}

	paramuser := models.GetUser{}
	paramuser.IdUser = params.IdUser
	getuser,errx := m.UserGet(paramuser)
	if errx != nil {
		return []models.UserGet{}, errx
	}
	return getuser, nil

}

func (m *User) UserForgotPassword(params models.CheckAkunUser) (models.UserGet, error) {

	updateverifikasi := models.UserGet{}
	user := models.UserUpdate{}

	user.TokenRepassword = m.helper.StringWithCharset()

	errs := databases.DatabaseSellPump.DB.Table("user").Where("email = ?", params.Email).Update(&user).Error

	if errs != nil {
		return models.UserGet{}, errs
	}

	errx := databases.DatabaseSellPump.DB.Table("user").Where("email = ?", params.Email).Find(&updateverifikasi).Error

	if errx != nil {
		return models.UserGet{}, errx
	}

	erry := m.helper.SendForgotPassword(updateverifikasi.Email,updateverifikasi.IdUser,user.TokenRepassword)
	if erry != nil {
		return updateverifikasi, erry
	}

	return updateverifikasi, nil

}

func (m *User) UserVerificationRepassword(params models.VerificationUpdate) ([]models.UserGet, error) {

	user := models.UserUpdate{}
	getuser := []models.UserGet{}

	user.UpdatedAt = m.helper.GetTimeNow()
	user.Password,_ = EncryptPassword(params.Password)

	err := databases.DatabaseSellPump.DB.Table("user").Where("id_user = ?", params.IdUser).Where("token_repassword = ?", params.TokenRepassword).Update(&user).Error

	if err != nil {
		return []models.UserGet{}, err
	}

	paramuser := models.GetUser{}
	paramuser.IdUser = params.IdUser
	getuser,errx := m.UserGet(paramuser)
	if errx != nil {
		return []models.UserGet{}, errx
	}
	return getuser, nil

}

func (m *User) LoginCheck(params models.UserToken) error {

	checkakun := models.UserGet{}
	var check bool

	check = databases.DatabaseSellPump.DB.Table("user").
		Where("email = ?", params.Email).Find(&checkakun).RecordNotFound()

	if check == true {
		err := errors.New("Email Tidak Ditemukan")
		return err
	}

	return error(nil)

}

func (m *User) UserGetLogin(params models.UserLogin) ([]models.UserGet, error) {

	user := []models.UserGet{}

	err := databases.DatabaseSellPump.DB.Table("user")
	if params.IdUser != "" {
		err = err.Where("id_user = ?", params.IdUser)
	}

	err = err.Find(&user)

	errx := err.Error


	if errx != nil {
		return []models.UserGet{}, errx
	}

	return user, nil
}

func (m *User) Signin(params models.UserToken) ([]models.UserGet, string, error) {

	userGet := models.GetUser{}
	userRead := []models.UserGet{}
	updateToken := models.UserUpdate{}
	var token string
	var er error

	err := m.LoginCheck(params)

	if err != nil {
		return userRead, "", err
	}

	if params.Email != "" {
		userGet.Email = params.Email
	}
	if params.Password != "" {
		userGet.Password = params.Password
	}

	userRead, err = m.UserGet(userGet)

	if err != nil {
		return userRead, "", err
	}

	if userRead[0].Verifikasi == "N" {
		err = errors.New("Akun Anda di Nonaktifkan, Tidak Dapat di Akses")
		return userRead, "", err
	}

	//token, er := m.helper.GetToken(userRead[0].IdUser)
	password,_ := DecryptPassword(userRead[0].Password)

	if userRead[0].Email == params.Email && params.Password == password && userRead[0].Verifikasi == "Y"  {
		fmt.Println("cocok")
		token, er = middleware.CreateAuth(userRead[0].IdUser, "user", "none", "none")

		if er != nil {
			return userRead, "", er
		}

		updateToken.Token = token

		err = databases.DatabaseSellPump.DB.Table("user").Where("id_user = ?", userRead[0].IdUser).Update(&updateToken).Error

		if err != nil {
			return userRead, "", err
		}
	}else {
		fmt.Println("tidak cocok")
	}



	return userRead, token, nil

}

func (m *User) ProfilCreate(params models.CreateProfil) (models.ProfilCreate, error) {

	profil := models.ProfilCreate{}

	if params.Npwp != nil {
		path := "/profil/"

		pathImage := "./files/"+path
		ext := filepath.Ext(params.Npwp.Filename)
		filename := strings.Replace(params.IdUser," ","_", -1)+"npwp"+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.Npwp, pathImage+filename)
		if errx != nil{
			return models.ProfilCreate{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		//profil.Npwp = new(string)
		profil.Npwp = url
	}
	if params.Siup != nil {
		path := "/profil/"

		pathImage := "./files/"+path
		ext := filepath.Ext(params.Siup.Filename)
		filename := strings.Replace(params.IdUser," ","_", -1)+"siup"+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.Siup, pathImage+filename)
		if errx != nil{
			return models.ProfilCreate{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		//profil.Npwp = new(string)
		profil.Siup = url
	}
	if params.Tdp != nil {
		path := "/profil/"

		pathImage := "./files/"+path
		ext := filepath.Ext(params.Tdp.Filename)
		filename := strings.Replace(params.IdUser," ","_", -1)+"tdp"+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.Tdp, pathImage+filename)
		if errx != nil{
			return models.ProfilCreate{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		//profil.Npwp = new(string)
		profil.Tdp = url
	}

	profil.IdProfil = m.helper.StringWithCharset()
	profil.Nama = params.Nama
	profil.Alamat = params.Alamat
	profil.IdUser = params.IdUser
	profil.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("profil").Create(&profil).Error

	if err != nil {
		return models.ProfilCreate{}, err
	}

	return profil, nil
}

func (m *User) ProfilGet(params models.GetUser) ([]models.ProfilCreate, error) {

	profil := []models.ProfilCreate{}

	err := databases.DatabaseSellPump.DB.Table("profil")
	if params.IdUser != "" {
		err = err.Where("id_user = ?", params.IdUser)
	}

	err = err.Find(&profil)

	errx := err.Error


	if errx != nil {
		return []models.ProfilCreate{}, errx
	}

	return profil, nil
}

func (m *User) ProfilUpdate(params models.UpdateProfil) ([]models.ProfilCreate, error) {

	profil := models.ProfilUpdate{}
	geprofil := []models.ProfilCreate{}

	if params.Npwp != nil {
		path := "/profil/"
		pathImage := "./files/"+path
		ext := filepath.Ext(params.Npwp.Filename)
		filename := strings.Replace(params.IdUser," ","_", -1)+"npwp"+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.Npwp, pathImage+filename)
		if errx != nil{
			return []models.ProfilCreate{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		profil.Npwp = new(string)
		*profil.Npwp = url
	}
	if params.Tdp != nil {
		path := "/profil/"
		pathImage := "./files/"+path
		ext := filepath.Ext(params.Tdp.Filename)
		filename := strings.Replace(params.IdUser," ","_", -1)+"tdp"+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.Tdp, pathImage+filename)
		if errx != nil{
			return []models.ProfilCreate{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		profil.Tdp = new(string)
		*profil.Tdp = url
	}
	if params.Siup != nil {
		path := "/profil/"
		pathImage := "./files/"+path
		ext := filepath.Ext(params.Siup.Filename)
		filename := strings.Replace(params.IdUser," ","_", -1)+"siup"+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.Siup, pathImage+filename)
		if errx != nil{
			return []models.ProfilCreate{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		profil.Siup = new(string)
		*profil.Siup = url
	}
	profil.UpdatedAt = m.helper.GetTimeNow()
	profil.Nama = params.Nama
	profil.Alamat = params.Alamat

	err := databases.DatabaseSellPump.DB.Table("profil").Where("id_user = ?", params.IdUser).Update(&profil).Error

	if err != nil {
		return []models.ProfilCreate{}, err
	}

	paramprofil := models.GetUser{}
	paramprofil.IdUser = params.IdUser
	geprofil,errx := m.ProfilGet(paramprofil)
	if errx != nil {
		return []models.ProfilCreate{}, errx
	}
	return geprofil, nil

}