package entities

type AccessDetails struct {
	AccessUUID string         `json:"access_uuid"`
	UserID     UniqueEntityID `json:"user_id"`
}

/**
** Essa estrutura contém os metadados ( access_uuide user_id) de que precisaremos para fazer
** uma consulta no Redis. Se houver algum motivo pelo qual não conseguimos obter os metadados
** deste token, a solicitação é interrompida com uma mensagem de erro.
**/
