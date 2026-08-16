|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

TimerSelector = Timer1;
TimerTrigger = Line1;
TimerTriggerActivation = LevelLow;
Register(Camera.EventLine1RisingEdge,CallbackDataObject,CallbackFunctionPtr)
EventSelector = Line1RisingEdge;
EventNotifications = On;
/* Wait for the event on the host to read the time. */
...
TimerSelector = Timer1;
PulseDuration = TimerValue;