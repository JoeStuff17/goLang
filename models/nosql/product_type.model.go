package nosql

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductType struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Attribute   string             `bson:"attribute,omitempty" json:"attribute"`
	Description string             `bson:"description,omitempty" json:"description"`
	CreatedBy   CreatedBy          `bson:"createdBy,omitempty" json:"createdBy"`
	CategoryIds []string           `bson:"categoryIds,omitempty" json:"categoryIds"`
	IsActive    bool               `bson:"isActive,omitempty" json:"isActive"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

type CreatedBy struct {
	UserId   uint16 `bson:"userId,omitempty" json:"userId"`
	UserName string `bson:"userName,omitempty" json:"userName"`
}
