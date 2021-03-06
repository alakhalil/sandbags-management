DROP TABLE IF EXISTS "user" CASCADE;
DROP TABLE IF EXISTS "order" CASCADE;
DROP TABLE IF EXISTS "order_equipment";
DROP TABLE IF EXISTS "equipment" CASCADE;
DROP TABLE IF EXISTS "status" CASCADE;
DROP TABLE IF EXISTS "action_type" CASCADE;
DROP TABLE IF EXISTS "log" CASCADE;
DROP TABLE IF EXISTS "comment" CASCADE;
DROP TABLE IF EXISTS "hierarchy" CASCADE;
DROP TABLE IF EXISTS "user_order_permission";
DROP TABLE IF EXISTS "permission" CASCADE;
DROP TABLE IF EXISTS "priority" CASCADE;
DROP TABLE IF EXISTS "driver" CASCADE;
DROP TABLE IF EXISTS "branch" CASCADE;
DROP TABLE IF EXISTS "otp" CASCADE;

DROP SEQUENCE IF EXISTS order_number_id_seq;