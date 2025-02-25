package modelos

type DadosAutenticacao struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}
