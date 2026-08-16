|  GEN<i>CAM |   | ![img-3.jpeg](img-3.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

### 1.4. Acronyms

|  Name | Description  |
| --- | --- |
|  BRM | Bootstrap Register Map  |
|  ABRM | Technology Agnostic Bootstrap Register Map  |
|  SBRM | Technology Specific Bootstrap Register Map  |
|  Device | Device to be controlled, can be any entity, may not be a camera  |
|  Host | Controlling Master, can be any entity, may not be a PC  |
|  Link | Connection between a device and a host.  |
|  Channel | Logic communication channel between two entities. A Channel is always unidirectional.  |
|  Datagram | A single GenCP packet.  |
|  Entity | Either the Device or the host  |
|  DRT | Device Response TimeThe time a device needs to process a command not including the transfer time for the packet containing the command.  |
|  PTT | Packet Transfer TimeTime to transfer a message/command over a link at a given link speed.  |
|  URL | Uniform Resource Locator  |
|  CCD | Common Command DataSection within a GenCP command packet which is common to all commands.  |
|  SCD | Specific Command DataSection within a GenCP command packet which is specific to a given command.  |

Table 1 – Acronyms