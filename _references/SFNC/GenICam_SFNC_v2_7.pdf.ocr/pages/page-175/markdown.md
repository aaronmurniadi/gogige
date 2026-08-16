|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

output line 2 to activate a light during the exposure time of each frame. The end of the sequence capture is signalled to the host with an acquisition end event.

*/

AcquisitionMode = MultiFrame;
AcquisitionFrameCount = 20;

TriggerSelector = FrameStart;
TriggerMode = On;
TriggerActivation = RisingEdge;
TriggerSource = Line1;

ExposureMode = Timed;
ExposureTime = 500;

LineSelector = Line2;
LineMode = Output;
LineInverter = True;
LineSource = ExposureActive

Register(Camera.EventAcquisitionEnd,CallbackDataObject,CallbackFunctionPtr)

EventSelector = AcquisitionEnd;
EventNotification = On;

AcquisitionStart();