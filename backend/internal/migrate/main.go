package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// migrationSequenceLength is the number of digits included in the prefix of created migrations.
const migrationSequenceLength = 6

// Migrate interacts with database migrations.
type Migrate struct {
	migrate  *migrate.Migrate
	filePath string
}

// New creates a new database migrator.
func New(migrationFilepath string, databaseURL string) (*Migrate, error) {
	cleanFilepath := filepath.Clean(migrationFilepath)
	sourceURL := fmt.Sprintf("file://%s", cleanFilepath)
	migrate, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		return nil, err
	}
	return &Migrate{
		migrate:  migrate,
		filePath: cleanFilepath,
	}, nil
}

// Create creates a migration.
func (m Migrate) Create(name string) error {
	ext := ".sql"
	seqGlob := filepath.Join(m.filePath, "*"+ext)
	matches, err := filepath.Glob(seqGlob)
	if err != nil {
		return err
	}

	version, err := nextSeqVersion(matches)
	if err != nil {
		return err
	}

	versionGlob := filepath.Join(m.filePath, version+"_*"+ext)
	matches, err = filepath.Glob(versionGlob)
	if err != nil {
		return err
	}
	if len(matches) > 0 {
		return fmt.Errorf("duplicate migration version: %s", version)
	}

	if err = os.MkdirAll(m.filePath, os.ModePerm); err != nil {
		return err
	}

	for _, direction := range []string{"up", "down"} {
		basename := fmt.Sprintf("%s_%s.%s%s", version, name, direction, ext)
		filename := filepath.Join(m.filePath, basename)
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			return err
		}
		f.WriteString("BEGIN;\n\n")
		if direction == "up" {
			f.WriteString("-- Replace with upgrade SQL\n\n")
		} else {
			f.WriteString("-- Replace with downgrade SQL\n\n")
		}
		f.WriteString("COMMIT;\n")
		f.Close()
	}
	return nil
}

// nextSeqVersion constructs the version string of the next migration to be created.
// Adapted from this non-exported golang-migrate function:
// https://github.com/golang-migrate/migrate/blob/511ae9f5b6bea5190b340243169359455281458b/internal/cli/commands.go#L23
func nextSeqVersion(matches []string) (string, error) {
	nextSeq := uint64(1)
	if len(matches) > 0 {
		filename := matches[len(matches)-1]
		matchSeqStr := filepath.Base(filename)
		idx := strings.Index(matchSeqStr, "_")
		if idx < 1 { //There should be at least 1 digit
			return "", fmt.Errorf("malformed migration filename: %s", filename)
		}
		var err error
		matchSeqStr = matchSeqStr[0:idx]
		nextSeq, err = strconv.ParseUint(matchSeqStr, 10, 64)
		if err != nil {
			return "", err
		}
		nextSeq++
	}

	version := fmt.Sprintf("%0[2]*[1]d", nextSeq, migrationSequenceLength)
	if len(version) > migrationSequenceLength {
		return "", fmt.Errorf("next sequence number %s too large. At most %d digits are allowed", version, migrationSequenceLength)
	}
	return version, nil
}
