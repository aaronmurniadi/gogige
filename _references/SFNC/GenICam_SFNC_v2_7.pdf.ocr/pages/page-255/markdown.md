|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Timers are readable and can be used to measure the duration of internal or external signals. Timers can be cascaded to increase their range if necessary. A Timer can also be used to generate a timed strobe pulse with an optional delay before activation.

To set the destination output Line of the Timer pulse, see for example **Timer1Active** entry of the **LineSource** feature.

Note that Timers can also be used to generate an Event when a predetermined maximum value (or duration) is reached. See the **EventSelector** feature.

Below we show examples of typical use cases of timer's control features in C/C++ pseudo-code.

For simplicity, the object name is omitted (e.g. **TimerValue** instead of **Camera.TimerValue**) and the default state of the device is assumed.

/* Generates a 300 us strobe pulse coming from the Timer 1 when a rising edge trigger is detected on the physical Line 2 of the device connector.

*/

TimerSelector = Timer1;
TimerDuration = 300;
TimerTriggerActivation = RisingEdge;
TimerTriggerSource = Line2;

/* Generates a 200us Timer pulse (Strobe) delayed by 100 us on the physical output Line 2.

The Timer pulse is started using a trigger coming from physical input Line 1.

*/

TimerSelector = Timer1;
TimerDuration = 200;
TimerDelay = 100;
TimerTriggerSource = Line1;
TimerTriggerActivation = RisingEdge;
LineSelector = Line2;
LineMode = Output;
LineSource = Timer1Active;

/* Use of a Timer to measure the length in microseconds of a negative pulse on the physical input Line1. An Event is also generated to the host application to signal the end of the pulse.

*/