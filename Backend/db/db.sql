-- Active: 1706712961341@@127.0.0.1@5432@main
CREATE DATABASE main;

\c main

CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY, first_name text not null, last_name text, phone_number text UNIQUE not NULL, username text UNIQUE not null, password text not null, image text UNIQUE, bio VARCHAR(100), created_at timestamp not null default now(), updated_at timestamp not null default now()
);

CREATE TABLE IF NOT EXISTS chats (
    id serial primary key , people int[] unique , created_at timestamp not null default now(), updated_at timestamp not null default now()
);

CREATE TABLE IF NOT EXISTS messages (
    id serial PRIMARY KEY, chat_id int NOT NULL, sender int NOT NULL, reciever int, content VARCHAR(300) NOT NULL, created_at timestamp not null default now(), updated_at timestamp not null default now()
);

CREATE TABLE IF NOT EXISTS contacts (
    user_id int not null, contact_id int not null, contact_name text not null, primary key (user_id, contact_id)
);