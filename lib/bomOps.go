package lib
import(
	"database/sql"
	"fmt"
)
func CreateBomTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS bom (
		parent_id INTEGER NOT NULL,
		parent_quantity INTEGER,
		child_id INTEGER,
		child_quantity INTEGER,
		byproduct_id INTEGER,
		byproduct_quantity INTEGER,
		FOREIGN KEY (parent_id) REFERENCES products(id) ON DELETE CASCADE,
		FOREIGN KEY (child_id) REFERENCES products(id) ON DELETE CASCADE,
		FOREIGN KEY (byproduct_id) REFERENCES products(id) ON DELETE CASCADE
	);`
	return db.Exec(sql)
}

func (b *BOM) InsertBOM(db *sql.DB) (int64, error) {
	var rowCount int
	var err error
	var result sql.Result

	sql := `INSERT INTO bom (parent_id, parent_quantity, child_id, child_quantity, byproduct_id, byproduct_quantity)
			VALUES (?, ?, ?, ?, ?, ?)`

	if b.Child_id == nil || b.Child_quantity == nil {
		return 0, fmt.Errorf("child_id and child_quantity are required")
	}

	parentCount := 0
	if b.Parent_id != nil && b.Parent_quantity != nil {
		if len(*b.Parent_id) != len(*b.Parent_quantity) {
			return 0, fmt.Errorf("parent_id and parent_quantity arrays must have the same length")
		}
		parentCount = len(*b.Parent_id)
	}

	byproductCount := 0
	if b.Byproduct_id != nil && b.Byproduct_quantity != nil {
		if len(*b.Byproduct_id) != len(*b.Byproduct_quantity) {
			return 0, fmt.Errorf("byproduct_id and byproduct_quantity arrays must have the same length")
		}
		byproductCount = len(*b.Byproduct_id)
	}

	if parentCount >= byproductCount {
		rowCount = parentCount
	} else {
		rowCount = byproductCount
	}

	for i := 0; i < rowCount; i++ {
		var parentID, parentQuantity, byproductID, byproductQuantity *int64

		if b.Parent_id != nil && i < len(*b.Parent_id) {
			parentID = &(*b.Parent_id)[i]
		}
		if b.Parent_quantity != nil && i < len(*b.Parent_quantity) {
			parentQuantity = &(*b.Parent_quantity)[i]
		}
		if b.Byproduct_id != nil && i < len(*b.Byproduct_id) {
			byproductID = &(*b.Byproduct_id)[i]
		}
		if b.Byproduct_quantity != nil && i < len(*b.Byproduct_quantity) {
			byproductQuantity = &(*b.Byproduct_quantity)[i]
		}

		result, err = db.Exec(sql,
			nullOrValue(parentID),
			nullOrValue(parentQuantity),
			*b.Child_id,
			*b.Child_quantity,
			nullOrValue(byproductID),
			nullOrValue(byproductQuantity),
		)
		if err != nil {
			return 0, fmt.Errorf("failed to insert BOM: %v", err)
		}
	}

	if result != nil {
		return result.LastInsertId()
	}
	return 0, nil
}

func nullOrValue(ptr *int64) interface{} {
	if ptr == nil {
		return nil
	}
	return *ptr
}

