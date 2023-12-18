-- +goose Up
create table dictionary (
    id serial primary key,
    word varchar(255) not null,
    origin text,
    fullWord text,
    definition text,
    etymology text,
    type varchar(255),
    sentence text,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

-- +goose Down


DROP TABLE dictionary;
```