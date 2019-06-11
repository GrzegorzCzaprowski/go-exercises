CREATE TABLE public.users
(
    id bigserial NOT NULL PRIMARY KEY,
    email text NOT NULL UNIQUE,
    password bytea,
    created_at timestamp
    with time zone DEFAULT timezone
    ('utc'::text, now
    ()) NOT NULL
    );
