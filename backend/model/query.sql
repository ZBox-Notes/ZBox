-- name: ListUsers :many
SELECT * FROM users;

-- name: UserById :one
SELECT * FROM users 
WHERE id = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, full_name)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET email = $2, full_name = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListNotes :many
SELECT * FROM notes;

-- name: NoteById :one
SELECT * FROM notes 
WHERE id = $1
LIMIT 1;

-- name: CreateNote :one
INSERT INTO notes (user_id, title, content)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateNote :one
UPDATE notes
SET user_id = $2, title = $3, content = $4
WHERE id = $1
RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;

-- name: ListBoxes :many
SELECT * FROM boxes;

-- name: BoxById :one
SELECT * FROM boxes 
WHERE id = $1
LIMIT 1;

-- name: CreateBox :one
INSERT INTO boxes (user_id, name)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateBox :one
UPDATE boxes
SET user_id = $2, name = $3
WHERE id = $1
RETURNING *;

-- name: DeleteBox :exec
DELETE FROM boxes
WHERE id = $1;

-- name: ListNotesBoxes :many
SELECT * FROM notes_boxes;

-- name: NotesBoxesByNoteId :many
SELECT * FROM notes_boxes 
WHERE note_id = $1;

-- name: NotesBoxesByBoxId :many
SELECT * FROM notes_boxes 
WHERE box_id = $1;

-- name: CreateNotesBox :one
INSERT INTO notes_boxes (note_id, box_id)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteNotesBox :exec
DELETE FROM notes_boxes
WHERE note_id = $1 AND box_id = $2;