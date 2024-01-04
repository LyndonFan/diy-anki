-- Create the `decks` table
CREATE TABLE IF NOT EXISTS decks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create the `cards` table
CREATE TABLE IF NOT EXISTS cards (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    deck_id INTEGER NOT NULL,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    last_review_at TIMESTAMP NOT NULL,
    next_review_at TIMESTAMP NOT NULL,
    FOREIGN KEY (deck_id) REFERENCES decks (id) ON DELETE CASCADE
);

-- Create v_decks view

CREATE VIEW IF NOT EXISTS v_decks AS
WITH latest_decks AS (
    SELECT
        id,
        MAX(updated_at) AS updated_at
    FROM
        decks
    GROUP BY
        id
)
SELECT
    decks.*
FROM
    decks
INNER JOIN
    latest_decks
ON
    decks.id = latest_decks.id
    AND decks.updated_at = latest_decks.updated_at
;

-- Create v_cards view
CREATE VIEW IF NOT EXISTS v_cards AS
WITH latest_cards AS (
    SELECT
        id,
        deck_id,
        MAX(updated_at) AS updated_at
    FROM
        cards
    GROUP BY
        id, deck_id
)
SELECT
    cards.*
FROM
    cards
INNER JOIN
    latest_cards
ON
    cards.id = latest_cards.id
    AND cards.deck_id = latest_cards.deck_id
    AND cards.updated_at = latest_cards.updated_at
;