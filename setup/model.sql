drop table if exists plan cascade;


create table plan(
	plan_id serial not null primary key,
	status bool not null default false,
	description text null
);

insert into plan(status, description) values(true, 'The plane landed safely in Uzbekistan from China');
insert into plan(status, description) values(true, 'The plane landed safely in Uzbekistan from USA');
insert into plan(status, description) values(false,' '), 
(false, ' '), 
(true, 'The plane landed safely in Russia from Uzbekistan');

	
	
