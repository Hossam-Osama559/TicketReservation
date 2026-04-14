package ticket

import (
	"database/sql"
)

type MysqlTicketRepo struct {
	db *sql.DB
}

/*
use cases for the ticket


1---create ticket but not paid yet with it will create session in stripe

2---after paid, and stripe sending the webhook , make the ticket paid


3---user can retreive tickets
*/

func NewMysqlTicketRepo(db *sql.DB) *MysqlTicketRepo {

	return &MysqlTicketRepo{db: db}
}

func (repo *MysqlTicketRepo) Create(t *Ticket) error {
	query := `
		INSERT INTO tickets (origin, arrival, price,  user_id,scanned)
		VALUES (?, ?, ?, ?,?)
	`

	result, err := repo.db.Exec(query,
		t.Origin,
		t.Arrival,
		t.Price,
		t.UserId,
		false,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	t.Id = int(id)
	return nil
}

func (repo *MysqlTicketRepo) GetById(id int) (*Ticket, error) {
	query := `
		SELECT id, origin, arrival, price,user_id,scanned
		FROM tickets
		WHERE id = ?
	`

	row := repo.db.QueryRow(query, id)

	t := &Ticket{}
	err := row.Scan(
		&t.Id,
		&t.Origin,
		&t.Arrival,
		&t.Price,

		&t.UserId,

		&t.Scanned,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return t, nil
}

func (repo *MysqlTicketRepo) Update(t *Ticket) error {
	query := `
		UPDATE tickets
		SET origin = ?, arrival = ?, price = ?,  user_id = ?,scanned= ?
		WHERE id = ?
	`

	_, err := repo.db.Exec(query,
		t.Origin,
		t.Arrival,
		t.Price,

		t.UserId,

		t.Scanned,

		t.Id,
	)

	return err
}

func (repo *MysqlTicketRepo) MarkAsPaid(paymentId string) error {
	query := `
		UPDATE tickets
		SET paid = true
		WHERE payment_id = ?
	`

	_, err := repo.db.Exec(query, paymentId)
	return err
}

func (repo *MysqlTicketRepo) Delete(id int) error {
	query := `DELETE FROM tickets WHERE id = ?`
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *MysqlTicketRepo) GetByUser(userId int) ([]Ticket, error) {
	query := `
		SELECT id, origin, arrival, price,  user_id
		FROM tickets t inner join payments p on p.ticket_id=t.id
		WHERE t.user_id = ? and p.completed=1
	`

	rows, err := repo.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []Ticket

	for rows.Next() {
		var t Ticket
		err := rows.Scan(
			&t.Id,
			&t.Origin,
			&t.Arrival,
			&t.Price,

			&t.UserId,
		)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, t)
	}

	return tickets, nil
}
