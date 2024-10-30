-- name: CreateUserRecord :one
INSERT INTO user_profiles(email, first_name, last_name, otp_secret)
VALUES ($1, $2, $3, $4)
RETURNING id;
-- name: CreateEmailRecord :exec
INSERT INTO email_account(user_id, password)
VALUES($1, $2);
-- name: CreateUserSession :one
INSERT INTO user_sessions(session_id, user_id, user_agent)
VALUES($1, $2, $3)
RETURNING id;
-- name: GetUserDataFromSessionID :one
SELECT s.session_id,
    s.user_agent,
    u.id,
    u.email,
    u.first_name
FROM user_sessions s
    JOIN user_profiles u ON s.user_id = u.id
WHERE s.is_valid = TRUE;
-- name: ChangeSessionStatus :exec
UPDATE user_sessions
SET is_valid = $2
WHERE user_id = $1;
-- name: GetUserByID :one
SELECT u.*
from user_profiles u
WHERE u.id = $1;
-- name: GetUserByEmail :one
SELECT u.id,
    u.email,
    e.password
FROM user_profiles u
    JOIN email_account e on u.id = e.user_id
WHERE u.email = $1;