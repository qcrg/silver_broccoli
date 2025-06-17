package postgres

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/qcrg/silver_broccoli/utils"
	"github.com/qcrg/silver_broccoli/utils/initiator"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MockConfig struct {
	mock.Mock
}

func (t *MockConfig) GetConnectionString() string {
	args := t.Called()
	return args.String(0)
}

func (t *MockConfig) GetTLSMod() string {
	args := t.Called()
	return args.String(0)
}

type StdOutLogger struct{}

func (StdOutLogger) Printf(format string, v ...any) {
	fmt.Printf(format, v...)
}

var test_config Config

func TestMain(m *testing.M) {
	initiator.DefaultInitAll()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	const (
		username = "guest"
		passwd   = "asdf;lkj"
		db_name  = "test"
	)

	pq_cont, err := postgres.Run(
		ctx,
		"postgres:alpine",
		postgres.WithDatabase(db_name),
		postgres.WithUsername(username),
		postgres.WithPassword(passwd),
		postgres.WithInitScripts(
			filepath.Join(utils.GetProjectDir(), "testdata/dev/postgres/init/init.sql"),
		),
		tc.WithHostConfigModifier(func(host_conf *container.HostConfig) {
			host_conf.Binds = append(
				host_conf.Binds,
				filepath.Join(utils.GetProjectDir(), "/testdata/data:/csv_import:ro"),
			)
		}),
		tc.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
		tc.WithLogger(StdOutLogger{}),
	)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run postgres container")
	}
	defer func() {
		err := pq_cont.Terminate(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to terminate container")
		}
	}()

	conn_str, err := pq_cont.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get connection string")
	}

	mock_conf := &MockConfig{}
	mock_conf.On("GetConnectionString").Return(conn_str)
	mock_conf.On("GetTLSMod").Return("disable")
	test_config = mock_conf
	log.Warn().Msg(conn_str)

	os.Exit(m.Run())
}

func get_test_database(t *testing.T) *DB {
	db, err := NewDatabase(test_config)
	require.NoError(t, err)
	return db
}
