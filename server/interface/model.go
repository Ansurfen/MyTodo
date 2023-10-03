package interfaces

type Po interface {
	Table
}

type Dao interface {
	Create(v Po) error
	Delete()
	FindOne()
	FineMany()
}
