using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Web;
using ChaosTasks.Data;
using ChaosTasks.Services;

var builder = WebApplication.CreateBuilder(new WebApplicationOptions {
  Args = args,
  WebRootPath = "/var/www/wwwroot"
});

System.Console.WriteLine($"WebRoot: {builder.Environment.WebRootPath}");
builder.Environment.WebRootPath = "";

// Add services to the container.
builder.Services.AddRazorPages();
builder.Services.AddServerSideBlazor();
builder.Services.AddSingleton<WeatherForecastService>();
builder.Services.AddSingleton<TodoFileService>();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Error");
}

app.UseStaticFiles();

app.UseRouting();

app.MapBlazorHub();
app.MapFallbackToPage("/_Host");

app.Run();
