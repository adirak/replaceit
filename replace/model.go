package replace

// ReplaceItem
type ReplaceItem struct {
	From string `structs:"from" json:"from" bson:"from"`
	To   string `structs:"to" json:"to" bson:"to"`
}
