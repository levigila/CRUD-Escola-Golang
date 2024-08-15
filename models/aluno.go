package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Aluno struct {
    id                   int
    nome                 string
    idade                int
    nota_primeiro_semestre float64
    nota_segundo_semestre  float64
    nome_professor        string
    numero_sala           int
}

// Função para buscar todos os alunos
func GetAlunos(db *sql.DB) ([]Aluno, error) {
    rows, err := db.Query("SELECT id, nome, idade, nota_primeiro_semestre, nota_segundo_semestre, nome_professor, numero_sala FROM alunos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var alunos []Aluno

    for rows.Next() {
        var aluno Aluno
        err = rows.Scan(&aluno.id, &aluno.nome, &aluno.idade, &aluno.nota_primeiro_semestre, &aluno.nota_segundo_semestre, &aluno.nome_professor, &aluno.numero_sala)
        if err != nil {
            return nil, err
        }
        alunos = append(alunos, aluno)
    }

    return alunos, nil
}

// Função para criar um novo aluno
func CreateAluno(db *sql.DB, aluno Aluno) error {
    sqlStatement := `
    INSERT INTO alunos (nome, idade, nota_primeiro_semestre, nota_segundo_semestre, nome_professor, numero_sala)
    VALUES ($1, $2, $3, $4, $5, $6)`
    _, err := db.Exec(sqlStatement, aluno.nome, aluno.idade, aluno.nota_primeiro_semestre, aluno.nota_segundo_semestre, aluno.nome_professor, aluno.numero_sala)
    if err != nil {
        return err
    }
    return nil
}

// Função para atualizar um aluno existente
func UpdateAluno(db *sql.DB, id string, aluno Aluno) error {
    sqlStatement := `
    UPDATE alunos
    SET nome = $2, idade = $3, nota_primeiro_semestre = $4, nota_segundo_semestre = $5, nome_professor = $6, numero_sala = $7
    WHERE id = $1`
    _, err := db.Exec(sqlStatement, id, aluno.nome, aluno.idade, aluno.nota_primeiro_semestre, aluno.nota_segundo_semestre, aluno.nome_professor, aluno.numero_sala)
    if err != nil {
        return err
    }
    return nil
}

// Função para deletar um aluno existente
func DeleteAluno(db *sql.DB, id string) error {
    sqlStatement := `
    DELETE FROM alunos
    WHERE id = $1`
    _, err := db.Exec(sqlStatement, id)
    if err != nil {
        return err
    }
    return nil
}
