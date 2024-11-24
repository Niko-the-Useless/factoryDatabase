package lib

type Product struct {
	Name string
	Production_time float32
}

type Machine struct {
	Name string
	Crafting_speed float32
	Polution float32
	Module_slot int
	Q_coef_a float32
	Q_coef_b float32
	Q5_mod float32
	Drain float32
	Energy_consumption float32
}

type BOM struct {
	Parent_id int
	Parent_quantity []int
	Child_id []int 
	Child_quantity []int
	Byproduct_id []int
	Byproduct_quantity []int
}

type Target struct{
	Id *int
	Name *string
}

//udaje że mam enumy
type ProductUpdate int
const(
	Name = iota
	Production_time
)
/*
type MachineUpdate int
const(
	Name = iota
	Crafting_speed 
	Polution 
	Module_slot 
	Q_coef_a 
	Q_coef_b 
	Q5_mod 
	Drain 
	Energy_consumption 
)
*/
