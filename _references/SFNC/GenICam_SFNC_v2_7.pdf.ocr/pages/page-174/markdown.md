/* Multi-Frame acquisition started by a single Software trigger delayed by 1 millisecond.
The Trigger starts the whole sequence acquisition. The Exposure time for each frame is
set to 500 us.

*/

AcquisitionMode = MultiFrame;
AcquisitionFrameCount = 20;
TriggerSelector = AcquisitionStart;
TriggerMode = On;
TriggerSource = Software;
TriggerDelay = 1000;
ExposureMode = Timed;
ExposureTime = 500;
AcquisitionStart();
TriggerSoftware();

/* Continuous acquisition in Hardware trigger mode. The Frame triggers are Rising Edge signals
coming from the physical Line 2. The Exposure time is 500us. An exposure end event is also
sent to the Host application after the exposure of each frame to signal that the inspected part
can be moved. The timestamp of the event is also read.

*/

AcquisitionMode = Continuous;
TriggerSelector = FrameStart;
TriggerMode = On;
TriggerActivation = RisingEdge;
TriggerSource = Line2;
ExposureMode = Timed;
ExposureTime = 500;
Register(Camera.EventExposureEnd, CallbackDataObject, CallbackFunctionPtr)
EventSelector = ExposureEnd;
EventNotification = On;
AcquisitionStart();
...
// In the callback of the ExposureEnd event, get the event timestamp:
Timestamp = EventExposureEndTimestamp;
...
AcquisitionStop();

/* Multi-Frame acquisition with each frame triggered by a Hardware trigger on Line 1.
A negative pulse of the exposure signal duration (500us) is also sent to the physical