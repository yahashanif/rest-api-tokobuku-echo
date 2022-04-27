package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Buku struct {
	ID           int    `json:"id"`
	IdKategori   int    `json:"id_kategori"`
	NamaKategori string `json:"nama_kategori"`
	Judul        string `json:"judul"`
	Pengarang    string `json:"pengarang"`
	Penerbit     string `json:"penerbit"`
	TahunTerbit  int    `json:"tahun_terbit"`
	JumlahBuku   int    `json:"jumlah_buku"`
	Harga        int    `json:"harga"`
	Gambar       string `json:"gambar"`
	Deskripsi    string `json:"deskripsi"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
