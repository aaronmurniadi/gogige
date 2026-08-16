|  GEN<i>CAM |   | ![img-29.jpeg](img-29.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

4.3.1. Command Packet Layout

|  Width (Bytes) | Offset (Bytes) | Description  |   |   |
| --- | --- | --- | --- | --- |
|  Prefix  |   |   |   |   |
|  2 | 0 | flagsFlags to enable/disable command options or to provide additional info on the specific command.  |   |   |
|   |   |  Bit offset (lsb << x) | Width (bits) | Description  |
|   |   |  0 | 14 | Reserved, set to 0  |
|   |   |  14 | 1 | RequestAckIf set the sender requests an acknowledge packet from the command receiver.  |
|   |   |  15 | 1 | CommandResendIf set the command is sent as a retry of a previous sent that failed.  |
|  2 | 2 | command_idcommand_id as specified in the Command ID chapter 4.3.3  |   |   |
|  2 | 4 | lengthLength of the Specific Command Data depending on the command ID not including Prefix, Postfix and CCD  |   |   |
|  2 | 6 | request_idSequential number to identify a single command. This id is provided by the command sender and incremented every time a new command is issued.  |   |   |
|  SCD  |   |   |   |   |
|  Postfix  |   |   |   |   |

Table 4 – Common Command Data