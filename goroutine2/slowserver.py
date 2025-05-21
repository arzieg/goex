#!/usr/bin/env python

# Includes
import getopt
import sys
import os.path
import subprocess
from http.server import HTTPServer
from http.server import BaseHTTPRequestHandler
import socketserver
import time

########  Predefined variables #########
helpstring = """Usage: {scriptname} args...
    Where args are:
        -h, --help
            Show help
        -p PORTNUMBER
            Port number to run on
        -d delay-in-seconds
            How long to wait before responding
"""

helpstring = helpstring.format(scriptname=sys.argv[0])

def beSlow(seconds):
    time.sleep(float(seconds))


########  Functions and classes #########
class SlowserverRequestHandler(BaseHTTPRequestHandler):
    def do_GET(s):
        if s.path == "/slow":
            # Check status
            # Assume fail
            code = 200
            status = ""

            # Be slow for a while
            beSlow(seconds)

            s.send_response(200)
            s.send_header("Content-type", "text/html")
            s.end_headers()
            s.wfile.write(b"I'm a slow response LOL\n")

        else:
            s.send_response(200)
            s.send_header("Content-type", "text/html")
            s.end_headers()
            s.wfile.write(b"slowserver - reporting for duty. Slowly...\n")


# Parse args
try:
    options, remainder = getopt.getopt(sys.argv[1:], "hp:d:", ['help'])
except:
    print("Invalid args. Use -h or --help for help.")
    raise
    sys.exit(1)

HTTPPORT = 8000
for opt, arg in options:                                                  
    if opt in ('-h', '--help'):                                           
        print(helpstring)                                                 
        sys.exit(0)                                                       
    elif opt in ('-p'):                                                   
        HTTPPORT = int(arg)                                               
    elif opt in ('-d'):                                                
        seconds = arg                                                  
                                                                       
# Start HTTP service                                                   
server_class=HTTPServer                                                
handler_class=SlowserverRequestHandler               
server_address = ('', HTTPPORT)                      
httpd = server_class(server_address, handler_class)
try:                                               
    httpd.serve_forever()                          
except KeyboardInterrupt:                          
    pass                                           
httpd.server_close()
