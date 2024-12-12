package tools

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Number interface {
	int32 | int64 | float32 | float64
}

func ConvertUUIDToString(uuidVal pgtype.UUID) (string, error) {
	if len(uuidVal.Bytes) != 16 {
		return "", fmt.Errorf("invalid UUID value")
	}

	// Properly format the UUID bytes into the standard UUID string format
	s := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuidVal.Bytes[0:4],   // First 4 bytes
		uuidVal.Bytes[4:6],   // Next 2 bytes
		uuidVal.Bytes[6:8],   // Next 2 bytes
		uuidVal.Bytes[8:10],  // Next 2 bytes
		uuidVal.Bytes[10:16], // Remaining 6 bytes
	)

	return s, nil
}

func ConvertStringToPgtypeText(stringValue string) pgtype.Text {
	var pgtext pgtype.Text
	pgtext.String = stringValue
	pgtext.Valid = true

	return pgtext
}

func ConvertStringToUUID(idString string) (pgtype.UUID, error) {
	if len(idString) == 0 {
		return pgtype.UUID{}, fmt.Errorf("invalid id")
	}
	decoded, err := hex.DecodeString(strings.ReplaceAll(idString, "-", ""))

	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("invalid UUID format: %w", err)
	}

	var pgUUID pgtype.UUID
	copy(pgUUID.Bytes[:], decoded)
	pgUUID.Valid = true

	return pgUUID, nil
}

func ConvertNumberTypeInPgType[V Number](number V) (interface{}, error) {
	switch any(number).(type) {
	case int16, int8:
		return pgtype.Int2{
			Int16: int16(number),
			Valid: true,
		}, nil
	case int32:
		return pgtype.Int4{
			Int32: int32(number),
			Valid: true,
		}, nil
	case int, int64:
		return pgtype.Int8{
			Int64: int64(number),
			Valid: true,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported number type: %T", number)
	}
}

func ConvertStringToDate(dateString string) (pgtype.Date, error) {
	if len(dateString) == 0 {
		return pgtype.Date{}, fmt.Errorf("invalid input: date string is empty")
	}
	var dateFormat pgtype.Date
	parseDate, err := time.Parse("2006-01-02", dateString)

	if err != nil {
		return pgtype.Date{}, fmt.Errorf("error parsing date string '%s': expected format 'YYYY-MM-DD', got: %v", dateString, err)
	}

	dateFormat.Time = parseDate
	dateFormat.Valid = true

	return dateFormat, nil
}

// ConvertToNumeric converts a generic numeric value to pgtype.Numeric
func ConvertToNumeric(value interface{}) (pgtype.Numeric, error) {
	var numeric pgtype.Numeric

	// Handle supported numeric types
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		numeric.Int = big.NewInt(reflect.ValueOf(v).Int())
		numeric.Valid = true
	case float32, float64:
		f := big.NewFloat(reflect.ValueOf(v).Float())
		intPart := new(big.Int)

		// extract integer part and exponent
		f.Int(intPart)
		exp := f.MantExp(nil)
		numeric.Int = intPart
		numeric.Exp = int32(exp)
		numeric.Valid = true
	default:
		return numeric, errors.New("unsupported type for pgtype.Numeric conversion")
	}

	return numeric, nil
}

// ConvertToInt2 converts a generic numeric value to pgtype.Int2
func ConvertToInt2(value interface{}) (pgtype.Int2, error) {
	var int2 pgtype.Int2

	// Handle int types specifically for Int2
	switch v := value.(type) {
	case int8, int16, int32, int:
		intValue := reflect.ValueOf(v).Int()
		if intValue < -32768 || intValue > 32767 {
			return int2, errors.New("value out of range for pgtype.Int2")
		}
		int2.Int16 = int16(intValue)
		int2.Valid = true
	default:
		return int2, errors.New("unsupported type for pgtype.Int2 conversion")
	}

	return int2, nil
}
