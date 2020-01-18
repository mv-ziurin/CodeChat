using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.EntityFrameworkCore;
using CodeChatApp.Database;
using CodeChatApp.Hubs;
using CodeChatApp.Controllers.JsonRPC;
using CodeChatApp.Services;
using Microsoft.Extensions.Logging;

namespace CodeChatApp
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        public void ConfigureServices(IServiceCollection services)
        {
            var connectionString = Configuration.GetConnectionString("CodeChat");

            services.AddEntityFrameworkNpgsql().AddDbContext<CodeChatContext>(options => options.UseNpgsql(connectionString));
            services.AddScoped<IRepository, Repository>();
            services.AddScoped<IValidator, Validator>();
            services.AddScoped<IImplementor, Implementor>();

            services.AddCors(o => o.AddPolicy("CorsPolicy", builder =>
            {
                builder.AllowAnyOrigin()
                       .AllowCredentials()
                       .AllowAnyMethod()
                       .AllowAnyHeader();
            }));
            services.AddSignalR();
            services.AddMvc().SetCompatibilityVersion(CompatibilityVersion.Version_2_1);
        }

        public void Configure(IApplicationBuilder app, IHostingEnvironment env, ILoggerFactory loggerFactory)
        {
            loggerFactory.AddConsole(LogLevel.Trace);
            loggerFactory.AddDebug(LogLevel.Trace);

            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseCors("CorsPolicy");
            app.UseSignalR(routes =>
            {
                routes.MapHub<CodeChatHub>("/mainChat");
            });
            app.UseMvc();
        }
    }
}
