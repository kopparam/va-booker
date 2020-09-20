package main

type getBookableClassesRequest struct {
	Ampm     string `json:"AMPM"`
	Category int
	ISODate  string
	SiteID   string
}
