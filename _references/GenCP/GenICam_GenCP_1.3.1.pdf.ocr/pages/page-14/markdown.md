|  GEN<i>CAM |   | ![img-8.jpeg](img-8.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

## 3. Operation

### 3.1. Protocol

#### 3.1.1. Command & Acknowledge Mechanism

The protocol uses a command/acknowledge pattern. On each channel each entity has a defined role of being either a “command sender and acknowledge receiver” or a “command receiver and acknowledge sender”. It is defined in the BRM which channel acts as a command channel from the host to the device, and which channel is used for the opposite direction from the device to the host. The command sender sends a command and waits for the acknowledge packet. The command receiver receives the command, acts according to the command, and sends the acknowledge packet with the result.

The communication on the default communication channel defines the role of an entity. The sender of a command on the default communication channel is called the host. The command receiver on the default communication channel is called the (remote) device.

A command packet contains a number called command_id, which specifies the action to be executed by the receiver and some additional data to be used when executing the command. The command receiver is expected to process the command and return the result to the sender of the command using an acknowledge packet.

There are commands which always need an acknowledge packet (for example ReadMem) and commands where the acknowledge packet is optional (for example WriteMem). The demand for an acknowledge packet is indicated by a bit in the command packet. In case no acknowledge packet is requested, it is recommended for the command sender to wait the Maximum Device Response Time before the next command is sent.

All commands on a channel are sent sequentially. After a command has been sent, the command sender must wait for an acknowledge packet if requested or wait for a timeout and process the failure before the next command may be sent.

Each command is sent with a sequentially incremented request id. This id allows resending a command in case of a failure. A successful communication would follow this schema: