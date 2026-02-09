package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 1. On crée l'application principale
	app := NewApp()

	// 2. On FORCE l'initialisation manuelle de la DB/Crypto avant le Bind
	// (Sinon app.db est nil jusqu'à ce que Wails appelle OnStartup)
	app.startup(context.Background())

	// 3. MAINTENANT on peut créer le handler, car app.db n'est plus nil
	evalHandler := NewEvalHandler(app.db, app.cryptoSvc)

	wails.Run(&options.App{
		Title:  "SGBD Psychosocial",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// On garde OnStartup pour les autres besoins de Wails
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
			evalHandler,
		},
	})
}
