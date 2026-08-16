### 5.7.1 ExposureMode

|  Name | ExposureMode  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Off Timed TriggerWidth TriggerControlled  |

Sets the operation mode of the Exposure.

Possible values are:

- **Off**: Disables the Exposure and let the shutter open.
- **Timed**: Timed exposure. The exposure duration time is set using the ExposureTime or ExposureAuto features and the exposure starts with the FrameStart or LineStart.
- **TriggerWidth**: Uses the width of the current Frame or Line trigger signal(s) pulse to control the exposure duration. Note that if the Frame or Line **TriggerActivation** is RisingEdge or LevelHigh, the exposure duration will be the time the trigger stays High. If **TriggerActivation** is FallingEdge or LevelLow, the exposure time will last as long as the trigger stays Low.
- **TriggerControlled**: Uses one or more trigger signal(s) to control the exposure duration independently from the current Frame or Line triggers. See **ExposureStart**, **ExposureEnd** and **ExposureActive** of the **TriggerSelector** feature.

Note also that **ExposureMode** has priority over the Exposure Trigger settings defined using **TriggerSelector=Exposure...** and defines which trigger (if any) is active.

For example, if:

ExposureMode = Timed;
ExposureTime = 200;

Then the Exposure will be controlled using the **ExposureTime** Feature, even if the following code is done: