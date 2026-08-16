/* Counts the number of missed triggers during an Acquisition.*/

CounterSelector         = Counter1;
CounterTriggerSource    = AcquisitionStart;
CounterResetSource      = CounterTrigger;
CounterEventSource       = FrameTriggerMissed;
AcquisitionMode          = Continuous;
AcquisitionStart();
...
AcquisitionStop();
NbPulses                = CounterValue;
printf("%ld Triggers missed during the acquisition.", NbPulses);

/* Counts the number of rising edge signals received on Line1 during a frame.*/

CounterSelector         = Counter1;
CounterTriggerSource    = FrameStart;
CounterResetSource      = CounterTrigger;
CounterEventSource       = Line1;
CounterEventActivation   = RisingEdge;
AcquisitionMode          = SingleFrame;
AcquisitionStart();
...
NbPulses                = CounterValue;
printf("%ld Input Pulses received during the frame.", NbPulses);

/* Use a counter to generate an event at line 200 of each captured Frame
in a continuous acquisition.

*/