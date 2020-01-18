      CREATE TABLE "__EFMigrationsHistory" (
          "MigrationId" character varying(150) NOT NULL,
          "ProductVersion" character varying(32) NOT NULL,
          CONSTRAINT "PK___EFMigrationsHistory" PRIMARY KEY ("MigrationId")
      );
	  
	  SELECT EXISTS (SELECT 1 FROM pg_catalog.pg_class c JOIN pg_catalog.pg_namespace n ON n.oid=c.relnamespace WHERE c.relname='__EFMigrationsHistory');
	  
	  SELECT "MigrationId", "ProductVersion"
      FROM "__EFMigrationsHistory"
      ORDER BY "MigrationId";
	  
	  CREATE EXTENSION IF NOT EXISTS citext;
	  
	  CREATE TABLE public."Chats" (
          "Id" bigserial NOT NULL,
          "Name" text NULL,
          CONSTRAINT "PK_Chats" PRIMARY KEY ("Id")
      );
	  
	  CREATE TABLE public.received_email (
          email citext NOT NULL,
          CONSTRAINT "PK_received_email" PRIMARY KEY (email)
      );
	  
	  CREATE TABLE public.users (
          username citext NOT NULL,
          email citext NOT NULL,
          hash character varying(128) NOT NULL,
          "PhotoUrl" text NULL,
          CONSTRAINT "PK_users" PRIMARY KEY (username),
          CONSTRAINT "AK_users_email" UNIQUE (email)
      );
