|  DeviceMessageChannelCount | O | IInteger | R | - | I | This feature is deprecated (See DeviceEventChannelCount).  |
| --- | --- | --- | --- | --- | --- | --- |
|  DeviceCharacterSet | O | IEnumeration | R | - | G | Character set used by the strings of the device.  |
|  DeviceReset | R | ICommand | W | - | G | Resets the device to its power up state.  |
|  DeviceIndicatorMode | O | IEnumeration | R/W | - | E | Controls the behavior of the indicators (such as LEDs) showing the status of the Device.  |
|  DeviceFeaturePersistenceStart | O | ICommand | (R)/W | - | G | Indicate to the device and GenICam XML to get ready for persisting of all streamable features.  |
|  DeviceFeaturePersistenceEnd | O | ICommand | (R)/W | - | G | Indicate to the device the end of feature persistence.  |
|  DeviceRegistersStreamingStart | R | ICommand | (R)/W | - | G | Prepare the device for registers streaming without checking for consistency.  |
|  DeviceRegistersStreamingEnd | R | ICommand | (R)/W | - | G | Announce the end of registers streaming.  |
|  DeviceRegistersCheck | R | ICommand | (R)/W | - | E | Perform the validation of the current register set for consistency.  |
|  DeviceRegistersValid | R | IBoolean | R | - | E | Returns if the current register set is valid and consistent.  |
|  DeviceRegistersEndianness | O | IEnumeration | R/(W) | - | G | Endianness of the registers of the device.  |
|  DeviceTemperatureSelector | O | IEnumeration | R/W | - | E | Selects the location within the device, where the temperature will be measured.  |
|  DeviceTemperature [DeviceTemperatureSelector] | O | IFloat | R | C | E | Device temperature in degrees Celsius (C).  |
|  DeviceClockSelector | O | IEnumeration | R/(W) | - | E | Selects the clock frequency to access from the device.  |
|  DeviceClockFrequency [DeviceClockSelector] | O | IFloat | R/(W) | Hz | E | Returns the frequency of the selected Clock.  |
|  DeviceSerialPortSelector | R | IEnumeration | R/(W) | - | E | Selects which serial port of the device to control.  |
|  DeviceSerialPortBaudRate [DeviceSerialPortSelector] | R | IEnumeration | R/(W) | - | E | This feature controls the baud rate used by the selected serial port.  |
|  Timestamp | R | IInteger | R | ns | E | Reports the current value of the device timestamp counter.  |
|  TimestampReset | O | ICommand | (R)/W | - | E | Resets the current value of the device timestamp counter.  |
|  TimestampLatch | O | ICommand | (R)/W | - | E | Latches the current timestamp counter into TimestampLatchValue.  |