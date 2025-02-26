CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) NOT NULL,
        full_name VARCHAR(255) NOT NULL,
        is_active BOOLEAN NOT NULL DEFAULT TRUE,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
    );

CREATE TABLE
    notes (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        keep_in_inbox BOOLEAN NOT NULL DEFAULT TRUE,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE
    boxes (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL,
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW (),
        FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE
    notes_boxes (
        note_id INTEGER NOT NULL,
        box_id INTEGER NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        PRIMARY KEY (note_id, box_id),
        FOREIGN KEY (note_id) REFERENCES notes (id),
        FOREIGN KEY (box_id) REFERENCES boxes (id)
    );