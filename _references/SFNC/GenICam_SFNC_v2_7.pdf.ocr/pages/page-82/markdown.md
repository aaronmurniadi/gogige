|  Selector] |  |  |  |  |  |   |
| --- | --- | --- | --- | --- | --- | --- |
|  GevFirstURL | O | IString | R | - | I | Indicates the first URL to the GenICam XML device description file.  |
|  GevSecondURL | O | IString | R | - | I | Indicates the second URL to the GenICam XML device description file.  |
|  GevNumberOfInterfaces | O | IInteger | R | - | I | This feature is deprecated (See DeviceLinkSelector).  |
|  GevPersistentIPAddress[GevInterfaceSelector] | O | IInteger | R/W | - | B | Controls the Persistent IP address for this logical link.  |
|  GevPersistentSubnetMask[GevInterfaceSelector] | O | IInteger | R/W | - | B | Controls the Persistent subnet mask associated with the Persistent IP address on this logical link.  |
|  GevPersistentDefaultGateway[GevInterfaceSelector] | O | IInteger | R/W | - | B | Controls the persistent default gateway for this logical link.  |
|  GevMessageChannelCount | O | IInteger | R | - | I | This feature is deprecated (See DeviceEventChannelCount).  |
|  GevStreamChannelCount | O | IInteger | R | - | I | This feature is deprecated (See DeviceStreamChannelCount).  |
|  GevHeartbeatTimeout | O | IInteger | R/W | ms | G | This feature is deprecated (See DeviceLinkHeartbeatTimeout).  |
|  GevTimestampTickFrequency | O | IInteger | R | Hz | I | This feature is deprecated (See the increment of the TimestampLatchValue feature).  |
|  GevTimestampControlLatch | O | ICommand | W | - | I | This feature is deprecated (See TimestampLatch).  |
|  GevTimestampControlReset | O | ICommand | W | - | I | This feature is deprecated (See TimestampReset).  |
|  GevTimestampValue | O | IInteger | R | - | I | This feature is deprecated (See TimestampLatchValue).  |
|  GevDiscoveryAckDelay | O | IInteger | R/(W) | ms | E | Indicates the maximum randomized delay the device will wait to acknowledge a discovery command.  |
|  GevIEEE1588 | O | IBoolean | R/W | - | I | This feature is deprecated (See PtpEnable).  |
|  GevIEEE1588ClockAccuracy | O | IEnumeration | R/(W) | - | I | This feature is deprecated (See PtpClockAccuracy).  |
|  GevIEEE1588Status | O | IEnumeration | R | - | I | This feature is deprecated (See PtpStatus).  |
|  GevGVCPExtendedStatusCodesSelector | O | IEnumeration | R/W | - | G | Selects the GigE Vision version to control extended status codes for.  |
|  GevGVCPExtendedStatusCodes[GevGVCPExtendedStatusCodesSelector] | O | IBoolean | R/W | - | G | Enables the generation of extended status codes.  |
|  GevGVCPPendingAck | O | IBoolean | R/W | - | G | Enables the generation of PENDING_ACK.  |