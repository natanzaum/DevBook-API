insert into usuarios(nome, nick, email, senha)
values
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$tB59fAVhpTyhp5xcCmhuJuFmhS7XvUfrE8cWqKE8vDKrVn7mxMhqO"),
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$tB59fAVhpTyhp5xcCmhuJuFmhS7XvUfrE8cWqKE8vDKrVn7mxMhqO"),
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$tB59fAVhpTyhp5xcCmhuJuFmhS7XvUfrE8cWqKE8vDKrVn7mxMhqO"),
("Usuario 4", "usuario_4", "usuario4@gmail.com", "$2a$10$tB59fAVhpTyhp5xcCmhuJuFmhS7XvUfrE8cWqKE8vDKrVn7mxMhqO"),
("Usuario 5", "usuario_5", "usuario5@gmail.com", "$2a$10$tB59fAVhpTyhp5xcCmhuJuFmhS7XvUfrE8cWqKE8vDKrVn7mxMhqO");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(1, 3),
(1, 4),
(3, 1);

insert into publicacoes(titulo, conteudo, autor_id) 
values
("Publicação do usuário 1", "Esta é a publicação do usuário 1", 1),
("Publicação do usuário 2", "Esta é a publicação do usuário 2", 2),
("Publicação do usuário 3", "Esta é a publicação do usuário 3", 3),
("Publicação do usuário 4", "Esta é a publicação do usuário 4", 4),
("Publicação do usuário 5", "Esta é a publicação do usuário 5", 5);