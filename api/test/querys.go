package test


var InicialTableUsersQuery string = `
CREATE TABLE users (
    id serial NOT NULL,
    full_name VARCHAR(150) NOT NULL,
    password VARCHAR(256) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    picture VARCHAR(256) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT pk_users PRIMARY KEY(id)
);
`
var InicialTableTasksQuery string = `
CREATE TABLE tasks (
    id serial NOT NULL,
    user_id int NOT NULL,
    description TEXT NOT NULL,
	fulfilled BOOL DEFAULT false,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT pk_tasks PRIMARY KEY(id),
    CONSTRAINT fk_tasks_users FOREIGN KEY(user_id) REFERENCES users(id)
);
`

var DestroyTables string = `
	DROP TABLE tasks;
	DROP TABLE users;
`

var GetAllUsersQuery string = `SELECT * FROM users`
var GetAllTasksQuery string = `SELECT * FROM tasks`

var FulfillTask string = `UPDATE tasks SET fulfilled = true WHERE id = $1`