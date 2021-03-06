package models

import "mime/multipart"

type CreateUser struct {
	Nama        string `json:"nama" form:"nama"`
	NoTelp		string `json:"no_telp" form:"no_telp"`
	Email       string `json:"email" form:"email"`
	Foto      	*multipart.FileHeader `json:"foto" form:"foto"`
	Password    string `json:"password" form:"password"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}
type UserCreate struct {
	IdUser		string `json:"id_user" form:"id_user"`
	Nama        string `json:"nama" form:"nama"`
	NoTelp		string `json:"no_telp" form:"no_telp"`
	Email       string `json:"email" form:"email"`
	Foto      	string `json:"foto" form:"foto"`
	Password    string `json:"password" form:"password"`
	Status    	string `json:"status" form:"status"`
	Verifikasi 	string   `json:"verifikasi" form:"verifikasi"`
	CreatedAt   string `json:"created_at" form:"created_at"`
	TokenRegister  	string `json:"token_register" form:"token_register"`
	TokenRepassword string `json:"token_repassword" form:"token_repassword"`
}

type VerifikasiUser struct {
	IdVerifikasi	string `json:"id_verifikasi" form:"id_verifikasi"`
	IdUser			string `json:"id_user" form:"id_user"`
	Email       	string `json:"email" form:"email"`
	Status    		string `json:"status" form:"status"`
	CreatedAt   	string `json:"created_at" form:"created_at"`
	ExpiredAt   	string `json:"expired_at" form:"expired_at"`
	TokenRegister  	string `json:"token_register" form:"token_register"`
}

type ProfilUser struct {
	IdProfil		string `json:"id_profil" form:"id_profil"`
	IdUser			string `json:"id_user" form:"id_user"`
	CreatedAt   	string `json:"created_at" form:"created_at"`
}

type GetUser struct {
	IdUser       		string `json:"id_user" form:"id_user"`
	Email      		string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Search       		string `json:"search" form:"search"`
	Limit     			*int64 `json:"limit" form:"limit"`
	Offset    			*int64 `json:"offset" form:"offset"`
}

type UserGet struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	Nama         	string `json:"nama" form:"nama"`
	Email      		string `json:"email" form:"email"`
	NoTelp         	string `json:"no_telp" form:"no_telp"`
	Foto      		string `json:"foto" form:"foto"`
	Status         	string `json:"status" form:"status"`
	Password    	string `json:"password" form:"password"`
	Verifikasi  	string `json:"verifikasi" form:"verifikasi"`
	CreatedAt       string `json:"created_at" form:"created_at"`
	UpdatedAt       string `json:"updated_at" form:"updated_at"`
	DeletedAt       string `json:"deleted_at" form:"deleted_at"`
}

type UserLogin struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	Nama         	string `json:"nama" form:"nama"`
	Email      		string `json:"email" form:"email"`
	NoTelp         	string `json:"no_telp" form:"no_telp"`
	Foto      		string `json:"foto" form:"foto"`
	Password    	string `json:"password" form:"password"`
	Status         	string `json:"status" form:"status"`
	Verifikasi  	string `json:"verifikasi" form:"verifikasi"`
	CreatedAt       string `json:"created_at" form:"created_at"`
	UpdatedAt       string `json:"updated_at" form:"updated_at"`
	DeletedAt       string `json:"deleted_at" form:"deleted_at"`
}

type UserRead struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	Nama         	string `json:"nama" form:"nama"`
	Email      		string `json:"email" form:"email"`
	NoTelp         	string `json:"no_telp" form:"no_telp"`
	Foto      		string `json:"foto" form:"foto"`
	Status         	string `json:"status" form:"status"`
	Verifikasi  	string `json:"verifikasi" form:"verifikasi"`
}

type UpdateToken struct {
	Token    string `json:"token" form:"token"`
	FcmToken string `json:"fcm_token" form:"fcm_token"`
}

type UpdateUser struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	Nama         	string `json:"nama" form:"nama"`
	Email      		string `json:"email" form:"email"`
	NoTelp         	string `json:"no_telp" form:"no_telp"`
	Password    	string `json:"password" form:"password"`
	Foto      		*multipart.FileHeader `json:"foto" form:"foto"`
}

type VerificationUpdate struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	TokenRegister  	string `json:"token_register" form:"token_register"`
	Password    	string `json:"password" form:"password"`
	TokenRepassword string `json:"token_repassword" form:"token_repassword"`
}

type UserUpdate struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	Token    		string `json:"token" form:"token"`
	Nama         	string `json:"nama" form:"nama"`
	Email      		string `json:"email" form:"email"`
	NoTelp         	string `json:"no_telp" form:"no_telp"`
	Password    	string `json:"password" form:"password"`
	Foto      		*string `json:"foto" form:"foto"`
	Status         	string `json:"status" form:"status"`
	Verifikasi  	string `json:"verifikasi" form:"verifikasi"`
	CreatedAt       string `json:"created_at" form:"created_at"`
	UpdatedAt       string `json:"updated_at" form:"updated_at"`
	TokenRegister  	string `json:"token_register" form:"token_register"`
	TokenRepassword string `json:"token_repassword" form:"token_repassword"`
}

type DeleteUser struct {
	IdUser 		string `json:"id_user" form:"id_user"`
	DeletedAt  	string `json:"deleted_at" form:"deleted_at"`
}

type CheckAkunUser struct {
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
}

type CheckAkunRead struct {
	IdVerifikasi	string `json:"id_verifikasi" form:"id_verifikasi"`
	IdUser       	string `json:"id_user" form:"id_user"`
	Email       	string `json:"email" form:"email"`
	Status      	string `json:"status" form:"status"`
	ExpiredAt   	string `json:"expired_at" form:"expired_at"`
	UpdatedAt       string `json:"updated_at" form:"updated_at"`
	TokenRegister  	string `json:"token_register" form:"token_register"`
}

type ForgotPasswordUser struct {
	IdForgotPassword	string `json:"id_forgot_password" form:"id_forgot_password"`
	IdUser       		string `json:"id_user" form:"id_user"`
	Email       		string `json:"email" form:"email"`
	CreatedAt   		string `json:"created_at" form:"created_at"`
}

type UserToken struct {
	IdToken     *string `json:"id_token" form:"id_token"`
	FcmToken    string  `json:"fcm_token" form:"fcm_token"`
	IdUser      *string `json:"id_user" form:"id_user" binding:"-"`
	Email       string `json:"email" form:"email" binding:"-"`
	Password    string  `json:"password" form:"password" binding:"-"`
}

type CreateProfil struct {
	IdUser		string `json:"id_user" form:"id_user"`
	Nama        string `json:"nama" form:"nama"`
	Alamat      string `json:"alamat" form:"alamat"`
	Npwp      	*multipart.FileHeader `json:"npwp" form:"npwp"`
	Siup      	*multipart.FileHeader `json:"siup" form:"siup"`
	Tdp      	*multipart.FileHeader `json:"tdp" form:"tdp"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}
type ProfilCreate struct {
	IdProfil	string `json:"id_profil" form:"id_profil"`
	IdUser		string `json:"id_user" form:"id_user"`
	Nama        string `json:"nama" form:"nama"`
	Alamat      string `json:"alamat" form:"alamat"`
	Npwp      	string `json:"npwp" form:"npwp"`
	Siup      	string `json:"siup" form:"siup"`
	Tdp      	string `json:"tdp" form:"tdp"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

type UpdateProfil struct {
	IdProfil	string `json:"id_profil" form:"id_profil"`
	IdUser		string `json:"id_user" form:"id_user"`
	Nama        string `json:"nama" form:"nama"`
	Alamat      string `json:"alamat" form:"alamat"`
	Npwp      	*multipart.FileHeader `json:"npwp" form:"npwp"`
	Siup      	*multipart.FileHeader `json:"siup" form:"siup"`
	Tdp      	*multipart.FileHeader `json:"tdp" form:"tdp"`
}

type ProfilUpdate struct {
	IdProfil	string `json:"id_profil" form:"id_profil"`
	IdUser		string `json:"id_user" form:"id_user"`
	Nama        string `json:"nama" form:"nama"`
	Alamat      string `json:"alamat" form:"alamat"`
	Npwp      	*string `json:"npwp" form:"npwp"`
	Siup      	*string `json:"siup" form:"siup"`
	Tdp      	*string `json:"tdp" form:"tdp"`
	CreatedAt   string `json:"created_at" form:"created_at"`
	UpdatedAt   string `json:"updated_at" form:"updated_at"`
}