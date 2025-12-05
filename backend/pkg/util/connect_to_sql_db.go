package util

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/clr"

	"github.com/glebarez/sqlite"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func log_level() logger.LogLevel {
	log_db_mode := os.Getenv("LOG_DB_MODE")
	log_level := logger.Silent
	switch log_db_mode {
	case "silent":
		log_level = logger.Silent
	case "error":
		log_level = logger.Error
	case "warn":
		log_level = logger.Warn
	case "info":
		log_level = logger.Info
	case "debug":
		log_level = logger.Info
	}
	return log_level
}
func ConnectToSQLDB(dbName, dbHost, dbPort, dbUser, dbPass string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if ConvertStringTo(dbPort, 0) > 0 {
		if db, err = InitMySqlDB(dbUser, dbPass, dbHost, dbPort, dbName); err == nil {
			return db, nil
		} else if db, err = InitPostgreSqlDB(dbUser, dbPass, dbHost, dbPort, dbName); err == nil {
			return db, nil
		} else if db, err = InitMsSqlDB(dbUser, dbPass, dbHost, dbPort, dbName); err == nil {
			return db, nil
		}
	} else if db, err = InitSqlLiteDB(dbName); err == nil {
		return db, nil
	}
	return nil, fmt.Errorf("failed to connect to any database: %v", err)
}
func InitMySqlDB(dbUser, dbPass, dbHost, dbPort, dbName string) (*gorm.DB, error) {
	// Try connecting to information_schema
	infoSchemaURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/information_schema?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort,
	)
	infoSchemaDB, err := gorm.Open(mysql.Open(infoSchemaURI), &gorm.Config{Logger: logger.Default.LogMode(log_level())})
	if err == nil {
		// Check if the database exists
		var dbExists bool
		query := fmt.Sprintf("SELECT EXISTS(SELECT SCHEMA_NAME FROM SCHEMATA WHERE SCHEMA_NAME = '%s')", dbName)
		if err = infoSchemaDB.Raw(query).Scan(&dbExists).Error; err == nil {
			if !dbExists {
				createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
				if err = infoSchemaDB.Exec(createDBQuery).Error; err == nil {
					fmt.Printf("Database %s created successfully\n", dbName)
				}
			}
		}

		// Close the connection to information_schema
		dbSQL, _ := infoSchemaDB.DB()
		dbSQL.Close()
	}

	// Connect directly to the specified database
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{Logger: logger.Default.LogMode(log_level())})
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Get the underlying sql.DB object
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to get db instance: %v", err)
	}

	// Set connection pool parameters
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to the database successfully : " + dbName)
	return db, nil
}
func InitPostgreSqlDB(dbUser, dbPass, dbHost, dbPort, dbName string) (*gorm.DB, error) {
	// Connect to the PostgreSQL information_schema
	infoSchemaURI := fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass,
	)
	infoSchemaDB, err := gorm.Open(postgres.Open(infoSchemaURI), &gorm.Config{Logger: logger.Default.LogMode(log_level())})
	if err == nil {
		// Check if the database exists
		var dbExists bool
		query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
		err = infoSchemaDB.Raw(query).Scan(&dbExists).Error
		if err != nil {
			logrus.Error(err)
			return nil, fmt.Errorf("failed to check if database exists: %v", err)
		}

		// Create the database if it does not exist
		if !dbExists {
			createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
			err = infoSchemaDB.Exec(createDBQuery).Error
			if err != nil {
				logrus.Error(err)
				return nil, fmt.Errorf("failed to create database: %v", err)
			}
			fmt.Printf("Database %s created successfully\n", dbName)
		}

		// Close the connection to information_schema
		dbSQL, err := infoSchemaDB.DB()
		if err != nil {
			logrus.Error(err)
			return nil, fmt.Errorf("failed to get database instance: %v", err)
		}
		dbSQL.Close()
	}

	// Connect to the specified database
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPass,
	)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{Logger: logger.Default.LogMode(log_level())})
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Get the underlying sql.DB object
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to get db instance: %v", err)
	}

	// Set connection pool parameters
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to the PostgreSQL database successfully")
	return db, nil
}

func InitSqlLiteDB(dbName string) (*gorm.DB, error) {
	info := ""
	if dbName == "" {
		dbName = ":memory:?_pragma=foreign_keys(1)"
		info = clr.BgWhite("SQLITE MEMORY")
	} else if strings.HasSuffix(dbName, ".db") {
		dbName = dbName + "?_pragma=foreign_keys(1)"
		info = clr.BgWhite("SQLITE DB FILE " + dbName)
	} else if strings.HasSuffix(dbName, ".sqlite") {
		dbName = dbName + "?_pragma=foreign_keys(1)"
		info = clr.BgWhite("SQLITE DB FILE " + dbName)
	} else if !strings.Contains(dbName, ".") && dbName != ":memory:" {
		dbName = dbName + ".db?_pragma=foreign_keys(1)"
		info = clr.BgWhite("SQLITE DB FILE " + dbName)
	}
	info = clr.TextBlack(info)

	// Connect to the SQLite database
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to connect to SQLite database: %v", err)
	}

	// Get the underlying sql.DB object
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to get db instance: %v", err)
	}

	// Set connection pool parameters
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("DB : " + info)
	return db, nil
}

func InitMsSqlDB(dbUser, dbPass, dbHost, dbPort, dbName string) (*gorm.DB, error) {
	// Validasi parameter penting
	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return nil, fmt.Errorf("missing one or more required database connection parameters")
	}

	// Buat DSN
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Koneksi ke SQL Server
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to connect to SQL Server: %v", err)
	}

	// Ambil *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to get underlying DB instance: %v", err)
	}

	// Atur connection pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Tes koneksi
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping SQL Server: %v", err)
	}

	fmt.Println("Connected to SQL Server successfully")
	return db, nil
}

func InitWebDB(dbURI string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// InitAndCheckDB checks if the specified database exists, creates it if not, and returns a connection to the specified database
func InitAndCheckDB(dbUser, dbPass, dbHost, dbPort, dbName string) (*gorm.DB, error) {
	// Connect to information_schema
	infoSchemaURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/information_schema?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
	)
	var dbExists bool
	infoSchemaDB, err := gorm.Open(mysql.Open(infoSchemaURI), &gorm.Config{})
	if err == nil {
		// Check if the database exists
		query := fmt.Sprintf("SELECT EXISTS(SELECT SCHEMA_NAME FROM SCHEMATA WHERE SCHEMA_NAME = '%s')", dbName)
		if err = infoSchemaDB.Raw(query).Scan(&dbExists).Error; err == nil {
			if !dbExists {
				createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
				if err = infoSchemaDB.Exec(createDBQuery).Error; err == nil {
					fmt.Printf("Database %s created successfully\n", dbName)
				}
			}
		}

		// Create the database if it does not exist
		if !dbExists {
			createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
			err = infoSchemaDB.Exec(createDBQuery).Error
			if err != nil {
				logrus.Error(err)
				return nil, fmt.Errorf("failed to create database: %v", err)
			}
			fmt.Printf("Database %s created successfully\n", dbName)
		}

		// Close the connection to information_schema
		dbSQL, _ := infoSchemaDB.DB()
		dbSQL.Close()
	}

	// Connect to the specified database
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Get the underlying sql.DB object
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to get db instance: %v", err)
	}

	// Set connection pool parameters
	sqlDB.SetMaxIdleConns(10)           // Set the maximum number of idle connections
	sqlDB.SetMaxOpenConns(100)          // Set the maximum number of open connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Set the maximum lifetime of a connection
	return db, nil
}

func InitAndCheckDBViaSSH(dbUser, dbPass, dbHost, dbPort, dbName string,
	sshUser, sshPassword, sshHost string, sshPort int) (*gorm.DB, error) {

	// Register a custom dial function using SSH
	sshDialerName := "mysql+ssh"
	mysqlDriver.RegisterDialContext(sshDialerName, func(ctx context.Context, addr string) (net.Conn, error) {
		sshConfig := &ssh.ClientConfig{
			User: sshUser,
			Auth: []ssh.AuthMethod{
				ssh.Password(sshPassword),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), // You may want to validate the host key in production
			Timeout:         5 * time.Second,
		}

		sshAddr := fmt.Sprintf("%s:%d", sshHost, sshPort)
		sshClient, err := ssh.Dial("tcp", sshAddr, sshConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to dial SSH: %v", err)
		}

		// Forward DB connection through SSH
		dbAddr := fmt.Sprintf("%s:%s", dbHost, dbPort)
		return sshClient.Dial("tcp", dbAddr)
	})

	// Step 1: Connect to information_schema via SSH
	infoSchemaURI := fmt.Sprintf("%s:%s@%s(%s:%s)/information_schema?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, sshDialerName, dbHost, dbPort)

	var dbExists bool
	infoSchemaDB, err := gorm.Open(mysql.Open(infoSchemaURI), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to information_schema: %v", err)
	}

	query := fmt.Sprintf("SELECT EXISTS(SELECT SCHEMA_NAME FROM SCHEMATA WHERE SCHEMA_NAME = '%s')", dbName)
	if err = infoSchemaDB.Raw(query).Scan(&dbExists).Error; err != nil {
		return nil, fmt.Errorf("failed to check database existence: %v", err)
	}

	if !dbExists {
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
		if err = infoSchemaDB.Exec(createDBQuery).Error; err != nil {
			return nil, fmt.Errorf("failed to create database: %v", err)
		}
		fmt.Printf("Database %s created successfully\n", dbName)
	}

	// Close info_schema connection
	dbSQL, _ := infoSchemaDB.DB()
	dbSQL.Close()

	// Step 2: Connect to actual target DB
	dbURI := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, sshDialerName, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s via SSH tunnel: %v", dbName, err)
	}

	// Connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get db instance: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func InitPostgreSqlDBViaSSH(
	dbUser, dbPass, dbHost, dbPort, dbName string,
	sshUser, sshPassword, sshHost string, sshPort int,
) (*gorm.DB, error) {

	// SSH config
	sshConfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // use secure method in production!
		Timeout:         5 * time.Second,
	}

	// Connect to SSH
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshHost, sshPort), sshConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to dial SSH: %v", err)
	}

	// Custom dialer for PostgreSQL to use SSH tunnel
	customDialer := func(ctx context.Context, network, addr string) (net.Conn, error) {
		return sshClient.Dial(network, fmt.Sprintf("%s:%s", dbHost, dbPort))
	}

	// Create pgx config and override dialer
	pgxConfig, err := pgx.ParseConfig(fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPass, dbName))
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx config: %v", err)
	}
	pgxConfig.DialFunc = customDialer

	// Use stdlib for GORM
	sqlDB := stdlib.OpenDB(*pgxConfig)

	// Connect with GORM
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed to connect to PostgreSQL over SSH: %v", err)
	}

	return db, nil
}
