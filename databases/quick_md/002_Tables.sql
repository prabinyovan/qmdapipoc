CREATE TABLE IF NOT EXISTS "User"."Users"
(
    user_id                 BIGSERIAL,
    user_global_key         UUID          NOT NULL DEFAULT gen_random_uuid(),
    first_name              VARCHAR(200)  NOT NULL,
    last_name               VARCHAR(200)  NOT NULL,
    password                VARCHAR(200)  NOT NULL,
    email                   VARCHAR(200)  NOT NULL,
	
	created_date_time_utc   timestamp without time zone default (now() at time zone 'utc'),
	updated_date_time_utc   timestamp without time zone default (now() at time zone 'utc'),
	
    CONSTRAINT      pk_user_users_id        PRIMARY KEY (user_id),
    CONSTRAINT      u_user_users_email      UNIQUE (email)
)