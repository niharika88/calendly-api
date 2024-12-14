package dbmate

import (
	"context"
	"log/slog"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
)

func Migrate(ctx context.Context, dbDns string, debug bool) {
	slog.InfoContext(ctx, "Migrating database...")
	u, _ := url.Parse(dbDns)
	db := dbmate.New(u)

	if err := db.CreateAndMigrate(); err != nil {
		slog.ErrorContext(ctx, "Error migrating database", "error", err)
		panic(err)
	}
	slog.InfoContext(ctx, "Migrating database Done...")
}
