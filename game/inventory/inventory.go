package inventory

type InventoryType struct {
	Items   []ItemType
	BagSize int
}

type ItemType struct {
	Name     string
	Quantity int
}
