package service

import (
	"context"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type CharacterService struct {
	db *database.Database
}

func NewCharacterService(db *database.Database) *CharacterService {
	return &CharacterService{db: db}
}

func (s *CharacterService) GetAllCharacters() ([]models.AccidentCharacter, error) {
	var characters []models.AccidentCharacter

	query := `SELECT id, name, id_specialization FROM accident_character`
	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var character models.AccidentCharacter
		err := rows.Scan(&character.ID, &character.Name, &character.Specialization.ID)
		if err != nil {
			return nil, err
		}

		characters = append(characters, character)
	}

	return characters, nil
}
