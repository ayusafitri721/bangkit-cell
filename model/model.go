package model

import "time"

type Repair struct {
	ID        int       
	Customer  string    
	Phone     string    
	Issue     string    
	Price     int      
	Status    string    
	CreatedAt time.Time 
	UpdatedAt time.Time 
}