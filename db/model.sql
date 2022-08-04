CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

drop table if exists news CASCADE;
drop table if exists tags CASCADE;
drop table if exists news_tags CASCADE;
drop table if exists files CASCADE;

create table public.news (
    id uuid DEFAULT public.uuid_generate_v4() primary key,
    title text not null,
    author text not null,
    active boolean not null,
    activeFrom timestamp without time zone not null,
    text text not null,
    textJSON text not null,
    userId uuid not null,
    isImportant boolean not null
);

create table public.tags (
    id uuid DEFAULT public.uuid_generate_v4() primary key,
    name text,
    unique(name)
);

create table public.news_tags (
    news_id uuid references news (id),
    tag_id uuid references tags (id)
);

create table public.files (
    id uuid DEFAULT public.uuid_generate_v4() primary key,
    name text not null,
    ext text not null,
    base64 text not null,
    dateCreate timestamp without time zone not null,
    userId uuid not null,
    newsId uuid references news (id)
);

insert into news (id, title,author,active,activeFrom,text,textJSON,userId,isImportant) values 
('da25b569-467d-4ddb-92b6-a7d20981eb6b', 'Hello',  'pupa',true, timestamp '2022-07-25 12:00:00','Hello i am here',      '{"text":"Hello i am here"}',      '4cae26ac-d681-4364-88fb-deeb135e6ca2',false),
('8bb33810-ab41-44d1-8939-e8091d9ecded', 'Alive!', 'ilya',false,timestamp '2022-08-04 08:00:00','I am alive!',          '{"text":"I am alive!"}',          'fe195582-4581-4435-8624-64ba8299f629',true),
('fbfd4d1a-ef35-40f2-8a67-3723eaf7b46e', 'Taxes',  'vova',true, timestamp '2008-05-05 04:00:00','PAY TAXES!!!11!',      '{"text":"PAY TAXES!!!11!"}',      '8c7049b6-bb9e-4b20-a821-b717d2773848',true),
('0a87ab57-b300-4b76-bb6f-a9e340c52d5f', 'Arch',   'ilya',true, timestamp '2012-02-23 15:00:00','Clean Arch is garbage','{"text":"Clean Arch is garbage"}','f43b34f8-04c2-4fcf-ac43-283f149d3387',false);

insert into tags (id, name) values
('513b236a-a32b-4303-9e75-357e10de8d8b', 'Fun'),
('582167e7-0f77-4719-96fa-c21a7674e95d', 'Hard'),
('c2a1bda6-d222-4b1c-87a3-99f511b47652', 'Truth'),
('8756eb2f-028c-480c-8253-2f1dc814f732', 'No choice');


insert into news_tags (news_id, tag_id) values 
('da25b569-467d-4ddb-92b6-a7d20981eb6b', 'c2a1bda6-d222-4b1c-87a3-99f511b47652'),
('8bb33810-ab41-44d1-8939-e8091d9ecded', 'c2a1bda6-d222-4b1c-87a3-99f511b47652'),
('8bb33810-ab41-44d1-8939-e8091d9ecded', '582167e7-0f77-4719-96fa-c21a7674e95d'),
('fbfd4d1a-ef35-40f2-8a67-3723eaf7b46e', '8756eb2f-028c-480c-8253-2f1dc814f732');

insert into files (id,name,ext,base64,dateCreate,userId, newsId) values
('d5412101-4421-4098-857d-60ef15cd090f','Doc42',              'txt', 'base01',timestamp '2012-02-23 15:00:00','4cae26ac-d681-4364-88fb-deeb135e6ca2','da25b569-467d-4ddb-92b6-a7d20981eb6b'),
('2ea42060-5dff-47d1-8aa7-c5c2ec7aaea6','Raschetnaya_vypiska','docx','base02',timestamp '2012-02-23 15:00:00','8c7049b6-bb9e-4b20-a821-b717d2773848','fbfd4d1a-ef35-40f2-8a67-3723eaf7b46e'),
('6e28c91f-9006-4adb-a9b9-d7464127a078','Annot_',             'pptx','base03',timestamp '2012-02-23 15:00:00','8c7049b6-bb9e-4b20-a821-b717d2773848','fbfd4d1a-ef35-40f2-8a67-3723eaf7b46e'),
('67ff03d1-afbe-48b6-8e9a-5ca283471547','Screen125',          'png', 'base03',timestamp '2012-02-23 15:00:00','f43b34f8-04c2-4fcf-ac43-283f149d3387','0a87ab57-b300-4b76-bb6f-a9e340c52d5f'),
('ced8ba16-03f3-4a04-bfa7-6425d1e12cfa','main',               'cpp', 'base01',timestamp '2012-02-23 15:00:00','f43b34f8-04c2-4fcf-ac43-283f149d3387','0a87ab57-b300-4b76-bb6f-a9e340c52d5f');