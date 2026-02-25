package models

// AppResult est l'objet unique qui traverse le pont Wails
type AppResult struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`       // N'importe quel objet (Client, Note, etc.)
	Message    string      `json:"message"`    // Message pour l'UI (ex: "Client créé")
	ErrorCode  string      `json:"errorCode"`  // Code technique (ex: "AUTH_001")
	ErrorTrace string      `json:"errorTrace"` // Pour le debug (nom du fichier/ligne)
}
