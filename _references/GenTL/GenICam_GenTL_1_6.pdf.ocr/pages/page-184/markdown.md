|  ![img-276.jpeg](img-276.jpeg)CAN |   | ![img-277.jpeg](img-277.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|  DeviceID | IString | R | Interface wide unique identifier of this device.  |
|  DeviceVendorName | IString | R | Name of the device vendor.  |
|  DeviceModelName | IString | R | Name of the device model.  |
|  DeviceType | IEnumeration | R | Identifies the transport layer technology of the device. See chapter 6.6.1 for possible values.  |

Table 7-10: Stream enumeration features

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|  StreamSelector | IInteger | R/W | Selector for the different stream channels.The selector is 0 based in order to match the index of the C interface.  |
|  StreamID [StreamSelector] | IString | R | Device unique ID for the stream, e.g., a GUID.  |

#### 7.1.4 Data Stream Module

This section lists all features which must be available in the stream module: Port functions use the DS_HANDLE to access the features. The Port access for this module is mandatory.

Table 7-11: Data Stream information features

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|  StreamID | IString | R | Device unique ID for the data stream, e.g., a GUID.  |
|  StreamAnnouncedBufferCount | IInteger | R | Number of announced (known) buffers on this stream. This value is volatile. It may change if additional buffers are announced and/or buffers are revoked by the GenTL Consumer.  |
|  StreamAcquisitionModeSelector | IEnumeration | R/W | Available buffer handling modes of this Stream. Deprecated. Use “StreamBufferHandlingMode” instead. Value: “Default” (see chapter 5 Acquisition Engine page 42ff)  |
|  StreamBufferHandlingMode | IEnumeration | R/W | Available buffer handling modes of this Stream. Value: “Default” (see chapter 5 Acquisition Engine page 42ff)  |
|  StreamAnnounceBufferMinimum | IInteger | R | Minimal number of buffers to announce to enable selected buffer handling mode.  |