create table docker_container_resources (
    container_id char(64) not null,
    resource_type varchar(15) not null,
    value numeric not null,
    created_at timestamptz NOT NULL DEFAULT NOW()
);

create table docker_container_names (
    container_id char(64) not null,
    name text
);

create unique index on docker_container_names (container_id, name);

create table docker_container_logs (
    id bigint primary key auto increment,
    service text,
    action text,
    data text,
    created_at timestamptz not null
)