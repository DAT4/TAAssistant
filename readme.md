# Registration bot

**Usecase:** Register or unregister students to grant them permission to the discord channels.

+ The bot sits in the registration room and receives messages. 
+ All messages will be deleted after 5 seconds
+ If the message has a valid ID format then the bot will check if a student with this ID exists in the database
+ If the ID is validated correcly and the ID is not already registered to another discordUser then the bot will register the discordUserID to the Student in the database, and change the discordUserNickName to `studentFirstName - studentID` and add the discordUser to the discordRole `Student` which will give the student permission to see the hidden channels in the server.

## Installation
### Configuration
To run the program you need to set a configuration file.
```json
#conf.json
{
    "token": "discord.token",
    "botStatus": "can be empty",
    "regChan": "the name of the channel you use for registrations",
    "studRole": "the name of the role registerd students have",
    "mongo" : {
        "uri": "mongodb+srv://username:password@example.com",
        "db": "the database",
        "col": "the collection in your database"
    }
}
```
You should not have another channel with the same name as your registration channel.

### Start the program

To start the program you simply run this command

```bash
./TAAssistant -c path-to-config-file.json
```

if you dont put a path the program will look for `conf.json` in the same directory as the program itself.

### As a service

You can set the program up in systemd with this service file

```
#/etc/systemd/system/TAAssistant.service
[Unit]
Description="TAAssistant for Discord server at UNI"

[Service]
Type=simple
User=YourUser
ExecStart=/dir/to/TAAssistant -c /dir/to/conf.json

[Install]
WantedBy=multi-user.target
```
