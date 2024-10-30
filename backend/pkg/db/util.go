package db

import "github.com/jackc/pgx/v5/pgtype"

func Text(value string) pgtype.Text {
	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}

func BigInt(value int64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: value,
		Valid: true,
	}
}

// func
