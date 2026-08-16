/* Continuous Acquisition of frames in bursts of 10 frames. Each burst is triggered by a Hardware trigger on Line 1. The end of each burst capture is signalled to the host with a FrameBurstEnd event.

*/

AcquisitionMode = Continuous;
AcquisitionBurstFrameCount = 10;

TriggerSelector = FrameBurstStart;
TriggerMode = On;
TriggerActivation = RisingEdge;
TriggerSource = Line1;

Register(Camera.EventFrameBurstEnd,CallbackDataObject,CallbackFunctionPtr)
EventSelector = FrameBurstEnd;
EventNotification = On;

AcquisitionStart();

...

// In the callback of the end of burst event, get the event timestamp:
Timestamp = EventExposureEndTimestamp;

...

AcquisitionStop();

/* Multi-Frame Acquisition of 50 frames in 5 bursts of 10 frames. Each burst is triggered by a Hardware trigger on Line 1.

*/

AcquisitionMode = MultiFrame;
AcquisitionFrameCount = 50;
AcquisitionBurstFrameCount = 10;

TriggerSelector = FrameBurstStart;
TriggerMode = On;
TriggerActivation = RisingEdge;
TriggerSource = Line1;

AcquisitionStart();