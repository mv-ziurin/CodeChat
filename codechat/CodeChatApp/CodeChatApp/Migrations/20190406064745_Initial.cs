using System;
using Microsoft.EntityFrameworkCore.Migrations;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;

namespace CodeChatApp.Migrations
{
    public partial class Initial : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.EnsureSchema(
                name: "public");

            migrationBuilder.AlterDatabase()
                .Annotation("Npgsql:PostgresExtension:citext", "'citext', '', ''");

            migrationBuilder.CreateTable(
                name: "Chats",
                schema: "public",
                columns: table => new
                {
                    Id = table.Column<long>(nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.SerialColumn),
                    Name = table.Column<string>(nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Chats", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "received_email",
                schema: "public",
                columns: table => new
                {
                    email = table.Column<string>(type: "citext", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_received_email", x => x.email);
                });

            migrationBuilder.CreateTable(
                name: "users",
                schema: "public",
                columns: table => new
                {
                    username = table.Column<string>(type: "citext", nullable: false),
                    email = table.Column<string>(type: "citext", nullable: false),
                    hash = table.Column<string>(maxLength: 128, nullable: false),
                    PhotoUrl = table.Column<string>(nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_users", x => x.username);
                    table.UniqueConstraint("AK_users_email", x => x.email);
                });

            migrationBuilder.CreateTable(
                name: "CodeChats",
                schema: "public",
                columns: table => new
                {
                    Id = table.Column<long>(nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.SerialColumn),
                    Name = table.Column<string>(nullable: true),
                    ChatId = table.Column<long>(nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_CodeChats", x => x.Id);
                    table.ForeignKey(
                        name: "FK_CodeChats_Chats_ChatId",
                        column: x => x.ChatId,
                        principalSchema: "public",
                        principalTable: "Chats",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "event_log",
                schema: "public",
                columns: table => new
                {
                    id = table.Column<int>(nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.SerialColumn),
                    @event = table.Column<string>(name: "event", maxLength: 16, nullable: false),
                    username = table.Column<string>(type: "citext", nullable: true),
                    email = table.Column<string>(type: "citext", nullable: true),
                    ip = table.Column<string>(maxLength: 16, nullable: false),
                    user_agent = table.Column<string>(nullable: false),
                    event_time = table.Column<DateTime>(nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_event_log", x => x.id);
                    table.ForeignKey(
                        name: "event_log_email_fkey",
                        column: x => x.email,
                        principalSchema: "public",
                        principalTable: "users",
                        principalColumn: "email",
                        onDelete: ReferentialAction.Restrict);
                    table.ForeignKey(
                        name: "event_log_username_fkey",
                        column: x => x.username,
                        principalSchema: "public",
                        principalTable: "users",
                        principalColumn: "username",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateTable(
                name: "Messages",
                schema: "public",
                columns: table => new
                {
                    Id = table.Column<long>(nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.SerialColumn),
                    UserName = table.Column<string>(nullable: true),
                    ChatId = table.Column<long>(nullable: false),
                    Text = table.Column<string>(nullable: true),
                    Time = table.Column<DateTime>(nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Messages", x => x.Id);
                    table.ForeignKey(
                        name: "FK_Messages_Chats_ChatId",
                        column: x => x.ChatId,
                        principalSchema: "public",
                        principalTable: "Chats",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_Messages_users_UserName",
                        column: x => x.UserName,
                        principalSchema: "public",
                        principalTable: "users",
                        principalColumn: "username",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateTable(
                name: "UserChats",
                schema: "public",
                columns: table => new
                {
                    Id = table.Column<long>(nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.SerialColumn),
                    UserName = table.Column<string>(nullable: true),
                    ChatId = table.Column<long>(nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_UserChats", x => x.Id);
                    table.ForeignKey(
                        name: "FK_UserChats_Chats_ChatId",
                        column: x => x.ChatId,
                        principalSchema: "public",
                        principalTable: "Chats",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_UserChats_users_UserName",
                        column: x => x.UserName,
                        principalSchema: "public",
                        principalTable: "users",
                        principalColumn: "username",
                        onDelete: ReferentialAction.Restrict);
                });

            migrationBuilder.CreateIndex(
                name: "IX_CodeChats_ChatId",
                schema: "public",
                table: "CodeChats",
                column: "ChatId");

            migrationBuilder.CreateIndex(
                name: "IX_event_log_email",
                schema: "public",
                table: "event_log",
                column: "email");

            migrationBuilder.CreateIndex(
                name: "IX_event_log_username",
                schema: "public",
                table: "event_log",
                column: "username");

            migrationBuilder.CreateIndex(
                name: "IX_Messages_ChatId",
                schema: "public",
                table: "Messages",
                column: "ChatId");

            migrationBuilder.CreateIndex(
                name: "IX_Messages_UserName",
                schema: "public",
                table: "Messages",
                column: "UserName");

            migrationBuilder.CreateIndex(
                name: "IX_UserChats_ChatId",
                schema: "public",
                table: "UserChats",
                column: "ChatId");

            migrationBuilder.CreateIndex(
                name: "IX_UserChats_UserName",
                schema: "public",
                table: "UserChats",
                column: "UserName");

            migrationBuilder.CreateIndex(
                name: "email_idx",
                schema: "public",
                table: "users",
                column: "email",
                unique: true);
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "CodeChats",
                schema: "public");

            migrationBuilder.DropTable(
                name: "event_log",
                schema: "public");

            migrationBuilder.DropTable(
                name: "Messages",
                schema: "public");

            migrationBuilder.DropTable(
                name: "received_email",
                schema: "public");

            migrationBuilder.DropTable(
                name: "UserChats",
                schema: "public");

            migrationBuilder.DropTable(
                name: "Chats",
                schema: "public");

            migrationBuilder.DropTable(
                name: "users",
                schema: "public");
        }
    }
}
