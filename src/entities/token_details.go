package entities

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"access_uuid"`
	RefreshUuid  string `json:"refresh_uuid"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

/**
** Em nossa solução proposta, em vez de apenas criar um token, precisaremos criar dois JWTs:
** o token de acesso, o token de atualização. A estrutura TokenDetails contém as informações
** de tokens, seus períodos de expiração e uuids. O período de expiração e os uuids são muito
** úteis porque serão usados ​​ao salvar metadados de token no redis.
**/
