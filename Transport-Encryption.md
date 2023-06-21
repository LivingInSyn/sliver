 __⚠️ NOTE:__ This document does not apply when mTLS or WireGuard are used.

# Versions 1.5.40+ 

The following keys are embedded in each implant at compile time, the server also stores these values in its database in addition to the SHA2-256 hash of the implant's peer public key:

1. Age public key of server 
2. Age implant "peer" public key
3. Age implant "peer" private key
4. A minisign signature of the implant's age peer public key (signed by server's private key)
5. The server's minisign public key 

### Server to Implant Key Exchange

1. Implant generates random 256-bit symmetric "session key"
2. Implant generates:
 * Calculate the HMAC-SHA2-256 of the session key, where the HMAC key is the SHA2-256 hash of the peer private key
 * Uses Age to encrypt session key and HMAC-SHA256 with server's public key
3. Implant sends `[SHA2-256 Hash of Public Key | Age Ciphertext ]` to server.
5. Server decrypts message with its private key
6. Server generates a session ID, encrypts it with the session key using ChaCha20Poly1305, and sends it back
7. All messages are encrypted with the session key using ChaCha20Poly1305 and associated with via the session ID
8. Each side stores a SHA2-256 hash of each message's ciphertext to detect replayed messages

### Implant to Implant Key Exchange (Pivots)

1. An implant starts a pivot listener (the listener)
2. Another implant connects to the listener (the initiator) 
3. The initiator sends its Age public key and the minisign signature of its public key
4. The listener verifies the initiator's public key is signed by the listener's server's minisign public key
5. The listener generates a random session key and encrypts it with the initiator's verified public key
6. The listener sends its Age public key, the minisign signature of its public key, and the encrypted session key back to the initiator
7. The initiator verifies the listener's public key is signed by the initiator's server's minisign public key
8. The initiator decrypts the session key using Age
9. All messages are encrypted with the session key using ChaCha20Poly1305
10. Each side stores a SHA2-256 hash of each message's ciphertext to detect replayed messages
11. The initiator then performs a key exchange with the server, using the server's embedded ECC public key, using the same pattern as described above. This prevents upstream implants (i.e. the listener) from being able to decrypt any traffic sent between the initiator and the server.

### Known Limitations

There are some known limitations, if you spot any or have ideas on improvements please file a ticket or contact us.

1. Perfect Forward Secrecy: We do get some forward secrecy, since only the public key encrypted version of the session key is sent over the wire; therefore recovery of the hard coded implant keys from the binary should not result in recovery of the session key. Only when the server's private key is compromised can the session key be recovered, which we currently don't consider to be a problem.
2. Implants can potentially be tracked via the hash of their public key. However, this value is implant specific, so in order to track the implant this way you'd have to already have a copy of the specific implant you want to track. At which point more effective tracking mechanisms like YARA rules could be employed.
3. While messages cannot be replayed, valid messages can potentially be re-ordered.

# Versions 1.5.0+

⚠️ This version of the key exchange is [vulnerable to MitM attack](https://github.com/BishopFox/sliver/security/advisories/GHSA-8jxm-xp43-qh3q)

The following keys are embedded in each implant at compile time, the server also stores these values in its database in addition to the SHA2-256 hash of the implant's public key:

1. ECC public key of server 
2. ECC implant public key
3. ECC implant private key
4. A minisign signature of the implant ECC public key (signed by server's private key)
5. The server's minisign public key 
6. TOTP shared secret (server-wide shared secret)

### Server to Implant Key Exchange

1. Implant generates random 256-bit symmetric "session key"
2. Implant generates:
 * Current TOTP code using SHA2-256, Unix UTC, 8-digit numeric code
 * SHA2-256 hash of its own ECC public key
 * Uses Nacl Box (Curve25519, XSalsa20, and Poly1305) to encrypt session key with server's public ECC key
3. Implant sends `[ TOTP Code | SHA2-256 Hash of Public Key | Nacl Box Ciphertext ]` to server, note: in this scheme no ECC keys (even public keys) are ever sent over the wire, instead we only send the hash of the public key.
4. Server verifies TOTP Code (can optionally be disabled per listener server-side).
5. Server uses the SHA2-256 hash of public key to look up the implant's full ECC public key in its database
6. Decrypts Nacl with sender public key + server private key
7. Server generates a session ID, encrypts it with the session key using ChaCha20Poly1305, and sends it back
8. All messages are encrypted with the session key using ChaCha20Poly1305 and associated with via the session ID
9. Each side stores a SHA2-256 hash of each message's ciphertext to detect replayed messages

### Implant to Implant Key Exchange (Pivots)

1. An implant starts a pivot listener (the listener)
2. Another implant connects to the listener (the initiator) 
3. The initiator sends its ECC public key and the minisign signature of its public key
4. The listener verifies the initiator's public key is signed by the listener's server's minisign public key
5. The listener generates a random session key and encrypts it with the initiator's verified public key
6. The listener sends its ECC public key, the minisign signature of its public key, and the encrypted session key back to the initiator
7. The initiator verifies the listener's public key is signed by the initiator's server's minisign public key
8. The initiator decrypts the session key using Nacl Box (Curve25519, XSalsa20, and Poly1305) 
9. All messages are encrypted with the session key using ChaCha20Poly1305
10. Each side stores a SHA2-256 hash of each message's ciphertext to detect replayed messages
11. The initiator then performs a key exchange with the server, using the server's embedded ECC public key, using the same pattern as described above but without the TOTP code. This prevents upstream implants (i.e. the listener) from being able to decrypt any traffic sent between the initiator and the server.

### Known Limitations

There are some known limitations, if you spot any or have ideas on improvements please file a ticket or contact us.

1. Perfect Forward Secrecy: We do get some forward secrecy, since only the public key encrypted version of the session key is sent over the wire; therefore recovery of the hard coded implant keys from the binary should not result in recovery of the session key. Only when the server's private key is compromised can the session key be recovered, which we currently don't consider to be a problem.
2. Implants can potentially be tracked via the hash of their public key. However, this value is implant specific, so in order to track the implant this way you'd have to already have a copy of the specific implant you want to track. At which point more effective tracking mechanisms like YARA rules could be employed.
3. Session initialization messages can be replayed within the validity period of the TOTP value. TOTP values are valid for 30 seconds + 30 second margin of error, so the session initialization message can be replayed within about ~60 second period without obtaining a new TOTP code. However, the implant must use the session key to register itself post-key exchange so replayed session initialization does not appear to be a security risk even outside of the restrictive window.
4. While messages cannot be replayed, valid messages can potentially be re-ordered.

# Versions <= 1.4.x

## Transport Encryption for HTTP(S) / DNS

When the implant cannot directly route TCP traffic back to the C2 server or redirector the implant may, when configured to do so, connect back to the C2 server over HTTP, HTTPS, or DNS. In an operational environment it may not be possible to establish a trusted HTTPS connect back to the server and HTTP/DNS do not implement any transport encryption so Sliver "brings it's own crypto" to all of these protocols including HTTPS. This allows us to establish secure connections even if the only way out of the network is over a HTTPS interception proxy.

The following keys are embedded in each implant at compile time:

1. ECC Public Key of Server CA 
2. ECC Client Public Key _(used for mTLS)_
3. ECC Client Private Key _(used for mTLS)_

In this context "client" refers to the implant. The CA public key (#1) is shared among all implants generated by a given server. The client certificate pair (#2, and #3) are unique per-binary.

### Key Exchange

1. Implant requests a public RSA key from the server in the "clear"
2. Server responds with a public RSA key in the "clear" (X.509 2048-bit)
3. Implant verifies RSA public key is signed by the trusted authority embedded at compile-time
4. Implant generates AES session key, encrypts it with the public RSA key, and sends it to the server
5. Server generates a session ID, encrypts it with the session key using AES-GCM-256, and sends it back
6. All messages are encrypted with the session key using AES-GCM-256 and associated with via the session ID
7. Each side stores a SHA2-256 hash of each message's ciphertext to detect replayed messages.

__Note:__ "Clear" data may still be encoded/obfuscated but is considered public. Session IDs are also considered public parameters. It's up to the protocol specific `transports` to handle any data encoding/obfuscation. Session keys are only stored in memory

```
[implant] ---[ Do you have an RSA Key? ]-----> [server]
[implant] <--[ Here is my RSA Key ]----------- [server]
[implant] ---[ RSA encrypted AES key ]-------> [server]
[implant] <--[ AES encrypted session ID ]----- [server]

[implant] <--[ AES encrypted session data ]--> [server]
```

### Security Goals
The security goals of this key exchange/setup are:

* Robust data secrecy, integrity, and authenticity
* Protection against MitM attacks
* Protection against replay attacks
* Protection against bit-flipping, padding oracle, etc. attacks

### Known Limitations

The key exchange should provide a reasonable amount of security, however when possible we still recommend using mTLS or WireGuard. There are some known limitations in this scheme that we plan to address in the future, if you spot any or have ideas on improvements please file a ticket or contact us.

1. __No implant-to-server authentication__: One of the nice properties of mutual TLS is that the implant authenticates itself to the server in addition to the server authenticating itself to the implant. In this model the implant never proves to the server it was generated by this server. Since the implant does have ECC keys known to the server in the future we plan to have the implant sign some type of nonce or timestamp during the key exchange.
2. __Session Initialization Replays__: While there is replay protection against any messages sent between the server and the implant the initial message that establishes the session can be replayed against the server, but should only result in another (un-associated) session getting created.
3. __Perfect Forward Secrecy__: We do get some forward secrecy, since only the public key encrypted version of the session key is sent over the wire; therefore recovery of the hard coded implant keys from the binary should not result in recovery of the session key. Only when the server's private key is compromised is forward secrecy broken. However, it my be desirable to do a full DH-exchange to protected against recovery of the server certificates too.