insert into users (id, name, surname, patronymic)
values (1, 'Valera', 'Vedernikov', ''),
       (2, 'Sasha', 'Popov', ''),
       (3, 'Vasya', 'Smith', ''),
       (4, 'Test1', 'Test1', ''),
       (5, 'Test2', 'Test2', ''),
       (6, 'Test3', 'Test3', ''),
       (7, 'Test4', 'Test4', ''),
       (8, 'Test5', 'Test5', ''),
       (9, 'Test6', 'Test6', ''),
       (10, 'Test7', 'Test7', ''),
       (11, 'Test8', 'Test8', '');
insert into publishers (id, name, addr, code)
values (1, 'PITER', 'Moscow', 123),
       (2, 'PECHAT', 'London', 124),
       (3, 'TEST', 'TEST', 111);
insert into books (id, isbn, title, publisher_id, edition, description, author, year)
values (1, 11111123, 'War and Peace', 1, 1, 'book about life', 'L. N. Tolstoy', 1867),
       (2, 11111124, 'Kolobok', 2, 4, 'book about life', 'folklore', null),
       (3, 11111111, 'Holy Bible', 3, 2, 'book about life', 'God', null),
       (4, 11112123, 'Test1', 1, 1, 'book about life', 'Tester', 2022),
       (5, 11113124, 'Test2', 2, 1, 'book about life', 'Tester', 2022),
       (6, 11114111, 'Test3', 3, 1, 'book about life', 'Tester', 2022),
       (7, 11115123, 'Test4', 1, 1, 'book about life', 'Tester', 2022),
       (8, 11116124, 'Test5', 2, 1, 'book about life', 'Tester', 2022),
       (9, 11117111, 'Test6', 3, 1, 'book about life', 'Tester', 2022),
       (10, 11118123, 'Test7', 1, 1, 'book about life', 'Tester', 2022),
       (11, 11119124, 'Test8', 2, 1, 'book about life', 'Tester', 2022),
       (12, 11121111, 'Test9', 3, 1, 'book about life', 'Tester', 2022),
       (13, 11131123, 'Test10', 1, 1, 'book about life', 'Tester', 2022),
       (14, 11141124, 'Test11', 2, 1, 'book about life', 'Tester', 2022),
       (15, 11151111, 'Test12', 3, 1, 'book about life', 'Tester', 2022),
       (16, 11161123, 'Test13', 1, 1, 'book about life', 'Tester', 2022),
       (17, 11171124, 'Test14', 2, 1, 'book about life', 'Tester', 2022),
       (18, 11181111, 'Test15', 3, 1, 'book about life', 'Tester', 2022),
       (19, 11191123, 'Test16', 1, 1, 'book about life', 'Tester', 2022),
       (20, 11211124, 'Test17', 2, 1, 'book about life', 'Tester', 2022),
       (21, 11311111, 'Test18', 3, 1, 'book about life', 'Tester', 2022),
       (22, 11411123, 'Test19', 1, 1, 'book about life', 'Tester', 2022),
       (23, 11511124, 'Test20', 2, 1, 'book about life', 'Tester', 2022),
       (24, 11611111, 'Test21', 3, 1, 'book about life', 'Tester', 2022),
       (25, 11711123, 'Test22', 1, 1, 'book about life', 'Tester', 2022),
       (26, 11811124, 'Test23', 2, 1, 'book about life', 'Tester', 2022),
       (27, 11911111, 'Test24', 3, 1, 'book about life', 'Tester', 2022);
insert into book_entities (id, book_id, is_taken, taker_id, place)
values (1, 1, True, 1, 'shelf 1'),
       (2, 1, True, 2, 'shelf 1'),
       (3, 1, True, 3, 'shelf 1'),
       (4, 1, True, 4, 'shelf 1'),
       (5, 1, True, 5, 'shelf 1'),
       (6, 2, True, 2, 'shelf 1'),
       (7, 2, True, 3, 'shelf 1'),
       (8, 3, True, 1, 'shelf 1'),
       (9, 4, True, 2, 'shelf 1'),
       (10, 5, True, 3, 'shelf 1'),
       (11, 6, True, 1, 'shelf 1'),
       (12, 6, True, 3, 'shelf 1'),
       (13, 6, True, 2, 'shelf 1'),
       (14, 10, True, 3, 'shelf 1'),
       (15, 10, True, 1, 'shelf 1'),
       (16, 10, True, 2, 'shelf 1'),
       (17, 10, True, 4, 'shelf 1'),
       (18, 11, True, 1, 'shelf 1'),
       (19, 12, True, 3, 'shelf 1'),
       (20, 13, True, 2, 'shelf 1'),
       (21, 14, True, 3, 'shelf 1'),
       (22, 15, True, 1, 'shelf 1'),
       (23, 16, True, 2, 'shelf 1'),
       (24, 17, True, 3, 'shelf 1'),
       (25, 18, True, 1, 'shelf 1'),
       (26, 19, True, 3, 'shelf 1'),
       (27, 20, True, 2, 'shelf 1'),
       (28, 20, True, 3, 'shelf 1'),
       (29, 20, True, 4, 'shelf 1'),
       (30, 20, True, 5, 'shelf 1'),
       (31, 20, True, 6, 'shelf 1'),
       (32, 20, True, 7, 'shelf 1'),
       (33, 20, True, 8, 'shelf 1'),
       (34, 21, True, 2, 'shelf 1'),
       (35, 22, True, 3, 'shelf 1'),
       (36, 23, True, 1, 'shelf 1'),
       (37, 24, True, 2, 'shelf 1'),
       (38, 25, True, 3, 'shelf 1'),
       (39, 26, True, 1, 'shelf 1'),
       (40, 26, True, 3, 'shelf 1'),
       (41, 26, True, 2, 'shelf 1'),
       (42, 26, True, 5, 'shelf 1'),
       (43, 26, True, 6, 'shelf 1'),
       (44, 26, True, 8, 'shelf 1'),
       (45, 27, False, null, 'shelf 1'),
       (46, 7, False, null, 'shelf 1'),
       (47, 8, False, null, 'shelf 1'),
       (48, 9, False, null, 'shelf 1');

insert into book_entity_users (book_entity_id, user_id)
values (1, 2),
       (1, 3),
       (1, 4),
       (1, 5),
       (2, 7),
       (3, 8);
