## MCCAP
Merry Christmas! Chicken Art Protcol's implementation example both side of client & server.

## Protocol
All messages are encoded in UTF-8 & send via TCP.
No encryption, no compression, no security.
If not specified, all messages are ended with '\n' character.

### Request & Response

1. Request to server with 'Merry'.
2. Server response with 'Christmas!'.
3. Request to server with 'Chicken'.
4. Server response with fried chicken art with "C" delimiter.
5. End of protocol, close connection.

## Appendix
<img width="316" alt="image" src="https://github.com/doew/MCCAP/assets/39424676/aa033e9f-918e-4bfd-8398-d1d56bcc2249">
