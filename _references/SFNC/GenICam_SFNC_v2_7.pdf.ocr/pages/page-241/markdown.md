|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

# 10 Counter and Timer Control

This chapter lists all features that relates to control and monitoring of Counters and Timers.

Note that Counters and Timers can also be used to generate an Event when a predetermined maximum count (or duration) is reached. See the EventSelector feature.

## 10.1 Counter and Timer category

This section define the category that includes Counter's and Timer's features.

### 10.1.1 CounterAndTimerControl

|  Name | CounterAndTimerControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains the Counter and Timer control features.

## 10.2 Counter usage model

This section describes the Counters usage model.

A Counter is used to count internal events (FrameStart, Timer1End, ...), I/O external events (Input Line rising edge, ...) and even clock ticks. It can be reset, read or written at anytime. Counters can also be cascaded to increase their range if necessary.

To set the destination output Line of a Counter, see for example Counter1Active entry of the LineSource feature.

Note that Counters can also be used to generate an Event when a predetermined maximum count (or duration) is reached. See the EventSelector feature.

Below we show examples of typical use cases of counter's control features in C/C++ pseudo-code.

For simplicity, the object name is omitted (e.g. CounterValue instead of Camera.CounterValue) and the default state of the device is assumed.