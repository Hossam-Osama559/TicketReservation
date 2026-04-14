package session

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/aarondl/authboss/v3"
	"github.com/google/uuid"
)

type MysqlSessionRepo struct {
	db *sql.DB
}

func NewMysqlSessionRepo(db *sql.DB) *MysqlSessionRepo {

	return &MysqlSessionRepo{db: db}
}

func (repo *MysqlSessionRepo) ReadState(req *http.Request) (authboss.ClientState, error) {
	cookie, err := req.Cookie("sessionid")
	if err != nil {
		// No session → user not logged in
		return nil, nil
	}

	sessionID := cookie.Value

	var pid string
	var userid int

	query := `SELECT pid,userid FROM sessions WHERE session_id = ?`
	err = repo.db.QueryRow(query, sessionID).Scan(&pid, &userid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Create Authboss session
	userSession := NewSession()
	userSession.Put("uid", pid)
	userSession.Put("userid", strconv.Itoa(userid))

	return userSession, nil
}

func (repo *MysqlSessionRepo) WriteState(w http.ResponseWriter, state authboss.ClientState, events []authboss.ClientStateEvent) error {
	// Get PID from session ,currently the email
	pid := events[0].Value

	// Generate session ID
	sessionID := uuid.New().String()

	// Store in DB

	query := `
	INSERT INTO sessions (session_id, pid, expire_at, userid)
	SELECT ?, ?, DATE_ADD(NOW(), INTERVAL 24 HOUR), id
	FROM users
	WHERE email = ?
	`

	_, err := repo.db.Exec(query, sessionID, pid, pid)
	if err != nil {
		return err
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
