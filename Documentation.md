# Unofficial official otwol documentation



## Connection sequence

1. Client connects to the server and sends an Auth_Req packet.
	- If they do not send an Auth_Req packet the server will close the connection.

2. Server responds with Auth_Req_ACK

3. Client responds with Token and Password.
	- If the Token is invalid the server will respond with Auth_No_Token_Invalid
	
4. Server responds with Auth_Yes packet and server <-> client conversation can start.




