import discord

bot = discord.Bot()


@bot.event
async def on_ready():
    print("[*] Developer Badge Unlocker online")


@bot.slash_command()
async def bagde(ctx):
    await ctx.respond(
        "Check here to see if your badge is available https://discord.com/developers/active-developer"
    )


bot.run("")
