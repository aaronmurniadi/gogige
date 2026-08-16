- Line0AnyEdge: The event will be generated when a Falling or Rising Edge is detected on the Line 0.
- Line1AnyEdge: The event will be generated when a Falling or Rising Edge is detected on the Line 1.
- LinkTrigger0RisingEdge: The event will be generated when a Rising Edge is detected on the LinkTrigger 0.
- LinkTrigger1RisingEdge: The event will be generated when a Rising Edge is detected on the LinkTrigger 1.
- LinkTrigger0FallingEdge: The event will be generated when a Falling Edge is detected on the LinkTrigger 0.
- LinkTrigger1FallingEdge: The event will be generated when a Falling Edge is detected on the LinkTrigger 1.
- LinkTrigger0AnyEdge: The event will be generated when a Falling or Rising Edge is detected on the LinkTrigger 0
- LinkTrigger1AnyEdge: The event will be generated when a Falling or Rising Edge is detected on the LinkTrigger 1.
- LinkSpeedChange: The event will be generated when the link speed has changed.
- ActionLate: The event will be generated when a valid scheduled action command is received and is scheduled to be executed at a time that is already past.
- Error: The event will be generated when the device encounter an error.
- Test: The test event will be generated when the device receives the TestEventGenerate command (EventNotification for the Test event is always On).
- PrimaryApplicationSwitch: The event will be generated when a primary application switchover has been granted (GigE Vision Specific).

### 15.3 EventNotification

|  Name | EventNotification[EventSelector]  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off On Once  |