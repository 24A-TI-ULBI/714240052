package model

type Absensi struct {
	ID         string `json:"id" bson:"_id"`
	NPM        string `json:"npm" bson:"npm"`
	MatkulCode string `json:"matkulCode" bson:"matkulCode"`
	Tanggal    string `json:"tanggal" bson:"tanggal"`
	Status     string `json:"status" bson:"status"`
	Timestamp  string `json:"timestamp" bson:"timestamp"`
}
