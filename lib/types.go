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
	Parent_id int
	Parent_quantity []int
	Child_id []int 
	Child_quantity []int
	Byproduct_id []int
	Byproduct_quantity []int
}
