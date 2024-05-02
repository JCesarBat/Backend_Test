CREATE TYPE "recordStatus" AS ENUM (
    'app.status.Active',
    'app.status.Inactive',
    'app.status.Blocked',
    ''
    );
CREATE TABLE "users" (
            "account_id"  BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
            "user_id" BIGSERIAL not null ,
            "create_date" timestamp DEFAULT (now()),
            "last_update" timestamp  DEFAULT CURRENT_TIMESTAMP,
            "record_status" "recordStatus" DEFAULT 'app.status.Active'::"recordStatus",
            active "recordStatus" DEFAULT 'app.status.Active'::"recordStatus",
            "identification" varchar(50) NULL,
            "password" varchar(50) NOT NULL ,
            "company_name" varchar(50),
            "first_name" varchar(50) NOT NULL ,
            "last_name" varchar(50) NOT NULL ,
            email varchar(50) unique NOT NULL ,
            phone varchar(50),
            emergencyphone varchar(50),
            "i18n" varchar(50),
            address varchar(200),
            "auth_menu" varchar NULL DEFAULT '[]'::text,
            "auth_keys" varchar NULL DEFAULT '[]'::text,
            "auth_groups" varchar NULL DEFAULT '[]'::text,
            "sites" varchar NULL DEFAULT '[]'::text,
            "clients" varchar NULL DEFAULT '[]'::text
        );
