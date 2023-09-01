package entities

import (
	vo "go_mongo_api/src/entities/value_objects"
)

/**
* As tags json e bson indicam que um User pode ser serializado nestes formatos
**/

type User struct {
	Base
	Name     vo.Name     `json:"name" bson:"name"`
	Email    vo.Email    `json:"email" bson:"email"`
	Password vo.Password `json:"-" bson:"password"`
}

func BuildUser(name string, email string, password string) (*User, error) {
	user := &User{
		Name:     vo.Name{Value: name},
		Email:    vo.Email{Value: email},
		Password: vo.Password{Value: password},
	}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

/*
1 - O trecho user *User antes do nome da função representa uma amarração entre a função
Prepare e a struc User, ou seja, é como se User fosse uma classe e a função Prepare
fosse um método público.
2 - É usado o * no User, pois todas as vezes que o user for alterado ele será atualizado
em todos os objetos pois esta utilizando o mesmo local na memória
3 - O retorno da função é um error que pode ter valor nil, ou seja, se o erro for nil quer dizer
que a função funcionou corretamente. Esta é uma forma de validação.
4 - A função Prepare começa com letra maiuscula pois é um método publico. Funções que começam com
letra minuscula são funcões privadas.
*/

func (u *User) validate() error {
	if err := u.Name.Validate(); err != nil {
		return err
	}

	if err := u.Email.Validate(); err != nil {
		return err
	}

	if err := u.Password.Validate(); err != nil {
		return err
	}

	return nil
}
