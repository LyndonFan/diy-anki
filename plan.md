# DIY Anki

## Plan

-   app MVP

    -   can create, read, update, and delete cards/decks
        -   just prompt and answer
    -   review cards and update progress

-   Stack
    -   backend in Go
    -   frontend w/ JS (or Go + HTMX?)
    -   DB with sqlite
        -   works locally, no need to host extra stuff

## DB

-   deck
    -   deck id
    -   deck name
    -   created at?
    -   updated at?
-   card
    -   card id
    -   deck id
    -   created at
    -   updated at
    -   prompt
    -   answer
    -   info related to reviewing
        -   next to review
        -   last reviewed?
        -   review length?
        -   no. of times reviewed?
-   single review
    -   card id

## Backend routes

-   Operations:
    -   add new card
    -   delete card
    -   search / find cards
    -   edit card (content)
    -   review card
