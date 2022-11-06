package remodel

type AppProfile struct {
	Amenities           []Amenity            `json:"amenities"`
	RealEstateTypes     []RealEstateType     `json:"real_estate_types" `
	RealEstatePostTypes []RealEstatePostType `json:"real_estate_post_types" `
}
