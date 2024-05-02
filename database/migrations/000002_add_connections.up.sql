create table "connections" (
"id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
"connection_id" varchar(50) ,
"account_id" BIGINT not null ,
"user_id" BIGINT NOT NULL ,
"create_date" timestamp DEFAULT (now()),
"last_update" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
"record_status" "recordStatus" DEFAULT 'app.status.Active'::"recordStatus",
active "recordStatus" DEFAULT 'app.status.Active'::"recordStatus" NOT NULL ,
“connected” varchar(50),
“disconnected” varchar(50),
"user_data" varchar DEFAULT '{}'::text NULL,
"account_data" varchar DEFAULT '{}'::text NULL,
"auth_menu" varchar NULL DEFAULT '[]'::text,
"auth_keys" varchar NULL DEFAULT '[]'::text,
"auth_groups" varchar NULL DEFAULT '[]'::text,
"sites" varchar NULL DEFAULT '[]'::text,
"clients" varchar NULL DEFAULT '[]'::text
);