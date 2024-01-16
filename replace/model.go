package replace

// ReplaceItem
type ReplaceItem struct {
	From string `structs:"fo" json:"fo" bson:"fo"`
	To   string `structs:"to" json:"to" bson:"to"`
}
