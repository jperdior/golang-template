package bootstrap

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	"golang-template/internal/platform/bus/inmemory"
	"golang-template/internal/platform/database/mysql"
	"golang-template/internal/platform/mailer"
	"golang-template/internal/platform/server"
	"time"
)

func Run() error {

	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return err
	}

	mysql.ConnectDB(mysql.DatabaseConfig{
		User:     cfg.DatabaseUser,
		Password: cfg.DatabasePassword,
		Host:     cfg.DatabaseHost,
		Port:     cfg.DatabasePort,
		Name:     cfg.DatabaseName,
	})

	mailer.NewMailer(mailer.MailerConfig{
		Host:     cfg.MailerHost,
		Port:     cfg.MailerPort,
		User:     cfg.MailerUser,
		Password: cfg.MailerPassword,
	})

	var (
		commandBus = inmemory.NewCommandBus()
		queryBus   = inmemory.NewQueryBus()
		eventBus   = inmemory.NewEventBus()
	)

	ctx, srv := server.New(
		context.Background(),
		cfg.Host,
		cfg.Port,
		cfg.ShutdownTimeout,
		commandBus,
		queryBus,
		eventBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:""`
	Port            uint          `default:"9091"`
	ShutdownTimeout time.Duration `default:"10s"`

	// JWT configuration
	JwtSecret     string `required:"true"`
	JwtExpiration int    `default:"15"`

	// Database configuration
	DatabaseUser     string `required:"true"`
	DatabasePassword string `required:"true"`
	DatabaseHost     string `required:"true"`
	DatabasePort     int    `required:"true"`
	DatabaseName     string `required:"true"`

	// Mailer configuration
	MailerHost     string `required:"true"`
	MailerPort     int    `required:"true"`
	MailerUser     string `required:"true"`
	MailerPassword string `required:"true"`
}
