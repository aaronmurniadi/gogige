/* Linescan continuous acquisition with Hardware Frame and Line trigger. */

AcquisitionMode = Continuous;
TriggerSelector = FrameStart;
TriggerMode = On;
TriggerActivation = RisingEdge;
TriggerSource = Line1;
TriggerSelector = LineStart;
TriggerMode = On;
TriggerActivation = RisingEdge;
TriggerSource = Line2;
AcquisitionStart();
...
AcquisitionStop();

/* Framescan continuous acquisition with Hardware Frame trigger and the Exposure duration controlled by the Trigger pulse width.

*/

AcquisitionMode = Continuous;
TriggerSelector = FrameStart;
TriggerMode = On;
TriggerActivation = RisingEdge;
TriggerSource = Line1;
ExposureMode = TriggerWidth;
AcquisitionStart();
...
AcquisitionStop();