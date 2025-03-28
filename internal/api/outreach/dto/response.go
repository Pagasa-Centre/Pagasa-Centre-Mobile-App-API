package dto

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/domain"
)

type Outreach struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	Postcode     string `json:"postcode"`
	City         string `json:"city"`
	Country      string `json:"country"`
	VenueName    string `json:"venue_name"`
	Region       string `json:"region"`
	ThumbnailURL string `json:"thumbnail_url"`
}

func ToResponse(outreach *domain.Outreach) *Outreach {
	return &Outreach{
		ID:           outreach.ID,
		Name:         outreach.Name,
		AddressLine1: outreach.AddressLine1,
		AddressLine2: outreach.AddressLine2,
		Postcode:     outreach.Postcode,
		City:         outreach.City,
		Country:      outreach.Country,
		VenueName:    outreach.VenueName,
		Region:       outreach.Region,
		ThumbnailURL: outreach.ThumbnailURL,
	}
}

type GetAllOutreachesResponse struct {
	Outreaches []*Outreach `json:"outreaches"`
}

func ToGetAllOutreachesResponse(outreaches []*domain.Outreach) GetAllOutreachesResponse {
	var outreachesResp GetAllOutreachesResponse

	for _, ministry := range outreaches {
		outreachesResp.Outreaches = append(outreachesResp.Outreaches, ToResponse(ministry))
	}

	return outreachesResp
}
