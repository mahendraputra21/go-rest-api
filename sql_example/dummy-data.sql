--incase goose db version not created
--INSERT INTO public.goose_db_version(id, version_id, is_applied)VALUES (1, 0, true);

insert into public.brand(id, brand_name) VALUES (1, 'Cosmos');
insert into public.brand(id, brand_name) VALUES (2, 'Miyako');
insert into public.brand(id, brand_name) VALUES (3, 'Polytron');
insert into public.brand(id, brand_name) VALUES (4, 'Sanken');
insert into public.brand(id, brand_name) VALUES (5, 'Maspion');
insert into public.brand(id, brand_name) VALUES (6, 'Samsung');


insert into public.product(product_name, brand_id) VALUES ('Dispenser CWD-1060', 1);
insert into public.product(product_name, brand_id) VALUES ('Slow Cooker Miyako SC630', 2);
insert into public.product(product_name, brand_id) VALUES ('Smart Android TV 43" PLD 43AG9953', 3);
insert into public.product(product_name, brand_id) VALUES ('Mesin Cuci QW-S125DDSL', 4);
insert into public.product(product_name, brand_id) VALUES ('Kompor 2 TK 709 C (TF)', 5);
insert into public.product(product_name, brand_id) VALUES ('Teflon Frypan Set 2in1', 5);
insert into public.product(product_name, brand_id) VALUES ('Blender CB282G', 1);
insert into public.product(product_name, brand_id) VALUES ('Duo Galon HWD-z88', 4);
insert into public.product(product_name, brand_id) VALUES ('Gas Water Heater/Pemanas', 5);
insert into public.product(product_name, brand_id) VALUES ('Gas WASSER WH-506A', 3);
