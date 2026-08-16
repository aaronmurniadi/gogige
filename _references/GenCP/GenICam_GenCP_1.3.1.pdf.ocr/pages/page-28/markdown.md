|  GEN<i>CAM |   | ![img-28.jpeg](img-28.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

If compared to UDP/IP, a prefix would be omitted since everything is covered by the underlying protocol. For a serial connection, we would not need to cover addressing because it is not part of the technology. We need to identify a communication channel (by channel_id) and we need a CRC and we need a preamble to identify the protocol.

- The Common Command Data section contains data which describes the command. For example, this section contains the actual command identifier and the request_id.
- The Command Specific Data section is technology agnostic. It carries data which is specific for a given command. For example, for a read command it would contain the address to read from and the number of bytes to read.
- The Postfix section is technology specific. It carries for example a CRC checksum in case it is needed for a given technology. This section is only mandatory if defined for a given technology.