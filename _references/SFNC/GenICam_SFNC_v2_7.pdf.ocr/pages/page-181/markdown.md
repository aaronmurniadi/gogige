|   | ImmediateWithPadding  |
| --- | --- |

Controls how the AcquisitionStop command and the acquisition stopped using a trigger (e.g. AcquisitionActive, FrameBurstActive, FrameActive or FrameEnd trigger), ends an ongoing frame. This feature is mainly used in Linescan devices where each line in a frame is acquired sequentially.

- **Complete**: When stopped during a frame, the device will continue acquisition of lines until the specified Height is reached to deliver a complete default size frame.

Note that if each line is triggered from an external source and this line trigger stops no frame is delivered, and an AcquisitionAbort is needed.

- **Immediate**: Acquisition stops immediately even during a frame and only the lines acquired at the time are delivered.

- **ImmediateWithPadding**: Acquisition stops immediately even during a frame but the remaining of the frame will be padded with data to deliver a complete default Height frame.

Note: How the receiver knows which data is valid is beyond the scope of this feature, it can for example be done using chunk information.

### 5.5.2 AcquisitionAbort

|  Name | AcquisitionAbort  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Aborts the Acquisition immediately. This will end the capture without completing the current Frame or waiting on a trigger. If no Acquisition is in progress, the command is ignored.

### 5.5.3 AcquisitionArm

|  Name | AcquisitionArm  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |