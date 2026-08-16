|  GEN<i>CAM |   | ![img-17.jpeg](img-17.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

##### 3.1.3.1. Event ID

The source of an event on the Message Channel is identified by an event_id. An event_id is a 16-bit value. The bits in this value have the following meaning:

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 12 | Event ID  |
|  12 | 2 | ReservedSet to 0  |
|  14 | 2 | Namespace0 = GenCP Event ID1 = Technology specific Event ID2 = Device specific Event ID  |

Table 2 – Event ID

##### 3.1.3.2. GenCP Event ID Codes

|  Event ID (Hex) | Name | Description  |
| --- | --- | --- |
|  0x0000 | Error | Generic Error Event  |

Table 3 – GenCP Event IDs

#### 3.1.4. Failure

A failure on the Command Channel or the Message Channel is discovered through

- a corrupt CCD of a command or acknowledge packet
• a timeout waiting for an acknowledge
- an invalid (too short) packet (timeout waiting for the complete arrival)
• an incorrect packet header

##### 3.1.4.1. Corrupt Packet

A packet is corrupt if the transmission of the packet failed (e.g. a transmission failure caused the