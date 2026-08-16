|  GevGVCPHeartbeatDisable | O | IBoolean | R/W | - | I | This feature is deprecated (See DeviceLinkHeartbeatMode).  |
| --- | --- | --- | --- | --- | --- | --- |
|  GevGVCPPendingTimeout | O | IInteger | R | - | I | This feature is deprecated (See DeviceLinkCommandTimeout).  |
|  GevPrimaryApplicationSwitchoverKey | O | IInteger | W-O | - | G | Controls the key to use to authenticate primary application switchover requests.  |
|  GevGVSPExtendedIDMode | O | IEnumeration | R/(W) | - | E | Enables the extended IDs mode.  |
|  GevCCP | O | IEnumeration | R/W | - | G | Controls the device access privilege of an application.  |
|  GevPrimaryApplicationSocket | O | IInteger | R | - | G | Returns the UDP source port of the primary application.  |
|  GevPrimaryApplicationIPAddress | O | IInteger | R | - | G | Returns the address of the primary application.  |
|  GevMCPHostPort | O | IInteger | R/W | - | G | Controls the port to which the device must send messages.  |
|  GevMCDA | O | IInteger | R/W | - | G | Controls the destination IP address for the message channel.  |
|  GevMCTT | O | IInteger | R/W | ms | G | Provides the transmission timeout value in milliseconds.  |
|  GevMCRC | O | IInteger | R/W | - | G | Controls the number of retransmissions allowed when a message channel message times out.  |
|  GevMCSP | O | IInteger | R | - | G | This feature indicates the source port for the message channel.  |
|  GevStreamChannelSelector | O | IInteger | R/W | - | E | Selects the stream channel to control.  |
|  GevSCCFGPacketResendDestination[GevStreamChannelSelector] | O | IBoolean | R/W | - | G | Enables the alternate IP destination for stream packets resent due to a packet resend request.  |
|  GevSCCFGAllInTransmission[GevStreamChannelSelector] | O | IBoolean | R/W | - | G | Enables the selected GVSP transmitter to use the single packet per data block All-in Transmission mode.  |
|  GevSCCFGUnconditionalStreaming[GevStreamChannelSelector] | O | IBoolean | R/W | - | G | Enables the camera to continue to stream, for this stream channel, if its control channel is closed or regardless of the reception of any ICMP messages (such as destination unreachable messages).  |
|  GevSCCFGExtendedChunkData[GevStreamChannelSelector] | O | IBoolean | R/W | - | G | Enables cameras to use the extended chunk data payload type for this stream channel.  |
|  GevSCPDirection[GevStreamChannelSelector] | O | IEnumeration | R | - | I | This feature is deprecated (See DeviceStreamChannelType).  |
|  GevSCPInterfaceIndex[GevStreamChannelSelector] | O | IInteger | R/(W) | - | G | Index of the logical link to use.  |