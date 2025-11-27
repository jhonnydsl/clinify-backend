package dtos

type AdminInput struct {
	FullName  		string 		`json:"full_name" binding:"required"`
	Email     		string 		`json:"email" binding:"required,email"`
	Password  		string 		`json:"password" binding:"required"`
	BirthDate 		string 		`json:"birth_date"`
	Crp 	  		string 		`json:"crp"`
	Bio 			string 		`json:"bio"`
	ProfileImage 	string 		`json:"profile_image"`
	OfficeAddress 	string 		`json:"office_address"`
	Phone 			string 		`json:"phone"`
	PublicSlug 		string 		`json:"public_slug"`
}

type LoginInput struct {
	Email 		string `json:"email" binding:"required,email"`
	Password 	string `json:"password" binding:"required"`
}