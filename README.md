# n2n-encryption
It's a randomly generated key encryption tool in go

It's best used on small strings - because the key will be as long as the string. 
For example, on a 30kb file the key will be 166kb and the encrypted file will be over 1GB in size.

1) type in a string
2) encode
3) this will encode the string using a randomly generated key that's as long as the string is in bytes
4) you can save the string and key (saved in seperate JSON files, one starting with 'string' and the other starting with 'key')
5) you can load them as well.
6) you can also load in strings (imported in correct JSON format below) and encode them.

{"storageString":"INSERTSTRINGHERE","storageBy":null,"storageInt":null}
save the file as stringimport.JSON
and when you want to load it - type in 'import'.
as an example

That's all folks for now

TODO:

Create API to go along with this.
