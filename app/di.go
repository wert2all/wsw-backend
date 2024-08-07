package app

import (
	"context"
	"net/http"
	"time"
	"wsw/backend/domain/gowitness"
	"wsw/backend/domain/token/generator"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
	"wsw/backend/lib/utils"
	"wsw/backend/model/token"
	"wsw/backend/model/url"

	"github.com/golobby/container/v3"
)

func initDi(config Config, appContext context.Context) {
	initService(func() context.Context { return appContext })
	initService(func() (*ent.Client, error) { return newDBClient(config.Postgres, appContext) })
	initService(func() App { return appImpl{router: newRouter()} })

	initService(func() generator.TokenGenerator { return generator.NewTokenGenerator() })
	initService(func() gowitness.Client {
		client := http.Client{
			Timeout: time.Second * 2, // Timeout after 2 seconds
		}
		return gowitness.NewClient(client, config.API.BaseURL, config.App.Host.Images)
	})

	initService(func(client *ent.Client, ctx context.Context) repository.Token {
		return repository.NewToken(client, ctx)
	})
	initService(func(client *ent.Client, ctx context.Context) repository.Url {
		return repository.NewUrl(client, ctx)
	})

	initService(func(generator generator.TokenGenerator, tokenRepository repository.Token) token.Token {
		return token.NewModel(generator, tokenRepository)
	})
	initService(func(urlRepository repository.Url, client gowitness.Client) url.Url {
		return url.NewUrl(urlRepository, client)
	})
}

func initService(resolver interface{}) {
	err := container.Singleton(resolver)
	if err != nil {
		utils.F("Couldnt inititalize service: %v", err)
	}
}
