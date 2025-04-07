package models

type ContactInfo struct {
	Organization      string `bson:"organization"`
	PostalAddress     string `bson:"postal_address"`
	Location          string `bson:"location"`
	ResponsiblePerson string `bson:"responsible_person"`
	Email             string `bson:"email"`
	Phone             string `bson:"phone"`
	Fax               string `bson:"fax"`
}
