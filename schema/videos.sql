create table if not exists videos (
    video_id varchar(100) primary key,
    title    varchar(100) null,
    account_id varchar(100) not null,
    video_state varchar(100) not null,
    filepath varchar(500) not null,
    create_timestamp timestamptz not null default now(),
    update_timestamp timestamptz not null default now() 
);