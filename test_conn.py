import socket
import sys  

host = 'localhost'
port = 9988  # web

# Hardcoded auth token
Token = "AAAA-1234"
Token_Auth = "%OTWOL_password"

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((host , port))

# Send data to remote server
print('# Sending data to server')
request = b"Auth_Req"

try:
    s.sendall(request)
except socket.error:
    print ('Send failed')
    sys.exit()

# Receive data
print('# Receive data from server')
reply = s.recv(4096)

print(reply)

request = b"Token: AAAA-1234" 
#% (bytes(Token, encoding="UTF-8"), bytes(Token_Auth, encoding="UTF-8"))