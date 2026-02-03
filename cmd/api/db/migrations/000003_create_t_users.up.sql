CREATE TABLE IF NOT EXISTS hr.users(
    user_id serial,
    user_name varchar(15),
    user_email varchar(80),
    user_password varchar(125),
    user_handphone varchar(15),
    created_on timestamp,
    constraint pk_user_id primary key(user_id),
    constraint uq_user_email unique (user_email),
    constraint uq_user_handphone unique (user_handphone)
 );


CREATE TABLE IF NOT EXISTS hr.roles(
    role_id serial,
    role_name varchar(35),
    constraint pk_role_id primary key(role_id),
    constraint uq_role_name unique (role_name)
 );


CREATE TABLE IF NOT EXISTS hr.user_roles(
    user_id integer,
    role_id integer,
    constraint fk_users_user_id foreign key(user_id) references hr.users(user_id),
    constraint fk_roles_role_id foreign key (role_id) references hr.roles(role_id)
 );