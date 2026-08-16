Version 2.7.1

Standard Features Naming Convention

|  CxpErrorCounterStatus[CxpConnection Selector][CxpErrorCounterSelector] | R | IEnumeration | R | - | E | Returns the current status of the selected Cxp Error Counter on the connection selected by CxpConnectionSelector.  |
| --- | --- | --- | --- | --- | --- | --- |
|  CxpPoCxpAuto | O | ICommand | W | - | E | Activate automatic control of the Power over CoaXPress (PoCXP) for the Link.  |
|  CxpPoCxpTurnOff | O | ICommand | W | - | E | Disable Power over CoaXPress (PoCXP) for the Link.  |
|  CxpPoCxpTripReset | O | ICommand | W | - | E | Reset the Power over CoaXPress (PoCXP) Link after an over-current trip on the Device connection(s).  |
|  CxpPoCxpStatus | O | IEnumeration | R | - | E | Returns the Power over CoaXPress (PoCXP) status of the Device.  |
|  CxpFirstLineTriggerWithFrameStart | O | IBoolean | R/(W) | - | E | Specifies if a FrameStart trigger also triggers the first LineStart at the same time.  |