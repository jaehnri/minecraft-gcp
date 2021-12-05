#!/usr/bin/python3

import os
from mcstatus import MinecraftServer

server = MinecraftServer.lookup("localhost:25565")

try:
    status = server.status()

    if status.players.online == 0:
        print('[SCRIPT]: Turning server off')
        os.system("/sbin/shutdown -h now")
    else:
        print('[SCRIPT]: There are players online. Not turning server off')
except:
    print('[SCRIPT]: Could not count players in server...')
    print('[SCRIPT]: Turning off anyway!!')
    os.system('shutdown -h')