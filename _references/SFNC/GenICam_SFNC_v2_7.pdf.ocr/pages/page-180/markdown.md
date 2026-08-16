Note that unless the AcquisitionArm was executed since the last feature change, the AcquisitionStart command must validate all the current features for consistency before starting the Acquisition. This validation will not be repeated for the subsequent acquisitions unless a feature is changed in the device.

If the AcquisitionStart feature is currently not writable (locked), the application must not start the acquisition and must avoid using the feature until the feature becomes writable again.

This feature is generally mandatory for transmitters and transceivers of most Transport Layers.

### 5.5.4 AcquisitionStop

|  Name | AcquisitionStop  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Stops the Acquisition of the device at the end of the current Frame. It is mainly used when AcquisitionMode is Continuous but can be used in any acquisition mode.

If the camera is waiting for a trigger, the pending Frame will be cancelled. If no Acquisition is in progress, the command is ignored.

This feature is generally mandatory for transmitters and transceivers of most Transport Layers.

### 5.5.1 AcquisitionStopMode

|  Name | AcquisitionStopMode  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IEnum  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Complete Immediate  |