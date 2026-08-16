|  GEN<i>CAM |   | ![img-7.jpeg](img-7.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

A change in the Major Version Number indicates a significant feature change and a potential break in backward compatibility.

A change in the Minor Version Number indicates minor feature changes, bug fixes, text clarifications and assures backward compatibility.

### 2.5. CRC

The CRC checksum used on the packets depends on the underlying technology. If the underlying technology already provides a CRC, that service is used. If the underlying technology does not provide a CRC, the checksum is defined in the Appendix.

### 2.6. Link

A link is the physical end to end connection between a device and a host used for control communication. For example, for Camera Link Medium, despite the fact that there are two cables carrying data, there is only one serial link for the RS232 communication.

Each link can carry multiple logical communication channels. GenCP assumes a single link between a host and a device.

### 2.7. Channel

A channel is a logical communication path between two entities communicating over a link. There may be multiple logical channels on a single link. Each channel is identified by a unique id number. This number is used in the communication between two entities to identify the channel a packet belongs to. This is either part of the protocol layers below the protocol described here or in the PacketPrefix (see chapter 4.2), depending on the technology. This number is called “channel_id”. A channel’s communication is unidirectional, meaning that on a single channel, the sender and receiver side for commands and the sender and receiver side for acknowledges are fixed. Different logical channels may have different directions. The protocol also defines packet layouts and the communication scheme between a device and a host. This document assumes that for the master control channel the host is the command sender and the device is the command receiver even though the roles may change in real live.

#### 2.7.1. Default Channel

The default channel (first control channel) is technology dependent. For example, on Ethernet this would be a port number. For another technology it might be an arbitrary number.