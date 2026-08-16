|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 5.4 Acquisition and Trigger features usage examples

This section shows examples of typical use cases of acquisition and control of the SFNC features in C/C++ pseudo-code.

For simplicity, the object name is omitted (e.g., AcquisitionStart() instead of Camera.AcquisitionStart()) and the default state of the camera is assumed (e.g. Ready for a continuous acquisition start without trigger).

/* Continuous acquisition when the camera is in its reset state. */

AcquisitionMode = Continuous;
AcquisitionStart();
...
AcquisitionStop();

/* Single Frame acquisition in Hardware trigger mode using the external I/O Line 3. */

AcquisitionMode = SingleFrame;
TriggerSelector = FrameStart;
TriggerMode = On;
TriggerActivation = RisingEdge;
TriggerSource = Line3;
AcquisitionStart();