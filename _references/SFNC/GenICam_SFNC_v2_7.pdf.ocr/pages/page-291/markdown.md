|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 14 Action Control

The Action chapter describes all features related to Action Signals in the device.

### 14.1 Action usage model

Action Signals are a method to trigger actions in multiple devices at the same time (depending on the specific transport layer). Action Signals are used in the device in the same way as e.g. digital input lines.

One possible use for action signals is to raise a FrameStart trigger in multiple devices at the same time.

On most transport layers Action Signals are implemented using broadcast protocol messages. To allow a finegrained control which devices are allowed to react on the broadcasted action protocol message, the features ActionDeviceKey, ActionGroupKey and ActionGroupMask define filter conditions.

Each action protocol message contains an action device key, action group key and an action group mask. If the device detects a match between this protocol information and one of the actions selected by ActionSelector it raises the corresponding Action Signal.

Usage Examples:

/* Triggered Single Frame acquisition using the Action Signal 1. */

AcquisitionMode = SingleFrame;
TriggerSelector = FrameStart;
TriggerMode = On;
TriggerSource = Action1;

ActionDeviceKey = 0x12345678;
ActionSelector = 1;
ActionGroupKey = 0x1;
ActionGroupMask = 0x1;

AcquisitionStart();

// Here the Device is ready to receive the Action Command
// from an external source.

/* Generates a 200us Timer pulse (Strobe) on the physical output Line 2.
The Timer pulse is started using a trigger coming from Action Signal 3.
*/