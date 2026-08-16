|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

The current rating is generally more useful as it allows safe overdriving. Effectively these values give the maximum DC values for the light.

Once the rating is known, the light can be driven at any brightness from 0% to 100% of its rating.

Sometime, lights can also be overdriven. In this case, these rating values can be exceeded when the light is pulsed for a short time at more than its 100% value. This has huge benefits in machine vision as it can give more than 10 times brightness from the light in some cases. When overdriving, setting a proper minimal delay between 2 consecutive pulses is very important to prevent the light being damaged.

Most lighting controllers drive the light with a DC constant current. Some other controller type use use Pulse Width Modulation (PWM) to control the brightness of the light. When PWM is used, this is generally managed internally within the controller and is not a parameter that the user needs to control.

### Lighting control scenarios:

A light can be controlled in a number of ways. Common ways are:

#### Static state

The light is On or Off and controlled statically. The main control parameter is the lighting controller brightness.

// Continuous current sent to the light.

LineSelector = Line1; // Select the external output Line 1.
LineSource[Line1] = LightController0; // Put the static UserOutput1 bit as its source.
LightControllerSelector = LightController0; // Select the Light Controller.
LightVoltageRating[LightController0] = 24.0; // Max Voltage the light supports.
LightBrightness[LightController0] = 100; // Normal 100 % brightness.
LightControllerSource[LightController0] = UserOutput1; // Put the static UserOutput1 bit as its source.
UserOutputSelector = UserOutput1; // Select a static User output bit.
UserOutputValue = True; // Put it On (or Off).

#### External Pulse

The light is controlled by the signal coming from an external digital input signal. The main configuration parameter is the brightness.

// A pulse received on an input pin is sent to the light (keeping the same pulse duration).

LightControllerSelector = LightController0; // Select the Light Controller.
LightControllerSource[LightController0] = Line8; // Put the input Line8 as source.
LightVoltageRating[LightController0] = 24.0; // Max Voltage the light supports in continuous.
LightBrightness[LightController0] = 100; // Normal 100 % Light brightness.
LineSelector = Line1; // Select the external output Line 1.
LineSource[Line1] = LightController0; // Put the Light controller 1 as its source.