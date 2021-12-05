	
#!/bin/bash
mount -o discard,defaults /dev/disk/by-id/google-minecraft-disk /home/minecraft
cd /home/minecraft
(crontab -l | echo "*/10 * * * * /home/minecraft/check-server-online.py")| crontab -
screen -d -m -S mcs java -Xms1G -Xmx7G -jar minecraft_server.1.18.jar nogui