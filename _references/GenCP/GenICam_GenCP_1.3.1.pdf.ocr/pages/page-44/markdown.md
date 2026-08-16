|  GEN<i>CAM |   | ![img-47.jpeg](img-47.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 4.4.10. Event Command

If the MultiEvent Supported bit is set in the Device Capability register and if the MultiEvent Enable bit is set in the Device Configuration register, a single Event Command can carry multiple separate events including their data. The host must parse a received Event Command to determine how many single events are contained in a given Event Command and to access one of them. If the packet is parsed, more events are expected until the length stated in the SCD section is exhausted. The first event is located at address 0 in the SCD section of the command. The event n would start at \(Offset(Bytes) = \sum_{k=0}^{n-1} event\_size(k)\) within the SCD section where n is the index of the event to access. In case a single event does not carry additional data, the \(event\_size\) field is to be set to 12. This way the upper software layers can see if an event packet carries multiple events. Even if the MultiEvent is supported and enabled, an Event Command packet can contain only one event. In this case, the size in the CCD section would match the \(event\_size\) field in the SCD section.

If MultiEvent is not supported or if the MultiEvent Enable bit in the Device Configuration register is not set the event_size field must be set to 0 (reserved) and the size of data is deduced from the SCD size as stored in the CCD section of the packet.

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  Prefix  |   |   |
|  CCD (command_id = EVENT_CMD)  |   |   |
|  2 | 0 | event_sizeIf the MultiEvent Supported bit is set in the Device Capability register and if the MultiEvent Enable bit is set in the Device Configuration register: Size of event data object in bytes including event_size, event_id, timestamp and optional data.Otherwise 0 to be backward compatible.  |
|  2 | 2 | event_idThe event_id is a number identifying an event source. The schema of the event_id follows the description in chapter 3.1.3.1  |
|  8 | 4 | timestamp64 bit timestamp value in ns as defined in the timestamp bootstrap register.  |
|  X | 12 | dataOptional event specific data.  |
|  Postfix  |   |   |

Table 17 – Event Command SCD-Fields