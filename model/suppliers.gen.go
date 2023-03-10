// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSupplier = "suppliers"

// Supplier mapped from table <suppliers>
type Supplier struct {
	SupplierID        int32   `gorm:"column:supplier_id;primaryKey;autoIncrement:true" json:"supplier_id"`
	SupplierName      string  `gorm:"column:supplier_name;not null" json:"supplier_name"`
	SupplierEmail     *string `gorm:"column:supplier_email" json:"supplier_email"`
	SupplierTelephone *string `gorm:"column:supplier_telephone" json:"supplier_telephone"`
	SupplierWeb       *string `gorm:"column:supplier_web" json:"supplier_web"`
	SupplierNpwp      *string `gorm:"column:supplier_npwp" json:"supplier_npwp"`
	SupplierAddress   *string `gorm:"column:supplier_address" json:"supplier_address"`
	SupplierType      string  `gorm:"column:supplier_type;not null" json:"supplier_type"`
}

// TableName Supplier's table name
func (*Supplier) TableName() string {
	return TableNameSupplier
}
