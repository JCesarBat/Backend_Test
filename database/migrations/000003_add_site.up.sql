CREATE  TABLE  "sites" (
                            "account_id" BIGINT NOT NULL,
                           "site_id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
                           "site_name" varchar(50) NOT NULL,
                           "create_date" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           "last_update" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            "record_status" "recordStatus" DEFAULT 'app.status.Active'::"recordStatus",
                            active "recordStatus" DEFAULT 'app.status.Active'::"recordStatus",
                            "des" varchar(200),
                           "description" varchar,
                           "operate_by" varchar(50),
                           logo varchar,
                           "rules_documents" varchar DEFAULT '{}'::text NULL,
                           "services_amentiles" varchar DEFAULT '{}'::text NULL,
                           type varchar(50) DEFAULT 'Yard'::character varying,
                           email varchar(50) NULL,
                           phone varchar(50) NULL,
                           address varchar(200) NULL,
                           website varchar(200) NULL,
                           geolocation varchar(200) NULL
);

