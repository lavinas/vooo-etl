use vooo_migration;

-- insert into aggregator_ref
insert into aggregator_ref (id) values (11);


-- insert into client_ref
insert into client_ref (id) values (1507); -- epag
insert into client_ref (id) values (2210); -- tradio


insert into job (id, name, type, action, base, object, field, last, status) values (100, 'cliente reference', 'table', 'reference', 'vooo_migration', 'client_ref', 'id', 0, 'pending');
insert into job (id, name, type, action, base, object, field, last, status) values (200, 'cliente copy', 'table', 'copy', 'vooo_prod_backend', 'client', 'id', 1076, 'pending');

commit;