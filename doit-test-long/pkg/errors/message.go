package errors

type ErrorMessage struct {
	EN string
	ID string
}

type ErrorMessages map[string]ErrorMessage

var (
	EM = ErrorMessages{
		"internal": ErrorMessage{
			EN: `Internal Server Error. Please Call Administrator.`,
			ID: `Terjadi Kendala Pada Server. Mohon Hubungi Administrator.`,
		},
		"notfound": ErrorMessage{
			EN: `Record Does Not Exist. Please Validate Your Input Or Contact Administrator.`,
			ID: `Data Tidak Diketemukan. Mohon Cek Kembali Masukkan Anda Atau Hubungi Administrator.`,
		},
		"badrequest": ErrorMessage{
			EN: `Invalid Input. Please Validate Your Input.`,
			ID: `Kesalahan Input. Mohon Cek Kembali Masukkan Anda.`,
		},
		"unauthorized": ErrorMessage{
			EN: `Unauthorized Access. You are not authorized to access this resource.`,
			ID: `Akses Ditolak. Anda Belum Diijinkan Untuk Mengakses Aplikasi.`,
		},
		"uniqueconst": ErrorMessage{
			EN: `Record has existed and must be unique. Please Validate Your Input Or Contact Administrator.`,
			ID: `Data sudah ada. Mohon Cek Kembali Masukkan Anda Atau Hubungi Administrator.`,
		},
	}
)

func (em ErrorMessages) Message(lang string, i string) string {
	if lang == "ID" {
		return em[i].ID
	}
	return em[i].EN
}
