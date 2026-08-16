CounterSelector = Counter1;
CounterEventSource = LineStart;
CounterDuration = 200;
CounterTriggerSource = FrameStart;
CounterResetSource = CounterTrigger;

Register(Camera.EventCounter1End,CallbackDataObject,CallbackFunctionPtr)
EventSelector = Counter1End;
EventNotification = On;
AcquisitionMode = Continuous;
AcquisitionStart();
...
AcquisitionStop();

### 10.3 Counter features

This section describes the Counter features.

#### 10.3.1 CounterSelector

|  Name | CounterSelector  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Counter0 (If 0 based), Counter1, Counter2, ...  |

Selects which Counter to configure.

Possible values are:

- Counter0: Selects the counter 0.
- Counter1: Selects the counter 1.
- Counter2: Selects the counter 2.