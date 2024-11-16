package main

type Product struct {
	name string
	production_time float32
}

type Machine struct {
	name string
	crafting_speed float32
	polution float32
	module_slot int
	q_coef_a float32
	q_coef_b float32
	q5_mod float32
	drain float32
	energy_consumption float32
}

type BOM struct {
	parent_id int
	parent_quantity []int
	child_id []int 
	child_quantity []int
	byproduct_id []int
	byproduct_quantity []int
}
