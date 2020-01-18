      CREATE TABLE IF NOT EXISTS password_history(
         email CITEXT NOT NULL REFERENCES users (email),
         old_password VARCHAR(128) NOT NULL,
         old_date TIMESTAMP NOT NULL
      );


      CREATE TABLE public."CodeChats" (
          "Id" bigserial NOT NULL,
          "Name" text NULL,
          "ChatId" bigint NOT NULL,
          CONSTRAINT "PK_CodeChats" PRIMARY KEY ("Id"),
          CONSTRAINT "FK_CodeChats_Chats_ChatId" FOREIGN KEY ("ChatId") REFERENCES public."Chats" ("Id") ON DELETE CASCADE
      );
	  
	  CREATE TABLE public.event_log (
          id serial NOT NULL,
          event character varying(16) NOT NULL,
          username citext NULL,
          email citext NULL,
          ip character varying(16) NOT NULL,
          user_agent text NOT NULL,
          event_time timestamp without time zone NOT NULL,
          CONSTRAINT "PK_event_log" PRIMARY KEY (id),
          CONSTRAINT event_log_email_fkey FOREIGN KEY (email) REFERENCES public.users (email) ON DELETE RESTRICT,
          CONSTRAINT event_log_username_fkey FOREIGN KEY (username) REFERENCES public.users (username) ON DELETE RESTRICT
      );
	  
	  CREATE TABLE public."Messages" (
          "Id" bigserial NOT NULL,
          "UserName" citext NULL,
          "ChatId" bigint NOT NULL,
          "Text" text NULL,
          "Time" timestamp without time zone NOT NULL,
          CONSTRAINT "PK_Messages" PRIMARY KEY ("Id"),
          CONSTRAINT "FK_Messages_Chats_ChatId" FOREIGN KEY ("ChatId") REFERENCES public."Chats" ("Id") ON DELETE CASCADE,
          CONSTRAINT "FK_Messages_users_UserName" FOREIGN KEY ("UserName") REFERENCES public.users (username) ON DELETE RESTRICT
      );
	  
	  CREATE TABLE public."UserChats" (
          "Id" bigserial NOT NULL,
          "UserName" citext NULL,
          "ChatId" bigint NOT NULL,
          CONSTRAINT "PK_UserChats" PRIMARY KEY ("Id"),
          CONSTRAINT "FK_UserChats_Chats_ChatId" FOREIGN KEY ("ChatId") REFERENCES public."Chats" ("Id") ON DELETE CASCADE,
          CONSTRAINT "FK_UserChats_users_UserName" FOREIGN KEY ("UserName") REFERENCES public.users (username) ON DELETE RESTRICT
      );
	  
	  CREATE INDEX "IX_CodeChats_ChatId" ON public."CodeChats" ("ChatId");	  
	  CREATE INDEX "IX_event_log_email" ON public.event_log (email);	  
	  CREATE INDEX "IX_event_log_username" ON public.event_log (username);
	  CREATE INDEX "IX_Messages_ChatId" ON public."Messages" ("ChatId");
	  CREATE INDEX "IX_Messages_UserName" ON public."Messages" ("UserName");
	  CREATE INDEX "IX_UserChats_ChatId" ON public."UserChats" ("ChatId");
	  CREATE INDEX "IX_UserChats_UserName" ON public."UserChats" ("UserName");
	  CREATE UNIQUE INDEX email_idx ON public.users (email);
	  
	  INSERT INTO "__EFMigrationsHistory" ("MigrationId", "ProductVersion")
      VALUES ('20181114121413_Iniial', '2.1.4-rtm-31024');
