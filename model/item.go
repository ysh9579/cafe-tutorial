package model

import "time"

type Items []Item

type Item struct {
	ItemSeq     int64     `json:"item_seq,omitempty"`
	AdminSeq    int64     `json:"admin_seq,omitempty"`
	Category    int       `json:"category,omitempty"`
	Barcode     string    `json:"barcode,omitempty"`
	Price       int64     `json:"price,omitempty"`
	Cost        int64     `json:"cost,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	ExpireDT    time.Time `json:"expire_dt"`
	Size        int       `json:"size,omitempty"`
	RegDT       time.Time `json:"reg_dt"`
	ModDT       time.Time `json:"mod_dt"`
}
