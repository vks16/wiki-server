package entity

type Creator struct {
	FirstName string `json:"firstName" binding:"min=3,max=15,required"`
	LastName  string `json:"lastName" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" validate:"required,email"`
}

//  xml:"title" form:"title" "validate":"email" binding:"required"
type Video struct {
	Title       string  `json:"title" binding:"min=3,max=55" validate:"is-cool"`
	Description string  `json:"description" binding:"min=5,max=150"`
	URL         string  `json:"url" binding:"required,url"`
	Creator     Creator `json:"creator" binding:"required"`
}
