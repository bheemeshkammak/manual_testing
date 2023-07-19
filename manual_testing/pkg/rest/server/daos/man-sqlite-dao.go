package daos

import (
	"database/sql"
	"errors"
	"github.com/bheemeshkammak/manual_testing/manual_testing/pkg/rest/server/daos/clients/sqls"
	"github.com/bheemeshkammak/manual_testing/manual_testing/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type ManDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateMen(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS men(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Age TEXT NOT NULL,
		Name TEXT NOT NULL,
		Verified INTEGER NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewManDao() (*ManDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateMen(sqlClient)
	if err != nil {
		return nil, err
	}
	return &ManDao{
		sqlClient,
	}, nil
}

func (manDao *ManDao) CreateMan(m *models.Man) (*models.Man, error) {
	insertQuery := "INSERT INTO men(Age, Name, Verified)values(?, ?, ?)"
	res, err := manDao.sqlClient.DB.Exec(insertQuery, m.Age, m.Name, m.Verified)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("man created")
	return m, nil
}

func (manDao *ManDao) UpdateMan(id int64, m *models.Man) (*models.Man, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	man, err := manDao.GetMan(id)
	if err != nil {
		return nil, err
	}
	if man == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE men SET Age = ?, Name = ?, Verified = ? WHERE Id = ?"
	res, err := manDao.sqlClient.DB.Exec(updateQuery, m.Age, m.Name, m.Verified, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("man updated")
	return m, nil
}

func (manDao *ManDao) DeleteMan(id int64) error {
	deleteQuery := "DELETE FROM men WHERE Id = ?"
	res, err := manDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("man deleted")
	return nil
}

func (manDao *ManDao) ListMen() ([]*models.Man, error) {
	selectQuery := "SELECT * FROM men"
	rows, err := manDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var men []*models.Man
	for rows.Next() {
		m := models.Man{}
		if err = rows.Scan(&m.Id, &m.Age, &m.Name, &m.Verified); err != nil {
			return nil, err
		}
		men = append(men, &m)
	}
	if men == nil {
		men = []*models.Man{}
	}

	log.Debugf("man listed")
	return men, nil
}

func (manDao *ManDao) GetMan(id int64) (*models.Man, error) {
	selectQuery := "SELECT * FROM men WHERE Id = ?"
	row := manDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Man{}
	if err := row.Scan(&m.Id, &m.Age, &m.Name, &m.Verified); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("man retrieved")
	return &m, nil
}
