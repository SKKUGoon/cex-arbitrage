// Code generated by ent, DO NOT EDIT.

package tradelog

const (
	// Label holds the string label denoting the tradelog type in the database.
	Label = "trade_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldExchange holds the string denoting the exchange field in the database.
	FieldExchange = "exchange"
	// FieldTicker holds the string denoting the ticker field in the database.
	FieldTicker = "ticker"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// FieldStrategy holds the string denoting the strategy field in the database.
	FieldStrategy = "strategy"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldLeverage holds the string denoting the leverage field in the database.
	FieldLeverage = "leverage"
	// Table holds the table name of the tradelog in the database.
	Table = "tradelog"
)

// Columns holds all SQL columns for tradelog fields.
var Columns = []string{
	FieldID,
	FieldDate,
	FieldExchange,
	FieldTicker,
	FieldPosition,
	FieldStrategy,
	FieldPrice,
	FieldQuantity,
	FieldLeverage,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
