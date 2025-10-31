package models

type Bioskop struct {
	ID int `json:"id"`
	Nama string `json:"nama"`
	Lokasi string `json:"lokasi"`
	Rating float64 `json:"rating"`
}

type BioskopInput struct {
	Nama string `json:"nama" binding:"required"`
	Lokasi string `json:"lokasi" binding:"required"`
	Rating float64 `json:"rating"`
}