package main

import (
	"time"

	"github.com/Wexlersolk/GriefBlades/internal/db"
	"github.com/Wexlersolk/GriefBlades/internal/env"
	"github.com/Wexlersolk/GriefBlades/internal/grief/cache"
	"github.com/Wexlersolk/GriefBlades/internal/ratelimiter"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

const version = "1.1.0"

func main() {
	cfg := config{
		addr:        env.GetString("ADDR", "8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:8181"),
		db: dbConfig{
			addr:         env.GetString("DBADDR", "you can enter your db adress here"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "5m"),
		},
		redisCfg: redisConfig{
			addr:    env.GetString("REDDIS_ADDR", "localhost:5173"),
			pw:      env.GetString("REDDIS_PW", ""),
			db:      env.GetInt("REDDIS_DB", 0),
			enabled: env.GetBool("REDDIS_ENABLED", false),
		},

		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:       time.Hour * 24,
			fromEmail: env.GetString("FROM_EMAIL", ""),
			sendGrid: sendGridConfig{
				apiKey: env.GetString("SEND_API_KEY", ""),
			},
		},
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "secret"),
				exp:    time.Hour * 24,
				iss:    "griefblades",
			},
		},
		rateLimiter: ratelimiter.Config{
			RequestPerTimeFrame: env.GetInt("RATELIMITER_REQUEST_COUNT", 20),
			TimeFrame:           time.Hour,
			Enambled:            env.GetBool("RATELIMITER_ENABLED", true),
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(
		cfg.db.addr,
		cfg.db.tool,
		cfg.db.maxIdleConns,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	var rdb *redis.Client
	if cfg.redisCfg.enabled {
		rdb = cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.pw, cfg.redisCfg.db)
		logger.Info("redis cashe connection established")

		defer rdb.Close()
	}

	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.rateLimiter.RequestPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)
	
	mailer := mailer.

}
