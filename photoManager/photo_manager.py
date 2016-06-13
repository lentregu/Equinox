mport serial
import subprocess
import sys
import os
import time
from datetime import datetime
import datetime
import requests


ser=serial.Serial('/dev/ttyACM0',9600)
while 1:
        cad=ser.readline().rstrip()
        print(str(cad))
        utc = datetime.datetime.utcnow()
        s = utc.strftime("%Y-%m-%d-%H-%M%SZ")
        filename = 'IMG_'+ s + '.jpg'
#       filename = "26.jpg"
        if cad is "1":
                print("OK")
                os.system("raspistill -o "+ filename )
                scp="scp "+filename + " pi@amazon:/usr/share/nginx/html/"
                print(scp)
                os.system(scp)
D
                headers = c
                curl -H "Content-Type: application/json" -X POST -d '{"url": "http://52.37.253.204/20.jpg", "faceListID": "autorizados"}' http://localhost:8080/find

        else:
                print("no match")
                print(cad)
        print("------------------")
