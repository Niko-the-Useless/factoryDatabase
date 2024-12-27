package lib

type Product struct {
	Id *int64
	Name *string
	Production_time *float32
}

type Machine struct {
	Id *int64
	Name *string
	Crafting_speed *float32
	Polution *float32
	Module_slot *int
	Q_coef_a *float32
	Q_coef_b *float32
	Q5_mod *float32
	Drain *float32
	Energy_consumption *float32
}

type BOM struct {
	Parent_id *[]int64
	Parent_quantity *[]int64
	Child_id *int64
	Child_quantity *int64
	Byproduct_id *[]int64
	Byproduct_quantity *[]int64
}
