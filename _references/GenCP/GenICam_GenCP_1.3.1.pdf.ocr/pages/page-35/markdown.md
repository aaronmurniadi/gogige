|  GEN<i>CAM |   | ![img-33.jpeg](img-33.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 4.3.3. Command IDs

This chapter describes the command_ids for the command field in the Common Command Data section of a GenCP command packet. The layout of a 16bit command_id is as follows:

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 1 | Acknowledge Flag- Set this bit to 0 if the command_id belongs to a command- Set this bit to 1 if the command_id is used for an acknowledgement  |
|  1 | 14 | Command ValueNumber identifying a single command/acknowledge  |
|  15 | 1 | Custom Command Identifier- Set this bit to 0 to identify a standardized command value- Set this bit to 1 to mark a custom command value  |

Command_ids can either identify a command or an acknowledge.

Command_ids identifying a command must have the LSB cleared.

Command_ids identifying an acknowledgement must have the LSB set to 1.

Custom command_ids must have the most significant bit set (Hex 8xxx) so that they do not collide with future standard extensions.

Standardized command_ids are: