+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|LI | VN  |Mode |    Stratum     |     Poll      |  Precision   |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Root Delay                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Root Dispersion                       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                          Reference ID                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                     Reference Timestamp (64)                  +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                      Origin Timestamp (64)                    +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                      Receive Timestamp (64)                   +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

+                      Transmit Timestamp (64)                  +

+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

.                    Extension Field 1 (variable)               .

+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

.                    Extension Field 2 (variable)               .

+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                          Key Identifier                       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

|                            dgst (128)                         |

+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Flags:
    LI Leap Indicator (leap): 2-bit integer
        0     | no warning                           
        1     | last minute of the day has 61 seconds
        2     | last minute of the day has 59 seconds
        3     | unknown (clock unsynchronized) 

    VN Version Number (version): 3-bit integer representing the NTP version number

    Mode (mode): 3-bit integer representing the mode
        0     | reserved                 
        1     | symmetric active         
        2     | symmetric passive        
        3     | client                   
        4     | server                   
        5     | broadcast                
        6     | NTP control message      
        7     | reserved for private use

    Stratum (stratum): 8-bit integer representing the stratum
        0      | unspecified or invalid                             
        1      | primary server (e.g., equipped with a GPS receiver)
        2-15   | secondary server (via NTP)                         
        16     | unsynchronized                                     
        17-255 | reserved

    Poll: 8-bit signed integer representing the maximum interval between successive messages

    Precision: 8-bit signed integer representing the precision of the system clock

Root Delay (rootdelay): Total round-trip delay to the reference clock

Root Dispersion (rootdisp): Total dispersion to the reference clock

Reference ID (refid): 32-bit code identifying the particular server or reference clock

Reference Timestamp: Time when the system clock was last set or corrected

Origin Timestamp (org): Time at the client when the request departed for the server

Receive Timestamp (rec): Time at the server when the request arrived from the client

Transmit Timestamp (xmt): Time at the server when the response left for the client (roundtrip)

Destination Timestamp (dst): Time at the client when the reply arrived from the server
