create table if not exists videos (
    video_id varchar(100) primary key,
    title    varchar(100) null,
    account_id varchar(100) not null,
    video_state varchar(100) not null,
    max_width integer not null,
    max_height integer not null,
    max_frame_rate integer not null,
    duration_milliseconds integer not null,
    filepath varchar(500) not null,
    create_timestamp timestamptz not null default now(),
    update_timestamp timestamptz not null default now()
);

create table if not exists playbacks (
    playback_id varchar(100) primary key,
    video_id varchar(100) not null,
    playback_policy varchar(100) not null,
    create_timestamp timestamptz not null default now()
);

create table if not exists accounts (
    account_id uuid primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) not null,
    hashed_password char(60) not null,
    create_timestamp timestamptz not null default now(),
    active BOOLEAN not null default True
);

ALTER TABLE accounts ADD CONSTRAINT accounts_uc_email UNIQUE (email);