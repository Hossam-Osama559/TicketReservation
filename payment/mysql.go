package payment

import "database/sql"

type MysqlPaymentRepo struct {
	DB *sql.DB
}

func NewMysqlPaymentRepo(db *sql.DB) *MysqlPaymentRepo {

	return &MysqlPaymentRepo{DB: db}
}

func (repo *MysqlPaymentRepo) Create(instance Payment) error {
	query := `
		INSERT INTO payments (session_id, ticket_id, completed)
		VALUES (?, ?, ?)
	`
	_, err := repo.DB.Exec(query,
		instance.SessionId,
		instance.TicketId,
		instance.Completed,
	)
	return err
}

func (repo *MysqlPaymentRepo) Get(sessionid string) (*Payment, error) {
	query := `
		SELECT session_id, ticket_id, completed
		FROM payments
		WHERE session_id = ?
	`

	p := NewPayment()
	err := repo.DB.QueryRow(query, sessionid).Scan(
		&p.SessionId,
		&p.TicketId,
		&p.Completed,
	)

	return p, err
}

func (repo *MysqlPaymentRepo) Update(sessionid string, newinstance Payment) error {
	query := `
		UPDATE payments
		SET ticket_id = ?, completed = ?
		WHERE session_id = ?
	`

	_, err := repo.DB.Exec(query,
		newinstance.TicketId,
		newinstance.Completed,
		sessionid,
	)

	return err
}

func (repo *MysqlPaymentRepo) MarkCompleted(sessionid string) error {
	query := `
		UPDATE payments
		SET completed = TRUE
		WHERE session_id = ?
	`
	_, err := repo.DB.Exec(query, sessionid)
	return err
}

func (repo *MysqlPaymentRepo) Delete(sessionid string) error {
	query := `
		DELETE FROM payments
		WHERE session_id = ?
	`
	_, err := repo.DB.Exec(query, sessionid)
	return err
}
