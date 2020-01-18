using CodeChatApp.Database.Models;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Database
{
    public class CodeChatContext : DbContext
    {
        public CodeChatContext(DbContextOptions<CodeChatContext> options) : base(options) { }

        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            if (!optionsBuilder.IsConfigured)
            {
                optionsBuilder.UseNpgsql(@"Host=localhost;Port=5432;Database=codechat;Username=postgres;Password=postgres");
            }
        }

        public DbSet<EventLog> EventLog { get; set; }
        public DbSet<ReceivedEmail> ReceivedEmail { get; set; }
        public DbSet<Users> Users { get; set; }
        public DbSet<UserChat> UserChats { get; set; }
        public DbSet<CodeChat> CodeChats { get; set; }
        public DbSet<Message> Messages { get; set; }
        public DbSet<Chat> Chats { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.HasDefaultSchema("public");

            // -------------------------- GENERATED -------------------------------
            modelBuilder.HasPostgresExtension("citext");

            modelBuilder.Entity<EventLog>(entity =>
            {
                entity.ToTable("event_log");

                entity.Property(e => e.Id).HasColumnName("id");

                entity.Property(e => e.Email)
                    .HasColumnName("email")
                    .HasColumnType("citext");

                entity.Property(e => e.Event)
                    .IsRequired()
                    .HasColumnName("event")
                    .HasMaxLength(16);

                entity.Property(e => e.EventTime).HasColumnName("event_time");

                entity.Property(e => e.Ip)
                    .IsRequired()
                    .HasColumnName("ip")
                    .HasMaxLength(16);

                entity.Property(e => e.UserAgent)
                    .IsRequired()
                    .HasColumnName("user_agent");

                entity.Property(e => e.Username)
                    .HasColumnName("username")
                    .HasColumnType("citext");

                entity.HasOne(d => d.EmailNavigation)
                    .WithMany(p => p.EventLogEmailNavigation)
                    .HasPrincipalKey(p => p.Email)
                    .HasForeignKey(d => d.Email)
                    .HasConstraintName("event_log_email_fkey");

                entity.HasOne(d => d.UsernameNavigation)
                    .WithMany(p => p.EventLogUsernameNavigation)
                    .HasForeignKey(d => d.Username)
                    .HasConstraintName("event_log_username_fkey");
            });

            modelBuilder.Entity<ReceivedEmail>(entity =>
            {
                entity.HasKey(e => e.Email);

                entity.ToTable("received_email");

                entity.Property(e => e.Email)
                    .HasColumnName("email")
                    .HasColumnType("citext")
                    .ValueGeneratedNever();
            });

            modelBuilder.Entity<Users>(entity =>
            {
                entity.HasKey(e => e.Username);

                entity.ToTable("users");

                entity.HasIndex(e => e.Email)
                    .HasName("email_idx")
                    .IsUnique();

                entity.Property(e => e.Username)
                    .HasColumnName("username")
                    .HasColumnType("citext")
                    .ValueGeneratedNever();

                entity.Property(e => e.Email)
                    .IsRequired()
                    .HasColumnName("email")
                    .HasColumnType("citext");

                entity.Property(e => e.Hash)
                    .IsRequired()
                    .HasColumnName("hash")
                    .HasMaxLength(128);
            });
            // -------------------------- GENERATED -------------------------------

            // primary keys
            modelBuilder.Entity<Users>().HasKey(p => p.Username);
            modelBuilder.Entity<Chat>().HasKey(p => p.Id);
            modelBuilder.Entity<UserChat>().HasKey(p => p.Id);
            modelBuilder.Entity<Message>().HasKey(p => p.Id);
            modelBuilder.Entity<CodeChat>().HasKey(p => p.Id);

            // one-to-many

            modelBuilder.Entity<Chat>()
                .HasMany(p => p.CodeChats)
                .WithOne(s => s.Chat);

            modelBuilder.Entity<Chat>()
                .HasMany(p => p.Messages)
                .WithOne(s => s.Chat);

            modelBuilder.Entity<Users>()
                .HasMany(p => p.Messages)
                .WithOne(s => s.User);

            // many-to-many

            modelBuilder.Entity<UserChat>()
            .HasOne(sc => sc.User)
            .WithMany(s => s.UserChats)
            .HasForeignKey(sc => sc.UserName);

            modelBuilder.Entity<UserChat>()
                .HasOne(sc => sc.Chat)
                .WithMany(c => c.UserChats)
                .HasForeignKey(sc => sc.ChatId);


        }
    }
}
