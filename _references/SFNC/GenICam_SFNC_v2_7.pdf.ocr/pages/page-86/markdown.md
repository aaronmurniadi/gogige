|  CxpLinkSharingEnable | O | IBoolean | R/W | - | E | Enable or disable the link sharing functionality of the device.  |
| --- | --- | --- | --- | --- | --- | --- |
|  CxpLinkSharingSubDeviceSelector | O | IInteger | R/W | - | E | Index of the sub device used in the Link Sharing.  |
|  CxpLinkSharingStatus[CxpLinkSharingSubDeviceSelector] | O | IEnumeration | R | - | E | This feature provides the data sharing status for the selected sub device.  |
|  CxpLinkSharingSubDeviceType | O | IEnumeration | R | - | E | This feature provides the type of sub device.  |
|  CxpLinkSharingHorizontalStripeCount | O | IInteger | R/(W) | - | E | This feature provides the number of horizontal stripes that the device implements.  |
|  CxpLinkSharingVerticalStripeCount | O | IInteger | R/(W) | - | E | This feature provides the number of vertical stripes that the device implements.  |
|  CxpLinkSharingHorizontalOverlap | O | IInteger | R/(W) | - | E | This feature provides the number of pixel overlap in the horizontal stripes that the device implements.  |
|  CxpLinkSharingVerticalOverlap | O | IInteger | R/(W) | - | E | This feature provides the number of pixel overlap in the vertical stripes that the device implements.  |
|  CxpLinkSharingDuplicateStripe | O | IInteger | R/(W) | - | E | This feature provides the duplicate count in striped system.  |
|  CxpConnectionSelector | R | IInteger | R/W | - | E | Selects the CoaXPress physical connection to control.  |
|  CxpConnectionTestMode[CxpConnectionSelector] | R | IEnumeration | R/W | - | E | Enables the test mode for an individual physical connection of the Device.  |
|  CxpConnectionTestErrorCount[CxpConnectionSelector] | R | IInteger | R/W | - | E | Reports the current connection error count for test packets received by the device on the connection selected by CxpConnectionSelector.  |
|  CxpSendReceiveSelector | R | IEnumeration | R/W | - | E | Selects which one of the send or receive features to control.  |
|  CxpConnectionTestPacketCount[CxpConnectionSelector][CxpSendReceiveSelector] | R | IInteger | R/W | - | E | Reports the current count for the test packets on the connection selected by CxpConnectionSelector.  |
|  CxpErrorCounterSelector | R | IEnumeration | R/W | - | E | Selects which Cxp Error Counter to read or reset.  |
|  CxpErrorCounterReset[CxpConnectionSelector][CxpErrorCounterSelector] | R | ICommand | (R)/W | - | E | Resets the selected Cxp Error Counter on the connection selected by CxpConnectionSelector.  |
|  CxpErrorCounterValue[CxpConnectionSelector][CxpErrorCounterSelector] | R | IInteger | R | - | E | Reads the current value of the selected Cxp Error Counter on the connection selected by CxpConnectionSelector.  |