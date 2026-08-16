### Timed Pulse

The light is normally off. When a trigger is received the light waits for a delay then pulses for a defined time. The main configuration parameters are the brightness, delay before pulse and pulse width.

// A pulse sent to the light upon reception of a delayed trigger.

TimerSelector = Timer0; // Select the Timer1.
TimerDelay[Timer0] = 20; // Set the delay before the pulse starts to 20us.
TimerDuration[Timer0] = 200; // Set the pulse duration to 200us.
TimerTriggerSource = Line8; // Set external input Line 8 as Trigger source for the Timer1.
TimerTriggerActivation = RisingEdge; // Choose rising edge to trigger the Timer.
TimerTriggerArmDelay[Timer0] = 500; // Set the minimum delay before next pulse to 500us.
LightControllerSelector=LightController0; // Select the Light Controller.
LightControllerSource[LightController0] = Timer1Active; // Put the Timer Output as source.
LightVoltageRating[LightController0] = 24.0; // Max Voltage the light supports.
LightBrightness[LightController0] = 250; // Overdrive to 250 % brightness for a short period.
LineSelector = Line1; // Select the external output Line 1.
LineSource[Line1]=LightController0; // Put the Light controller as its source.

Note that a dedicated Line can be hard-wired to a light controller (no setting required).

In the case of embedded lighting, the light controller would be connected directly to the device's embedded light (no setting required). The above features should still be present for information purpose but read only.