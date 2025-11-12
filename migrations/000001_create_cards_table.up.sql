CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(16) NOT NULL CHECK(char_length(user_name) >= 3),
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE decks (
    deck_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE cards (
    card_id SERIAL NOT NULL,
    deck_id INTEGER REFERENCES decks(deck_id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
    text1 VARCHAR(64) NOT NULL,
    text2 VARCHAR(64) NOT NULL,
    new_time TIMESTAMP DEFAULT now(),
    created_at IMESTAMP DEFAULT now()
);
