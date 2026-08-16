|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 25.3 TestEventGenerate

|  Name | TestEventGenerate  |
| --- | --- |
|  Category | TestControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Generates a Test Event.

If this feature is present, an EventTestData category containing the EventTest and EventTestTimestamp features must be implemented in the EventControl category.

Note: The Test event does not need to be included in EventSelector since the notification of this event is always enabled.

### 25.4 TestPayloadFormatMode

|  Name | TestPayloadFormatMode  |
| --- | --- |
|  Category | TestControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Off (default) MultiPart GenDC  |

This feature allows setting a device in test mode and to output a specific payload format for validation of data streaming. This feature is intended solely for test purposes. The data can be real acquired data or any test pattern.