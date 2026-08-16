### 5.5.2 AcquisitionMode

|  Name | AcquisitionMode  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | SingleFrame MultiFrame Continuous  |

Sets the acquisition mode of the device. It defines mainly the number of frames to capture during an acquisition and the way the acquisition stops.

Possible values are:

- SingleFrame: One frame is captured.
- MultiFrame: The number of frames specified by AcquisitionFrameCount is captured.
- Continuous: Frames are captured continuously until stopped with the AcquisitionStop command.

This feature is generally mandatory for transmitters and transceivers of most Transport Layers.

### 5.5.3 AcquisitionStart

|  Name | AcquisitionStart  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Starts the Acquisition of the device. The number of frames captured is specified by AcquisitionMode.

The Acquisition might be conditioned by various triggers (See the Trigger... features). An AcquisitionStart command must be sent to the device before the acquisition related triggers become effective.