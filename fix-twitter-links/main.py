import discord
import urllib.parse

intents = discord.Intents.default()
intents.message_content = True

client = discord.Client(intents=intents)

@client.event
async def on_ready():
    print(f'''
    
______          _      _____ _              ___  ___          _    
|  ___|        | |    |  ___| |             |  \/  |         | |   
| |_ _   _  ___| | __ | |__ | | ___  _ __   | .  . |_   _ ___| | __
|  _| | | |/ __| |/ / |  __|| |/ _ \| '_ \  | |\/| | | | / __| |/ /
| | | |_| | (__|   <  | |___| | (_) | | | | | |  | | |_| \__ \   < 
\_|  \__,_|\___|_|\_\ \____/|_|\___/|_| |_| \_|  |_/\__,_|___/_|\_\
                                                                   
                Logged in {client.user}
    ''')

@client.event
async def on_message(message):
    if "https://twitter.com/" in message.content:
        newuri = urllib.parse.urlparse(message.content)._replace(netloc="vxtwitter.com")
        await message.reply(urllib.parse.urlunparse(newuri))
        
client.run("")
