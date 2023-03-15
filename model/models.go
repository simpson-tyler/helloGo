package model

import ( 
    "database/sql" 
    "time"
)

type Entity struct {
    id  string   `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt sql.NullTime `json:"deleted_at"`
}

type Apple struct {
    id  string `json:"id"`
    color string `json:"color"`
    variety string `json:"variety"`
    PollinationDate time.Time   `json:"date_of_pollination"`
    tree_id string  `json:"tree_id"`
}

type AppleEntity struct {
    Entity
    Apple
}
