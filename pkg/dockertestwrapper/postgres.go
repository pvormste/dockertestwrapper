package dockertestwrapper

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// PostgresImageName is the image name of the postgres docker image
const PostgresImageName string = "postgres"

// Possible postgres image versions
const (
	PostgresImageVersion10 string = "10"
	PostgresImageVersion11 string = "11"
)

// Default postgres connection details
const (
	DefaultPostgresPort     string = "5432/tcp"
	DefaultPostgresUser     string = "postgres"
	DefaultPostgresPassword string = "postgres"
	DefaultPostgresDatabase string = "postgres"
)

func InitPostgresContainer(version string) (*WrapperInstance, error) {
	userEnv := fmt.Sprintf("POSTGRES_USER=%s", DefaultPostgresUser)
	passwordEnv := fmt.Sprintf("POSTGRES_PASSWORD=%s", DefaultPostgresPassword)
	databaseEnv := fmt.Sprintf("POSTGRES_DB=%s", DefaultPostgresDatabase)

	params := WrapperParams{
		ImageName:           PostgresImageName,
		ImageVersion:        version,
		EnvVariables:        []string{userEnv, passwordEnv, databaseEnv},
		ContainerPort:       DefaultPostgresPort,
		AfterInitActionFunc: postgresAfterInitAction,
	}

	return InitContainer(params)
}

func postgresAfterInitAction(dockerHost string, hostPort int) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dockerHost, hostPort, DefaultPostgresUser, DefaultPostgresPassword, DefaultPostgresDatabase)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	return db.Ping()
}
