|  ImageCompressionJPEGFormatOption | O | IEnumeration | R/W | - | E | When JPEG is selected as the compression format, a device might optionally offer better control over JPEG-specific options through this feature.  |
| --- | --- | --- | --- | --- | --- | --- |

### 2.3 Acquisition Control

Contains the features related to image acquisition, including trigger and exposure control (See the Acquisition Control chapter for details).

Table 2-3: Acquisition Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  AcquisitionControl | R | ICategory | R | - | B | Category for the acquisition and trigger control features.  |
|  AcquisitionMode | R | IEnumeration | R/(W) | - | B | Sets the acquisition mode of the device.  |
|  AcquisitionStart | R | ICommand | (R)/W | - | B | Starts the Acquisition of the device.  |
|  AcquisitionStop | R | ICommand | (R)/W | - | B | Stops the Acquisition of the device at the end of the current Frame.  |
|  AcquisitionStopMode | O | IEnum | R/W | - | E | Controls how the AcquisitionStop command and the acquisition stopped using a trigger (e.  |
|  AcquisitionAbort | R | ICommand | (R)/W | - | E | Aborts the Acquisition immediately.  |
|  AcquisitionArm | O | ICommand | (R)/W | - | E | Arms the device before an AcquisitionStart command.  |
|  AcquisitionFrameCount | R | IInteger | R/W | - | B | Number of frames to acquire in MultiFrame Acquisition mode.  |
|  AcquisitionBurstFrameCount | O | IInteger | R/W | - | B | Number of frames to acquire for each FrameBurstStart trigger.  |
|  AcquisitionFrameRate | R | IFloat | R/W | Hz | B | Controls the acquisition rate (in Hertz) at which the frames are captured.  |
|  AcquisitionFrameRateEnable | R | IBoolean | R/W | - | E | Controls if the AcquisitionFrameRate feature is writable and used to control the acquisition rate.  |
|  AcquisitionLineRate | R | IFloat | R/W | Hz | B | Controls the rate (in Hertz) at which the Lines in a Frame are captured.  |
|  AcquisitionLineRateEnable | R | IBoolean | R/W | - | E | Controls if the AcquisitionLineRate feature is writable and used to control the acquisition rate.  |
|  AcquisitionStatusSelector | R | IEnumeration | R/W | - | E | Selects the internal acquisition signal to read using AcquisitionStatus.  |