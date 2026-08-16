## 22 Light Control

The Lighting Control chapter describes the model and features related to the control of Lighting.

This features can be applied to dedicated Lighting Controller or to Cameras with integrated lighting control features.

Also, regular features such as the ones defined in Digital I/O Control and Counter and Timer Control can be used for lighting control (to generate strobe pulses, handle external IOs pins). Those features are already used in many cameras for basic Lighting Control (Ex: The sensor Exposure signal can be sent to an external output line to control a strobe light).

### 22.1 Existing Timer features for Light Control

Many “Counter and Timer Control” features can be used to control Strobe pulse generation.

For example, for pulse with duration in Time unit, the “Timer” features should be used.

TimerSelector = Timer1; // Select the Timer1.
TimerDelay[Timer1] = 20; // Set the delay before the pulse starts to 20us.
TimerDuration[Timer1] = 500; // Set the pulse duration to 500us.
TimerTriggerSource = Line8; // Set external input Line 8 as Trigger source for the Timer1.
TimerTriggerActivation = RisingEdge; // Choose rising edge to trigger the Timer1.
LineSelector = Line1; // Select the external output Line 1.
LineSource[Line1]=Timer1Active; // Put the Timer 1 output pulse its source.

“Counter” and “Rotary Encoder” features can also be used for discrete pulses counting or none time based pulse generation.

Note also that events can be generated and sent to the user for most possible state changes of the Timers and Counters.

For example:

EventSelector = Timer1End; // Select Timer 1 End of active period event.

EventNotification[Timer1End] = On; // Enable the event callback to the user.

### 22.2 Light Control Features usage

Rating of the Light and Overdriving:

Normally lights have one of the following characteristics specified:

Current rating: The current needed to illuminate the light at 100% brightness
Voltage rating: The voltage needed to illuminate the light at 100% brightness