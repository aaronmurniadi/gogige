|  GEN<i>CAM |   | ![img-73.jpeg](img-73.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

## Appendix

### 1. Serial Port Implementations

This section specializes the generic protocol for the use over a serial link.

#### 1.1. Byteorder

For devices communicating over a serial link, the byte order of bootstrap registers and protocol fields is big-endian.

#### 1.2. Channel ID

The default channel_id for the control channel on a serial link is channel_id = 0.

#### 1.3. Packet Size

In order to maintain reasonable response times even with low link speeds, the packets must not exceed 1024 Bytes per packet.

#### 1.4. Serial Parameters

##### 1.4.1. Default port parameters

The link uses 8Bit, No Parity, 1 Stop Bit encoding and 9600 Baud per default. The Link can be switched to other communication parameters and/or higher baud rates after a communication has been established using the transport layer specific bootstrap registers.

##### 1.4.2. Changing port parameters

When switching to other communication parameters the procedure is as follows: