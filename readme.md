an application that allows a user to log in and out, create a new account, and view their account information.
as well as CRUD operations on a list of TODO items.

the application is built with the following technologies:
- React
- TailwindCSS
- TypeScript
- PostgreSQL - user database
- MongoDB - TODO items database
- Clerk Authentication
- golang - backend
    - gin - web framework
        - logging
        - cors
        - rate limiting
        - jwt authentication
        - password hashing
        - database connection
        - database migrations
        - scheduling
    - sqlc - database ORM
    - goose - database migrations
    - bcrypt - password hashing
    - go-cron - scheduling
