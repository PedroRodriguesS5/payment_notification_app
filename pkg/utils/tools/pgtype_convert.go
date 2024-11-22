package tools

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

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
