#Export mongodb environment variable before running the rest api export 
MONGO_URI_FOR_GO="mongodb+srv://user:password@hosturl"

#generate keys: openssl req -x509 -nodes -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365